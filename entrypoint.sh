#!/bin/bash

echo "Starting start-location..."
/usr/local/bin/start-location &
START_LOCATION_PID=$!
wait $START_LOCATION_PID

echo "Starting start-meilisearch..."
/usr/local/bin/start-meilisearch &
START_MEILISEARCH_PID=$!
wait $START_MEILISEARCH_PID

echo "Starting start-bot..."
/usr/local/bin/start-bot &
START_BOT_PID=$!
wait $START_BOT_PID

echo "All applications have finished."
