
<p>
Простое REST API с одним единственным методом, который загружает изображения.
</p>

<p>
Реализовано на go 1.11
<br>
В качестве менеджера зависимостей используется go modules (разворачивать за пределами GOPATH)
</p>

<p>
Запуск:
<br>
go run api/main.go
<br>
или
<br>
docker-compose up
</p>

<p>
Отправить POST запрос на адрес localhost:8080/photo с полями:
file type="file"
url type="text"
base64image type="text"
</p>

<br>
коллекция в постмане для тестирования:
https://www.getpostman.com/collections/eeb06f4e4f2f15e80179
