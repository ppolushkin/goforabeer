# Go for a beer

https://go-for-beer.azurewebsites.net/

## Prepare
cd src/sidenis.com/goforbeer/   
sudo docker login   
az login

## Build app
go clean && CGO_ENABLED=0 go build -a -installsuffix cgo

## Build docker image
sudo docker build -t ppolushkin/goforbeer:0.1 .

## Push docker image to DockerHub
sudo docker push ppolushkin/goforbeer:0.1

## Restart az web-app
az webapp restart -n go-for-beer -g web-apps

## Tail az web-app logs
az webapp log tail -n go-for-beer -g web-apps

## Rebuild az web-app
go clean && CGO_ENABLED=0 go build -a -installsuffix cgo && sudo docker build -t ppolushkin/goforbeer:0.1 . && sudo docker push ppolushkin/goforbeer:0.1 && az webapp restart -n go-for-beer -g web-apps

# Run docker locally
sudo docker run -p 8080:8080 ppolushkin/goforbeer:0.1