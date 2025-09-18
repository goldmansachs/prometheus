#!/bin/bash

# Script to gzip all necessary files for embedding
# Based on the embed.go file requirements

set -e

echo "Gzipping mantine-ui assets..."

# Gzip mantine-ui files
gzip -f -9 -k static/mantine-ui/index.html
gzip -f -9 -k static/mantine-ui/favicon.svg
gzip -f -9 -k static/mantine-ui/assets/codicon-B_nZgZYP.ttf
gzip -f -9 -k static/mantine-ui/assets/index-BGNavk_Y.js
gzip -f -9 -k static/mantine-ui/assets/index-BcytBn6b.css

echo "Gzipping other required assets..."

# Gzip other required files (only if they exist and aren't already gzipped)
find static/css static/js static/react static/vendor -type f \( -name "*.css" -o -name "*.js" -o -name "*.json" -o -name "*.html" -o -name "*.ico" -o -name "*.svg" -o -name "*.ttf" -o -name "*.cjs" -o -name "*.txt" -o -name "*.map" -o -name "*.eot" -o -name "*.woff" -o -name "*.woff2" -o -name "*.less" \) ! -name "*.gz" -exec gzip -f -9 -k {} \; 2>/dev/null || true

echo "Done! All required files have been gzipped."
echo "Files ready for embedding."