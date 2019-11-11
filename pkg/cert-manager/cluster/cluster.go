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

package cluster

import (
	"github.com/onosproject/onos-certs/pkg/cert-manager/kube"
	"github.com/onosproject/onos-certs/pkg/cert-manager/util/k8s"
)

// Cluster k8s cluster struct
type Cluster struct {
	*client
}

// New returns a new cluster Env
func New(kube kube.API) *Cluster {
	certManagerClient, _ := k8s.GetCertManagerClientSet()
	client := &client{
		namespace:         kube.Namespace(),
		certManagerClient: certManagerClient,
	}
	return &Cluster{
		client: client,
	}
}
