# Todo
Not implemented yet:
- [ ] Transaction book in service transaction (should call grpc borrow/return books)
- [ ] Add redis in get category and get author
- [ ] Testing for each handler and services
- [ ] Deployment using Docker

# ERP
You can access the ERP here: [dbdiagram.io](https://dbdiagram.io/d/library-be-68417b95ba2a4ac57bf9e159)
![ERP](assets/img/erp.png)

# Run Program
pleases make in differen terminal in root project folder
- `make run-users-services`: for run service users
- `make run-books-services`: for run service books
  

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
### Get Author By ID
- Endpoint: `/v1/books/authors/:id`
- Request Param:
    - `id`: id of author
- Response: 
    ```
    {
        "code": 200,
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
### Create Category
- Endpoint: `/v1/books/categories`
- Request Body:
    ```
    {
        "name": "fiksi",
        "description": "ini adalah fiksi"
    }
    ```
    - `name`: name of category
    - `description`: description of category
- Response: 
    ```
    {
        "code": 201,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": {
            "category_id": "636de519-e721-462f-9311-be1e67b40f5f",
            "name": "fiksi",
            "description": "ini adalah fiksi"
        }
    }
    ```
    - `code`: http code of response
    - `message.en`: message response in english
    - `message.id`: message response in bahasa
    - `data.category_id`: category_id of category
    - `data.name`: name of category
    - `data.description`: description of category
### Get Category By ID
- Endpoint: `/v1/books/categories/:id`
- Request Param:
    - `id`: id of category
- Response: 
    ```
    {
        "code": 200,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": {
            "category_id": "636de519-e721-462f-9311-be1e67b40f5f",
            "name": "fiksi",
            "description": "ini adalah fiksi"
        }
    }
    ```
    - `code`: http code of response
    - `message.en`: message response in english
    - `message.id`: message response in bahasa
    - `data.category_id`: category_id of category
    - `data.name`: name of category
    - `data.description`: description of category
### Create Books
- Endpoint: `/v1/books`
- Request Body:
    ```
    {
        "title": "donxuito",
        "isbn": "1234567892",
        "stock": 10,
        "description": "ini adalah deskripsi buku",
        "author_id": "e4ceb944-e0a4-41c8-a53b-5e0f2f545439",
        "category_id": "636de519-e721-462f-9311-be1e67b40f5f"
    }
    ```
- Response: 
    ```
    {
        "code": 201,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": {
            "book_id": "4532ff53-d069-41f4-bdfd-27a7803bf5e0",
            "title": "cantik itu luka",
            "isbn": "1234567890",
            "stock": 10,
            "description": "ini adalah deskripsi buku",
            "author": {
                "author_id": "e4ceb944-e0a4-41c8-a53b-5e0f2f545439",
                "name": "eka kurniawan",
                "biography": ""
            },
            "category": {
                "category_id": "636de519-e721-462f-9311-be1e67b40f5f",
                "name": "fiksi",
                "description": "ini adalah fiksi"
            }
        }
    }
    ```

### Get Book By ID
- Endpoint: `/v1/books/:id`
- Request Param:
    - `id`: id of book
- Response: 
    ```
    {
        "code": 200,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": {
            "book_id": "4532ff53-d069-41f4-bdfd-27a7803bf5e0",
            "title": "cantik itu luka",
            "isbn": "1234567890",
            "stock": 10,
            "description": "ini adalah deskripsi buku",
            "author": {
                "author_id": "e4ceb944-e0a4-41c8-a53b-5e0f2f545439",
                "name": "eka kurniawan",
                "biography": ""
            },
            "category": {
                "category_id": "636de519-e721-462f-9311-be1e67b40f5f",
                "name": "fiksi",
                "description": "ini adalah fiksi"
            }
        }
    }
    ```
### Search Books
- Endpoint: `/v1/books/search?title=lau&isbn=1234567891&author_id&category_id`
- Query Param:
    - `title`: title of book
    - `isbn`: isbn of book
    - `author_id`: author_id of book
    - `category_id`: category_id of book
- Response: 
    ```
    {
        "code": 200,
        "message": {
            "en": "Success",
            "id": "Berhasil"
        },
        "data": [{
            "book_id": "4532ff53-d069-41f4-bdfd-27a7803bf5e0",
            "title": "cantik itu luka",
            "isbn": "1234567890",
            "stock": 10,
            "description": "ini adalah deskripsi buku",
            "author": {
                "author_id": "e4ceb944-e0a4-41c8-a53b-5e0f2f545439",
                "name": "eka kurniawan",
                "biography": ""
            },
            "category": {
                "category_id": "636de519-e721-462f-9311-be1e67b40f5f",
                "name": "fiksi",
                "description": "ini adalah fiksi"
            }
        }]
    }
    ```