echo starting..
docker start
address=./deployments/docker-compose.yaml
docker-compose -f $address build --parallel
docker-compose -f $address up