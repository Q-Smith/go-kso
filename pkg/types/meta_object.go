package types

import (
	"encoding/json"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v3"
)

type (
	MetaObject struct {
		Name              string            `json:"name,omitempty" yaml:"name,omitempty"`
		GenerateName      string            `json:"generateName,omitempty" yaml:"generateName,omitempty"`
		ClusterName       string            `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`
		Namespace         string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
		UID               string            `json:"uid,omitempty" yaml:"uid,omitempty"`
		ResourceVersion   string            `json:"resourceVersion,omitempty" yaml:"resourceVersion,omitempty"`
		Generation        int               `json:"generation,omitempty" yaml:"generation,omitempty"`
		CreationTimestamp *time.Time        `json:"creationTimestamp,omitempty" yaml:"creationTimestamp,omitempty"`
		DeletionTimestamp *time.Time        `json:"deletionTimestamp,omitempty" yaml:"deletionTimestamp,omitempty"`
		Labels            map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
		Annotations       map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
		Finalizers        []string          `json:"finalizers,omitempty" yaml:"finalizers,omitempty"`
	}
)

// ----------------------------------------------------------------------------------- //
// MetaObject
// ----------------------------------------------------------------------------------- //

func (o *MetaObject) Equals(other *MetaObject) bool {
	matchName := strings.EqualFold(o.Name, other.Name)
	matchNamespace := strings.EqualFold(o.Namespace, other.Namespace)
	matchResourceVersion := strings.EqualFold(o.ResourceVersion, other.ResourceVersion)
	return matchName && matchNamespace && matchResourceVersion
}

func (o *MetaObject) String() string {
	return o.ToJSON(false)
}

func (o *MetaObject) ToJSON(pretty bool) string {
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

func (o *MetaObject) ToYAML() string {
	bs, err := yaml.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}
