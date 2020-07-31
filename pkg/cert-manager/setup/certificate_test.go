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
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestNewCertificate(t *testing.T) {
	cert := NewCertificate().Build()
	assert.Equal(t, "default-certificate", cert.GetCommonName())
}

func TestCertificate_GetCommonName(t *testing.T) {
	cert := NewCertificate().SetCommonName("test-cert").Build()
	assert.Equal(t, "test-cert", cert.GetCommonName())

}
func TestCertificate_GetDNSNames(t *testing.T) {
	dnsNames := make([]string, 2)
	dnsNames = append(dnsNames, "example.com")
	cert := NewCertificate().SetDNSNames(dnsNames).Build()
	assert.Equal(t, dnsNames[0], cert.GetDNSNames()[0])
}

func TestCertificate_GetIsCa(t *testing.T) {
	cert := NewCertificate().SetIsCa(true).Build()
	assert.Equal(t, true, cert.GetIsCa())
}

func TestCertificate_GetDuration(t *testing.T) {
	cert := NewCertificate().SetDuration(time.Second * 100).Build()
	assert.Equal(t, time.Second*100, cert.GetDuration())
}

func TestCertificate_GetRenewBefore(t *testing.T) {
	cert := NewCertificate().SetRenewBefore(time.Second * 100).Build()
	assert.Equal(t, time.Second*100, cert.GetRenewBefore())

}

func TestCertificate_GetKeyAlgorithm(t *testing.T) {
	cert := NewCertificate().SetKeyAlgorithm("rsa").Build()
	assert.Equal(t, "rsa", cert.GetKeyAlgorithm())
}

func TestCertificate_GetKeyEncoding(t *testing.T) {
	cert := NewCertificate().SetKeyEncoding("pkcs1").Build()
	assert.Equal(t, "pkcs1", cert.GetKeyEncoding())
}

func TestCertificate_GetKeySize(t *testing.T) {
	cert := NewCertificate().SetKeySize(256).Build()
	assert.Equal(t, 256, cert.GetKeySize())
}

func TestCertificate_GetOrganization(t *testing.T) {
	org := make([]string, 1)
	org = append(org, "ONF")
	cert := NewCertificate().SetOrganization(org).Build()
	assert.Equal(t, org[0], cert.GetOrganization()[0])
}

func TestCertificate_GetIPAddresses(t *testing.T) {
	ipAddresses := make([]string, 2)
	ipAddresses = append(ipAddresses, "10.0.0.1", "10.0.0.2")
	cert := NewCertificate().SetIPAddresses(ipAddresses).Build()
	assert.Equal(t, ipAddresses[0], cert.GetIPAddresses()[0])
	assert.Equal(t, ipAddresses[1], cert.GetIPAddresses()[1])
}

func TestCertificate_GetURISans(t *testing.T) {
	uriSANs := make([]string, 1)
	uriSANs = append(uriSANs, "uriName-test")
	cert := NewCertificate().SetURISans(uriSANs).Build()
	assert.Equal(t, uriSANs[0], cert.GetURISans()[0])
}

func TestCertificate_SetIssuerRef(t *testing.T) {
	issRef := IssuerRef{
		name:  "test-issuer",
		group: "test-group",
		kind:  "Issuer",
	}
	cert := NewCertificate().SetIssuerRef(issRef.name, issRef.group, issRef.kind).Build()
	assert.Equal(t, issRef.name, cert.GetIssuerRef().name)
	assert.Equal(t, issRef.group, cert.GetIssuerRef().group)
	assert.Equal(t, issRef.kind, cert.GetIssuerRef().kind)
}

func TestCertificate_GetSecretName(t *testing.T) {
	cert := NewCertificate().SetSecretName("test-secret").Build()
	assert.Equal(t, "test-secret", cert.GetSecretName())
}
