# dig-api
dig-api is for checking if private dns is working correctly with serverless application.

## Usage
1. Deploy docker image `kyomo/dig-api:latest`

2. check ip address of a domain
```
curl https://{HOST}/lookup/ip?domain=google.com
```
