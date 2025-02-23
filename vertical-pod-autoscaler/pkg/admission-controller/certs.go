/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/fsnotify/fsnotify"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/klog/v2"
)

type certsConfig struct {
	clientCaFile, tlsCertFile, tlsPrivateKey *string
	reload                                   *bool
}

func readFile(filePath string) []byte {
	res, err := os.ReadFile(filePath)
	if err != nil {
		klog.ErrorS(err, "Error reading certificate file", "file", filePath)
		return nil
	}
	klog.V(3).InfoS("Successfully read bytes from file", "bytes", len(res), "file", filePath)
	return res
}

type certReloader struct {
	tlsCertPath           string
	tlsKeyPath            string
	clientCaPath          string
	cert                  *tls.Certificate
	mu                    sync.RWMutex
	mutatingWebhookClient admissionregistrationv1.MutatingWebhookConfigurationInterface
}

func (cr *certReloader) start(stop <-chan struct{}, wg *sync.WaitGroup) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	fmt.Println("Watching", cr.tlsCertPath, cr.tlsKeyPath, cr.clientCaPath)

	if cr.tlsCertPath != "" {
		if err = watcher.Add(cr.tlsCertPath); err != nil {
			return err
		}
	}
	if cr.tlsKeyPath != "" {
		if err = watcher.Add(cr.tlsKeyPath); err != nil {
			return err
		}
	}
	if cr.clientCaPath != "" {
		if err = watcher.Add(cr.clientCaPath); err != nil {
			return err
		}
	}

	go func() {
		defer watcher.Close()
		for {
			fmt.Println("In goroutine A")
			select {
			case event := <-watcher.Events:
				fmt.Println("In goroutine event")

				// we need to watch "Remove" events because Kubernetes uses symbolic links to point to ConfigMaps/Secrets volumes
				if !event.Has(fsnotify.Remove) && !event.Has(fsnotify.Create) && !event.Has(fsnotify.Write) {

					fmt.Println("In goroutine D", event.Op.String(), event.Name)

					continue
				}
				switch event.Name {
				case cr.tlsCertPath:
					fmt.Println("In goroutine B")

					klog.V(2).InfoS("New certificate found, reloading")
					if err := cr.load(); err != nil {
						klog.ErrorS(err, "Failed to reload certificate")

						// adrian
						data, err := os.ReadFile(cr.tlsCertPath)
						if err != nil {
							log.Fatalf("failed to read file: %v", err)
						}
						fmt.Println("BEFORE CERTIFICATE DEBUG")
						fmt.Println(string(data))
						fmt.Println("AFTER CERTIFICATE DEBUG")

					}
				case cr.tlsKeyPath:
					fmt.Println("In goroutine C")

					klog.V(2).InfoS("New certificate found, reloading")
					fmt.Println("In goroutine E")

					if err := cr.load(); err != nil {
						fmt.Println("In goroutine F")

						klog.ErrorS(err, "Failed to reload KEY")
						fmt.Println("In goroutine G")

						// adrian
						data, err := os.ReadFile(cr.tlsKeyPath)
						fmt.Println("In goroutine H")

						if err != nil {
							log.Fatalf("failed to read file: %v", err)
						}
						fmt.Println("BEFORE KEY DEBUG")
						fmt.Println(string(data))
						fmt.Println("AFTER KEY DEBUG")

					}

				case cr.clientCaPath:
					if err := cr.reloadWebhookCA(); err != nil {
						klog.ErrorS(err, "Failed to reload client CA")
					}
				default:
					continue
				}
				// watches get removed along with the symlinks, so we need to add them back
				if event.Has(fsnotify.Remove) {
					if err := watcher.Add(event.Name); err != nil {
						klog.ErrorS(err, "Failed to add watcher for file", "filename", event.Name)
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("In goroutine watcher")

				klog.Warningf("Error watching certificate files: %s", err)
			case <-stop:
				fmt.Println("Gorouting received shutdown")
				wg.Done()
				return
			}
		}
	}()
	return nil
}

func (cr *certReloader) load() error {
	cert, err := tls.LoadX509KeyPair(cr.tlsCertPath, cr.tlsKeyPath)
	if err != nil {
		return err
	}
	cr.mu.Lock()
	defer cr.mu.Unlock()
	cr.cert = &cert
	return nil
}

func (cr *certReloader) reloadWebhookCA() error {
	client := cr.mutatingWebhookClient
	webhook, err := client.Get(context.TODO(), webhookConfigName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if webhook == nil {
		return fmt.Errorf("webhook not found")
	}
	if len(webhook.Webhooks) == 0 {
		return fmt.Errorf("webhook configuration has no webhooks")
	}
	currentBundle := webhook.Webhooks[0].ClientConfig.CABundle[:]
	base64CurrentBundle := base64.StdEncoding.EncodeToString(currentBundle)
	newBundle := readFile(cr.clientCaPath)
	base64NewBundle := base64.StdEncoding.EncodeToString(newBundle)
	// make sure clientCA actually changed
	if base64CurrentBundle == base64NewBundle {
		klog.V(2).InfoS("Client CA did not change, skipping patch")
		return nil
	}
	klog.V(2).InfoS("New client CA found, reloading and patching webhook")
	patch := []byte(fmt.Sprintf(`{"webhooks":[{"name":"%s","clientConfig":{"caBundle":"%s"}}]}`, webhookName, base64NewBundle))
	_, err = client.Patch(context.TODO(), webhookConfigName, types.StrategicMergePatchType, patch, metav1.PatchOptions{})
	if err == nil {
		klog.V(2).InfoS("Successfully patched webhook with new client CA")
	}
	return err
}

func (cr *certReloader) getCertificate(_ *tls.ClientHelloInfo) (*tls.Certificate, error) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	return cr.cert, nil
}

/*

certs_test.go:175: I
certs_test.go:181: J
certs_test.go:175: I
certs_test.go:181: J
certs_test.go:175: I
certs_test.go:181: J
certs_test.go:175: I
In goroutine A
certs_test.go:181: J
certs_test.go:186: K
In goroutine event
In goroutine D CHMOD         "/var/folders/65/291bxggd7ql85hzwmn2td2500000gn/T/TestKeypairReloader1263719920/cert.key" /var/folders/65/291bxggd7ql85hzwmn2td2500000gn/T/TestKeypairReloader1263719920/cert.key
In goroutine A
In goroutine event
In goroutine C
In goroutine E
certs_test.go:196: M
certs_test.go:198: N
certs_test.go:96: Removing tempdir
In goroutine A
In goroutine event
--- PASS: TestKeypairReloader (2.22s)

*/
