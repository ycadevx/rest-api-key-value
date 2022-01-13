# Golang REST API

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/net/http)


<h2> In-memory, key-value store golang-rest-api <h3>

>It allows you to save requests made in API to a file in 'N' seconds..
## Clone the project

```
$ git clone https://github.com/ycak1r/rest-api-key-value
```
# Use-Case
>  /get-all :   GET method. Returns the value of temp/data.json file result </br>
>  /get/{key} : GET method. Returns the value of a key that is also in the file.</br>
>  /set :       POST method. Allows you to add the  key/value pairs in the data.json file</br>
>  /flush :     DELETE method. Delete  all keys/value data in file.


# How to run Standart
``` sh
    go run main.go  OR go run ./
 ```
# How to run Docker Container
``` sh

   docker build -t golang-restapi .

   docker run -p 8081:8085 golang-rest-api
 ```

# How to create Docker Image and Container

> Create Dockerfile <br>
>Follow the steps to create the image and the container.

``` sh

 docker images

 docker build -t golang-restapi .
 
 docker run -p 8081:8085 golang-rest-api
 
 docker run -d --publish 8081:8085 golang-rest-api

 ```


## Json Data Example :

{"81c84f95-6d83-4a5a-81f4-12dde63563e4":"FEF2020000000001", <br>"9246706f-90ea-4a6d-8ac2-793600aec311":"CMD-200001",<br>"9fb8cc65-1070-473e-92a4-d1b279debb32":"CMD-20220111YCA",<br>"d6fc80b1-382c-4c8a-870e-557dc5786341":"CMD-200005",<br>"ef1bcf33-ac0d-4194-b426-b2975cb4bd65":"CMD-202000000050"}

