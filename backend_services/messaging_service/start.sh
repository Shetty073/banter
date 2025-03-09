#!/bin/bash

echo "✅ Loading environment variables for messaging_service..."

# Load environment variables from .env
if [ -f ".env" ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo "❌ No .env file found in messaging_service. Exiting..."
  exit 1
fi

echo "🚀 Starting messaging_service..."
mix phx.server
