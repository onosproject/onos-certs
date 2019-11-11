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

	"gotest.tools/assert"
)

func TestResource_Build(t *testing.T) {
	resource := NewResource().Build()
	assert.Equal(t, resource.name, "default")
	assert.Equal(t, resource.apiVersion, "cert-manager.io/v1alpha2")

}

func TestResource_SetName(t *testing.T) {
	resource := NewResource().SetName("test-set-resource-name").Build()
	assert.Equal(t, resource.name, "test-set-resource-name")
}

func TestResource_SetAPIVersion(t *testing.T) {
	resource := NewResource().SetAPIVersion("cert-manager.io/v1").Build()
	assert.Equal(t, resource.apiVersion, "cert-manager.io/v1")

}

func TestResource_SetResourceType(t *testing.T) {

}
