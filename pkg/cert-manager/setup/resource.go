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
	"github.com/onosproject/onos-certs/pkg/cert-manager/cluster"
	"github.com/onosproject/onos-certs/pkg/cert-manager/kube"
)

// ResourceBuilder Resource builder interface
type ResourceBuilder interface {
	SetName(string) ResourceBuilder
	SetAPIVersion(string) ResourceBuilder
	SetResourceType(resourceType) ResourceBuilder
	SetCluster(*cluster.Cluster) ResourceBuilder
	Build() *Resource
}

// NewResource Creates an instance of repository builder
func NewResource() ResourceBuilder {
	return &Resource{
		name:         "default",
		apiVersion:   "cert-manager.io/v1alpha2",
		resourceType: nullResourceType,
		cluster:      cluster.New(kube.GetAPI("default")),
	}
}

// Build build a Resource instance
func (resource *Resource) Build() *Resource {
	return &Resource{
		name:         resource.name,
		apiVersion:   resource.apiVersion,
		resourceType: resource.resourceType,
		cluster:      resource.cluster,
	}
}

// Resource is the base resource for cert-manager resources
type Resource struct {
	cluster      *cluster.Cluster
	name         string
	apiVersion   string
	resourceType resourceType
}

// SetName set resource name
func (resource *Resource) SetName(name string) ResourceBuilder {
	resource.name = name
	return resource
}

// SetAPIVersion set resource api version
func (resource *Resource) SetAPIVersion(apiVersion string) ResourceBuilder {
	resource.apiVersion = apiVersion
	return resource
}

// SetResourceType set resource type
func (resource *Resource) SetResourceType(resourceType resourceType) ResourceBuilder {
	resource.resourceType = resourceType
	return resource
}

// SetCluster set cluster
func (resource *Resource) SetCluster(cluster *cluster.Cluster) ResourceBuilder {
	resource.cluster = cluster
	return resource
}
