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

runNew() {
  rm -rf run
  mv ready run
}

runNewWithBackup() {
  rm -rf last
  mv run last
  mv ready run
}

runLast() {
  rm -rf ready
  mv run ready
  mv last run
}

run_action() {
  case "$1" in
    new)
      runNew
      return 1
      ;;
  
    "new-backup")
      runNewWithBackup
      return 1
      ;;
  
    "last")
      runLast
      return 1
      ;;
  
    *)
      println "Please enter new, new-backup or last"
      ;;
  esac
  return 0
}

interactive() {
  println "enter new, new-backup or last"
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
  cd ~/web/www/html
  if [[ "$ARGS_COUNT" == 0 ]]; then
    interactive
  else
    run_action "$ACTION"
  fi
}

readonly ACTION="$1"

main
