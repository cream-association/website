#!/bin/bash

if [ "$(whoami)" = "pixi" ]; then
    cd "$HOME/creamrobotics"
    git pull
    docker compose down && docker compose up --build -d
else
    echo "This script must be run as user 'pixi'."
fi
