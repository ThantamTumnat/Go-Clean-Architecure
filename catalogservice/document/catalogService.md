# catalogService Documentation

## Introduction

Welcome to the catalogService. In this service, You can manage requests with the function userRead and this service can produce and consume about the topic of userService.

## API Data
An API joke stories .

``API : https://api.sampleapis.com/jokes/goodJokes`` -->

## Endpoint

### GET {url}/v1/catalog/read

In this endpoint, it will check your request that you have an account or not. If you have an account, you can get the joke storie with id of joke storie and this function will produce this data about topic : userHistory

**Request Body:**

```json
{
    "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
    "id": "380" //id of the joke stories
}
```

**Response**:

```json
{
    "status": "OK",
    "status_code": 200,
    "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
    "histories": {
        "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
        "id": "380",
        "type": "programming",
        "setup": "There are 10 kinds of people in this world.",
        "punchline": "Those who understand binary, those who don't, and those who weren't expecting a base 3 joke."
    }
}
```
**Event : userHistory**:
```json
{
    "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
    "id": "380",
    "type": "programming",
    "setup": "There are 10 kinds of people in this world.",
    "punchline": "Those who understand binary, those who don't, and those who weren't expecting a base 3 joke."
}
```

### GET {url}/v1/catalog/getCatalogs

In this endpoint, it will check your request that you have an account or not.
If you have an account, you can get all of joke stories .

**Request Body:**

```json
{
    "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de"
}
```

**Response**:

```json
{
    "status": "OK",
    "status_code": 200,
    "Message": "Successfully, Your can get all of the joke stories",
    "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
    "catalogs": [
        {
            "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
            "id": "1",
            "type": "general",
            "setup": "What did the fish say when it hit the wall?",
            "punchline": "Dam."
        },
        {
            "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
            "id": "2",
            "type": "general",
            "setup": "How do you make a tissue dance?",
            "punchline": "You put a little boogie on it."
        },
        {
            "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
            "id": "3",
            "type": "general",
            "setup": "Why did the tree go to the dentist?",
            "punchline": "It needed a root canal."
        }, // ......... all of the joke stories.
    ]
}
```

### POST {url}/v1/catalog/favorite

In this endpoint, it will check your request that you have an account or not.
If you have an account, you can mark the joke that you favorite.

**Request Body:**

```json
{
    "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
    "id": "380" //id of the joke stories
}
```
**Response**:

```json
{
    "status": "OK",
    "status_code": 200,
    "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
    "favorite": {
        "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
        "id": "380",
        "type": "programming",
        "setup": "There are 10 kinds of people in this world.",
        "punchline": "Those who understand binary, those who don't, and those who weren't expecting a base 3 joke."
    }
}
```

### GET {url}/v1/catalog/getFavorites

In this endpoint, it will check your request that you have an account or not.
If you have an account, you can get all of joke stories that you mark it favorite.

**Request Body:**

```json
{
    "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de"
}
```

**Response**:

```json
{
    "status": "ok",
    "status_code": 200,
    "favorites": [
        {
            "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
            "id": "1",
            "type": "general",
            "setup": "What did the fish say when it hit the wall?",
            "punchline": "Dam."
        },
        {
            "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
            "id": "2",
            "type": "general",
            "setup": "How do you make a tissue dance?",
            "punchline": "You put a little boogie on it."
        },
        {
            "user_id": "b4671803-edcd-4ea2-b926-9d3433a339de",
            "id": "350",
            "type": "general",
            "setup": "Why did the tree go to the dentist?",
            "punchline": "It needed a root canal."
        }
    ]
}
```

---
## Status Codes

<ul>
  <li>200 : OK. Request was successful.</li>
  <li>400 : Bad request. The request was invalid or cannot be served.</li>
  <li>404 : Not found. The request not found the content</li>
  <li>422 : Unprocessable Entity. may something broke on function and bad request body.</li>
</ul>

## Change Log

<ul>
  <li>2023-11-21 : build kafka and design data struct.</li>
  <li>2023-11-22 : make API {url}/user/.... ,{url}/catalog/....</li>
  <li>2023-11-22 : logging data</li>
  <li>2023-11-23 : update documentation</li>
</ul>

## Support

If you have questions, you can But if you can't You need to learn something more and you can contact me to say hello. This is my mail: Ittipat.l@kkumail.com

---





