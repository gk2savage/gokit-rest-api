# golang-rest-api

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://github.com/go-kit/kit)

Implementation of REST API with gokit framework

![Go](https://miro.medium.com/max/500/1*Gbi_XNOkPFbWkIkJC7LnBQ.gif)

https://github.com/go-kit/kit
- GIN
- GORM
- SQLITE

```
[GIN-debug] GET    /products                 --> main.getProducts (3 handlers)
[GIN-debug] GET    /products/:id             --> main.getProductById (3 handlers)
[GIN-debug] POST   /products                 --> main.insertProduct (3 handlers)
[GIN-debug] PUT    /products/:id             --> main.updateProduct (3 handlers)
[GIN-debug] DELETE /products/:id             --> main.deleteProduct (3 handlers)
[GIN-debug] Listening and serving HTTP on :1991
```
