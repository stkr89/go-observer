# goobserver

goobserver provides helper methods to implement observer pattern in go.

### Subscribe to topic

```go
import (
    "github.com/stkr89/goobserver"
)

var observer *goobserver.GoObserver

func GetObserver() *goobserver.GoObserver {
    if observer == nil {
        observer = goobserver.NewGoObserver()
    }

    return observer
}

func MyFunc(payload []byte) {
    fmt.println(string(payload))
}

err := GetObserver().Subscribe("my-topic", Myfunc)
if err != nil {
    fmt.println(err)
}
```

### Publish to topic

```go
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

err := GetObserver().Publish("my-topic", bytes)
```
