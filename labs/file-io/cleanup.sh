#!/usr/bin/env bash
# This script just exists to reset the files in the lab without having to use
#  Git commands for that purpose

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")

echo "Resetting my-file.txt"
echo "Hello there, I am a file" > "$SCRIPT_DIR/my-file.txt"

echo "Resetting my-task.txt"
echo "This is meaningless text and should be truncated" \
  > "$SCRIPT_DIR/my-task.txt"

if [[ -f "$SCRIPT_DIR/this-file-does-not-exist.txt" ]]; then
  echo "Removing this-file-does-not-exist.txt"
  rm "$SCRIPT_DIR/this-file-does-not-exist.txt"
fi

echo "Cleanup end"
echo
