/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type BehaviorOnMxFailure string

const (
	BehaviorOnMxFailure_USE_DEFAULT_VALUE BehaviorOnMxFailure = "USE_DEFAULT_VALUE"
	BehaviorOnMxFailure_REJECT_MESSAGE    BehaviorOnMxFailure = "REJECT_MESSAGE"
)

type BulkEmailStatus string

const (
	BulkEmailStatus_SUCCESS                          BulkEmailStatus = "SUCCESS"
	BulkEmailStatus_MESSAGE_REJECTED                 BulkEmailStatus = "MESSAGE_REJECTED"
	BulkEmailStatus_MAIL_FROM_DOMAIN_NOT_VERIFIED    BulkEmailStatus = "MAIL_FROM_DOMAIN_NOT_VERIFIED"
	BulkEmailStatus_CONFIGURATION_SET_NOT_FOUND      BulkEmailStatus = "CONFIGURATION_SET_NOT_FOUND"
	BulkEmailStatus_TEMPLATE_NOT_FOUND               BulkEmailStatus = "TEMPLATE_NOT_FOUND"
	BulkEmailStatus_ACCOUNT_SUSPENDED                BulkEmailStatus = "ACCOUNT_SUSPENDED"
	BulkEmailStatus_ACCOUNT_THROTTLED                BulkEmailStatus = "ACCOUNT_THROTTLED"
	BulkEmailStatus_ACCOUNT_DAILY_QUOTA_EXCEEDED     BulkEmailStatus = "ACCOUNT_DAILY_QUOTA_EXCEEDED"
	BulkEmailStatus_INVALID_SENDING_POOL_NAME        BulkEmailStatus = "INVALID_SENDING_POOL_NAME"
	BulkEmailStatus_ACCOUNT_SENDING_PAUSED           BulkEmailStatus = "ACCOUNT_SENDING_PAUSED"
	BulkEmailStatus_CONFIGURATION_SET_SENDING_PAUSED BulkEmailStatus = "CONFIGURATION_SET_SENDING_PAUSED"
	BulkEmailStatus_INVALID_PARAMETER                BulkEmailStatus = "INVALID_PARAMETER"
	BulkEmailStatus_TRANSIENT_FAILURE                BulkEmailStatus = "TRANSIENT_FAILURE"
	BulkEmailStatus_FAILED                           BulkEmailStatus = "FAILED"
)

type ContactLanguage string

const (
	ContactLanguage_EN ContactLanguage = "EN"
	ContactLanguage_JA ContactLanguage = "JA"
)

type ContactListImportAction string

const (
	ContactListImportAction_DELETE ContactListImportAction = "DELETE"
	ContactListImportAction_PUT    ContactListImportAction = "PUT"
)

type DataFormat string

const (
	DataFormat_CSV  DataFormat = "CSV"
	DataFormat_JSON DataFormat = "JSON"
)

type DeliverabilityDashboardAccountStatus string

const (
	DeliverabilityDashboardAccountStatus_ACTIVE             DeliverabilityDashboardAccountStatus = "ACTIVE"
	DeliverabilityDashboardAccountStatus_PENDING_EXPIRATION DeliverabilityDashboardAccountStatus = "PENDING_EXPIRATION"
	DeliverabilityDashboardAccountStatus_DISABLED           DeliverabilityDashboardAccountStatus = "DISABLED"
)

type DeliverabilityTestStatus string

const (
	DeliverabilityTestStatus_IN_PROGRESS DeliverabilityTestStatus = "IN_PROGRESS"
	DeliverabilityTestStatus_COMPLETED   DeliverabilityTestStatus = "COMPLETED"
)

type DimensionValueSource string

const (
	DimensionValueSource_MESSAGE_TAG  DimensionValueSource = "MESSAGE_TAG"
	DimensionValueSource_EMAIL_HEADER DimensionValueSource = "EMAIL_HEADER"
	DimensionValueSource_LINK_TAG     DimensionValueSource = "LINK_TAG"
)

type DkimSigningAttributesOrigin string

const (
	DkimSigningAttributesOrigin_AWS_SES  DkimSigningAttributesOrigin = "AWS_SES"
	DkimSigningAttributesOrigin_EXTERNAL DkimSigningAttributesOrigin = "EXTERNAL"
)

