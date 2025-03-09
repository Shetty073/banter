#!/bin/bash

echo "✅ Loading environment variables for user_service..."

# Load environment variables from .env
if [ -f ".env" ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo "❌ No .env file found in user_service. Exiting..."
  exit 1
fi

echo "🚀 Starting user_service..."
mix phx.server
