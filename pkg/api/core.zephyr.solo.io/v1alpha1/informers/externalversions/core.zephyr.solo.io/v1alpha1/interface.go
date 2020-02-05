/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// KubernetesClusters returns a KubernetesClusterInformer.
	KubernetesClusters() KubernetesClusterInformer
	// Meshes returns a MeshInformer.
	Meshes() MeshInformer
	// MeshGroups returns a MeshGroupInformer.
	MeshGroups() MeshGroupInformer
	// MeshServices returns a MeshServiceInformer.
	MeshServices() MeshServiceInformer
	// MeshWorkloads returns a MeshWorkloadInformer.
	MeshWorkloads() MeshWorkloadInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// KubernetesClusters returns a KubernetesClusterInformer.
func (v *version) KubernetesClusters() KubernetesClusterInformer {
	return &kubernetesClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Meshes returns a MeshInformer.
func (v *version) Meshes() MeshInformer {
	return &meshInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MeshGroups returns a MeshGroupInformer.
func (v *version) MeshGroups() MeshGroupInformer {
	return &meshGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MeshServices returns a MeshServiceInformer.
func (v *version) MeshServices() MeshServiceInformer {
	return &meshServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MeshWorkloads returns a MeshWorkloadInformer.
func (v *version) MeshWorkloads() MeshWorkloadInformer {
	return &meshWorkloadInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
