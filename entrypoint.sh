#!/bin/bash

CONFIG_PATH="/usr/local/kiwi-config/local.yml"

echo "Starting start-location..."
/usr/local/bin/start-location -config $CONFIG_PATH &
START_LOCATION_PID=$!
wait $START_LOCATION_PID

echo "Starting start-meilisearch..."
/usr/local/bin/start-meilisearch -config $CONFIG_PATH &
START_MEILISEARCH_PID=$!
wait $START_MEILISEARCH_PID

echo "Starting start-bot..."
/usr/local/bin/start-bot -config $CONFIG_PATH &
START_BOT_PID=$!
wait $START_BOT_PID

echo "All applications have finished."
