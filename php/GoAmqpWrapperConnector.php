<?php

declare(strict_types=1);

/*
 * GoAmqpWrapperConnector
 * https://github.com/mmalessa/go-amqp-wrapper
 */

class GoAmqpWrapperConnector
{
    const ACK = 0;
    const REJECT = 1;
    const REJECT_REQUEUE =2;

    private $msg;

    public function __construct()
    {
        $jsonMessage = file_get_contents("php://stdin");
        if (false === $jsonMessage) {
            return;
        }
        $this->msg = json_decode($jsonMessage, true);
    }

    public function getMessage(): array
    {
        return $this->msg;
    }

    public function getAcknowledger(): array
    {
        return $this->msg['Acknowledger'];

    }

    public function getHeaders(): array
    {
        return $this->msg['Headers'];

    }

    public function getContentType(): string
    {
        return $this->msg['ContentType'];
    }

    public function getContentEncoding(): string
    {
        $this->msg['ContentEncoding'];
    }

    public function getDeliveryMode(): int
    {
        $this->msg['DeliveryMode'];
    }

    public function getPriority(): int
    {
        return $this->msg['Priority'];
    }

    public function getCorrelationId()
    {
        return $this->msg['CorrelationId'];
    }

    public function getReplyTo()
    {
        return $this->msg['ReplyTo'];
    }

    public function getExpiration()
    {
        return $this->msg['Expiration'];
    }

    public function getMessageId()
    {
        return $this->msg['MessageId'];
    }

    public function getTimestamp()
    {
        return $this->msg['Timestamp'];
    }

    public function getType()
    {
        return $this->msg['Type'];
    }

    public function getUserId()
    {
        $this->msg['UserId'];
    }

    public function getAppId()
    {
        return $this->msg['AppId'];
    }

    public function getConsumerTag()
    {
        return $this->msg['ConsumerTag'];
    }

    public function getMessageCount()
    {
        return $this->msg['MessageCount'];
    }

    public function getDeliveryTag()
    {
        return $this->msg['DeliveryTag'];
    }

    public function getRedelivered()
    {
        return $this->msg['Redelivered'];
    }

    public function getExchange()
    {
        return $this->msg['Exchange'];
    }

    public function getRoutingKey()
    {
        return $this->msg['RoutingKey'];
    }

    public function getBody()
    {
        return base64_decode($this->msg['Body']);
    }

}
