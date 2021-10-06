#!/bin/bash -x

if [[ "$OSTYPE" == "linux"* ]]; then
    SRC=$HOME/go/bin
    DST=/usr/local/bin
    if [[ -f "$DST/go-auth-ss" ]]; then
        sudo systemctl stop GoAuthMs
        sudo $DST/go-auth-ss -service uninstall
        sleep 3
    fi
    go install
    (cd go-auth-ss; go install)
    sudo cp $SRC/go-auth-ms $DST
    sudo cp $SRC/go-auth-ss $DST
    sudo $DST/go-auth-ss -service install
    sudo systemctl restart GoAuthMs
fi
