# ERP
You can access the ERP here: [dbdiagram.io](https://dbdiagram.io/d/library-be-68417b95ba2a4ac57bf9e159)
![ERP](assets/img/erp.png)

# API Specs
## Users
### Health
- Endpoint: `/v1/users/health`
- Request: - 
- Response: 
    ```
    {
        "code": 200,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": null
    }
    ```
    - `code`: http code of response
    - `message.en`: message response in english
    - `message.id`: message response in bahasa
    - `data`: data response
### Register
- Endpoint: `/v1/users/register`
- Request Body:
    ```
    {
        "email": "dhany@gmail.com",
        "name": "dhany",
        "password": "123456789",
        "role": "admin"
    }
    ```
    - `email`: email of user
    - `name`: name of user
    - `password`: password user
    - `role`: role user, one of [admin, general]
- Response: 
    ```
    {
        "code": 201,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": {
            "user_id": "781cf26e-b7e7-4b63-8bb5-db040416e72b",
            "email": "dhany@gmail.com",
            "name": "dhany",
            "role": "admin"
        }
    }
    ```
    - `code`: http code of response
    - `message.en`: message response in english
    - `message.id`: message response in bahasa
    - `data.user_id`: user_id of user
    - `data.email`: email of user
    - `data.name`: name of user
    - `data.role`: role of user
### Login
- Endpoint: `/v1/users/login`
- Request Body:
    ```
    {
        "email": "dhany@gmail.com",
        "password": "123456789",
    }
    ```
    - `email`: email of user
    - `password`: password user
- Response: 
    ```
    {
        "code": 200,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRoYW55QGdtYWlsLmNvbSIsImV4cCI6ODgyMjk0MzQ1OSwiaWF0IjoxNzQ5MzAyNTI1LCJpZCI6ImQ3NzQxMDcyLTQwYzktNGE0Mi04MDU1LTUzOWRkNDcwZDkyZiIsInJvbGUiOiJhZG1pbiJ9.YgoV9BuJ4sCZ2WZokrN5hX78w6SvgUJ4Xmgu1Sfz3po"
        }
    }
    ```
    - `code`: http code of response
    - `message.en`: message response in english
    - `message.id`: message response in bahasa
    - `data.token`: access token of user
## Books
### Health
- Endpoint: `/v1/books/health`
- Request: - 
- Response: 
    ```
    {
        "code": 200,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": null
    }
    ```
    - `code`: http code of response
    - `message.en`: message response in english
    - `message.id`: message response in bahasa
    - `data`: data response
### Create Author
- Endpoint: `/v1/books/authors`
- Request Body:
    ```
    {
        "name": "eka kurniawan",
        "biography": ""
    }
    ```
    - `name`: name of author
    - `biography`: biography of author
- Response: 
    ```
    {
        "code": 201,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": {
            "author_id": "e4ceb944-e0a4-41c8-a53b-5e0f2f545439",
            "name": "eka kurniawan",
            "biography": ""
        }
    }
    ```
    - `code`: http code of response
    - `message.en`: message response in english
    - `message.id`: message response in bahasa
    - `data.author_id`: author_id of author
    - `data.name`: name of author
    - `data.biography`: biography of author