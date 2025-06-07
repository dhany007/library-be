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
        "password": "dhany",
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