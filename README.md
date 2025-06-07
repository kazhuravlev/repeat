# repeat

A simple Go tool to repeat commands multiple times with optional delays.

## Installation

```shell
go install github.com/kazhuravlev/repeat/cmd/repeat@latest
```

## Usage

```shell
repeat <count> [delay] <command> [args...]
```

## Examples

### Basic repetition
```shell
# Ping a server 5 times
repeat 5 ping -c 1 google.com

# Check system time 10 times
repeat 10 date
```

### With delay between executions
```shell
# Health check every 2 seconds, 10 times
repeat 10 2s curl -s https://api.example.com/health

# Monitor disk usage every minute, 5 times
repeat 5 1m df -h
```

### Real-world use cases
```shell
# Stress test an API endpoint
repeat 100 curl -X POST https://api.example.com/test

# Watch for a service to come online
repeat 30 1s nc -zv localhost 8080

# Collect performance metrics
repeat 20 500ms ps aux | grep nginx
```

## Notes

- Commands are executed sequentially
- Exit codes are ignored - repeat continues even if commands fail
- Delay format supports Go duration syntax: `1s`, `500ms`, `2m30s`, etc.
