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

## Example requests

```shell
$ curl -X POST "http://0.0.0.0:8080/v1/pets" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"name\":\"Floyd\",\"owner\":\"Justin\",\"species\": \"SPECIES_PIG\",\"picture_url\":\"https://i.imgur.com/fbzdvB7.jpg\"}"
$ curl -X POST "http://0.0.0.0:8080/v1/pets" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"name\":\"Penny\",\"owner\":\"Caroline\",\"species\": \"SPECIES_CAT\",\"picture_url\":\"https://i.imgur.com/i08E7uB.jpg\"}"
$ curl -X POST "http://0.0.0.0:8080/v1/pets" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"name\":\"Bear\",\"owner\":\"Irena\",\"species\": \"SPECIES_DOG\",\"picture_url\":\"https://i.imgur.com/3koghMg.jpg\"}"
```
