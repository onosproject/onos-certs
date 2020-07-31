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

type resourceType string
type issuerType string

func (s resourceType) name() string {
	return string(s)
}

func (s issuerType) name() string {
	return string(s)
}

const (
	// issuerResourceType issuer resource type
	issuerResourceType resourceType = "Issuer"
	// clusterIssuerResourceType cluster issuer resource type
	clusterIssuerResourceType resourceType = "ClusterIssuer"
	// nullResourceType null resource type
	nullResourceType resourceType = "null"
)

const (
	// ca CA issuer
	ca issuerType = "CA"
)
