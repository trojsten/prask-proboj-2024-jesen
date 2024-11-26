#!/usr/bin/env bash
set -euo pipefail

go build -o server_linux ./server/
GOOS=windows go build -o server_windows.exe ./server/
GOOS=darwin go build -o server_mac ./server/
