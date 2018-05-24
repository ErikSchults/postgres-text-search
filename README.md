# postgre full text search with golang

## Running project

* install glide on your machine: `go get github.com/Masterminds/glide`
* `glide install`
* `docker-compose up`

## Endpoints

### `GET /search`
Returns all books from database

### `GET /search?txt=searchstring`
Returns books whose combination of author, title and notes contain searchstring