# HashiPet

The API for pets of HashiCorp
## Requirements

Go 1.18+

## Running

Running `main.go` starts a web server on http://0.0.0.0:8080/.

```
$ go run main.go
```

## Example requests

```shell
curl -X POST "http://0.0.0.0:8080/v1/pets" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"name\":\"Floyd\",\"owner\":\"Justin\",\"species\": \"SPECIES_PIG\",\"picture_url\":\"https://i.imgur.com/fbzdvB7.jpg\"}"
curl -X POST "http://0.0.0.0:8080/v1/pets" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"name\":\"Penny\",\"owner\":\"Caroline\",\"species\": \"SPECIES_CAT\",\"picture_url\":\"https://i.imgur.com/i08E7uB.jpg\"}"
curl -X POST "http://0.0.0.0:8080/v1/pets" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"name\":\"Bear\",\"owner\":\"Irena\",\"species\": \"SPECIES_DOG\",\"picture_url\":\"https://i.imgur.com/3koghMg.jpg\"}"
```
