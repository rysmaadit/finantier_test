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
| ENCRYPTION_KEY | A key used for encryption | string | alphanumeric |

### Generate Encryption Key

Generate 256 -bit key and use Cipher Block Chaining (CBC)

```shell
openssl enc -aes-256-cbc -k secret -P -md sha1
```

You might get this following result
```text
salt=1D072B881875XXXX
key=B2363AC0318F46B00A167B0A715B74F4ABEFD3B651E98AC9E6F20533A5EEXXXX
iv =000625201358F0B061967760AA99XXXX
```
Set `key` as `ENCRYPTION_KEY` value

## TBA

- Add unit tests
