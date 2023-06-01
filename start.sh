echo Getting swag to update swagger.json
echo I promise...
sudo systemctl start docker
#sudo rm -rf . haha
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag
go test ./test/.
swag init -g cmd/notificationService/notification_service_main.go
address=./deployments/docker-compose.yaml
docker-compose -f $address build
docker-compose -f $address up