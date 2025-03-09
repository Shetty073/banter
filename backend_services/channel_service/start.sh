#!/bin/bash

echo "✅ Loading environment variables for channel_service..."

# Load environment variables from .env
if [ -f ".env" ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo "❌ No .env file found in channel_service. Exiting..."
  exit 1
fi

echo "🚀 Starting channel_service..."
mix phx.server
