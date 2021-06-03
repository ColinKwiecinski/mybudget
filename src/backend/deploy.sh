docker volume prune

docker rm -f rServe2
docker run --network customNet2 -d --name rServe2 redis

export DB_NAME=mybudgetsqlserver
export MYSQL_ROOT_PASSWORD=brucelee1

docker pull jjustinlim/mybudgetmysql
docker rm -f mybudgetmysql
docker run -d --network customNet2 -p 3306:3306 --name mybudgetmysql -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD -e MYSQL_DATABASE=$DB_NAME jjustinlim/mybudgetmysql

export TLSCERT=/etc/letsencrypt/live/api.justinlim.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/api.justinlim.me/privkey.pem
export SESSIONKEY=$(openssl rand -base64 18)
export REDISADDR=rServe2:6379
export DSN=root:$MYSQL_ROOT_PASSWORD@tcp\(mybudgetmysql:3306\)/$DB_NAME?parseTime=true

docker rm -f mybudgetgateway
docker pull jjustinlim/mybudgetgateway

docker run --network customNet2 -d \
    --name mybudgetgateway \
    -p 443:443 \
    -v /etc/letsencrypt:/etc/letsencrypt:ro \
    -e TLSCERT=$TLSCERT \
    -e TLSKEY=$TLSKEY \
    -e SESSIONKEY=$SESSIONKEY \
    -e REDISADDR=$REDISADDR \
    -e DSN=$DSN \
    jjustinlim/mybudgetgateway