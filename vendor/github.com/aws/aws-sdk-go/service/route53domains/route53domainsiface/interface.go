// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package route53domainsiface provides an interface for the Amazon Route 53 Domains.
package route53domainsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53domains"
)

// Route53DomainsAPI is the interface type for route53domains.Route53Domains.
type Route53DomainsAPI interface {
	CheckDomainAvailabilityRequest(*route53domains.CheckDomainAvailabilityInput) (*aws.Request, *route53domains.CheckDomainAvailabilityOutput)

	CheckDomainAvailability(*route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error)

	DeleteTagsForDomainRequest(*route53domains.DeleteTagsForDomainInput) (*aws.Request, *route53domains.DeleteTagsForDomainOutput)

	DeleteTagsForDomain(*route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error)

	DisableDomainAutoRenewRequest(*route53domains.DisableDomainAutoRenewInput) (*aws.Request, *route53domains.DisableDomainAutoRenewOutput)

	DisableDomainAutoRenew(*route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error)

	DisableDomainTransferLockRequest(*route53domains.DisableDomainTransferLockInput) (*aws.Request, *route53domains.DisableDomainTransferLockOutput)

	DisableDomainTransferLock(*route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error)

	EnableDomainAutoRenewRequest(*route53domains.EnableDomainAutoRenewInput) (*aws.Request, *route53domains.EnableDomainAutoRenewOutput)

	EnableDomainAutoRenew(*route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error)

	EnableDomainTransferLockRequest(*route53domains.EnableDomainTransferLockInput) (*aws.Request, *route53domains.EnableDomainTransferLockOutput)

	EnableDomainTransferLock(*route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error)

	GetDomainDetailRequest(*route53domains.GetDomainDetailInput) (*aws.Request, *route53domains.GetDomainDetailOutput)

	GetDomainDetail(*route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error)

	GetOperationDetailRequest(*route53domains.GetOperationDetailInput) (*aws.Request, *route53domains.GetOperationDetailOutput)

	GetOperationDetail(*route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error)

	ListDomainsRequest(*route53domains.ListDomainsInput) (*aws.Request, *route53domains.ListDomainsOutput)

	ListDomains(*route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error)

	ListDomainsPages(*route53domains.ListDomainsInput, func(*route53domains.ListDomainsOutput, bool) bool) error

	ListOperationsRequest(*route53domains.ListOperationsInput) (*aws.Request, *route53domains.ListOperationsOutput)

	ListOperations(*route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error)

	ListOperationsPages(*route53domains.ListOperationsInput, func(*route53domains.ListOperationsOutput, bool) bool) error

	ListTagsForDomainRequest(*route53domains.ListTagsForDomainInput) (*aws.Request, *route53domains.ListTagsForDomainOutput)

	ListTagsForDomain(*route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error)

	RegisterDomainRequest(*route53domains.RegisterDomainInput) (*aws.Request, *route53domains.RegisterDomainOutput)

	RegisterDomain(*route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error)

	RetrieveDomainAuthCodeRequest(*route53domains.RetrieveDomainAuthCodeInput) (*aws.Request, *route53domains.RetrieveDomainAuthCodeOutput)

	RetrieveDomainAuthCode(*route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error)

	TransferDomainRequest(*route53domains.TransferDomainInput) (*aws.Request, *route53domains.TransferDomainOutput)

	TransferDomain(*route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error)

	UpdateDomainContactRequest(*route53domains.UpdateDomainContactInput) (*aws.Request, *route53domains.UpdateDomainContactOutput)

	UpdateDomainContact(*route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error)

	UpdateDomainContactPrivacyRequest(*route53domains.UpdateDomainContactPrivacyInput) (*aws.Request, *route53domains.UpdateDomainContactPrivacyOutput)

	UpdateDomainContactPrivacy(*route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error)

	UpdateDomainNameserversRequest(*route53domains.UpdateDomainNameserversInput) (*aws.Request, *route53domains.UpdateDomainNameserversOutput)

	UpdateDomainNameservers(*route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error)

	UpdateTagsForDomainRequest(*route53domains.UpdateTagsForDomainInput) (*aws.Request, *route53domains.UpdateTagsForDomainOutput)

	UpdateTagsForDomain(*route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error)
}