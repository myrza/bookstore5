# Project bookstore5

Учебный проект - представляет из себя симуляцию хранилища информации о книгах и их авторах. 

## Getting Started

1. Клонировать https://github.com/myrza/NewBluePrint
2. Запустить Docker
3. Сформировать 2 контейнера командой make docker-run
4. Проверить работу контейтеров командой docker ps -a
5. Начать работу. Рекомендуемое приложение Postman
6. Запустить тесты из Кейсов тестирования (ниже)

## MakeFile

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```


## Кейсы тестирования 
Запускаем контейнеры  make docker-run
-------------------------------------
База данных: 
1. Проверим работу базы данных
    localhost:8000/health
2. Создаем структуру базы данных 
    localhost:8000/create_db
Авторы:
1. Создание автора 1
    POST localhost:8000/create_author
    JSON: 
    {
        "name": "Герман",
        "surname": "Гессе",
        "biography": "Немецкий писатель и художник, лауреат Нобелевской премии (1946).",
        "birthday": "1877-01-01T00:00:00Z"
    }
2. Создание автора-2 
    POST localhost:8000/create_author
    JSON: 
    {
        "name": "Франц",
        "surname": "Кафка",
        "biography": "Австрийский писатель еврейского происхождения, широко признаваемый как одна из ключевых фигур литературы XX века.",
        "birthday": "1886-07-03T00:00:00Z"
    }
3. Вывод информации по всем авторам
    GET localhost:8000/authors

4. Вывод информации по автору 1
    GET localhost:8000/author/1 

5. Обновление информации по автору 1 (Добавлена дата смерти)
    PUT localhost:8000/authors/1 
    JSON: 
    {
    "name": "Герман",
    "surname": "Гессе",
    "biography": "Немецкий писатель и художник, лауреат Нобелевской премии (1946). Умер  9 августа 1962",
    "birthday": "1877-01-01T00:00:00Z"
}
6. Удаление автора 2 (Кафка)
    DELETE localhost:8000/delete_author/2

7. Вывод информацию по всем авторам
    GET localhost:8000/authors

8. Негативный сценарий 
    GET localhost:8000/author/777

-----------------------------------------
Книги:
1. Создание первой книги автора 1 (Гессе)
    POST localhost:8000/create_book
    JSON:
    {

        "title": "Степной волк",
        "authorid": "1",
        "isbn": "isbn",
        "year": "1927"
    }
2. Создание второй книги автора 1
    POST localhost:8000/create_book
    JSON: 
     {

        "title": "Игра в бисер",
        "authorid": "1",
        "isbn": "isbn",
        "year": "1943"
    }
3. Вывод информации по всем книгам
    GET localhost:8000/books

4. Вывод информации по книге 1
    GET localhost:8000/book/1

5. Обновление информации по книге 1 (поправили ISBN)
    PUT localhost:8000/update_book/1
    JSON:
     {

        "title": "Игра в бисер",
        "authorid": "1",
        "isbn": "isbn-new",
        "year": "1943"
    }
6. Удаление книги 2 
    DELETE localhost:8000/delete_book/1

7. Вывод информации по всем книгам
    GET localhost:8000/books

8. Негативный сценарий 
    GET localhost:8000/book/777
------------------------------------------
Транзакции:
1. Обновление информации по первой книге автора 1
    PUT localhost:8000/update_books_author/1
    JSON: 
    {
    "author_id": 1,
    "name": "Герман-Германович",
    "surname": "Гессе",
    "biography": "No bio",
    "birthday": "2023-01-01T00:00:00Z",
    "title": "Степной волк укусит за бок",
    "isbn": "isbn",
    "year": "1927"
}
2. Вывод информации по первой книге автора 1
    PUT localhost:8000/book_author/1
3. Негативный сценарий. Неправильный формат даты - меняем также имя автора, чтобы проверить что транзакция откатилась
    PUT localhost:8000/update_books_author/1
    JSON: 
    {
    "author_id": 1,
    "name": "Герман-Германович-Transaction",
    "surname": "Гессе",
    "biography": "No bio",
    "birthday": "01.01.2020333",
    "title": "Степной волк укусит за бок",
    "isbn": "isbn",
    "year": "1927"
}
4. Убеждаемся, что имя автора не поменялось
    PUT localhost:8000/book_author/1
