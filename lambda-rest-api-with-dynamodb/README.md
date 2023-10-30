# aws-golang-rest-api-with-dynamodb

Build & Deploy
```
make deploy
```

POST API:

```
curl --location --request POST 'https://h8dvm4gbfa.execute-api.ap-southeast-1.amazonaws.com/dev/todos' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Duong",
    "details": "do something"
}'
```

GET LIST API:
```
curl --location --request GET 'https://h8dvm4gbfa.execute-api.ap-southeast-1.amazonaws.com/dev/todos'
```

GET ONE API:
```
curl --location --request GET 'https://h8dvm4gbfa.execute-api.ap-southeast-1.amazonaws.com/dev/todos/b14cbf0e-4f95-422f-8653-d7c04284551b'
```

UPDATE API:
```
curl --location --request PUT 'https://h8dvm4gbfa.execute-api.ap-southeast-1.amazonaws.com/dev/todos/88fb95ae-22ec-4c91-8785-b11881b32c10' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Duong",
    "details": "updated"
}'
```

DELETE API:
```
curl --location --request DELETE 'https://h8dvm4gbfa.execute-api.ap-southeast-1.amazonaws.com/dev/todos/88fb95ae-22ec-4c91-8785-b11881b32c10'
```