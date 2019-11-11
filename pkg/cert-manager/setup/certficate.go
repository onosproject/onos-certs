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
	"time"

	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type issuerRef struct {
	name  string
	kind  string
	group string
}

// Certificate certificate
type Certificate struct {
	commonName   string
	secretName   string
	dnsNames     []string
	ipAddresses  []string
	isCa         bool
	keyAlgorithm string
	keyEncoding  string
	keySize      int
	organization []string
	uriSANs      []string
	//usages       []string
	duration    time.Duration
	renewBefore time.Duration
	issuerRef   issuerRef
	*Resource
}

// CertificateBuilder is certificate builder interface
type CertificateBuilder interface {
	SetResource(*Resource) CertificateBuilder
	SetSecretName(string) CertificateBuilder
	SetCommonName(string) CertificateBuilder
	SetIssuerRef(string, string, string) CertificateBuilder
	Build() Certificate
}

// NewCertificate Creates an instance of certificate builder
func NewCertificate() CertificateBuilder {
	return &Certificate{
		Resource:    NewResource().SetResourceType(issuerResourceType).Build(),
		commonName:  "default-certificate",
		duration:    time.Duration(time.Hour * 10),
		renewBefore: time.Duration(time.Hour * 1),
	}
}

// Build build a certificate instance
func (ce *Certificate) Build() Certificate {
	return Certificate{
		Resource:     ce.Resource,
		commonName:   ce.commonName,
		secretName:   ce.secretName,
		dnsNames:     ce.dnsNames,
		ipAddresses:  ce.ipAddresses,
		isCa:         ce.isCa,
		keyAlgorithm: ce.keyAlgorithm,
		keyEncoding:  ce.keyEncoding,
		keySize:      ce.keySize,
		organization: ce.organization,
		uriSANs:      ce.uriSANs,
		duration:     ce.duration,
		renewBefore:  ce.renewBefore,
		issuerRef:    ce.issuerRef,
	}
}

// SetResource set resource name
func (ce *Certificate) SetResource(resource *Resource) CertificateBuilder {
	ce.Resource = resource
	return ce
}

// SetCommonName set common name
func (ce *Certificate) SetCommonName(commonName string) CertificateBuilder {
	ce.commonName = commonName
	return ce
}

// SetSecretName set secret name
func (ce *Certificate) SetSecretName(secretName string) CertificateBuilder {
	ce.secretName = secretName
	return ce
}

// SetIssuerRef set issuer ref
func (ce *Certificate) SetIssuerRef(name string, group string, kind string) CertificateBuilder {
	ref := issuerRef{
		name:  name,
		group: group,
		kind:  kind,
	}
	ce.issuerRef = ref
	return ce
}

// Create create a certificate
func (ce *Certificate) Create() error {
	certificate := cmapi.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ce.name,
			Namespace: ce.cluster.GetNameSpace(),
		},
		Spec: cmapi.CertificateSpec{
			CommonName:   ce.commonName,
			IPAddresses:  ce.ipAddresses,
			SecretName:   ce.secretName,
			KeySize:      ce.keySize,
			KeyAlgorithm: cmapi.KeyAlgorithm(ce.keyAlgorithm),
			KeyEncoding:  cmapi.KeyEncoding(ce.keyEncoding),
			Organization: ce.organization,
			IssuerRef: cmmeta.ObjectReference{
				Name:  ce.issuerRef.name,
				Group: ce.issuerRef.group,
				Kind:  ce.issuerRef.kind,
			},
			IsCA:     ce.isCa,
			DNSNames: ce.dnsNames,
			RenewBefore: &metav1.Duration{
				Duration: ce.renewBefore,
			},
			Duration: &metav1.Duration{
				Duration: ce.duration,
			},
			URISANs: ce.uriSANs,
			Usages:  cmapi.DefaultKeyUsages(),
		},
	}

	_, err := ce.cluster.GetCertManagerClient().
		CertmanagerV1alpha2().
		Certificates(ce.cluster.GetNameSpace()).
		Create(&certificate)

	return err

}
