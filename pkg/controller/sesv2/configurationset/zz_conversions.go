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

package configurationset

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/sesv2"

	svcapitypes "github.com/crossplane/provider-aws/apis/sesv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetConfigurationSetInput returns input for read
// operation.
func GenerateGetConfigurationSetInput(cr *svcapitypes.ConfigurationSet) *svcsdk.GetConfigurationSetInput {
	res := &svcsdk.GetConfigurationSetInput{}

	if cr.Spec.ForProvider.ConfigurationSetName != nil {
		res.SetConfigurationSetName(*cr.Spec.ForProvider.ConfigurationSetName)
	}

	return res
}

// GenerateConfigurationSet returns the current state in the form of *svcapitypes.ConfigurationSet.
func GenerateConfigurationSet(resp *svcsdk.GetConfigurationSetOutput) *svcapitypes.ConfigurationSet {
	cr := &svcapitypes.ConfigurationSet{}

	if resp.ConfigurationSetName != nil {
		cr.Spec.ForProvider.ConfigurationSetName = resp.ConfigurationSetName
	} else {
		cr.Spec.ForProvider.ConfigurationSetName = nil
	}
	if resp.DeliveryOptions != nil {
		f1 := &svcapitypes.DeliveryOptions{}
		if resp.DeliveryOptions.SendingPoolName != nil {
			f1.SendingPoolName = resp.DeliveryOptions.SendingPoolName
		}
		if resp.DeliveryOptions.TlsPolicy != nil {
			f1.TLSPolicy = resp.DeliveryOptions.TlsPolicy
		}
		cr.Spec.ForProvider.DeliveryOptions = f1
	} else {
		cr.Spec.ForProvider.DeliveryOptions = nil
	}
	if resp.ReputationOptions != nil {
		f2 := &svcapitypes.ReputationOptions{}
		if resp.ReputationOptions.ReputationMetricsEnabled != nil {
			f2.ReputationMetricsEnabled = resp.ReputationOptions.ReputationMetricsEnabled
		}
		cr.Spec.ForProvider.ReputationOptions = f2
	} else {
		cr.Spec.ForProvider.ReputationOptions = nil
	}
	if resp.SendingOptions != nil {
		f3 := &svcapitypes.SendingOptions{}
		if resp.SendingOptions.SendingEnabled != nil {
			f3.SendingEnabled = resp.SendingOptions.SendingEnabled
		}
		cr.Spec.ForProvider.SendingOptions = f3
	} else {
		cr.Spec.ForProvider.SendingOptions = nil
	}
	if resp.SuppressionOptions != nil {
		f4 := &svcapitypes.SuppressionOptions{}
		if resp.SuppressionOptions.SuppressedReasons != nil {
			f4f0 := []*string{}
			for _, f4f0iter := range resp.SuppressionOptions.SuppressedReasons {
				var f4f0elem string
				f4f0elem = *f4f0iter
				f4f0 = append(f4f0, &f4f0elem)
			}
			f4.SuppressedReasons = f4f0
		}
		cr.Spec.ForProvider.SuppressionOptions = f4
	} else {
		cr.Spec.ForProvider.SuppressionOptions = nil
	}
	if resp.Tags != nil {
		f5 := []*svcapitypes.Tag{}
		for _, f5iter := range resp.Tags {
			f5elem := &svcapitypes.Tag{}
			if f5iter.Key != nil {
				f5elem.Key = f5iter.Key
			}
			if f5iter.Value != nil {
				f5elem.Value = f5iter.Value
			}
			f5 = append(f5, f5elem)
		}
		cr.Spec.ForProvider.Tags = f5
	} else {
		cr.Spec.ForProvider.Tags = nil
	}
	if resp.TrackingOptions != nil {
		f6 := &svcapitypes.TrackingOptions{}
		if resp.TrackingOptions.CustomRedirectDomain != nil {
			f6.CustomRedirectDomain = resp.TrackingOptions.CustomRedirectDomain
		}
		cr.Spec.ForProvider.TrackingOptions = f6
	} else {
		cr.Spec.ForProvider.TrackingOptions = nil
	}

	return cr
}

