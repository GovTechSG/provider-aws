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

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// EmailIDentityParameters defines the desired state of EmailIDentity
type EmailIDentityParameters struct {
	// Region is which region the EmailIDentity will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The configuration set to use by default when sending from this identity.
	// Note that any configuration set defined in the email sending request takes
	// precedence.
	ConfigurationSetName *string `json:"configurationSetName,omitempty"`
	// If your request includes this object, Amazon SES configures the identity
	// to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, as
	// opposed to the default method, Easy DKIM (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
	//
	// You can only specify this object if the email identity is a domain, as opposed
	// to an address.
	DkimSigningAttributes *DkimSigningAttributes `json:"dkimSigningAttributes,omitempty"`
	// The email address or domain that you want to verify.
	// +kubebuilder:validation:Required
	EmailIDentity *string `json:"emailIDentity"`
	// An array of objects that define the tags (keys and values) that you want
	// to associate with the email identity.
	Tags                          []*Tag `json:"tags,omitempty"`
	CustomEmailIDentityParameters `json:",inline"`
}

// EmailIDentitySpec defines the desired state of EmailIDentity
type EmailIDentitySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EmailIDentityParameters `json:"forProvider"`
}

// EmailIDentityObservation defines the observed state of EmailIDentity
type EmailIDentityObservation struct {
	// An object that contains information about the DKIM attributes for the identity.
	DkimAttributes *DkimAttributes `json:"dkimAttributes,omitempty"`
	// The email identity type.
	IDentityType *string `json:"identityType,omitempty"`
	// Specifies whether or not the identity is verified. You can only send email
	// from verified email addresses or domains. For more information about verifying
	// identities, see the Amazon Pinpoint User Guide (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html).
	VerifiedForSendingStatus *bool `json:"verifiedForSendingStatus,omitempty"`
}

// EmailIDentityStatus defines the observed state of EmailIDentity.
type EmailIDentityStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EmailIDentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIDentity is the Schema for the EmailIDentities API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EmailIDentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailIDentitySpec   `json:"spec"`
	Status            EmailIDentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIDentityList contains a list of EmailIDentities
type EmailIDentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailIDentity `json:"items"`
}

// Repository type metadata.
var (
	EmailIDentityKind             = "EmailIDentity"
	EmailIDentityGroupKind        = schema.GroupKind{Group: Group, Kind: EmailIDentityKind}.String()
	EmailIDentityKindAPIVersion   = EmailIDentityKind + "." + GroupVersion.String()
	EmailIDentityGroupVersionKind = GroupVersion.WithKind(EmailIDentityKind)
)

func init() {
	SchemeBuilder.Register(&EmailIDentity{}, &EmailIDentityList{})
}
