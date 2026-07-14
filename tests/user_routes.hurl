# ==========================================
# 1. Test Registration Validation Failure
# ==========================================
POST http://localhost:8080/query
[Options]
variable: initialEmail=h123@example.com
```graphql
mutation CreateUser($input: CreateUser!) {
  createUser(input: $input) {
    id
    name
    email
    roles {
        id
        name
    }
    isActive
  }
}

variables {
    "input": {
        "name": "hor",
        "email": "{{initialEmail}}",
        "password": "Horkitu123",
        "confirmPassword": "Horkitu123",
        "roles": [1,2],
        "isActive": true
    }
}
```
HTTP 200
[Captures]
userTestID: jsonpath "$.data.createUser.id"
[Asserts]
jsonpath "$.data.createUser.name" == "hor"
jsonpath "$.data.createUser.email" == "{{initialEmail}}"

# =============== UPDATE
POST http://localhost:8080/query
[Options]
variable: newEmail=h456@example.com
```graphql
mutation UpdateUser($input: UpdateUser!) {
  updateUser(id:{{userTestID}}, input: $input) {
    id
    name
    email
    roles {
        id
        name
    }
    isActive
  }
}

variables {
    "input": {
        "name": "horkitu",
        "email": "{{newEmail}}",
        "password": "Horkitu123",
        "confirmPassword": "Horkitu123",
        "roles": [1,2],
        "isActive": false
    }
}
```
HTTP 200
[Asserts]
jsonpath "$.data.updateUser.name" == "horkitu"
jsonpath "$.data.updateUser.email" != "{{initialEmail}}"
jsonpath "$.data.updateUser.email" == "{{newEmail}}"
jsonpath "$.data.updateUser.isActive" == false


# # ==========================================

# # 3. Test Login & Capture JWT Token

# # ==========================================

# # Assuming you have a login mutation that returns a token

# POST http://localhost:8080/query
# [GraphQLVariables]
# {
# "email": "testuser_{{uuid}}@example.com",
# "password": "Password123!"
# }

# ```graphql
# mutation Login($email: String!,$password: String!) {
#   login(email: $email, password:$password) {
#     token
#   }
# }

# ```

# HTTP 200
# [Asserts]
# jsonpath "$.errors" not exists
# jsonpath "$.data.login.token" exists
# [Captures]

# # Save the JWT token into a variable for the next request!

# jwt_token: jsonpath "$.data.login.token"

# # ==========================================

# # 4. Authenticated Query (Using captured token)

# # ==========================================

# # Assuming you have a query to get the current authenticated user

# POST http://localhost:8080/query
# Authorization: Bearer {{jwt_token}}

# ```graphql
# query GetCurrentUser {
#   me {
#     id
#     email
#   }
# }

# ```

# HTTP 200
# [Asserts]
# jsonpath "$.errors" not exists
# jsonpath "$.data.me.email" contains "testuser_"

# =============== DELETE USER
POST http://localhost:8080/query
```graphql
mutation DeleteUser {
    deleteUser(id: {{userTestID}})
}
```
HTTP 200
[Asserts]
jsonpath "$.data.deleteUser" == "User with ID {{userTestID}} successfully deleted"