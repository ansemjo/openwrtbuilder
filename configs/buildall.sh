#!/usr/bin/env bash
# try to build all the owrtbuildconfs here
set -eu -o pipefail

# resolve path to builder script
openwrtbuilder="$(readlink -f ../openwrtbuilder)"

# iterate over all configs
for conf in **/owrtbuildconf; do (
  cd "$(dirname "$conf")"
  pwd
  "$openwrtbuilder"
); done
