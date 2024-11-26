#!/usr/bin/env bash
set -euo pipefail

LOC=$(pwd)

rm -f $LOC/template.zip
zip -r $LOC/template.zip observer player/*.py server_* maps config.json games.json

cd runner
zip -r $LOC/template.zip runner_*
