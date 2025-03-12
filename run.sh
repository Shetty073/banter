#!/bin/bash

# Determine the platform
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Set the default configuration file
CONFIG_FILE=".air.toml"

# Use the appropriate configuration file based on the platform
if [[ "$OS" == "darwin" || "$OS" == "linux" ]]; then
  CONFIG_FILE=".air.mac.toml"
elif [[ "$OS" == "mingw"* || "$OS" == "cygwin"* ]]; then
  CONFIG_FILE=".air.toml"
fi

# Check if the configuration file exists; fallback to default if it doesn't
if [[ ! -f "$CONFIG_FILE" ]]; then
  echo "Configuration file $CONFIG_FILE not found. Using default settings."
  CONFIG_FILE=".air.toml"
fi

# Run the air command with the selected configuration file
echo "Using configuration: $CONFIG_FILE"
"$(go env GOPATH)/bin/air" -c "$CONFIG_FILE"
