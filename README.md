# HTTP Server

- reference: https://github.com/nicewook/learn-go-with-tests-ko/blob/master/ko-KR/http-server.md
- iterations 를 기반으로 버전을 하나씩 만들어가며 주요 미션과 배운 점들을 정리해본다. 
- 참고. 실제 원문의 버전 별 구현과는 다르다. 

## v1. 가장 기본적 테스트

**구현 목표**

`GET /players/{name}` 를 보내면 "20" 을 받게 한다. 

**배운 것**

1. 전체 HTTP 서버나 핸들러를 구현하는 대신, `PlayerServer` 함수만을 테스트하자. 
2. 컴파일 성공, 테스트 통과를 최대한 빨리 통과하자. 그냥 원하는 값을 리턴해주게 하드코딩 하는거다. 
3. `http.Handler` 인터페이스는 `ServeHTTP(ResponseWriter, *Request)` 메서드를 구현하라는 것이다. 
4. `net/http/httptest` 패키지의 `ResponseRecorder`를 이용해서 response의 내용을 분석할 수 있다. 

## v2. 실제 동작 하는 애플리케이션/다른 플레이어에 대해 승점 리턴하기 

**구현 목표**

1. 동작하는 HTTP 서버와 핸들러를 만들자. 
2. v1의 `Pepper` 말고 `Floyd`의 승점도 받게 해보자 

**배운 것**

1. `ServeHTTP`
```
함수 시그니처가 `ServeHTTP` 메서드와 동일한 함수를 `http.HandlerFunc` 타입으로 컨버전하면 이 타입의 `ServeHTTP` 메서드를 사용할 수 있다. 
즉, `http.Handler` 인터페이스의 요구사항을 만족하게 된다. 다시 말해 `http.ListenAndServe` 함수의 인자로 들어갈 수 있다.
```
2. request `r.URL.Path` 에서 플레이어의 이름을 추출해내자. 이름으로 분기해서 리턴하게 하자. 
3. 리팩터링 하자. 함수 추출, 반복 줄이기 (DRY up)
4. t.Helper() 좀 더 알아보자. 

## v3. 



TIL

```
Convert to type http.HandleFunc and using its ServeHttp() method, we can easily make simple function to http.Handler
```

TEST

> Different response for the each player






