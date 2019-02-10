# Varnostalert

Alerting microservice. Reads from a kafka topic, sends alerts to email/slack/pageduty.

## Building

```bash
make all
```

## Installation

Executing `make install` will create a systemd unit and move Varnost-Alert to /opt/varnost/alert/ .

## Usage
By default, varnostalert will try to pull configurations from the configuration service upon startup. You can also call it with the `--config "foo.json"` flag to skip the configuration service

## Configuration
Configuration is either stored locally or in Varnostifig. Either way, it will look like this:

`{
"kafka": {
  "BrokerList": ["127.0.0.1"],
  "Topic": "Alert",
  "Partition": 1,
  "OffsetType": -1,
  "MessageCountStart": 0
},
  "alert": {
    "FatalEmail": "khodges42@gmail.com",
    "FatalSlackChannel": "alerts"
  }
}`