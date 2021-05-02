# Golang Fibonacci REST API

A HTTP service with a single endpoint that calculates *n*-th fibonacci sequence (1, 1, 2, 3, 5, 8, 13 ...) number.

## Run the solution
```
go run main.go
```
It should show something like:
```
ApiServer started at http port 8000
```

## Container
- Docker

## Build Steps

1. First we need to build the docker image for running the application

```docker build . -t golang-fibonacc-rest-api```

2. After the image has been built correctly we can proceed to run it as follows 

```docker run -p 8000:8000 golang-fibonacc-rest-api```

 ### Run the solution
The server will be running on [http://localhost:8000](http://localhost:8000) it should behave like this:
```sql
> curl http://localhost:8000/fib?n=1
1

> curl http://localhost:8000/fib?n=10
55

> curl http://localhost:8000/fib?n=72
498454011879264
```

## Tag image

```docker tag golang-fibonacc-rest-api thiagodevel/go-fib:0.0.1```

### Update Dockerhub

```docker push thiagodevel/go-fib:0.0.1```

## Container Desploy to Kubernetes

```` kubectl apply -f fib.yaml ```


### Testing the service On Kubernetes

``` kubectl get service ```

### Testing the Kubernetes Service on localserver (using ***minikube***)

``` minikube service go-fib-service --url ```
