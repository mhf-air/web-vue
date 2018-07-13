#!/usr/bin/env bash

# set -o xtrace
set -o errexit
set -o pipefail
readonly ARGS="$@"
readonly ARGS_COUNT="$#"
println() {
  printf "$@\n"
}
# ================================================================================

enable() {
  sudo ufw allow "$PORT"
}

disable() {
  sudo ufw delete allow "$PORT"
}

run_action() {
  case "$1" in
    e)
      enable
      println "port $PORT enabled"
      return 1
      ;;
  
    d)
      disable
      println "port $PORT disabled"
      return 1
      ;;
  
    *)
      println "Please enter e(enable) or d(disable)"
      ;;
  esac
  return 0
}

interactive() {
  println "enter e(enable) or d(disable) for port $PORT"
  while :
  do
    read action
    run_action "$action"
    local result="$?"
    if [[ "$result" > 0 ]]; then
      break
    fi
  done
}

main() {
  if [[ "$ARGS_COUNT" == 0 ]]; then
    interactive
  else
    run_action "$ACTION"
  fi
}

readonly ACTION="$1"
readonly PORT=9003

main
