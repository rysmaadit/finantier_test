# Stock Service

This is a repository for stock service, to get daily status of current stocks that the result will be encrypted with AES-256

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
| ENVIRONMENT | Application environment | string | development/local |
| POLYGON_API_KEY | Polygon stock monitoring API key | string | alphanumeric |
| POLYGON_BASE_URL | Polygon stock monitoring base URL | string | url |

## TBA

- Add unit tests
