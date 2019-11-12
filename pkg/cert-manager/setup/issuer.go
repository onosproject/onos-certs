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

// Issuer a generic issuer
type Issuer struct {
	*Resource
}

// IssuerBuilder is issuer builder interface
type IssuerBuilder interface {
	SetResource(*Resource) IssuerBuilder
	Build() *Issuer
}

// NewIssuer Creates an instance of ca issuer builder
func NewIssuer() IssuerBuilder {
	return &Issuer{
		Resource: NewResource().SetResourceType(issuerResourceType).Build(),
	}
}

// Build build an Issuer instance
func (is *Issuer) Build() *Issuer {
	return &Issuer{
		Resource: is.Resource,
	}
}

// SetResource set resource name
func (is *Issuer) SetResource(resource *Resource) IssuerBuilder {
	is.Resource = resource
	return is
}

// GetAllIssuers get list of all issuers
func (is *Issuer) GetAllIssuers() (*cmapi.IssuerList, error) {
	return is.cluster.GetCertManagerClient().
		CertmanagerV1alpha2().
		Issuers(is.cluster.GetNameSpace()).
		List(metav1.ListOptions{})
}
