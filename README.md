# MIJStter

## How to build

```sh
go get github.com/gin-gonic/gin
go get github.com/mattn/go-sqlite3
go get code.google.com/p/go-uuid/uuid
go build
```

## before go-sqlite install (for Windows)

```sh
1.install mingw
  URL: http://www.mingw.org/
2.Add the path to MinGW gcc to your Windows PATH environment variable
```

## API References

### GET /users

#### Parameter

```
/users?limit=10
```

Default is 20.

#### Results
```json
[
  {
    "id" : 1,
    "user_name" : "mijs_taro"
  }, ...
]
```

### POST /users

#### Parameters

```json
{
  "user_name" : "mijs_taro",
  "password" : "Xie9aek7",
  "password_confirm" : "Xie9aek7"
}
```

#### Results

ID

```json
1
```

### POST /login

#### Parameters

```json
{
  "user_name" : "mijs_taro",
  "password" : "Xie9aek7"
}
```

#### Results

```json
{
  "session_id" : "d8f1e91c-2764-4d8c-9fb0-0f1ced5dd77a"
}
```

### GET /posts

#### Parameter

```
/posts?limit=10
```

Default is 20.

#### Results

```json
[
  {
    "user_id" : 1,
    "user_name" : "mijs_taro",
    "message" : "Hello, Go !",
    "url" : "http://example.com/"
  }, ...
]
```

### POST /posts

#### Parameters

```json
{
  "session_id" : "d8f1e91c-2764-4d8c-9fb0-0f1ced5dd77a",
  "message" : "Hello, Go !",
  "url" : "http://example.com/"
}
```

#### Results

ID

```json
1
```

### POST /images

#### Parameters

Post `multipart/form-data` or post a image binary directly.

#### Results

Image url.
