# Token Generator API ⛓️

This project is used to generate secret tokens for internal learners. It is written in the GoLang programming language

## Example

#### Request
```http
GET /api/token/generate

# Default value: 10
token-length: 25
```

#### Response
```json
{
    "token": "b9c53dd4e308abbef7988447a55927ea7af496e378fb158b62"
}
```

Written by: Petr Němeček