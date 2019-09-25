<?php
$jsonMessage = file_get_contents("php://stdin");
if (false === $jsonMessage) {
    return;
}
$msg = json_decode($jsonMessage, true);
print_r($msg);
