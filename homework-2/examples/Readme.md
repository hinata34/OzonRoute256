### CURL запросы для проверки Server
curl -v "http://localhost:9000/"

curl -v "http://localhost:9000/?nice" -X DELETE

curl -X PUT "http://localhost:9000" -d "nice"

curl -X POST "http://localhost:9000" -H "hw-sum: 1a"

curl -X POST "http://localhost:9000" -H "hw-sum: -6"

### CURL запросы для проверки ServerWithData
curl -v -X POST "http://localhost:9001" -d '{ "d": 0, "value": "nice" }'

curl -v -X POST "http://localhost:9001" -d '{ "id": 2, "val": "lolkek" }'

curl -v -X POST "http://localhost:9001" -d '{ "id": -1, "value": "lolkek" }'

curl -v -X PUT "http://localhost:9001" -d '{ "d": 0, "value": "nice" }'

curl -v -X PUT "http://localhost:9001" -d '{ "id": 2, "val": "lolkek" }'

curl -v -X GET "http://localhost:9001/?id=-1"

curl -v -X DELETE "http://localhost:9001/?id=-1"