// GenerateCreateConfigurationSetInput returns a create input.
func GenerateCreateConfigurationSetInput(cr *svcapitypes.ConfigurationSet) *svcsdk.CreateConfigurationSetInput {
	res := &svcsdk.CreateConfigurationSetInput{}

	if cr.Spec.ForProvider.ConfigurationSetName != nil {
		res.SetConfigurationSetName(*cr.Spec.ForProvider.ConfigurationSetName)
	}
	if cr.Spec.ForProvider.DeliveryOptions != nil {
		f1 := &svcsdk.DeliveryOptions{}
		if cr.Spec.ForProvider.DeliveryOptions.SendingPoolName != nil {
			f1.SetSendingPoolName(*cr.Spec.ForProvider.DeliveryOptions.SendingPoolName)
		}
		if cr.Spec.ForProvider.DeliveryOptions.TLSPolicy != nil {
			f1.SetTlsPolicy(*cr.Spec.ForProvider.DeliveryOptions.TLSPolicy)
		}
		res.SetDeliveryOptions(f1)
	}
	if cr.Spec.ForProvider.ReputationOptions != nil {
		f2 := &svcsdk.ReputationOptions{}
		if cr.Spec.ForProvider.ReputationOptions.ReputationMetricsEnabled != nil {
			f2.SetReputationMetricsEnabled(*cr.Spec.ForProvider.ReputationOptions.ReputationMetricsEnabled)
		}
		res.SetReputationOptions(f2)
	}
	if cr.Spec.ForProvider.SendingOptions != nil {
		f3 := &svcsdk.SendingOptions{}
		if cr.Spec.ForProvider.SendingOptions.SendingEnabled != nil {
			f3.SetSendingEnabled(*cr.Spec.ForProvider.SendingOptions.SendingEnabled)
		}
		res.SetSendingOptions(f3)
	}
	if cr.Spec.ForProvider.SuppressionOptions != nil {
		f4 := &svcsdk.SuppressionOptions{}
		if cr.Spec.ForProvider.SuppressionOptions.SuppressedReasons != nil {
			f4f0 := []*string{}
			for _, f4f0iter := range cr.Spec.ForProvider.SuppressionOptions.SuppressedReasons {
				var f4f0elem string
				f4f0elem = *f4f0iter
				f4f0 = append(f4f0, &f4f0elem)
			}
			f4.SetSuppressedReasons(f4f0)
		}
		res.SetSuppressionOptions(f4)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f5 := []*svcsdk.Tag{}
		for _, f5iter := range cr.Spec.ForProvider.Tags {
			f5elem := &svcsdk.Tag{}
			if f5iter.Key != nil {
				f5elem.SetKey(*f5iter.Key)
			}
			if f5iter.Value != nil {
				f5elem.SetValue(*f5iter.Value)
			}
			f5 = append(f5, f5elem)
		}
		res.SetTags(f5)
	}
	if cr.Spec.ForProvider.TrackingOptions != nil {
		f6 := &svcsdk.TrackingOptions{}
		if cr.Spec.ForProvider.TrackingOptions.CustomRedirectDomain != nil {
			f6.SetCustomRedirectDomain(*cr.Spec.ForProvider.TrackingOptions.CustomRedirectDomain)
		}
		res.SetTrackingOptions(f6)
	}

	return res
}

// GenerateDeleteConfigurationSetInput returns a deletion input.
func GenerateDeleteConfigurationSetInput(cr *svcapitypes.ConfigurationSet) *svcsdk.DeleteConfigurationSetInput {
	res := &svcsdk.DeleteConfigurationSetInput{}

	if cr.Spec.ForProvider.ConfigurationSetName != nil {
		res.SetConfigurationSetName(*cr.Spec.ForProvider.ConfigurationSetName)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}