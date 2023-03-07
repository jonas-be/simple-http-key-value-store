# Simple Http key value store

This application is a key value store, controlled over http requests.

> **Note**
> This project was mainly for learning Go, especially the `net/http` and `net/http/httptest` libraries.


## Request

```http://localhost:8080/data?key=a&value=YourValue```

### GET
Returns the value of your given key

### PUT
Create or update the entry for your given key, with the given value

### DELETE
Delete the entry for your given key