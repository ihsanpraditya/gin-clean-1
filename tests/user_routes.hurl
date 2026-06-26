# ==============================================================================
# HURL TEST SUITE FOR USER ROUTE
# ==============================================================================

# 1. POST - Register User Baru
POST http://localhost:8080/users
Content-Type: application/json
{
  "name": "Ihsan Praditya",
  "email": "ihsan@example.com",
  "password": "secretpassword123"
}

HTTP 201
[Asserts]
header "Content-Type" contains "application/json"
jsonpath "$.message" == "User registered successfully"
jsonpath "$.data.name" == "Ihsan Praditya"
jsonpath "$.data.email" == "ihsan@example.com"
jsonpath "$.data.id" exists
# Tangkap ID user dari response untuk dipakai di request berikutnya
[Captures]
created_user_id: jsonpath "$.data.id"


# 2. POST - Register dengan Email yang Sama (Ekspektasi: Conflict 409)
POST http://localhost:8080/users
Content-Type: application/json
{
  "name": "Cloning Ihsan",
  "email": "ihsan@example.com",
  "password": "anotherpassword"
}

HTTP 409
[Asserts]
jsonpath "$.error" == "email already taken"


# 3. GET - Mengambil Semua List Users
GET http://localhost:8080/users

HTTP 200
[Asserts]
jsonpath "$" isCollection
jsonpath "$[0].id" exists


# 4. GET - Mengambil Data User Berdasarkan ID yang Baru Dibuat
GET http://localhost:8080/users/{{created_user_id}}

HTTP 200
[Asserts]
jsonpath "$.id" == {{created_user_id}}
jsonpath "$.name" == "Ihsan Praditya"


# 5. PUT - Memperbarui Data User (Update)
PUT http://localhost:8080/users/{{created_user_id}}
Content-Type: application/json
{
  "name": "Ihsan Praditya Updated",
  "email": "ihsan.new@example.com"
}

HTTP 200
[Asserts]
jsonpath "$.message" == "User updated successfully"
jsonpath "$.data.name" == "Ihsan Praditya Updated"
jsonpath "$.data.email" == "ihsan.new@example.com"


# 6. DELETE - Menghapus User Berdasarkan ID
DELETE http://localhost:8080/users/{{created_user_id}}

HTTP 200
[Asserts]
jsonpath "$.message" == "User deleted successfully"


# 7. GET - Memastikan User Sudah Terhapus (Ekspektasi: Not Found 404)
GET http://localhost:8080/users/{{created_user_id}}

HTTP 404
[Asserts]
jsonpath "$.error" == "user not found"
