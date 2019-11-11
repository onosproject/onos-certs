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

func main() {

	/*defaultCluster := cluster.New(kube.GetAPI("default"))
	resource := setup.NewResource().
		SetName("ca-issuer-2").
		SetCluster(defaultCluster).
		SetResourceType("issuer").
		Build()
	caIssuer := setup.NewCaIssuer().SetResource(resource).
		SetSecretName("ca-key-pair").
		Build()
	_ = caIssuer.Create()
	issuer, _ := caIssuer.GetCaIssuer()
	fmt.Println(issuer.Name, issuer.Namespace, issuer.Status)*/
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
