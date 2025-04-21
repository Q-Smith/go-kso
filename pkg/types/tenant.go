package types

import (
	"encoding/json"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type (
	Tenants []*Tenant
	Tenant  struct {
		MetaType   `json:",inline" yaml:",inline"`
		MetaObject `json:"metadata,omitempty" yaml:"metadata,omitempty"`
		Spec       *TenantSpec   `json:"spec" yaml:"spec"`
		Status     *TenantStatus `json:"status,omitempty" yaml:"status,omitempty"`
	}

	TenantSpec struct {
		Name string `json:"name" yaml:"name"`
	}

	TenantStatus struct {
		ObservedGeneration int `json:"observedGeneration,omitempty" yaml:"observedGeneration,omitempty"`
	}
)

// ----------------------------------------------------------------------------------- //
// Tenant
// ----------------------------------------------------------------------------------- //

func (o *Tenant) Equals(other *Tenant) bool {
	matchMetaType := o.MetaType.Equals(&other.MetaType)
	matchMetaObject := o.MetaObject.Equals(&other.MetaObject)
	matchSpec := o.Spec.Equals(other.Spec)

	return matchMetaType && matchMetaObject && matchSpec
}

func (o *Tenant) String() string {
	return o.ToJSON(false)
}

func (o *Tenant) ToJSON(pretty bool) string {
	var bs []byte
	var err error

	if pretty {
		bs, err = json.MarshalIndent(o, "", "  ")
	} else {
		bs, err = json.Marshal(o)
	}

	if err != nil {
		return err.Error()
	}
	return string(bs)
}

func (o *Tenant) ToYAML() string {
	bs, err := yaml.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}

// ----------------------------------------------------------------------------------- //
// Tenants
// ----------------------------------------------------------------------------------- //

func (o Tenants) Len() int {
	return len(o)
}

func (o Tenants) IsEmpty() bool {
	return len(o) == 0
}

func (o Tenants) First() *Tenant {
	if len(o) > 0 {
		return o[0]
	}
	return nil
}

func (o Tenants) Last() *Tenant {
	if len(o) > 0 {
		return o[len(o)-1]
	}
	return nil
}

func (o Tenants) Remove(idx int) Tenants {
	// re-slicing (does not maintain order of remaining elements)
	return append(o[:idx], o[idx+1:]...)
}

func (o Tenants) String() string {
	return o.ToJSON(false)
}

func (o Tenants) ToJSON(pretty bool) string {
	var bs []byte
	var err error

	if pretty {
		bs, err = json.MarshalIndent(o, "", "  ")
	} else {
		bs, err = json.Marshal(o)
	}

	if err != nil {
		return err.Error()
	}
	return string(bs)
}

func (o Tenants) ToYAML() string {
	bs, err := yaml.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}

// ----------------------------------------------------------------------------------- //
// TenantSpec
// ----------------------------------------------------------------------------------- //

func (o *TenantSpec) Equals(other *TenantSpec) bool {
	matchName := strings.EqualFold(o.Name, other.Name)

	return matchName
}

func (o *TenantSpec) String() string {
	return o.ToJSON(false)
}

func (o *TenantSpec) ToJSON(pretty bool) string {
	var bs []byte
	var err error

	if pretty {
		bs, err = json.MarshalIndent(o, "", "  ")
	} else {
		bs, err = json.Marshal(o)
	}

	if err != nil {
		return err.Error()
	}
	return string(bs)
}

func (o *TenantSpec) ToYAML() string {
	bs, err := yaml.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}

// ----------------------------------------------------------------------------------- //
// TenantStatus
// ----------------------------------------------------------------------------------- //

func (o *TenantStatus) Equals(other *TenantStatus) bool {
	matchObservedGeneration := o.ObservedGeneration == other.ObservedGeneration

	return matchObservedGeneration
}

func (o *TenantStatus) String() string {
	return o.ToJSON(false)
}

func (o *TenantStatus) ToJSON(pretty bool) string {
	var bs []byte
	var err error

	if pretty {
		bs, err = json.MarshalIndent(o, "", "  ")
	} else {
		bs, err = json.Marshal(o)
	}

	if err != nil {
		return err.Error()
	}
	return string(bs)
}

func (o *TenantStatus) ToYAML() string {
	bs, err := yaml.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}
