#!/usr/bin/env php
<?php

// Read the metadata from fd3.
$metadata = file_get_contents("php://fd/3");
if (false === $metadata) {
  fwrite(STDERR, "failed to read metadata from fd3\n");
  exit(1);
}

// Decode the metadata.
$metadata = json_decode($metadata, true);
if (JSON_ERROR_NONE != json_last_error()) {
  fwrite(STDERR, "failed to decode metadata\n");
  fwrite(STDERR, json_last_error_msg() . PHP_EOL);
  exit(1);
}

// Read the body from STDIN.
$body = file_get_contents("php://stdin");
if (false === $body) {
  fwrite(STDERR, "failed to read body from STDIN\n");
  exit(1);
}


