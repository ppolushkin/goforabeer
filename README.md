# Go for a beer

https://go-for-beer.azurewebsites.net/

## Prepare
cd src/sidenis.com/goforbeer/   
sudo docker login   
az login

## Build app
go clean && CGO_ENABLED=0 go build -a -installsuffix cgo

## Build docker image
sudo docker build -t ppolushkin/goforbeer:1.0 .

## Push docker image to DockerHub
sudo docker push ppolushkin/goforbeer:1.0

## Restart az web-app
az webapp restart -n go-for-beer -g web-apps

## Tail az web-app logs
az webapp log tail -n go-for-beer -g web-apps

## Rebuild az web-app
go clean && CGO_ENABLED=0 go build -a -installsuffix cgo && sudo docker build -t ppolushkin/goforbeer:1.0 . && sudo docker push ppolushkin/goforbeer:1.0 && az webapp restart -n go-for-beer -g web-apps

# Run docker locally
sudo docker run -p 8080:8080 ppolushkin/goforbeer:1.0



2018-10-02T13:49:50.478603395Z JWT token validated: isValid= true, Error = %!s(<nil>)VerifyAccessToken error request for metadata was not successful: 
Get https://identity-np.swissre.com/oauth2/default/.well-known/openid-configuration: x509: certificate signed by unknown authority/secured/beers called
2018-10-02T13:49:50.478810795Z authHeader 
