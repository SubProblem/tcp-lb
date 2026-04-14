#!/bin/bash

REQUESTS=${1:-10}

for i in $(seq 1 $REQUESTS); do
    echo "request $i" | timeout 1 nc localhost 8080 > /dev/null 2>&1
    echo "sent request $i"
done

echo ""
echo "Stats:"
curl -s localhost:9091/stats