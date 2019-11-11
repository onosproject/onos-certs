// Copyright 2019-present Open Networking Foundation.
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

package setup

import (
	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SelfSignedIssuer CA issuer
type SelfSignedIssuer struct {
	*Resource
}

// SelfSignedIssuerBuilder is issuer builder interface
type SelfSignedIssuerBuilder interface {
	SetResource(*Resource) SelfSignedIssuerBuilder
	Build() SelfSignedIssuer
}

// NewSelfSignedIssuer Creates an instance of self signed issuer builder
func NewSelfSignedIssuer() SelfSignedIssuerBuilder {
	return &SelfSignedIssuer{
		Resource: NewResource().SetResourceType(issuerResourceType).Build(),
	}
}

// Build build a Self Signed Issuer instance
func (ss *SelfSignedIssuer) Build() SelfSignedIssuer {
	return SelfSignedIssuer{
		Resource: ss.Resource,
	}
}

// SetResource set resource name
func (ss *SelfSignedIssuer) SetResource(resource *Resource) SelfSignedIssuerBuilder {
	ss.Resource = resource
	return ss
}

// Create creates a self signed issuer resource
func (ss *SelfSignedIssuer) Create() error {
	switch ss.resourceType {
	case issuerResourceType:
		issuer := cmapi.Issuer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      ss.name,
				Namespace: ss.cluster.GetNameSpace(),
			},
			Spec: cmapi.IssuerSpec{
				IssuerConfig: cmapi.IssuerConfig{
					SelfSigned: &cmapi.SelfSignedIssuer{},
				},
			},
		}
		_, err := ss.cluster.GetCertManagerClient().
			CertmanagerV1alpha2().
			Issuers(ss.cluster.GetNameSpace()).
			Create(&issuer)
		return err

	case clusterIssuerResourceType:
		clusterIssuer := cmapi.ClusterIssuer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      ss.name,
				Namespace: ss.cluster.GetNameSpace(),
			},
			Spec: cmapi.IssuerSpec{
				IssuerConfig: cmapi.IssuerConfig{
					SelfSigned: &cmapi.SelfSignedIssuer{},
				},
			},
		}
		_, err := ss.cluster.GetCertManagerClient().
			CertmanagerV1alpha2().
			ClusterIssuers().
			Create(&clusterIssuer)
		return err
	case nullResourceType:
	}

	return nil
}

// GetSelfSignedIssuer get an issuer based on a given and options
func (ss *SelfSignedIssuer) GetSelfSignedIssuer() (*cmapi.Issuer, error) {
	certManagerClient := ss.cluster.GetCertManagerClient()
	return certManagerClient.CertmanagerV1alpha2().
		Issuers(ss.cluster.GetNameSpace()).
		Get(ss.name, metav1.GetOptions{})
}
