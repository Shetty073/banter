#!/bin/bash

echo "✅ Loading environment variables for media_service..."

# Load environment variables from .env
if [ -f ".env" ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo "❌ No .env file found in media_service. Exiting..."
  exit 1
fi

echo "🚀 Starting media_service..."
mix phx.server
