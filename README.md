# 👓 **GENV**

**GE**neric e**NV** is a zero dependency package for effortless interaction with env variables.

## 📦 Install

```bash
go get github.com/abecodes/genv
```

## 💻 Usage

**GENV** provides a `Get` and a `GetWithDefault` function.

Using them is pretty straight forward, but there are two constraints:

- If the value for a given env variable does not match the requested type, the \*requested types **default value\*** will be returned
- If a value of `time.Time` is requested, the string in the env needs to be in RFC1123 format ("Thu, 30 May 2024 21:01:37 GMT")

**GENV** supports all basic types as well as `time.Time` (RFC1123) and `time.Duration`.

### Get

```go
package main

import (
    "fmt"

    "github.com/abecodes/genv"
)

func main() {
    // This will retrieve the value set for env var "HOME" and return the value as string
    home := genv.Get[string]("HOME")

    // prints: "(string) /your/home/path"
    fmt.Printf("(%T) %v\n", home, home)

    // This will retrieve the value set for env var "KITTY_PID" and return the value as int
    pid := genv.Get[int]("KITTY_PID")

    // prints: "(int) 1234"
    fmt.Printf("(%T) %v\n", pid, pid)

    // Since everything in the env is a string, we can return the value for "KITTY_PID" as string as well
    pidStr := genv.Get[string]("KITTY_PID")

    // prints: "(string) 1234"
    fmt.Printf("(%T) %v\n", pidStr, pidStr)

    // This does not work the other way around. If we try to turn a `string` type into an `int` type,
    // we will just receive the `int` type default value
    homeInt := genv.Get[int]("HOME")

    // prints: "(int) 0"
    fmt.Printf("(%T) %v\n", homeInt, homeInt)
}
```

### GetWithDefault

```go
package main

import (
    "fmt"

    "github.com/abecodes/genv"
)

func main() {
    // Since no value exists for this key, the default will be returned
    str := genv.GetWithDefault[string]("DOES_NOT_EXIST", "hello")

    // prints: "(string) hello"
    fmt.Printf("(%T) %v\n", str, str)
}
```
