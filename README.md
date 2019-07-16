# Sentry Errors Prometheus Exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/nonamenix/sentry-telegraf-json)](https://goreportcard.com/report/github.com/nonamenix/sentry-telegraf-json)

```bash
./sentry-telegraf-json --help

Usage of ./sentry-telegraf-json:
      --organization string    Organization name in sentry (default "XXX")
      --query string           Sentry query for projects filtering
      --sentry-url string      The sentry url (default "https://sentry.io")
      --stats-period string    Sentry stats period (default "24h")
      --token string           Sentry API authorization token

```

## Build

Work with 1.12 and modules

```bash
go build
```

## Usage Example

Take sentry token from https://sentry.io/api

```bash
./sentry2prometheus --sentry-url=https://sentry.io \
    --organization=XXX \
    --query=team:web \
    --token=token_from_sentry
```
