// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"sort"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewRbacConfig(namespace, name string) *RbacConfig {
	rbacconfig := &RbacConfig{}
	rbacconfig.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return rbacconfig
}

func (r *RbacConfig) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *RbacConfig) SetStatus(status core.Status) {
	r.Status = status
}

func (r *RbacConfig) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Mode,
		r.Inclusion,
		r.Exclusion,
		r.EnforcementMode,
	)
}

type RbacConfigList []*RbacConfig
type RbacconfigsByNamespace map[string]RbacConfigList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list RbacConfigList) Find(namespace, name string) (*RbacConfig, error) {
	for _, rbacConfig := range list {
		if rbacConfig.GetMetadata().Name == name {
			if namespace == "" || rbacConfig.GetMetadata().Namespace == namespace {
				return rbacConfig, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find rbacConfig %v.%v", namespace, name)
}

func (list RbacConfigList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, rbacConfig := range list {
		ress = append(ress, rbacConfig)
	}
	return ress
}

func (list RbacConfigList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, rbacConfig := range list {
		ress = append(ress, rbacConfig)
	}
	return ress
}

func (list RbacConfigList) Names() []string {
	var names []string
	for _, rbacConfig := range list {
		names = append(names, rbacConfig.GetMetadata().Name)
	}
	return names
}

func (list RbacConfigList) NamespacesDotNames() []string {
	var names []string
	for _, rbacConfig := range list {
		names = append(names, rbacConfig.GetMetadata().Namespace+"."+rbacConfig.GetMetadata().Name)
	}
	return names
}

func (list RbacConfigList) Sort() RbacConfigList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list RbacConfigList) Clone() RbacConfigList {
	var rbacConfigList RbacConfigList
	for _, rbacConfig := range list {
		rbacConfigList = append(rbacConfigList, resources.Clone(rbacConfig).(*RbacConfig))
	}
	return rbacConfigList
}

func (list RbacConfigList) Each(f func(element *RbacConfig)) {
	for _, rbacConfig := range list {
		f(rbacConfig)
	}
}

func (list RbacConfigList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *RbacConfig) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace RbacconfigsByNamespace) Add(rbacConfig ...*RbacConfig) {
	for _, item := range rbacConfig {
		byNamespace[item.GetMetadata().Namespace] = append(byNamespace[item.GetMetadata().Namespace], item)
	}
}

func (byNamespace RbacconfigsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace RbacconfigsByNamespace) List() RbacConfigList {
	var list RbacConfigList
	for _, rbacConfigList := range byNamespace {
		list = append(list, rbacConfigList...)
	}
	return list.Sort()
}

func (byNamespace RbacconfigsByNamespace) Clone() RbacconfigsByNamespace {
	cloned := make(RbacconfigsByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &RbacConfig{}

// Kubernetes Adapter for RbacConfig

func (o *RbacConfig) GetObjectKind() schema.ObjectKind {
	t := RbacConfigCrd.TypeMeta()
	return &t
}

func (o *RbacConfig) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*RbacConfig)
}

var RbacConfigCrd = crd.NewCrd("rbac.istio.io",
	"rbacconfigs",
	"rbac.istio.io",
	"v1alpha1",
	"RbacConfig",
	"rbacconfig",
	false,
	&RbacConfig{})
