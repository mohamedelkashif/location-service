# Flink hiring task - Backend project (GoLang)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


## General Info
1. The task is implemented using GO as programming language and [Gorilla/mux](https://github.com/gorilla/mux) as a webframework.
2. I used [vscode](https://code.visualstudio.com/) as IDE.



## Installation and setup
1. Unzip the compressed file to any where on your computer if you download or just clone the repository.
2. Navigate to the cloned/downloaded folder and open it using your favourite IDE.

## Running
```
cd $GOPATH/src/github.com/mohamedelkashif/store-location-service
go build
./store-location-service
```


## Usage & available end-points
There are three end-points available in this project, one for the main POST request and the other is for health check
```
{
    "store_id": "germy-templohof",
    "Name": "germany111",
    "Country": "germany",
    "country_code": "DE",
    "Location": {"lat": 11.34, "lng": 10.12},
    "slow_service": false
}
```

| Method        | Endpoint                                        | Body            |
| ------------- |:-----------------------------------------------:| --------------- |
| GET           | localhost:8080/api/v1/stores                    | no body         |
| POST          | localhost:8080/api/v1/stores/?country=DE&max=10 | Object above    |
| GET           | localhost:8080/api/v1/stores/?country=DE&max=10 | country, max    |

## TODO
1. Adding swagger documentation
2. Adding unit testing
3. Dockerizing


## Author
[Mohamed Mahmoud  ElKashif](mailto:uhammedmahmmoudd@gmail.com)
