GOOS=linux go build
docker build -t jjustinlim/mybudgetgateway .
go clean
docker image push jjustinlim/mybudgetgateway