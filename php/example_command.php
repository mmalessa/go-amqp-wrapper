<?php

require_once ('GoAmqpWrapperConnector.php');

$connector = new GoAmqpWrapperConnector();

$body = $connector->getBody();
print_r($body);

exit(GoAmqpWrapperConnector::ACK);
