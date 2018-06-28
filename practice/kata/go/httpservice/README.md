# HTTP Service Kata

My own exercise, inspired by [Mat Ryer's article](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831).

Use TDD to write a web server package called `webapp`, 2-3 pomodoros :)

It should have the following end points:

## /notfound

* does not exist, returns 404

## /

* redirects to /hello, returns 301

## /hello

* Says 'hello'

## /hello/{name}

* Says 'hello {name}'

## /hello.json

* Responds with the following JSON

```json
{
    "response": "hello"
}
```

## /admin/abc123

* fake auth token `abc123` passed as a url param
* invokes middleware that validates the fake token
* responds with 'admin authorized`

## /admin/abc124

* as above but with incorrrect token `abc124`
* middleware responds with 401 Unauthorized code

