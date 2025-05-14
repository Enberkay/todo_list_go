| Method | Endpoint                  | คำอธิบาย                       |
| ------ | ------------------------- | ------------------------------ |
| POST   | `/api/users`              | สร้างผู้ใช้ใหม่                |
| GET    | `/api/users`              | ดูผู้ใช้ทั้งหมด                |
| GET    | `/api/users/:id`          | ดูผู้ใช้คนเดียว                |
| PUT    | `/api/users/:id`          | แก้ไขผู้ใช้                    |
| DELETE | `/api/users/:id`          | ลบผู้ใช้                       |
| POST   | `/api/todos`              | สร้าง Todo (พร้อม user\_id)    |
| GET    | `/api/todos`              | ดู Todo ทั้งหมด                |
| GET    | `/api/todos/:id`          | ดู Todo รายการเดียว            |
| PUT    | `/api/todos/:id`          | แก้ไข Todo                     |
| DELETE | `/api/todos/:id`          | ลบ Todo                        |
| GET    | `/api/todos/user/:userId` | ดู Todo ทั้งหมดของผู้ใช้คนนั้น |
