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

func TestNewResource(t *testing.T) {
	resource := NewResource().Build()
	assert.Equal(t, "default", resource.GetName())
	assert.Equal(t, "cert-manager.io/v1alpha2", resource.GetAPIVersion())

}

func TestResource_SetName(t *testing.T) {
	resource := NewResource().SetName("test-resource").Build()
	assert.Equal(t, "test-resource", resource.GetName())

}

func TestResource_SetAPIVersion(t *testing.T) {
	resource := NewResource().SetAPIVersion("cert-manager.io/v1").Build()
	assert.Equal(t, "cert-manager.io/v1", resource.GetAPIVersion())
}

func TestResource_GetName(t *testing.T) {
	resource := NewResource().Build()
	assert.Equal(t, "default", resource.GetName())
}

func TestResource_GetAPIVersion(t *testing.T) {
	resource := NewResource().Build()
	assert.Equal(t, "cert-manager.io/v1alpha2", resource.GetAPIVersion())
}
