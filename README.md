# go-amqp-wrapper
Go AMQP wrapper  
-- experiments only --

Based on https://github.com/streadway/amqp  
Inspired by https://github.com/corvus-ch/rabbitmq-cli-consumer

# Install
```shell script
go get -u github.com/mmalessa/go-amqp-wrapper
```

# Install
...to ~/go/bin directory
```shell script
go install github.com/mmalessa/go-amqp-wrapper
```

# Copy o Link to /bin directory
```shell script
sudo cp ~/go/bin/go-amqp-wrapper /bin/
```

# Run in terminal
```shell script
go-amqp-wrapper go-amqp-wrapper --config=config.yaml
```

# Run as service - systemd
//FIXME / TODO
#### Create config directory and files ie:
```shell script
sudo mkdir /etc/go-amqp-wrapper
sudo cp ~/go/src/github.com/mmalessa/go-amqp-wrapper/config.yaml /etc/go-amqp-wrapper/config1.yaml
```
...and modify '.yaml' file

#### Create systemd unit file
```shell script
sudo cp ~/go/src/github.com/mmalessa/go-amqp-wrapper/linux/go-amqp-wrapper.service /etc/systemd/system/go-amqp-wrapper-1.service
```
...and modify '.service' file

#### Start / stop service
```shell script
sudo systemctl start go-amqp-wrapper-1
sudo systemctl stop go-amqp-wrapper-1
```
