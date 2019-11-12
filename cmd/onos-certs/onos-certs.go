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

package main

import (
	"fmt"
	"time"

	"github.com/onosproject/onos-certs/pkg/cert-manager/cluster"
	"github.com/onosproject/onos-certs/pkg/cert-manager/kube"
	"github.com/onosproject/onos-certs/pkg/cert-manager/setup"
)

func main() {

	defaultCluster := cluster.New(kube.GetAPI("default"))
	resource := setup.NewResource().
		SetName("ca-issuer-1").
		SetCluster(defaultCluster).
		SetResourceType("Issuer").
		Build()
	caIssuer := setup.NewCaIssuer().SetResource(resource).
		SetSecretName("ca-key-pair").
		Build()
	_ = caIssuer.Create()
	issuer, _ := caIssuer.GetCaIssuer()
	fmt.Println(issuer.Name, issuer.Namespace, issuer.Status)

	resource2 := setup.NewResource().
		SetName("certificate-3").
		SetCluster(defaultCluster).
		Build()

	cert := setup.NewCertificate().SetResource(resource2).
		SetSecretName("ca-key-pair").
		SetCommonName("cert-3").
		SetDuration(time.Minute*62).
		SetRenewBefore(time.Hour*1).
		SetIssuerRef("ca-issuer-1", "", "Issuer").
		Build()

	err := cert.Create()
	fmt.Println(err)

	/*defaultCluster := cluster.New(kube.GetAPI("default"))
	resource := setup.NewResource().
		SetName("self-signed").
		SetCluster(defaultCluster).
		SetResourceType("issuer").
		Build()
	selfSignedIssuer := setup.NewSelfSignedIssuer().SetResource(resource).
		Build()
	_ = selfSignedIssuer.Create()
	issuer, _ := selfSignedIssuer.GetSelfSignedIssuer()
	fmt.Println(issuer.Name, issuer.Namespace, issuer.Status)*/

}
