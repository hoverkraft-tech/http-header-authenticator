#!/bin/ash

# shellcheck shell=dash

set -eu
set -o pipefail

[ -z "${HEADER}" ] && echo "HEADER is required" && exit 1
[ -z "${VALUE}" ] && echo "VALUE is required" && exit 1

/app/http-header-authenticator -H "${HEADER}" -v "${VALUE}"
