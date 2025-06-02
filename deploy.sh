#!/bin/bash

set -e

if [ -z "$DROPLET" ]; then
    echo "DROPLET env var is not set. Aborting."
    exit 1
fi

echo "[*] Building sninja..."
go build -o sninja

echo "[*] Copying to remote server..."
scp sninja $DROPLET:~/.local/bin/sninja

echo "[✓] Deployed to $DROPLET"
