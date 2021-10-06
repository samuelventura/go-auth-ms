# go-auth-ms

Authentication micro service

## api

```bash
# add app
curl -X POST http://127.0.0.1:31651/api/app/appName
# show app
curl -X GET http://127.0.0.1:31651/api/app/appName
# delete app
curl -X DELETE http://127.0.0.1:31651/api/app/appName
# list apps
curl -X GET http://127.0.0.1:31651/api/app

#post app+dev+email to receive email with code
curl -X POST http://127.0.0.1:31651/api/access \
  -H "Auth-App: appName" \
  -H "Auth-Dev: devId" \
  -H "Auth-DevRt: devRt" \
  -H "Auth-Email: user@domain.tld"
#post app+dev+code to be granted a token
curl -X POST http://127.0.0.1:31651/api/login \
  -H "Auth-App: appName" \
  -H "Auth-Dev: devId" \
  -H "Auth-DevRt: devRt" \
  -H "Auth-Email: user@domain.tld" \
  -H "Auth-Code: 083430"
#post app+dev+token to retrieve token status/email (server side)
curl -X POST http://127.0.0.1:31651/api/token \
  -H "Auth-App: appName" \
  -H "Auth-Dev: devId" \
  -H "Auth-Email: user@domain.tld" \
  -H "Auth-Token: d0d08a5a-0fd7-47a1-ba11-14776eec3e86"
#post app+dev to disable all codes/tokens for that dev
curl -X POST http://127.0.0.1:31651/api/logout/dev \
  -H "Auth-App: appName" \
  -H "Auth-Dev: devId"
#post app+email to disable all codes/tokens for any dev
curl -X POST http://127.0.0.1:31651/api/logout/email \
  -H "Auth-App: appName" \
  -H "Auth-Email: user@domain.tld"
```

## helpers

```bash
#AUTH_LOGS=/var/log
#AUTH_ENDPOINT=127.0.0.1:31651
#AUTH_DB_DRIVER=sqlite|postgres
#AUTH_DB_SOURCE=<driver dependant>
#https://gorm.io/docs/connecting_to_the_database.html
go install && go-auth-ms
go install && go-auth-ss
sqlite3 ~/go/bin/go-auth-ms.db3 ".tables"
sqlite3 ~/go/bin/go-auth-ms.db3 ".schema app_dros"
sqlite3 ~/go/bin/go-auth-ms.db3 ".schema code_dros"
sqlite3 ~/go/bin/go-auth-ms.db3 ".schema token_dros"
sqlite3 ~/go/bin/go-auth-ms.db3 "select * from app_dros"
sqlite3 ~/go/bin/go-auth-ms.db3 "select * from code_dros"
sqlite3 ~/go/bin/go-auth-ms.db3 "select * from token_dros"
#for go-sqlite in linux
sudo apt install build-essentials
```

# security

- for internal network only (api is unsafe)
- tokens/codes are disabled on logout
- tokens/codes disable previous ones when generated
- tokens/codes must be used in same device that requested them
- tokens never expire (logout left to client side)
- codes self disable when used
- codes expire in 5 minutes
