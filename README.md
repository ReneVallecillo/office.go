
Test authorized url

```
curl -X GET http://localhost:8080/profile
```

Authenticate
```
curl -v -X POST -c cookie.txt -F "email=email@test.com" -F "password=test123" http://localhost:8080/api/v1/login
```

Test With Token in Cookie

```
curl -v -X GET \
-b cookie.txt \
http://localhost:8080/profile
```

If -b does not work use:

curl -v --cookie \
"Auth=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImVtYWlsQHRlc3QuY29tIiwiZXhwIjoxNDc0NzU4MDM1LCJpc3MiOiJsb2NhbGhvc3Q6OTAwMCJ9.-HcQHvhYgn5ZUhnTmzHgEh-UGd7ms0tktR9LlnZKCN8" 
\-X GET http://localhost:8080/profile

remember to change the token string

TODO:
Old cookies are working, check how to deal with old tokens or expired ones
