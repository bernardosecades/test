# Test

## Docker

`docker-compose up`

This will stop all containers and remove them (if you have some lock file with mysql problem):

`docker stop $(docker ps -a -q)`
`docker rm $(docker ps -a -q)`

BloomRPC:

`localhost:3333`

