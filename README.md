# binance-wallet
convenient endpoints for managing binance wallet

## get started
```bash
make
```

## development
```bash
make dev
```

## deployment
```bash
make docker-push
```

## endpoints
1. `/deposits`
    parameters:
    - `asset string`: currency initials 
    - `days int`: since how many days behind
```bash
# e.g:
curl --request GET \
  --url http://localhost:2999/deposits \
  --header 'Content-Type: application/json' \
  --data '{ "asset": "ADA", "days": 90 }'
```
