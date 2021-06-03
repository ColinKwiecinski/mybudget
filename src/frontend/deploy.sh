docker rm -f mybudgetclient
docker pull jjustinlim/mybudgetclient
docker run -d --name mybudgetclient -p 80:80 -p 443:443 -v /etc/letsencrypt:/etc/letsencrypt:ro jjustinlim/mybudgetclient