type DkimStatus string

const (
	DkimStatus_PENDING           DkimStatus = "PENDING"
	DkimStatus_SUCCESS           DkimStatus = "SUCCESS"
	DkimStatus_FAILED            DkimStatus = "FAILED"
	DkimStatus_TEMPORARY_FAILURE DkimStatus = "TEMPORARY_FAILURE"
	DkimStatus_NOT_STARTED       DkimStatus = "NOT_STARTED"
)

type EventType string

const (
	EventType_SEND              EventType = "SEND"
	EventType_REJECT            EventType = "REJECT"
	EventType_BOUNCE            EventType = "BOUNCE"
	EventType_COMPLAINT         EventType = "COMPLAINT"
	EventType_DELIVERY          EventType = "DELIVERY"
	EventType_OPEN              EventType = "OPEN"
	EventType_CLICK             EventType = "CLICK"
	EventType_RENDERING_FAILURE EventType = "RENDERING_FAILURE"
	EventType_DELIVERY_DELAY    EventType = "DELIVERY_DELAY"
	EventType_SUBSCRIPTION      EventType = "SUBSCRIPTION"
)

type IDentityType string

const (
	IDentityType_EMAIL_ADDRESS  IDentityType = "EMAIL_ADDRESS"
	IDentityType_DOMAIN         IDentityType = "DOMAIN"
	IDentityType_MANAGED_DOMAIN IDentityType = "MANAGED_DOMAIN"
)

type ImportDestinationType string

const (
	ImportDestinationType_SUPPRESSION_LIST ImportDestinationType = "SUPPRESSION_LIST"
	ImportDestinationType_CONTACT_LIST     ImportDestinationType = "CONTACT_LIST"
)

type JobStatus string

const (
	JobStatus_CREATED    JobStatus = "CREATED"
	JobStatus_PROCESSING JobStatus = "PROCESSING"
	JobStatus_COMPLETED  JobStatus = "COMPLETED"
	JobStatus_FAILED     JobStatus = "FAILED"
)

type MailFromDomainStatus string

const (
	MailFromDomainStatus_PENDING           MailFromDomainStatus = "PENDING"
	MailFromDomainStatus_SUCCESS           MailFromDomainStatus = "SUCCESS"
	MailFromDomainStatus_FAILED            MailFromDomainStatus = "FAILED"
	MailFromDomainStatus_TEMPORARY_FAILURE MailFromDomainStatus = "TEMPORARY_FAILURE"
)

type MailType string

const (
	MailType_MARKETING     MailType = "MARKETING"
	MailType_TRANSACTIONAL MailType = "TRANSACTIONAL"
)

type ReviewStatus string

const (
	ReviewStatus_PENDING ReviewStatus = "PENDING"
	ReviewStatus_FAILED  ReviewStatus = "FAILED"
	ReviewStatus_GRANTED ReviewStatus = "GRANTED"
	ReviewStatus_DENIED  ReviewStatus = "DENIED"
)

type SubscriptionStatus string

const (
	SubscriptionStatus_OPT_IN  SubscriptionStatus = "OPT_IN"
	SubscriptionStatus_OPT_OUT SubscriptionStatus = "OPT_OUT"
)

type SuppressionListImportAction string

const (
	SuppressionListImportAction_DELETE SuppressionListImportAction = "DELETE"
	SuppressionListImportAction_PUT    SuppressionListImportAction = "PUT"
)

type SuppressionListReason string

const (
	SuppressionListReason_BOUNCE    SuppressionListReason = "BOUNCE"
	SuppressionListReason_COMPLAINT SuppressionListReason = "COMPLAINT"
)

type TLSPolicy string

const (
	TLSPolicy_REQUIRE  TLSPolicy = "REQUIRE"
	TLSPolicy_OPTIONAL TLSPolicy = "OPTIONAL"
)

type WarmupStatus string

const (
	WarmupStatus_IN_PROGRESS WarmupStatus = "IN_PROGRESS"
	WarmupStatus_DONE        WarmupStatus = "DONE"
)
