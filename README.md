# rlld.co

URL shortner that either sends you to the correct url or to some random meme.

## Run this locally
Follow these steps to run this locally

### Database
To run this locally, ensure that you have docker installed and run this:
``` docker
docker run --name mongodb -d -p 27017:27017 mongo
```

### Server
To run the server locally on port 8080 and with that mongo connection do:
```terminal
cd backend; go run .
```