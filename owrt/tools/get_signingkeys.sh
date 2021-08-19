#!/usr/bin/env bash

# docs page where keys are parsed from
KEYS="https://openwrt.org/docs/guide-user/security/signatures"

# print file header, comment and variable initialization
cat <<HEADER
// go generated file
package builder

// SigningKeys is a list of usign/signify public keys used to sign downloads
// on the OpenWRT mirrors. It was parsed and downloaded automatically with
// \`go generate\` from $KEYS
// on $(date --iso --utc) by $(id -un)@$(hostname).
var SigningKeysTimestamp = "$(date --iso=seconds --utc)"
var SigningKeys = map[string]string{
HEADER

# parse the wiki page for a list of keys, download and format as string array
curl -s "$KEYS" |\
  sed -n 's/.*"\(https:[^"]\+;a=blob_plain;f=usign\/[a-f0-9]\+\)".*/\1/p' |\
  while read key; do
    printf '  "%s": `%s`,\n' "$(basename "$key")" "$(curl -s "$key")"
  done;

# print closing bracket
printf '}\n'
