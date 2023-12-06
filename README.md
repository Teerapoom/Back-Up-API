# ProJect-Api-Dormitory

##  DB Diagram  
  ![Logo](https://github.com/Teerapoom/ProJect-Api-Dormitory/blob/main/IMG/Api_Dromitory002%20(1).png)

## User level  
  ![Logo](https://github.com/Teerapoom/ProJect-Api-Dormitory/blob/main/IMG/User.drawio.png)

## API Reference

### User 🙋‍♂️

#### Login 

```http
  GET /auth/user/login
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `UserName` | `string` | **Required**. UserName such as "T0001" |
| `Password` | `string` | **Required**. Password |

### Admin 🙋‍♂️