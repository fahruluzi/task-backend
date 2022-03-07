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
  "phone" : "081320243889",
  "role" : "admin"
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
  "success" : true
}
```

### Login
`POST /user/login`

*Request*
> Body Params : JSON
```json
{
  "phone": "081320243889",
  "password" : "ZA10"
}
```

*Response*

> HTTP 200 : Example Response
```json
{
  "message": "User Authenticated!",
  "data": {
    "claims": {
      "name" : "Fahrul Fauzi",
      "phone" : "081320243889",
      "role" : "admin",
      "authenticated_at" : "1646669672444"
    },
    "token" : "ey.."
  },
  "success" : true
}
```

> HTTP 400 : Example Response
```json
{
  "message": "Invalid Phone Number or Password!",
  "data": null,
  "success" : false
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
    "claims": {
      "name" : "Fahrul Fauzi",
      "phone" : "081320243889",
      "role" : "admin",
      "authenticated_at" : "1646669672444"
    },
    "token" : "ey.."
  },
  "success" : true
}
```

> HTTP 401 : Example Response
```json
{
  "message": "Invalid Token!",
  "data": null,
  "success" : false
}
```