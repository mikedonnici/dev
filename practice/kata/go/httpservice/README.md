# HTTP Service Kata

This is my own, inspired by [Mat Ryer's article](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831).

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

## /adminok

* fake auth token `abc123` passed as an argument to handler
* middleware that validates the fake token
* responds with 'admin authenticated`

## /adminfail

* as above but with incorrrect token `abc124`
* middleware that validates a the fake token is correct
* responds with 404






