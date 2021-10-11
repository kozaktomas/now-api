# now-api

Application with one important endpoint which returns current unix timestamp in text format. Port 8080 is hardcoded in
the app.

## Request example

```
GET http://localhost:8080/now
```

## Response example

```
1633979524
```

## Additional info

- there is an endpoint with Prometheus metrics - `/metrics`

## Docker

```
docker pull kozaktomas:now-api:latest
```