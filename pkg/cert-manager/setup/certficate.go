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

// IssuerRef IssuerRef certificate field
type IssuerRef struct {
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
	issuerRef   IssuerRef
	*Resource
}

// CertificateBuilder is certificate builder interface
type CertificateBuilder interface {
	SetResource(*Resource) CertificateBuilder
	SetSecretName(string) CertificateBuilder
	SetCommonName(string) CertificateBuilder
	SetDNSNames([]string) CertificateBuilder
	SetIPAddresses([]string) CertificateBuilder
	SetIsCa(bool) CertificateBuilder
	SetKeyAlgorithm(string) CertificateBuilder
	SetKeyEncoding(string) CertificateBuilder
	SetKeySize(int) CertificateBuilder
	SetOrganization([]string) CertificateBuilder
	SetURISans([]string) CertificateBuilder
	SetDuration(time.Duration) CertificateBuilder
	SetRenewBefore(time.Duration) CertificateBuilder
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

// GetCommonName get certificate common name
func (ce *Certificate) GetCommonName() string {
	return ce.commonName
}

// SetCommonName set common name
func (ce *Certificate) SetCommonName(commonName string) CertificateBuilder {
	ce.commonName = commonName
	return ce
}

// GetSecretName get secret name
func (ce *Certificate) GetSecretName() string {
	return ce.secretName
}

// SetSecretName set secret name
func (ce *Certificate) SetSecretName(secretName string) CertificateBuilder {
	ce.secretName = secretName
	return ce
}

// GetDNSNames get certificate dns names
func (ce *Certificate) GetDNSNames() []string {
	return ce.dnsNames
}

// SetDNSNames set dns names
func (ce *Certificate) SetDNSNames(dnsNames []string) CertificateBuilder {
	ce.dnsNames = dnsNames
	return ce
}

// GetIPAddresses get certificate IP addresses
func (ce *Certificate) GetIPAddresses() []string {
	return ce.ipAddresses
}

// SetIPAddresses set ip addresses
func (ce *Certificate) SetIPAddresses(ipAddresses []string) CertificateBuilder {
	ce.ipAddresses = ipAddresses
	return ce
}

// GetIsCa get IsCa certificate field
func (ce *Certificate) GetIsCa() bool {
	return ce.isCa
}

// SetIsCa set isCa field
func (ce *Certificate) SetIsCa(isCa bool) CertificateBuilder {
	ce.isCa = isCa
	return ce
}

// GetKeyAlgorithm get certificate  key algorithm
func (ce *Certificate) GetKeyAlgorithm() string {
	return ce.keyAlgorithm
}

// SetKeyAlgorithm set keyAlgorithm certificate field
func (ce *Certificate) SetKeyAlgorithm(keyAlgorithm string) CertificateBuilder {
	ce.keyAlgorithm = keyAlgorithm
	return ce
}

// GetKeyEncoding get keyEncoding certificate field
func (ce *Certificate) GetKeyEncoding() string {
	return ce.keyEncoding
}

// SetKeyEncoding set KeyEncoding field
func (ce *Certificate) SetKeyEncoding(keyEncoding string) CertificateBuilder {
	ce.keyEncoding = keyEncoding
	return ce
}

// GetKeySize get keySize certificate field
func (ce *Certificate) GetKeySize() int {
	return ce.keySize
}

// SetKeySize set KeySize field
func (ce *Certificate) SetKeySize(keySize int) CertificateBuilder {
	ce.keySize = keySize
	return ce
}

// GetOrganization get organization certificate field
func (ce *Certificate) GetOrganization() []string {
	return ce.organization
}

// SetOrganization set organization certificate field
func (ce *Certificate) SetOrganization(organization []string) CertificateBuilder {
	ce.organization = organization
	return ce
}

// GetURISans get UriSAN certificate field
func (ce *Certificate) GetURISans() []string {
	return ce.uriSANs
}

// SetURISans set UriSANs certificate field
func (ce *Certificate) SetURISans(uriSANs []string) CertificateBuilder {
	ce.uriSANs = uriSANs
	return ce
}

// GetDuration get duration certificate field
func (ce *Certificate) GetDuration() time.Duration {
	return ce.duration
}

// SetDuration set Duration field
func (ce *Certificate) SetDuration(duration time.Duration) CertificateBuilder {
	ce.duration = duration
	return ce
}

// GetRenewBefore get renewBefore certificate field
func (ce *Certificate) GetRenewBefore() time.Duration {
	return ce.renewBefore
}

// SetRenewBefore set renewBefore field
func (ce *Certificate) SetRenewBefore(renewBefore time.Duration) CertificateBuilder {
	ce.renewBefore = renewBefore
	return ce
}

// GetIssuerRef get issuerRef certificate field
func (ce *Certificate) GetIssuerRef() IssuerRef {
	return ce.issuerRef
}

// SetIssuerRef set issuer ref
func (ce *Certificate) SetIssuerRef(name string, group string, kind string) CertificateBuilder {
	ref := IssuerRef{
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
