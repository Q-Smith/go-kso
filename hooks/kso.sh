#!/usr/bin/env bash

# Lifecycle Order
# 1 - OnStartup hooks.
# 2 - Synchronization hooks.
# 3 - Event hooks (kubernetes, schedule).

# Loops split values when it sees any whitespace like space, tab, or newline.
# So, you should use IFS (Internal Field Separator) to control how the split is handled.
IFS=''

hook::main() {
  if [ ! -e /usr/bin/kso ]; then
    echo "The 'kso' file not found!"
    exit 1
  fi

  if [[ $1 == "--config" ]] ; then
    hook::on_config
    exit 0
  else
    # Debug
    hook::on_debug

    if [[ -z "${BINDING_CONTEXT_PATH}" ]]; then
      echo "The 'BINDING_CONTEXT_PATH' env var not found!"
      exit 1
    fi

    # Binding, Type & Group
    ctx_type=$(jq -r '.[0].type' $BINDING_CONTEXT_PATH) # [Synchronization, Schedule, Event]
    ctx_binding=$(jq -r '.[0].binding' $BINDING_CONTEXT_PATH) # [onStartup, kubernetes, "<binding-name>"]

    # OnStartup
    # [{ "binding": "onStartup" }]
    if [[ $ctx_binding == "onStartup" ]] ; then
      hook::on_startup "Bootup", $ctx_binding
      exit 0
    fi

    # Synchronization
    # [{ "binding": "<binding-name>", "type":"Synchronization", "objects": [{}, ...], "snapshots": {"<binding-name>": [{}, ...]} }]
    if [[ $ctx_type == "Synchronization" ]] ; then
      hook::on_sync $ctx_type, $ctx_binding
      exit 0
    fi

    # Schedule
    # [{ "binding": "<binding-name>", "type":"Schedule", "snapshots": {"<binding-name>": [{}, ...]} }]
    if [[ $ctx_type == "Schedule" ]] ; then
      hook::on_schedule $ctx_type, $ctx_binding
      exit 0
    fi

    # Event
    # [{ "binding": "<binding-name>", "type":"Event", "watchEvent": "Added", "object": {...}, "snapshots": {"<binding-name>": [{}, ...]} }]
    if [[ $ctx_type == "Event" ]] ; then
      hook::on_event $ctx_type, $ctx_binding
      exit 0
    fi
  fi
}

hook::on_config() {
  /usr/bin/kso --config
}

hook::on_debug() {
  /usr/bin/kso --debug
}

hook::on_startup() {
  ctx_type=$1
  ctx_binding=$2

  /usr/bin/kso --startup --type=$ctx_type --binding=$ctx_binding
}

hook::on_sync() {
  ctx_type=$1
  ctx_binding=$2

  objects=()
  for object in $(jq -c -r '.[]' $BINDING_CONTEXT_PATH | jq -c -r '.objects[].object'); do
    objects+=($object)
  done

  for object in ${objects[@]}; do
    /usr/bin/kso --sync --type=$ctx_type --binding=$ctx_binding --object=$object
  done
}

hook::on_schedule() {
  ctx_type=$1
  ctx_binding=$2

  objects=()
  for object in $(jq -c -r '.[]' $BINDING_CONTEXT_PATH | jq -c -r '.snapshots."monitor-tenant"[].object'); do
    objects+=($object)
  done

  for object in ${objects[@]}; do
    /usr/bin/kso --schedule --type=$ctx_type --binding=$ctx_binding --object=$object
  done
}

hook::on_event() {
  ctx_type=$1
  ctx_binding=$2

  objects=()
  for object in $(jq -c -r '.[]' $BINDING_CONTEXT_PATH | jq -c -r '.object'); do
    objects+=($object)
  done

  for object in ${objects[@]}; do
    /usr/bin/kso --event --type=$ctx_type --binding=$ctx_binding --object=$object
  done
}

hook::main "$@"
