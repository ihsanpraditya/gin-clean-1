# ==========================================
# CREATE
# ==========================================
POST http://localhost:8080/query
[Options]
variable: initialName=hor
variable: initialEmail=h123@example.com
```graphql
mutation CreateUsers($input1: CreateUser!, $input2: CreateUser!) {
  firstUser: createUser(input: $input1) {
    id
    name
    email
    roles {
        id
        name
    }
    isActive
  }
  
  secondUser: createUser(input: $input2) {
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
    "input1": {
        "name": "{{initialName}}",
        "email": "{{initialEmail}}",
        "password": "Horkitu123",
        "confirmPassword": "Horkitu123",
        "roles": [1,2],
        "isActive": true
    },
    "input2": {
        "name": "secondUser",
        "email": "second@example.com",
        "password": "Horkitu123",
        "confirmPassword": "Horkitu123",
        "roles": [1,2],
        "isActive": true
    }
}
```
HTTP 200
[Captures]
userTestID: jsonpath "$.data.firstUser.id"
secondUserTestID: jsonpath "$.data.secondUser.id"
[Asserts]
jsonpath "$.data.firstUser.name" == "{{initialName}}"
jsonpath "$.data.firstUser.email" == "{{initialEmail}}"

# ==========================================
# READ ONE
# ==========================================
POST http://localhost:8080/query
```graphql
query ReadUser {
    user(id: {{userTestID}}) {
        name
        email
    }
}
```
HTTP 200
[Asserts]
jsonpath "$.data.user.name" == "{{initialName}}"
jsonpath "$.data.user.email" == "{{initialEmail}}"

# ==========================================
# READ MANY
# ==========================================
POST http://localhost:8080/query
```graphql
query ReadUsers {
    users {
        name
        email
    }
}
```
HTTP 200
[Asserts]
jsonpath "$.data.users[0].email" == "{{initialEmail}}"
jsonpath "$.data.users[*].name" contains "{{initialName}}"

# ==========================================
# UPDATE
# ==========================================
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

# ==========================================
# DELETE ONE
# ==========================================
POST http://localhost:8080/query
```graphql
mutation DeleteUser {
    deleteUser(id: {{userTestID}})
}
```
HTTP 200
[Asserts]
jsonpath "$.data.deleteUser" == "User with ID {{userTestID}} successfully deleted"

# ==========================================
# DELETE MANY
# ==========================================
POST http://localhost:8080/query
```graphql
mutation DeleteUsers {
    deleteUsers(ids: [{{secondUserTestID}}])
}
```
HTTP 200
[Asserts]
jsonpath "$.data.deleteUsers" == "Users with IDs [{{secondUserTestID}}] successfully deleted"