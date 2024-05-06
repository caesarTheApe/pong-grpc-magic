#!/bin/bash

# Set the script to exit immediately if any command exits with a non-zero status.
set -e

# Ensure correct number of arguments are provided.
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <Path to JSON file> <Contract name>"
    exit 1
fi

JSON_FILE=$1
CONTRACT_NAME=$2

# Define file names for the ABI and BIN
ABI_FILE="./$CONTRACT_NAME.abi"
BIN_FILE="./$CONTRACT_NAME.bin"

# Extract the ABI and BIN from the JSON file
jq -r '.abi' $JSON_FILE > $ABI_FILE
if [ $? -ne 0 ]; then
    echo "Failed to extract ABI"
    exit 1
fi

jq -r '.bytecode' $JSON_FILE > $BIN_FILE
if [ $? -ne 0 ]; then
    echo "Failed to extract bytecode"
    exit 1
fi

# Generate Go binding
abigen --abi=$ABI_FILE --bin=$BIN_FILE --pkg=$CONTRACT_NAME --out="./$CONTRACT_NAME.go"
if [ $? -eq 0 ]; then
    echo "Go binding generated successfully at ./$CONTRACT_NAME.go"
else
    echo "Failed to generate Go binding"
fi
