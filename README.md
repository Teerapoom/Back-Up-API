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

#### register
```http
  POST /admin/register
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `UserName` | `string` | **Required**. UserName such as "T0001" |
| `Password` | `string` | **Required**. Password |


#### View All User
```http
  GET /admin/users
```

#### View User BY ID
```http
  GET /admin/users/:id
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `ID` | `string` | **Required**. ID of user |