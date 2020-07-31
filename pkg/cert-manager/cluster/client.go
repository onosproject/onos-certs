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
	clientset "github.com/jetstack/cert-manager/pkg/client/clientset/versioned"
)

// client is the base for all Kubernetes cluster objects
type client struct {
	namespace         string
	certManagerClient clientset.Interface
}

// GetNameSpace get cluster namespace
func (c *client) GetNameSpace() string {
	return c.namespace
}

// GetCertManagerClient get cert manager client
func (c *client) GetCertManagerClient() clientset.Interface {
	return c.certManagerClient
}
