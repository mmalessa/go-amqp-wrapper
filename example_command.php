<?php
echo "This is Example command\n";
// // Read the metadata from fd3.
// $metadata = file_get_contents("php://fd/3");
// if (false === $metadata) {
//   fwrite(STDERR, "failed to read metadata from fd3\n");
//   exit(1);
// }

// // Decode the metadata.
// $metadata = json_decode($metadata, true);
// if (JSON_ERROR_NONE != json_last_error()) {
//   fwrite(STDERR, "failed to decode metadata\n");
//   fwrite(STDERR, json_last_error_msg() . PHP_EOL);
//   exit(1);
// }

// Read the body from STDIN.
$message = file_get_contents("php://stdin");
if (false === $message) {
  fwrite(STDERR, "failed to read body from STDIN\n");
  exit(1);
}
echo $message . PHP_EOL . PHP_EOL;

$msg = json_decode($message, true);
print_r($msg);
echo PHP_EOL;

echo base64_decode($msg['Body']) . PHP_EOL;