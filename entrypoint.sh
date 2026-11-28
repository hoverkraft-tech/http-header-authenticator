#!/bin/sh
set -eu

if [ -z "${HEADER:-}" ]; then
  echo "HEADER is required" >&2
  exit 1
fi
if [ -z "${VALUE:-}" ]; then
  echo "VALUE is required" >&2
  exit 1
fi

exec /usr/local/bin/http-header-authenticator check -H "${HEADER}" -V "${VALUE}"
