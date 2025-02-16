# ipfs-svc

## Description

This service provides IPFS uploading functionality

## Install

  ```
  git clone github.com/Dmytro-Hladkykh/ipfs-svc
  cd ipfs-svc
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  ./main run service
  ```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.

Note: if you are using Gitlab for building project `docs/spec/paths` folder must not be
empty, otherwise only `Build and Publish` job will be passed.  

## Running from docker 
  
Make sure that docker installed.

use `docker run ` with `-p 8080:80` to expose port 80 to 8080

  ```
  docker build -t github.com/Dmytro-Hladkykh/ipfs-svc .
  docker run -e KV_VIPER_FILE=/config.yaml github.com/Dmytro-Hladkykh/ipfs-svc
  ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `run service` command

## Testing and Local Development

### Setting up IPFS locally

1. **Install IPFS:**

```bash
# MacOS
brew install ipfs

# Linux
wget https://dist.ipfs.tech/kubo/v0.22.0/kubo_v0.22.0_linux-amd64.tar.gz
tar -xvzf kubo_v0.22.0_linux-amd64.tar.gz
cd kubo
sudo bash install.sh
```

2. **Initialize and start IPFS:**

```bash
ipfs init

ipfs daemon
```

3. **Upload endpoint:**

```bash
curl -X POST \
  http://localhost:8000/integrations/ipfs-svc/v1/public/upload \
  -H 'Content-Type: application/vnd.api+json' \
  -d '{
    "data": {
      "type": "upload_json",
      "attributes": {
        "metadata": {
          "name": "Test",
          "description": "Test IPFS upload"
        }
      }
    }
  }'
```

4. **Verify uploaded content:** 

```bash 
ipfs cat <received_hash>
```

### Third-party services


## Contact

Responsible //
The primary contact for this project is  //
