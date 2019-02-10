# Varnostifig

Configuration microservice. Reads from a flat file upon startup and provides a simple API to provide configs.

## Building

```bash
make all
```

## Installation
Executing `make install` will create a systemd unit and move Varnost-Config to /opt/varnost/config/ .

## Usage
The config service will by default read from `/opt/varnost/config/configs/` upon startup.

To get an applicatons configuration, hit `/api/v1/config/{app_id}`