# HTTP Server

- reference: https://github.com/nicewook/learn-go-with-tests-ko/blob/master/ko-KR/http-server.md

## v1. Initial test and pass

TEST

> Send `GET /players/{name}` and get "20"

TIL

```
Test just the `PlayerServer` function, not the whole server and handler operation
Success compile and Pass test whatever it cost
```

## v2. Real application 

Do

> Make working HTTP server application

TIL

```
Convert to type http.HandleFunc and using its ServeHttp() method, we can easily make simple function to http.Handler
```

TEST

> Different response for the each player






