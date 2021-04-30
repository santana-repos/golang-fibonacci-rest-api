# Golang Fibonacci REST API

A HTTP service with a single endpoint that calculates *n*-th fibonacci sequence (1, 1, 2, 3, 5, 8, 13 ...) number.

### Run the solution
```
go run main.go
```
It should show something like:
```
ApiServer started at http port 8000
```

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