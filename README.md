# returnr - Small webserver written in Go that returns requests

## How-to:

1. Build Docker image:
```
docker build . -t returnr
```

2. Run container:
```
docker run --rm -p 8080:8080 returnr:latest
```

3. Test with `curl` or similar tools:
```
curl http://localhost:8080/returnr

GET /returnr HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/7.58.0
```

```
curl -X POST -H 'Content-Type: application/json' http://localhost:8080/returnr

POST /returnr HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Type: application/json
User-Agent: curl/7.58.0
```

```
echo "Hello World" > data.txt
curl -X POST -H 'Content-Type: application/text' -d @data.txt http://localhost:8080/return

POST /return HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 11
Content-Type: application/text
User-Agent: curl/7.58.0

Hello World
```

