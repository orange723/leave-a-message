#!/bin/bash

# Define the target URL
URL="http://localhost:3000/api/v1/message" # Use localhost instead of 0.0.0.0 for curl

# Number of messages to insert
NUM_MESSAGES=1000

echo "Starting to insert $NUM_MESSAGES random messages..."

for i in $(seq 1 $NUM_MESSAGES); do
  # Generate a random string for the message (20 characters alphanumeric)
  RANDOM_MESSAGE=$(openssl rand -base64 32 | tr '+/=' '___' | head -c 20)

  # Create the JSON payload
  PAYLOAD="{\"message\": \"$RANDOM_MESSAGE\"}"

  # Send the POST request using curl
  # We use -s for silent mode, -o /dev/null to discard output, and -w "%{http_code}" to print only the status code
  HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Content-Type: application/json" -d "$PAYLOAD" "$URL")

  if [ "$HTTP_STATUS" -ne 200 ]; then
    echo "Error inserting message $i (HTTP Status: $HTTP_STATUS). Payload: $PAYLOAD"
    # Optionally, break the loop on error
    # break
  else
    echo "Successfully inserted message $i (Status: $HTTP_STATUS)"
  fi
done

echo "Finished inserting $NUM_MESSAGES messages."
