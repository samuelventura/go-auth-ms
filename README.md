# go-auth-ms

Auth micro service

## api

```bash
# add app
curl -X POST http://127.0.0.1:31650/api/app/appID
# delete app
curl -X DELETE http://127.0.0.1:31650/api/app/appID
# list apps
curl -X GET http://127.0.0.1:31650/api/app

#post app+dev+email to receive email with code
curl -X POST http://127.0.0.1:31651/api/access \
    --data '{"app":"appID","dev":"devID","email":"user@domain.tld"}'
#post app+dev+code to be granted a token
curl -X POST http://127.0.0.1:31651/api/login \
    --data '{"app":"appID","dev":"devID","code":"012345"}'
#post app+dev+token to retrieve token status/email (server side)
curl -X POST http://127.0.0.1:31651/api/token \
    --data '{"app":"appID","dev":"devID","token":"k8D3jY"}'
#post app+dev to disable all codes/tokens for that dev
curl -X POST http://127.0.0.1:31651/api/logout \
    --data '{"app":"appID","dev":"devID"}'
#post app+email to disable all codes/tokens for any dev
curl -X POST http://127.0.0.1:31651/api/logout \
    --data '{"app":"appID","email":"user@domain.tld"}'
```

# security

- tokens/codes expire in 5 minutes
- tokens/codes must be used in same device that requested them
- tokens/codes disable previous ones when generated
