sudo systemctl start docker
echo Getting swag to update swagger.json
echo I promise...
#sudo rm -rf . haha
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag
swag init -g cmd/notificationService/notification_service_main.go
address=./deployments/docker-compose.yaml
docker-compose -f $address build --parallel
docker-compose -f $address up