package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

type (
	Options struct {
		Debug       bool
		Config      bool
		Startup     bool
		Sync        bool
		Schedule    bool
		Event       bool
		BindingType string
		BindingName string
		Object      string
	}
)

const (
	DebugUsage       = "Print out debug information"
	ConfigUsage      = ""
	StartupUsage     = ""
	SyncUsage        = ""
	ScheduleUsage    = ""
	EventUsage       = ""
	BindingTypeUsage = ""
	BindingNameUsage = ""
	ObjectUsage      = ""
)

func main() {
	opts := new(Options)

	flag.BoolVar(&opts.Debug, "d", false, DebugUsage)
	flag.BoolVar(&opts.Debug, "debug", false, DebugUsage)
	flag.BoolVar(&opts.Config, "config", false, ConfigUsage)
	flag.BoolVar(&opts.Startup, "startup", false, StartupUsage)
	flag.BoolVar(&opts.Sync, "sync", false, SyncUsage)
	flag.BoolVar(&opts.Schedule, "schedule", false, ScheduleUsage)
	flag.BoolVar(&opts.Event, "event", false, EventUsage)

	flag.StringVar(&opts.BindingType, "type", "", BindingTypeUsage)
	flag.StringVar(&opts.BindingName, "binding", "", BindingNameUsage)
	flag.StringVar(&opts.Object, "object", "", ObjectUsage)

	flag.Parse()

	ctx := context.Background()

	hook, err := NewHook(opts)
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	err = hook.Execute(ctx)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
