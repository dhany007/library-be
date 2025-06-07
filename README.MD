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