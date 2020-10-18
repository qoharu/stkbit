# Stockbit interview project
Author: Salman Abdulqohar

The code architecture you see in this project is my experiment on DDD approach with clean-architecture. you may want to see my writings in medium:
https://medium.com/@qoharu/clean-architecture-dan-pendekatan-ddd-domain-driven-design-pada-go-golang-d8f236b47096

## Before you start
- Please make sure you have docker installed on your system.

## How to start the service locally
1. Make sure you have docker installed on your machine.
2. run unit test with `make test`
3. run `make docker`
4. The application should be running at :55501

## Testing the APIs
- These services use OMDB API

### Search Movie
 

`GET localhost:55501/v1/search?searchword=[keyword]&pagination=[page]`

Request: none

**curl example**
```shell script
curl --location --request GET 'localhost:55501/v1/search?searchword=batman&pagination=1'
```

Response:success [200]
```json
{
  "data": {
    "movies": [
      {
        "title": "Batman Begins",
        "year": "2005",
        "imdbID": "tt0372784",
        "type": "movie",
        "posterURL": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
      }
    ],
    "totalData": 393
  },
  "error": "",
  "success": true
}
```
