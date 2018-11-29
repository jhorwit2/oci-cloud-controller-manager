// Copyright 2017 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestValidateConfig(t *testing.T) {
	testCases := []struct {
		name string
		in   *Config
		errs field.ErrorList
	}{
		{
			name: "valid configuration",
			in: &Config{
				Auth: AuthConfig{
					Region:        "us-phoenix-1",
					TenancyID:     "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					UserID:        "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					PrivateKey:    "-----BEGIN RSA PRIVATE KEY----- (etc)",
					Fingerprint:   "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1: "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2: "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
				},
			},
			errs: field.ErrorList{},
		},
		{
			name: "valid minimal configuration with instance principals auth",
			in: &Config{
				Auth: AuthConfig{
					UseInstancePrincipals: true,
				},
			},
			errs: field.ErrorList{},
		}, {
			name: "valid with non default security list management mode",
			in: &Config{
				Auth: AuthConfig{
					Region:        "us-phoenix-1",
					TenancyID:     "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					UserID:        "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					PrivateKey:    "-----BEGIN RSA PRIVATE KEY----- (etc)",
					Fingerprint:   "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1:                    "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2:                    "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
					SecurityListManagementMode: ManagementModeFrontend,
				},
			},
			errs: field.ErrorList{},
		}, {
			name: "missing region",
			in: &Config{
				Auth: AuthConfig{
					TenancyID:     "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					UserID:        "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					PrivateKey:    "-----BEGIN RSA PRIVATE KEY----- (etc)",
					Fingerprint:   "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1: "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2: "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
				},
			},
			errs: field.ErrorList{
				&field.Error{Type: field.ErrorTypeRequired, Field: "auth.region", BadValue: ""},
			},
		}, {
			name: "missing tenancy",
			in: &Config{
				Auth: AuthConfig{
					Region:        "us-phoenix-1",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					UserID:        "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					PrivateKey:    "-----BEGIN RSA PRIVATE KEY----- (etc)",
					Fingerprint:   "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1: "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2: "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
				},
			},
			errs: field.ErrorList{
				&field.Error{Type: field.ErrorTypeRequired, Field: "auth.tenancy", BadValue: ""},
			},
		}, {
			name: "missing compartment",
			in: &Config{
				Auth: AuthConfig{
					Region:      "us-phoenix-1",
					TenancyID:   "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					UserID:      "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					PrivateKey:  "-----BEGIN RSA PRIVATE KEY----- (etc)",
					Fingerprint: "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1: "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2: "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
				},
			},
			errs: field.ErrorList{},
		}, {
			name: "missing user",
			in: &Config{
				Auth: AuthConfig{
					Region:        "us-phoenix-1",
					TenancyID:     "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					PrivateKey:    "-----BEGIN RSA PRIVATE KEY----- (etc)",
					Fingerprint:   "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1: "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2: "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
				},
			},
			errs: field.ErrorList{
				&field.Error{Type: field.ErrorTypeRequired, Field: "auth.user", BadValue: ""},
			},
		}, {
			name: "missing key",
			in: &Config{
				Auth: AuthConfig{
					Region:        "us-phoenix-1",
					TenancyID:     "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					UserID:        "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					Fingerprint:   "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1: "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2: "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
				},
			},
			errs: field.ErrorList{
				&field.Error{Type: field.ErrorTypeRequired, Field: "auth.key", BadValue: ""},
			},
		}, {
			name: "missing figerprint",
			in: &Config{
				Auth: AuthConfig{
					Region:        "us-phoenix-1",
					TenancyID:     "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					UserID:        "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					PrivateKey:    "-----BEGIN RSA PRIVATE KEY----- (etc)",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1: "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2: "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
				},
			},
			errs: field.ErrorList{
				&field.Error{Type: field.ErrorTypeRequired, Field: "auth.fingerprint", BadValue: ""},
			},
		}, {
			name: "missing vcnid",
			in: &Config{
				Auth: AuthConfig{
					Region:        "us-phoenix-1",
					TenancyID:     "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					UserID:        "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					PrivateKey:    "-----BEGIN RSA PRIVATE KEY----- (etc)",
					Fingerprint:   "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{},
			},
			errs: field.ErrorList{
				&field.Error{Type: field.ErrorTypeRequired, Field: "vcn", BadValue: "", Detail: "VCNID configuration must be provided if configuration for subnet1 is not provided"},
			},
		}, {
			name: "invalid security list management mode",
			in: &Config{
				Auth: AuthConfig{
					Region:        "us-phoenix-1",
					TenancyID:     "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					CompartmentID: "ocid1.compartment.oc1..aaaaaaaa3um2atybwhder4qttfhgon4j3hcxgmsvnyvx4flfjyewkkwfzwnq",
					UserID:        "ocid1.user.oc1..aaaaaaaai77mql2xerv7cn6wu3nhxang3y4jk56vo5bn5l5lysl34avnui3q",
					PrivateKey:    "-----BEGIN RSA PRIVATE KEY----- (etc)",
					Fingerprint:   "8c:bf:17:7b:5f:e0:7d:13:75:11:d6:39:0d:e2:84:74",
				},
				LoadBalancer: &LoadBalancerConfig{
					Subnet1:                    "ocid1.tenancy.oc1..aaaaaaaatyn7scrtwtqedvgrxgr2xunzeo6uanvyhzxqblctwkrpisvke4kq",
					Subnet2:                    "ocid1.subnet.oc1.phx.aaaaaaaahuxrgvs65iwdz7ekwgg3l5gyah7ww5klkwjcso74u3e4i64hvtvq",
					SecurityListManagementMode: "invalid",
				},
			},
			errs: field.ErrorList{
				&field.Error{
					Type:     field.ErrorTypeInvalid,
					Field:    "loadBalancer.securityListManagementMode",
					BadValue: "invalid",
					Detail:   "invalid security list management mode",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.Complete()
			result := ValidateConfig(tt.in)
			if !reflect.DeepEqual(result, tt.errs) {
				t.Errorf("ValidateConfig(%#v)\n=>        %q \nExpected: %q", tt.in, result, tt.errs)
			}
		})
	}
}
