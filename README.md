---
marp: true
---

# go-json-bson-tutorial

---

# Omitempty State
## `common` case

```go
type User struct {
	Firstname	string   `json:"firstname,omitempty"`
	IsActive  bool    `json:"is_active,omitempty"`
}
```

| Field     | Value                                | Actual Result        |
|-----------|--------------------------------------|----------------------|
| Firstname | ``` &User{} ```                     | {}                   |
| Firstname | ``` &User {Firstname: ""} ``` | {} |
| Firstname | ``` &User {Firstname: "golf"} ``` | {"firstname": "golf"} |
| Firstname | ``` &User{Firstname: "golf", IsActive:  true} ``` | {"firstname": "golf", "is_active": true} |
| **Firstname** | ``` &User{Firstname: "golf", IsActive:  false} ``` | **{"firstname": "golf"}** |

---

## Simulate User Request Payload

---

### User: Checked IsActive
```
&User{
	Firstname: "golf",
	IsActive:  true,
}

// output
{"firstname":"golf","IsActive":true}
```

---

### User: Unchecked IsActive
In this case you will see that your field `is_active` is missing. Since omitempty will ignore you field once reach zero-value.

> A value of bool type has 2 possible values: false and true. And you want to "communicate" 3 different states with a bool field, namely to not update the field, to set the field to false and to set the field to true. This is obviously not possible.


```
&User{
	Firstname: "golf",
	IsActive:  false,
}

// output
{"firstname":"golf"}
```

---

### To solve above case
**Add pointer to your type**

```
type User struct {
	Firstname string   `json:"firstname,omitempty"`
	IsActive  *bool    `json:"is_active,omitempty"`
}
```

---

**`true` case**

```
isActive := true
&User{
	Firstname: "golf",
	IsActive:  &isActive,
}

// output
{"firstname":"golf","is_active":true}
```

---

**`false` case**

```
isActive := false
&User{
	Firstname: "golf",
	IsActive:  &isActive,
}

// output
{"firstname":"golf","is_active":false}
```

---

**`zero value` case**

hit `omitempty`

```
&User{
	Firstname: "golf",
}

// output
{"firstname":"golf"}
```

---

## omitempty tag
The purpose of `omitempty` tag is ignoring empty field.

## bson
You need to understand `bson` type cleary form https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/
And also learn more once apply to mongodb repository https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial-part-1-connecting-using-bson-and-crud-operations