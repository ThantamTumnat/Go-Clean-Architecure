# userService Documentation

## Introduction

Welcome to the UserService this service, You can manage requests with functions userReaded, userUpdated, userDeleted, and userCreate. The last in this service can produce and consume about the topic of catalogService.

## Endpoints 

### POST {url}/v1/user/created

Create a user account with topic "userCreated".

**Request Body:**

```json
{
    "username": "Example",
    "password": "H0xI51XI8PyylMF",
    "address": "Khon kaen"
}
```
**Response**:

```json
{
    "status": "OK",
    "status_code": 200,
    "message": "Successfully, Your account was created",
    "user_id": "7a779ce2-a96b-457f-a90f-ac979ce2640b",
    "username": "Example"
}
```
**Event : userCreated**:
```json
{
    "user_id": "7a779ce2-a96b-457f-a90f-ac979ce2640b",
    "username": "Example",
    "password": "H0xI51XI8PyylMF",
    "address": "Khon kaen"
}
```

### PUT {url}/v1/user/updated

Update a user account with topic "userUpdated".

**Request Body:**

```json
{
    "user_id" : "7a779ce2-a96b-457f-a90f-ac979ce2640b",
    "username": "Example",
    "password": "F0xI51XI8PyylMF",
    "address": "Mukdahan"
}
```

**Response**:

```json
{
    "status": "OK",
    "status_code": 200,
    "message": "Successfully, Your account was updated",
    "user_id": "7a779ce2-a96b-457f-a90f-ac979ce2640b",
    "password": "F0xI51XI8PyylMF",
    "address": "Mukdahan",
}
```

**Event : userUpdated**:
```json
{
    "user_id" : "7a779ce2-a96b-457f-a90f-ac979ce2640b",
    "username": "Example",
    "password": "F0xI51XI8PyylMF",
    "address": "Mukdahan"
}
```

### GET {url}/v1/user/readed

In this endpoint, you can get all the joke stories that you have ever read

**Request Body:**

```json
{
    "UserID": "b4671803-edcd-4ea2-b926-9d3433a339de"
}
```

**Response**:

```json
{
    "status": "ok",
    "status_code": 200,
    "History": [
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

### DELETE {url}/v1/userDeleted

Delete a user account with topic "userDeleted".

**Request Body:**

```json
{
    "UserID": "e919d31f-dc2d-40fa-bc0c-09172fdd7b4b"
}
```


**Response**:

```json
{
    "status": "OK",
    "status_code": 200,
    "message": "Successfully, Your account was deleted",
    "user_id": "e919d31f-dc2d-40fa-bc0c-09172fdd7b4b"
}
```

**Event : userDeleted**:
```json
{
    "UserID": "e919d31f-dc2d-40fa-bc0c-09172fdd7b4b"
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
<!-- ## Consume Topic

### userHistory

when catalogService was produce data with this topic, it will keep the data into the database.

**Request Massage**:

```json
{
    "UserID": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "id" : 388, //id of joke storie
    "type" : "Programming",
    "setup" : "My code is dead",
    "punchline": "ahh I forgot ; "
}
``` -->