#!/bin/bash

# TrainLoop SDK Examples Setup Script
echo "Setting up TrainLoop SDK Examples..."

# Copy .env.example to .env if .env doesn't exist
if [ ! -f .env ]; then
  echo "Creating .env file from template..."
  cp .env.example .env
  echo "Please update the .env file with your actual TrainLoop API key"
fi

# Install all dependencies with npm (TypeScript + dev dependencies)
echo "Installing npm dependencies..."
npm install

# Install Python dependencies
echo "Installing Python dependencies..."
# Create Python virtual environment if it doesn't exist
if [ ! -d "venv" ]; then
  echo "Creating Python virtual environment..."
  python3 -m venv venv
fi
# Activate virtual environment and install dependencies
echo "Activating virtual environment and installing dependencies..."
source venv/bin/activate
pip install -r python/requirements.txt
# Deactivate virtual environment
deactivate

# Install Go dependencies
echo "Setting up Go dependencies..."
cd go
go get github.com/TrainLoop/sdk@v0.0.7
go mod tidy
cd ..

echo "Setup complete! You can now run examples using the following commands:"
echo "  npm run run:ts    # Run TypeScript example"
echo "  npm run run:py    # Run Python example"
echo "  npm run run:go    # Run Go example"
