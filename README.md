# goobserver

goobserver provides helper methods to implement observer pattern in go.

### Subscribe to topic(s)

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

func MyFunc1(payload []byte) {
    fmt.println(string(payload))
}

func MyFunc2(payload []byte) {
    fmt.println(string(payload))
}

func main() {
    err := GetObserver().Subscribe("my-topic-1", Myfunc1)
    if err != nil {
        fmt.println(err)
    }
    
    err := GetObserver().Subscribe("my-topic-2", Myfunc2)
    if err != nil {
        fmt.println(err)
    }
}
```

### Publish to topic(s)

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

err := GetObserver().Publish("my-topic-1", bytes)
if err != nil {
    fmt.println(err)
}
```
