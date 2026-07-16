# login, in case of super_admin has the role in policy.csv
POST http://localhost:8080/query
```graphql
mutation {
  login(email: "superadmin@example.com", password: "password123") {
    token
  }
}
```
HTTP 200
[Captures]
token: jsonpath "$.data.login.token"

# ==========================================
# CREATE
# ==========================================
POST http://localhost:8080/query
Authorization: Bearer {{token}}
[Options]
variable: firstRole=hor
variable: secondRole=sec
```graphql
mutation CreateRole {
  firstRole: createRole(name: "{{firstRole}}") {
    id
    name
  }
  
  secondRole: createRole(name: "{{secondRole}}") {
    id
    name
  }
}
```
HTTP 200
[Captures]
roleTestID: jsonpath "$.data.firstRole.id"
secondRoleTestID: jsonpath "$.data.secondRole.id"
[Asserts]
jsonpath "$.data.firstRole.name" == "{{firstRole}}"

# ==========================================
# READ ONE
# ==========================================
POST http://localhost:8080/query
Authorization: Bearer {{token}}
```graphql
query ReadRole {
    role(id: {{roleTestID}}) {
        name
    }
}
```
HTTP 200
[Asserts]
jsonpath "$.data.role.name" == "{{firstRole}}"

# ==========================================
# READ MANY
# ==========================================
POST http://localhost:8080/query
Authorization: Bearer {{token}}
```graphql
query ReadRoles {
    roles {
        name
    }
}
```
HTTP 200
[Asserts]
jsonpath "$.data.roles[*].name" contains "{{firstRole}}"

# ==========================================
# UPDATE
# ==========================================
POST http://localhost:8080/query
Authorization: Bearer {{token}}
[Options]
variable: newRoleName=petok
```graphql
mutation UpdateRole {
  updateRole(id:{{roleTestID}}, name: "{{newRoleName}}") {
    id
    name
  }
}
```
HTTP 200
[Asserts]
jsonpath "$.data.updateRole.name" == "{{newRoleName}}"

# ==========================================
# DELETE ONE
# ==========================================
POST http://localhost:8080/query
Authorization: Bearer {{token}}
```graphql
mutation DeleteRole {
    firstRole: deleteRole(id: {{roleTestID}})
    secondRole: deleteRole(id: {{secondRoleTestID}})
}
```
HTTP 200
[Asserts]
jsonpath "$.data.firstRole" == "Role with ID {{roleTestID}} successfully deleted"
jsonpath "$.data.secondRole" == "Role with ID {{secondRoleTestID}} successfully deleted"