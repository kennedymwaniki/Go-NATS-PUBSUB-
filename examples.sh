#!/bin/bash

# Example script to demonstrate the notification service

echo "========================================"
echo "Rental Management Notification Service"
echo "========================================"
echo ""

# Check if NATS is running
if ! docker ps | grep -q nats; then
    echo "⚠️  NATS server is not running. Starting it now..."
    docker-compose up -d
    sleep 3
    echo "✓ NATS server started"
else
    echo "✓ NATS server is already running"
fi

echo ""
echo "To use the notification service:"
echo ""
echo "1. Start the subscriber (in a new terminal):"
echo "   go run cmd/subscriber/main.go"
echo ""
echo "2. Start the publisher (in another terminal):"
echo "   go run cmd/publisher/main.go"
echo ""
echo "Or use Make commands:"
echo "   make run-subscriber"
echo "   make run-publisher"
echo ""
echo "Example notifications:"
echo "- Payment: Tenant T001 pays $1500 for Property P123"
echo "- Complaint: Tenant T002 reports noise issue in Property P124"
echo "- Repair: Tenant T003 requests plumbing repair in Property P125"
echo ""
echo "NATS Monitoring: http://localhost:8222"
echo ""
