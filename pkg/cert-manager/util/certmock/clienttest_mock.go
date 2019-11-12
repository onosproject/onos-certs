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

package certmock

import (
	certMgrClient "github.com/jetstack/cert-manager/pkg/client/clientset/versioned"
	clientset "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/fake"
)

// CertManagerMock cert manager mock data structure
type CertManagerMock struct {
	Client certMgrClient.Interface
}

// NewCertManagerMock create a cert manager client mock
func NewCertManagerMock() *CertManagerMock {
	return &CertManagerMock{
		Client: clientset.NewSimpleClientset(),
	}
}
