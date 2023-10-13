#!/bin/ash

# shellcheck shell=dash

set -eu
set -o pipefail

[ -z "${HEADER}" ] && echo "HEADER is required" && exit 1
[ -z "${VALUE}" ] && echo "VALUE is required" && exit 1

/usr/local/bin/http-header-authenticator check -H "${HEADER}" -V "${VALUE}"
