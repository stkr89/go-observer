# go-observer

go-observer provides helper methods to implement observer pattern in go.

### Subscribe to topic

```go
import "github.com/stkr89/liger"

func MyFunc(payload []byte) {
    fmt.println(string(payload))
}

err := liger.Subscribe("my-topic", Myfunc)
if err != nil {
    fmt.println(err)
}
```

### Publish to topic

```go
import "github.com/stkr89/liger"

type User struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
}

bytes, _ := json.Marshal(User{
    FirstName: "Foo",
    LastName:  "Bar",
    Email: "foo.bar@email.com"
})

err := liger.Publish("my-topic", bytes)
```
