# rlld.co

URL shorter that either sends you to the correct url or to some random meme.

## Run this locally
Follow these steps to run this locally:

### Setup Environment
Create a .env using the one in the `.env.example` and set the variables to something that makes sense to you :)

### Run 
To run this locally, ensure that you have docker installed on your machine, and it's up and running.
Go to the root directory and run this:
```docker
docker-compose -f docker-compose.yaml up   
```

Access `http://localhost:8080` on your browser and enjoy the api.

### Test the App
To run some *pristine* unit tests do:
```terminal
go test ./...
```
