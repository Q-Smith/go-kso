package types

import (
	"encoding/json"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type (
	GVKNN struct {
		MetaType  `json:",inline" yaml:",inline"`
		Name      string `json:"name,omitempty" yaml:"name,omitempty"`
		Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
		UID       string `json:"uid,omitempty" yaml:"uid,omitempty"`
	}
)

// ----------------------------------------------------------------------------------- //
// GVKNN (Group, Version, Kind, Namespace, Name)
// ----------------------------------------------------------------------------------- //

func (o *GVKNN) Equals(other *GVKNN) bool {
	matchKind := strings.EqualFold(o.Kind, other.Kind)
	matchVersion := strings.EqualFold(o.APIVersion, other.APIVersion)
	matchName := strings.EqualFold(o.Name, other.Name)
	matchNamespace := strings.EqualFold(o.Namespace, other.Namespace)
	return matchKind && matchVersion && matchName && matchNamespace
}

func (o *GVKNN) String() string {
	return o.ToJSON(false)
}

func (o *GVKNN) ToJSON(pretty bool) string {
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

func (o *GVKNN) ToYAML() string {
	bs, err := yaml.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}
