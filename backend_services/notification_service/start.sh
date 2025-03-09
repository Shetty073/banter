#!/bin/bash

echo "✅ Loading environment variables for notification_service..."

# Load environment variables from .env
if [ -f ".env" ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo "❌ No .env file found in notification_service. Exiting..."
  exit 1
fi

echo "🚀 Starting notification_service..."
mix phx.server
