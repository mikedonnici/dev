# HTTP Service Kata

My own exercise, inspired by [Mat Ryer's article](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831).

Write a web server package called `webapp` with the following end points:

## /

* redirects to /hello (302)

## /hello

* Says 'hello'

## /hello/mike

* Says 'hello mike'

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


Include tests, try with TDD :)






