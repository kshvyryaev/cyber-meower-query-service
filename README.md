# Query service

Service for searching meow messages

## Run elasticsearch docker

`docker network create elasticsearch`
`docker run -d --name elasticsearch --network elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.14.2`
