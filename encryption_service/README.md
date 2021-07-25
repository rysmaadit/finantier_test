# Encryption Service

This is a repository for encryption service, to encrypt stock data to AES-256

## Prerequisite

To run this program, you will need

### App Dependencies

```$xslt
- Golang 1.12+
- Go mod enabled
```

## How to Run

### Setup App Config

```
cp .env.example .env
cp .env.example .env.testing
```

### Run Application

```
make run
```

## How to Test

```
make test
```

## How to Lint

```
make lint
```

## Deployment

### Build

```
make build
```

## Configuration

| NAME | DESCRIPTION | TYPE | VALUE
| ------ | ------ | ------ | ------ |
| APP_NAME | Application name | string | alphabet |
| APP_PORT | Application port | int | number |
| LOG_LEVEL | Mode for log level configuration | string | debug/info |
| ENVIRONMENT | Application environment | string | development |
| ENCRYPTION_KEY | A key used for encryption, must be 32 characters long | string | alphanumeric |

## TBA

- Add unit tests on handler
