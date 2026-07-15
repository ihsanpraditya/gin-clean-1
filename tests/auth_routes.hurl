# ==========================================
# LOGIN
# ==========================================

# creating a user
POST http://localhost:8080/query
[Options]
variable: initialName=hor
variable: initialEmail=hor@example.com
variable: initialPassword=Horkitu123
```graphql
mutation CreateUsers($input: CreateUser!) {
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
        "name": "{{initialName}}",
        "email": "{{initialEmail}}",
        "password": "{{initialPassword}}",
        "confirmPassword": "{{initialPassword}}",
        "roles": [1,2],
        "isActive": true
    }
}
```
HTTP 200
[Captures]
userTestID: jsonpath "$.data.createUser.id"

# REAL LOGIN
POST http://localhost:8080/query
```graphql
mutation {
  login(email: "{{initialEmail}}", password: "{{initialPassword}}") {
    token
  }
}
```
HTTP 200
[Asserts]
jsonpath "$.data.login.token" exists
# [Captures]
# jwt_token: jsonpath "$.data.login.token"

# deleting user
POST http://localhost:8080/query
```graphql
mutation DeleteUser {
    deleteUser(id: {{userTestID}})
}
```
HTTP 200
[Asserts]
jsonpath "$.data.deleteUser" == "User with ID {{userTestID}} successfully deleted"