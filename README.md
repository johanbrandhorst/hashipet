# HashiPet

The API for pets of HashiCorp
## Requirements

Go 1.16+

## Running

Running `main.go` starts a web server on https://0.0.0.0:11000/. You can configure
the port used with the `$PORT` environment variable, and to serve on HTTP set
`$SERVE_HTTP=true`.

```
$ go run main.go
```

An OpenAPI UI is served on https://0.0.0.0:11000/.
