# ProJect-Api-Hotel API

## 🛠 Skills
[![My Skills](https://skillicons.dev/icons?i=go,mysql,postman)](https://skillicons.dev)
- Go lang
- MySql
- Postman

##  DB Diagram  
  ![Logo](https://github.com/Teerapoom/ProJect-Api-Hotel-API/blob/main/IMG/Api_Dromitory002%20(2).png)

## User level  
  ![Logo](https://github.com/Teerapoom/ProJect-Api-Dormitory/blob/main/IMG/User_level.jpg)

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


## JSON Response

### Login

```json
{
    "message": "Successfully logged in",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE3MDE4NTg3NTIsImlhdCI6MTcwMTg0MDc1MiwiaWQiOjMsInJvbGUiOjJ9.BkOlJbJZUcJJDTG_iS9LWbMp7Z7khcvEYQ-WVt1afnc",
    "username": "T0001"
}
```
### Register and User

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
### Room

```json
    {
        "CreatedAt": "2023-12-06T09:44:00.079+07:00",
        "UpdatedAt": "2023-12-06T09:53:38.254+07:00",
        "DeletedAt": null,
        "ID": 1,
        "user_id": 1,
        "name": "101",
        "description": "อาคาร 6",
        "selestatus": "ไม่ว่าง",
        "rent": 5500,
        "bed_combo_mattress": true,
        "table": false,
        "wardrobe": true,
        "tv_shelf": true,
        "shoe_rack": true,
        "statusID": 2,
        "StatusRoom": {
            "CreatedAt": "2023-12-06T09:43:39.884+07:00",
            "UpdatedAt": "2023-12-06T12:38:05.115+07:00",
            "DeletedAt": null,
            "ID": 2,
            "statusname": "ไม่ว่าง"
        }
    }
```

### Checkin

<details>
<summary> Show Details JSON </summary>

```json
       {
        "CreatedAt": "2023-12-06T09:44:13.265+07:00",
        "UpdatedAt": "2023-12-06T09:44:13.265+07:00",
        "DeletedAt": null,
        "ID": 1,
        "user_id": 1,
        "room_id": 1,
        "room_name": "101",
        "user_name_checkin": "T0001",
        "user_name_checkinid": 3,
        "deposit": 5000,
        "rentrate": 1000,
        "contractdate": "2023-11-21T07:00:00+07:00",
        "fullname": "จิรายุ แสงกระจ่าง",
        "dirth_date": "1990-01-01T07:00:00+07:00",
        "issued_by": "Government",
        "card_number": "5820821950430",
        "issued_date": "2020-01-01T07:00:00+07:00",
        "card_copyimg": "path/to/image.jpg",
        "phone1": "0801234567",
        "addr1": "123 Street, City",
        "place1": "Office",
        "renter2": "นวิยา โพธิ์สำราญ",
        "birth_date2": "2023-11-21T07:00:00+07:00",
        "issued_by2": "ภูวนัย ประเสริญวงศ์",
        "card_number2": "1461010111200",
        "issued_date2": "2020-01-01T07:10:00+07:00",
        "card_copyimg2": "path/to/image2.jpg",
        "phone2": "0898464642",
        "addr2": "48/1 หมู่4 ถ.บางกรวย-ไทรนอ้ยอ.บางบวัทองจ.นนทบุรี 11110",
        "place2": "อาคาร1",
        "User": {
            "CreatedAt": "2023-12-06T09:43:50.482+07:00",
            "UpdatedAt": "2023-12-06T09:43:50.482+07:00",
            "DeletedAt": null,
            "ID": 3,
            "role_id": 2,
            "userID": "T0001",
            "seleuser": "ผู้เช่า",
            "username": "ธีรภูมิ คูศิริวานิชกร",
            "email": "user1@example.com",
            "password": "$2a$10$j3LGnsIrArQSZIxsJzJLbeEIMBzH1SBkMXAig79P8cEDZYn9GK85W",
            "numberphone": "07520136432"
        },
        "Room": {
            "CreatedAt": "2023-12-06T09:44:00.079+07:00",
            "UpdatedAt": "2023-12-06T12:40:53.161+07:00",
            "DeletedAt": null,
            "ID": 1,
            "user_id": 1,
            "name": "101",
            "description": "อาคาร 6",
            "selestatus": "ไม่ว่าง",
            "rent": 5500,
            "bed_combo_mattress": true,
            "table": false,
            "wardrobe": true,
            "tv_shelf": true,
            "shoe_rack": true,
            "statusID": 2,
            "StatusRoom": {
                "CreatedAt": "2023-12-06T09:43:39.884+07:00",
                "UpdatedAt": "2023-12-06T12:38:05.115+07:00",
                "DeletedAt": null,
                "ID": 2,
                "statusname": "ไม่ว่าง"
            }
        }
    }
```
</details>

