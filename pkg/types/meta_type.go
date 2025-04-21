package types

import (
	"encoding/json"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type (
	MetaType struct {
		Kind       string `json:"kind,omitempty" yaml:"kind,omitempty"`
		APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	}
)

// ----------------------------------------------------------------------------------- //
// MetaType
// ----------------------------------------------------------------------------------- //

func (o *MetaType) Equals(other *MetaType) bool {
	matchKind := strings.EqualFold(o.Kind, other.Kind)
	matchVersion := strings.EqualFold(o.APIVersion, other.APIVersion)
	return matchKind && matchVersion
}

func (o *MetaType) String() string {
	return o.ToJSON(false)
}

func (o *MetaType) ToJSON(pretty bool) string {
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

func (o *MetaType) ToYAML() string {
	bs, err := yaml.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}
