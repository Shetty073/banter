#!/bin/bash

# Install direnv if not present
if ! command -v direnv &> /dev/null; then
  echo "💿 Installing direnv..."
  brew install direnv
fi

# Enable direnv
for service in channel_service media_service messaging_service notification_service presence_service user_service
do
  echo "🚀 Creating $service database..."

  cd $service

  # Ensure direnv is working correctly
  if [ ! -f ".envrc" ]; then
    echo "source_env .env" > .envrc
    direnv allow
    mix deps.get
  fi

  # Load .env and run migrations
  set -a && source .env && set +a
  mix ecto.create || {
    echo "❌ Failed to create database for $service."
    cd ..
    continue
  }

  cd ..
done
