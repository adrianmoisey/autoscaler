package model

import (
"testing"
"time"

"k8s.io/apimachinery/pkg/labels"
)

func TestGetControllingVPACache(t *testing.T) {
clusterState := NewClusterState(time.Hour)

// Create a test VPA
vpaID := VpaID{Namespace: "test-ns", VpaName: "test-vpa"}
selector := labels.SelectorFromSet(labels.Set{"app": "test"})
vpa := NewVpa(vpaID, selector, time.Now())
clusterState.vpas[vpaID] = vpa

// Create a test pod that matches the VPA
podID := PodID{Namespace: "test-ns", PodName: "test-pod"}
podLabels := labels.Set{"app": "test"}
clusterState.AddOrUpdatePod(podID, podLabels, "Running")
pod := clusterState.pods[podID]

// First call should be a cache miss
result1 := clusterState.GetControllingVPA(pod)
if result1 == nil {
t.Error("Expected to find controlling VPA")
}
if result1 != vpa {
t.Errorf("Expected VPA %v, got %v", vpa, result1)
}

// Verify it's in the cache
if cachedVpa, exists := clusterState.podToVpaCache[podID]; !exists {
t.Error("Expected pod to be in cache")
} else if cachedVpa != vpa {
t.Errorf("Expected cached VPA %v, got %v", vpa, cachedVpa)
}

// Second call should be a cache hit
result2 := clusterState.GetControllingVPA(pod)
if result2 != result1 {
t.Error("Expected same VPA from cache")
}

// Change pod labels - cache should be invalidated
newLabels := labels.Set{"app": "different"}
clusterState.AddOrUpdatePod(podID, newLabels, "Running")

// Cache should be cleared for this pod
if _, exists := clusterState.podToVpaCache[podID]; exists {
t.Error("Expected cache to be invalidated after pod label change")
}

// Next call should find no VPA (labels don't match)
pod = clusterState.pods[podID]
result3 := clusterState.GetControllingVPA(pod)
if result3 != nil {
t.Error("Expected no controlling VPA after label change")
}

// Nil results should NOT be cached
if _, exists := clusterState.podToVpaCache[podID]; exists {
t.Error("Expected nil results to not be cached")
}

// Delete pod - cache should be cleared
clusterState.AddOrUpdatePod(podID, podLabels, "Running") // Add back matching labels
clusterState.GetControllingVPA(clusterState.pods[podID]) // Cache it
clusterState.DeletePod(podID)

if _, exists := clusterState.podToVpaCache[podID]; exists {
t.Error("Expected cache to be cleared after pod deletion")
}
}

func TestGetControllingVPACacheInvalidationOnVPAChange(t *testing.T) {
clusterState := NewClusterState(time.Hour)

// Create test pods in two namespaces
pod1ID := PodID{Namespace: "ns1", PodName: "pod1"}
pod2ID := PodID{Namespace: "ns2", PodName: "pod2"}
podLabels := labels.Set{"app": "test"}

clusterState.AddOrUpdatePod(pod1ID, podLabels, "Running")
clusterState.AddOrUpdatePod(pod2ID, podLabels, "Running")

// Create VPA in ns1
vpa1ID := VpaID{Namespace: "ns1", VpaName: "vpa1"}
selector := labels.SelectorFromSet(labels.Set{"app": "test"})
vpa1 := NewVpa(vpa1ID, selector, time.Now())
clusterState.vpas[vpa1ID] = vpa1

// Cache both pods
result1 := clusterState.GetControllingVPA(clusterState.pods[pod1ID])
result2 := clusterState.GetControllingVPA(clusterState.pods[pod2ID])

if result1 != vpa1 {
t.Error("Expected pod1 to match vpa1")
}
if result2 != nil {
t.Error("Expected pod2 to not match any VPA")
}

// Verify pod1 is cached (nil not cached for pod2)
if _, exists := clusterState.podToVpaCache[pod1ID]; !exists {
t.Error("Expected pod1 to be cached")
}

// Delete VPA - should only invalidate cache for pods in same namespace
err := clusterState.DeleteVpa(vpa1ID)
if err != nil {
t.Errorf("Failed to delete VPA: %v", err)
}

// Cache for ns1 pods should be cleared
if _, exists := clusterState.podToVpaCache[pod1ID]; exists {
t.Error("Expected cache for pod1 (ns1) to be cleared after VPA deletion")
}

// Cache for ns2 pods should be unaffected (though it wasn't cached anyway since nil)
}
