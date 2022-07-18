// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package macieiface provides an interface to enable mocking the Amazon Macie service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package macieiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/macie"
)

// MacieAPI provides an interface to enable mocking the
// macie.Macie service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Macie.
//    func myFunc(svc macieiface.MacieAPI) bool {
//        // Make svc.AssociateMemberAccount request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := macie.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMacieClient struct {
//        macieiface.MacieAPI
//    }
//    func (m *mockMacieClient) AssociateMemberAccount(input *macie.AssociateMemberAccountInput) (*macie.AssociateMemberAccountOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMacieClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MacieAPI interface {
	AssociateMemberAccount(*macie.AssociateMemberAccountInput) (*macie.AssociateMemberAccountOutput, error)
	AssociateMemberAccountWithContext(aws.Context, *macie.AssociateMemberAccountInput, ...request.Option) (*macie.AssociateMemberAccountOutput, error)
	AssociateMemberAccountRequest(*macie.AssociateMemberAccountInput) (*request.Request, *macie.AssociateMemberAccountOutput)

	AssociateS3Resources(*macie.AssociateS3ResourcesInput) (*macie.AssociateS3ResourcesOutput, error)
	AssociateS3ResourcesWithContext(aws.Context, *macie.AssociateS3ResourcesInput, ...request.Option) (*macie.AssociateS3ResourcesOutput, error)
	AssociateS3ResourcesRequest(*macie.AssociateS3ResourcesInput) (*request.Request, *macie.AssociateS3ResourcesOutput)

	DisassociateMemberAccount(*macie.DisassociateMemberAccountInput) (*macie.DisassociateMemberAccountOutput, error)
	DisassociateMemberAccountWithContext(aws.Context, *macie.DisassociateMemberAccountInput, ...request.Option) (*macie.DisassociateMemberAccountOutput, error)
	DisassociateMemberAccountRequest(*macie.DisassociateMemberAccountInput) (*request.Request, *macie.DisassociateMemberAccountOutput)

	DisassociateS3Resources(*macie.DisassociateS3ResourcesInput) (*macie.DisassociateS3ResourcesOutput, error)
	DisassociateS3ResourcesWithContext(aws.Context, *macie.DisassociateS3ResourcesInput, ...request.Option) (*macie.DisassociateS3ResourcesOutput, error)
	DisassociateS3ResourcesRequest(*macie.DisassociateS3ResourcesInput) (*request.Request, *macie.DisassociateS3ResourcesOutput)

	ListMemberAccounts(*macie.ListMemberAccountsInput) (*macie.ListMemberAccountsOutput, error)
	ListMemberAccountsWithContext(aws.Context, *macie.ListMemberAccountsInput, ...request.Option) (*macie.ListMemberAccountsOutput, error)
	ListMemberAccountsRequest(*macie.ListMemberAccountsInput) (*request.Request, *macie.ListMemberAccountsOutput)

	ListMemberAccountsPages(*macie.ListMemberAccountsInput, func(*macie.ListMemberAccountsOutput, bool) bool) error
	ListMemberAccountsPagesWithContext(aws.Context, *macie.ListMemberAccountsInput, func(*macie.ListMemberAccountsOutput, bool) bool, ...request.Option) error

	ListS3Resources(*macie.ListS3ResourcesInput) (*macie.ListS3ResourcesOutput, error)
	ListS3ResourcesWithContext(aws.Context, *macie.ListS3ResourcesInput, ...request.Option) (*macie.ListS3ResourcesOutput, error)
	ListS3ResourcesRequest(*macie.ListS3ResourcesInput) (*request.Request, *macie.ListS3ResourcesOutput)

	ListS3ResourcesPages(*macie.ListS3ResourcesInput, func(*macie.ListS3ResourcesOutput, bool) bool) error
	ListS3ResourcesPagesWithContext(aws.Context, *macie.ListS3ResourcesInput, func(*macie.ListS3ResourcesOutput, bool) bool, ...request.Option) error

	UpdateS3Resources(*macie.UpdateS3ResourcesInput) (*macie.UpdateS3ResourcesOutput, error)
	UpdateS3ResourcesWithContext(aws.Context, *macie.UpdateS3ResourcesInput, ...request.Option) (*macie.UpdateS3ResourcesOutput, error)
	UpdateS3ResourcesRequest(*macie.UpdateS3ResourcesInput) (*request.Request, *macie.UpdateS3ResourcesOutput)
}

var _ MacieAPI = (*macie.Macie)(nil)