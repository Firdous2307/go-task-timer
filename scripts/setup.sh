#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "🚀 Setting up Go-Task-Timer development environment..."

# Check for required tools
check_tool() {
    if ! command -v $1 &> /dev/null; then
        echo -e "${RED}❌ $1 is not installed. Please install it first.${NC}"
        exit 1
    fi
}

check_tool "go"
check_tool "docker"


# Install Go dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

echo -e "${GREEN}✅ Setup complete! You can now run:${NC}"
echo "   go run cli/main.go"