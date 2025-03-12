#!/bin/bash

# Determine the operating system
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Set the output binary file based on the OS
OUTPUT="build/main"
if [[ "$OS" == "mingw"* || "$OS" == "cygwin"* ]]; then
  OUTPUT="build\\main.exe"
fi

# Ensure the build directory exists
mkdir -p build

# Build the Go project using the vendor directory and optimization flags
echo "Building the project using vendor dependencies..."
GOFLAGS="-trimpath -ldflags '-s -w'" # Optimization: trim debug info and paths
CMD="go build -mod=vendor $GOFLAGS -o $OUTPUT ."

echo "Running: $CMD"
eval $CMD

# Check if the build succeeded
if [[ $? -eq 0 ]]; then
  echo "Build succeeded! Binary is located at: $OUTPUT"
else
  echo "Build failed! Please check the output for errors."
  exit 1
fi
