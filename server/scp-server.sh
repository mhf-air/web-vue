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

main() {
  if [[ "$ARGS_COUNT" == 0 ]]; then
    println "please provide a remote"
    exit 1
  fi

  println "building server..."
  go build -o web server.go
  println "server done.\n"
  
  println "building deploy-server..."
  go build -o deploy deploy-server/a.go
  println "deploy-server done.\n"
  
  println "building cert-server..."
  go build -o cert cert-server/a.go
  println "cert server done.\n"
  
  println "scp..."
  scp \
    web deploy cert pre-deploy.sh run.sh \
    "${REMOTE}:~/web/bin"
  
  rm web
  rm deploy
  rm cert
  
  println "\ndone.\n"
}

readonly REMOTE="$1"

main
