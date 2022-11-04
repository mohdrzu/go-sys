# Go-SYS
### Example of Dockerized golang application.

## How to run

1. Build the image

 `docker build -t go-sys-image .`

2. Create container

`docker container create --name go-sys-container -p 8080:9000 go-sys-image`

3. Start the container

`docker container start go-sys-container `

4. Open web browser and navigate to 

`localhost:8080`

