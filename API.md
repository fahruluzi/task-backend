# Microservice Task Backend API Specs

Documentation API for Microservice Auth app and Fetch app

## Authentication

- HTTP Authorization, Bearer Schema

## Auth App

### Register

`POST /user/register`

*Request*
> Body Params : JSON

```json
{
  "username": "fahruluzi",
  "name": "Fahrul Fauzi",
  "phone": "081320243889",
  "role": "admin"
}
```

*Response*

> HTTP 201 : Example Response

```json
{
  "message": "User created!",
  "data": {
    "uuid": null,
    "username": "fahruluzi",
    "name": "Fahrul Fauzi",
    "phone": "081320243889",
    "role": "admin",
    "password": "ZA10"
  },
  "success": true
}
```

### Login

`POST /user/login`

*Request*
> Body Params : JSON

```json
{
  "phone": "081320243889",
  "password": "ZA10"
}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "User Authenticated!",
  "data": {
    "claims": {
      "name": "Fahrul Fauzi",
      "phone": "081320243889",
      "role": "admin",
      "authenticated_at": "1646669672444"
    },
    "token": "ey.."
  },
  "success": true
}
```

> HTTP 400 : Example Response

```json
{
  "message": "Invalid Phone Number or Password!",
  "data": null,
  "success": false
}
```

### Profile

`POST /user/profile`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "User Authenticated!",
  "data": {
    "name": "Fahrul Fauzi",
    "phone": "081320243889",
    "role": "admin",
    "authenticated_at": "1646669672444"
  },
  "success": true
}
```

> HTTP 401 : Example Response

```json
{
  "message": "Invalid Token!",
  "data": null,
  "success": false
}
```

## Fetch App 
`localhost:4000`

### Validate
`POST /validate`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "User Authenticated!",
  "data": {
    "name": "Fahrul Fauzi",
    "phone": "081320243889",
    "role": "admin",
    "authenticated_at": "1646669672444"
  },
  "success": true
}
```

> HTTP 401 : Example Response

```json
{
  "message": "Invalid Token!",
  "data": null,
  "success": false
}
```

### Fetch
`POST /fetch`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "Success Fetch Data!",
  "data": [
    {
      "uuid": "383e524e-5d7d-4e2a-93f7-21e0271206ca",
      "komoditas": "Cupang",
      "area_provinsi": "KALIMANTAN TIMUR",
      "area_kota": "BORNEO",
      "size": "120",
      "price": "200000",
      "price_usd" : "12.12312",
      "tgl_parsed": "2022-03-08T19:15:42Z",
      "timestamp": null
    }
  ],
  "success": true
}
```

> HTTP 401 : Example Response

```json
{
  "message": "Invalid Token!",
  "data": null,
  "success": false
}
```

> HTTP 403 : Example Response
```json
{
  "message": "Role not valid to access!",
  "data": null,
  "success": false
}
```

### Fetch
`POST /aggregation`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "message": "Success Aggregation Data!",
  "data": [
    {
      "year": 2022,
      "month": 3,
      "week": 2,
      "province": "KALIMANTAN TIMUR",
      "commodity" : [
        {
          "name" : "Cupang",
          "size" : {
            "maximal" : 12,
            "minimal" : 1,
            "median" : 10,
            "average" : 6
          },
          "price" : {
            "maximal" : 12,
            "minimal" : 1,
            "median" : 10,
            "average" : 6
          }
        }
      ]
    }
  ],
  "success": true
}
```

> HTTP 401 : Example Response

```json
{
  "message": "Invalid Token!",
  "data": null,
  "success": false
}
```

> HTTP 403 : Example Response
```json
{
  "message": "Role not valid to access!",
  "data": null,
  "success": false
}
```