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

// CaIssuer CA issuer
type CaIssuer struct {
	secretName string
	*Issuer
}

// CaIssuerBuilder is issuer builder interface
type CaIssuerBuilder interface {
	SetSecretName(string) CaIssuerBuilder
	SetIssuer(*Issuer) CaIssuerBuilder
	Build() CaIssuer
}

// NewCaIssuer Creates an instance of ca issuer builder
func NewCaIssuer() CaIssuerBuilder {
	return &CaIssuer{
		Issuer:     NewIssuer().Build(),
		secretName: "default",
	}
}

// Build build a CA Issuer instance
func (ca *CaIssuer) Build() CaIssuer {
	return CaIssuer{
		Issuer:     ca.Issuer,
		secretName: ca.secretName,
	}
}

// GetSecretName get secret name
func (ca *CaIssuer) GetSecretName() string {
	return ca.secretName
}

// SetSecretName set secret name
func (ca *CaIssuer) SetSecretName(secretName string) CaIssuerBuilder {
	ca.secretName = secretName
	return ca

}

// SetIssuer set a generic issuer
func (ca *CaIssuer) SetIssuer(issuer *Issuer) CaIssuerBuilder {
	ca.Issuer = issuer
	return ca
}

func (ca *CaIssuer) createIssuer() cmapi.Issuer {
	return cmapi.Issuer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ca.name,
			Namespace: ca.cluster.GetNameSpace(),
		},
		Spec: cmapi.IssuerSpec{
			IssuerConfig: cmapi.IssuerConfig{
				CA: &cmapi.CAIssuer{
					SecretName: ca.secretName,
				},
			},
		},
	}
}

func (ca *CaIssuer) createClusterIssuer() cmapi.ClusterIssuer {
	return cmapi.ClusterIssuer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ca.name,
			Namespace: ca.cluster.GetNameSpace(),
		},
		Spec: cmapi.IssuerSpec{
			IssuerConfig: cmapi.IssuerConfig{
				CA: &cmapi.CAIssuer{
					SecretName: ca.secretName,
				},
			},
		},
	}
}

// Create creates a CA issuer resource
func (ca *CaIssuer) Create() error {
	switch ca.resourceType {
	case issuerResourceType:
		issuer := ca.createIssuer()
		_, err := ca.cluster.GetCertManagerClient().
			CertmanagerV1alpha2().
			Issuers(ca.cluster.GetNameSpace()).
			Create(&issuer)
		return err

	case clusterIssuerResourceType:
		clusterIssuer := ca.createClusterIssuer()
		_, err := ca.cluster.GetCertManagerClient().
			CertmanagerV1alpha2().
			ClusterIssuers().
			Create(&clusterIssuer)
		return err
	case nullResourceType:
	}

	return nil
}

// GetCaIssuer get an issuer based on a given name and options
func (ca *CaIssuer) GetCaIssuer() (*cmapi.Issuer, error) {
	certManagerClient := ca.cluster.GetCertManagerClient()
	return certManagerClient.CertmanagerV1alpha2().
		Issuers(ca.cluster.GetNameSpace()).
		Get(ca.name, metav1.GetOptions{})
}
