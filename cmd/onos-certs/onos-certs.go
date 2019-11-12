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
		SetName("ca-issuer-2").
		SetCluster(defaultCluster).
		SetResourceType("Issuer").
		Build()
	Issuer := setup.NewIssuer().SetResource(resource).Build()
	caIssuer := setup.NewCaIssuer().SetIssuer(Issuer).
		SetSecretName("ca-key-pair-2").
		Build()
	_ = caIssuer.Create()
	issuer, _ := caIssuer.GetCaIssuer()
	issuerList, _ := caIssuer.Issuer.GetAllIssuers()
	fmt.Println(issuerList.Items, "----", issuer.Namespace, issuer.Kind)

	resource2 := setup.NewResource().
		SetName("certificate-1").
		SetCluster(defaultCluster).
		Build()

	cert := setup.NewCertificate().SetResource(resource2).
		SetSecretName("ca-key-pair-2").
		SetCommonName("cert-1").
		SetDuration(time.Hour*2).
		SetRenewBefore(time.Hour*1).
		SetIssuerRef("ca-issuer-2", "", "Issuer").
		Build()

	err := cert.Create()
	fmt.Println(err)

}
