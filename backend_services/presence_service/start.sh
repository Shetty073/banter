#!/bin/bash

echo "✅ Loading environment variables for presence_service..."

# Load environment variables from .env
if [ -f ".env" ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo "❌ No .env file found in presence_service. Exiting..."
  exit 1
fi

echo "🚀 Starting presence_service..."
mix phx.server
