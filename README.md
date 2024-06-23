# Book API

## Установка и запуск

1. Клонирование репозитория:
    ```sh
    git clone https://github.com/YarikGuk/BookApi.git
    cd LibraryProject
    ```

2. Сборка и запуск:
    ```sh
    make build
    make run
    ```

3. Адрес - `http://localhost:8080`

## Endpoints

### Создать автора
curl -X POST -H "Content-Type: application/json" -d '{"first_name": "Александр", "last_name": "Пушкин", "biography": "Русский поэт, драматург и прозаик, заложивший основы русского реалистического направления.", "birth_date": "1799-06-06"}' http://localhost:8080/authors

### Получить всех авторов
curl http://localhost:8080/authors

### Получить автора по ID
curl http://localhost:8080/authors/1

### Обновить автора по ID
curl -X PUT -H "Content-Type: application/json" -d '{"first_name": "Фёдор", "last_name": "Достоевский", "biography": "Русский писатель, мыслитель, философ и публицист.", "birth_date": "1821-11-11"}' http://localhost:8080/authors/1

### Удалить автора по ID
curl -X DELETE http://localhost:8080/authors/1

----------------------------

### Создать книгу
curl -X POST -H "Content-Type: application/json" -d '{"title": "Война и мир", "author_id": 1, "year": 1869, "isbn": "978-5-17-087484-6"}' http://localhost:8080/books

### Получить все книги
curl http://localhost:8080/books

### Получить книгу по ID
curl http://localhost:8080/books/1

### Обновить книгу по ID
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Преступление и наказание", "author_id": 1, "year": 1866, "isbn": "978-5-17-087485-3"}' http://localhost:8080/books/1

### Удалить книгу по ID
curl -X DELETE http://localhost:8080/books/1

----------------------------

### Обновить одновременно название книги и биографию автора

curl -X PUT -H "Content-Type: application/json" -d '{  "book_title": "Абсолютно новое название книги", "author_bio": "Обновленная биография автора" }' http://localhost:8080/books/1/authors/1
