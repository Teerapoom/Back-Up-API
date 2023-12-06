# ProJect-Api-Dormitory

##  DB Diagram  
  ![Logo](https://github.com/Teerapoom/ProJect-Api-Dormitory/blob/main/IMG/Api_Dromitory002%20(1).png)

## User level  
  ![Logo](https://github.com/Teerapoom/ProJect-Api-Dormitory/blob/main/IMG/User.drawio.png)

## API Reference

### Role User 🙋‍♂️

#### Login 

```http
  POST /auth/user/login
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `UserName` | `string` | **Required**. UserName such as "T0001" |
| `Password` | `string` | **Required**. Password |

### Role Admin 👩‍💻

[Documentation Role Admin](https://docs.google.com/document/d/1iVSmTnf7N_W1tlCR89I_UlMG7tMIZSfzgtpL8nS0Ml0/edit?usp=sharing)


## Test JSON 

### Login

```json
    "message": "Successfully logged in",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE3MDE4NTg3NTIsImlhdCI6MTcwMTg0MDc1MiwiaWQiOjMsInJvbGUiOjJ9.BkOlJbJZUcJJDTG_iS9LWbMp7Z7khcvEYQ-WVt1afnc",
    "username": "T0001"
```
### Register

```json
{
    "user": {
        "CreatedAt": "2023-12-06T12:35:40.409+07:00",
        "UpdatedAt": "2023-12-06T12:35:40.409+07:00",
        "DeletedAt": null,
        "ID": 11,
        "role_id": 2,
        "userID": "T0003",
        "seleuser": "ผู้เช่า",
        "username": "พงศภัค โชคชัย",
        "email": "user3@example.com",
        "password": "$2a$10$MuxIKTY.ZN5QMXDG1vg5/.xC/wQ7rC6n4WHxo9XeiKJu88QdgbSZO",
        "numberphone": "07520136432"
    }
}
```
