# rlld.co

URL shortner that either sends you to the correct url or to some random meme.

## Run this locally
Follow these steps to run this locally

> TODO: create a docker-compose for the whole application. For now it only runs the database. Bit of an 

### Database
To run this locally, ensure that you have docker installed on your machine and it's running.
Go to the root directory and run this:
```docker
docker-compose -f docker-compose.yaml up   
```

### Server
To run the server locally on port 8080 and Postgres connection do:
```terminal
cd backend; go run .
```