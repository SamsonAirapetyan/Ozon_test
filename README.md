# Ozon_test:<p>ShortLink Service

Проект представляет собой сервис по укорачиванию ссылок

**Используемые технологии:**
* `gRPC-gateway`
* `PostgreSQL`
* `Mock Testing`
* `Docker`

## API Endpoints ##
<h3>Create Short Link<h3>
<h6>Endpoint: `POST /create`</h6>
  
`POST /create body{"link": "ozon.ru"}`
```shell
#POST
curl -X POST localhost:8080/create -H "Content-Type: application/json" -d '{"link": "ozon.ru"}'
```
<h6>Описание</h6>

Создает короткую ссылку для предоставленной исходной ссылки

<h3>Retrieve Original Link</h3>
<h6>Endpoint: GET /get/{shortLink}</h6>

`GET /get/{shortLink}`
```shell
#GET
curl -X GET localhost:8080/get/Lw1XBy9jH5
```

<h6>Описание</h6>
Получает исходную ссылку, соответствующую заданной короткой ссылке

## Usage ##

```shell 
#Выбор PostgreSQL в качестве хранилища
make psql
#Выбор ii-memory в качестве хранилища
make in-memory
#Запуск тестов
make tests
```

