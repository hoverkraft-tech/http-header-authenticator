#!/bin/ash

# shellcheck shell=dash

set -eu
set -o pipefail

HEADER="${HEADER:X-auth-header}"
VALUE="${VALUE:secret}"

/app/http-header-authenticator -H "${HEADER}" -v "${VALUE}"
