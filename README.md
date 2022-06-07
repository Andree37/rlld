# rlld.co

URL shorter that either sends you to the correct url or to some random meme.

## Run this locally
Follow these steps to run this locally:

### Setup Environment
Create a .env using the one in the `.env.example` and set the variables to something that makes sense to you :)

### Run 
To run this locally, ensure that you have docker installed on your machine, and it's up and running.

Also ensure you have rlld-frontend (https://github.com/Andree37/rlld-frontend) pulled on a directory neighboring this.

Go to the root directory and run this:
```docker
docker-compose -f docker-compose.yaml up   
```

Access `http://localhost:80` on your browser and enjoy the whole app.

