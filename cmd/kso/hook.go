package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Q-Smith/go-kso/pkg/types"
)

type (
	Hook struct {
		opts     *Options
		bindCtx  string
		resource *types.Tenant
	}
)

func NewHook(opts *Options) (*Hook, error) {
	instance := new(Hook)
	instance.opts = opts

	ctxFile := os.Getenv("BINDING_CONTEXT_PATH")
	if _, err := os.Stat(ctxFile); err == nil {
		bs, err := os.ReadFile(ctxFile)
		if err != nil {
			return nil, err
		}
		instance.bindCtx = string(bs)
	}

	if len(opts.Object) != 0 {
		resource := new(types.Tenant)
		err := json.Unmarshal([]byte(opts.Object), &resource)
		if err != nil {
			return nil, err
		}
		instance.resource = resource
	}

	return instance, nil
}

func (h *Hook) Execute(ctx context.Context) error {
	if h.opts.Debug {
		h.onDebug(ctx)
	}

	if h.opts.Config {
		return h.onConfig(ctx)
	} else if h.opts.Startup {
		return h.onStartup(ctx)
	} else if h.opts.Sync {
		return h.onSync(ctx)
	} else if h.opts.Schedule {
		return h.onSchedule(ctx)
	} else if h.opts.Event {
		return h.onEvent(ctx)
	}

	return nil
}

func (h *Hook) onDebug(ctx context.Context) {
	bs := []byte(h.bindCtx)
	b64Ctx := base64.StdEncoding.EncodeToString(bs)
	fmt.Printf("%s\n", b64Ctx)
}

func (h *Hook) onConfig(ctx context.Context) error {
	// NOTE: Make sure no "\t" characters exist in YAML
	fmt.Printf(`
configVersion: v1
onStartup: 1
schedule:
- name: "refresh-auth-token"
  crontab: "*/30 * * * *"
  allowFailure: true
  includeSnapshotsFrom: ["monitor-tenant"]
kubernetes:
- name: "monitor-tenant"
  apiVersion: qsmith.com/v1alpha1
  kind: Tenant
  executeHookOnSynchronization: true
  executeHookOnEvent: ["Added", "Modified", "Deleted"]
`)

	return nil
}

func (h *Hook) onStartup(ctx context.Context) error {
	fmt.Printf("Start Up!\n")
	return nil
}

func (h *Hook) onSync(ctx context.Context) error {
	return h.onLogic(ctx)
}

func (h *Hook) onSchedule(ctx context.Context) error {
	return h.onLogic(ctx)
}

func (h *Hook) onEvent(ctx context.Context) error {
	return h.onLogic(ctx)
}

func (h *Hook) onLogic(ctx context.Context) error {
	fmt.Println(h.resource)
	return nil
}
