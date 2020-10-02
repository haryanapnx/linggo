# Linggo v1.1.1

This is a simple go library for configuring displays messages into multi languages.

### func Tr()
```go
func Tr(file string, key string) (message interface{}, err error)
```
### func Set()
```go
func Set(file string) (jsonFile JFile, err error)
```
### func Ts()
```go
func (f JFile) Ts(key string) (message interface{})
```

### Installation
`go get github.com/michaelwp/linggo` atau
`go get github.com/michaelwp/linggo@v1.1.1`

### Example
A. Create a Json file that contain the language used.
- id.json (Indonesia)
```json 
{
  "hello world": "halo dunia",
  "welcome": "selamat datang",
  "please": "silahkan"
}
```
- en.json (English)
```json 
{
  "hello world": "hello world"
}
```
- fr.json (French)
```json 
{
  "hello world": "Bonjour le monde"
}
```
- cn.json (Chinese)
```json 
{
  "hello world": "你好，世界"
}
```
B. Call the json file in the first linggo.Tr() parameter.

```go 
package main

import (
    "github.com/michaelwp/linggo"  // import linggo package
    "log"
)

func main(){
    // example translation into Chinese
    msg, err := linggo.Tr("ch", "hello world")
    // handle errors (recommended to handle errors in a better way)
    if err != nil {
        log.Fatal(err)
    }
    // log the result
    log.Println(msg)
}
```
#### Result
```text
你好，世界
```

C. Init json configuration globally.

```go 
package main

import (
    "github.com/michaelwp/linggo"  // import linggo package
    "log"
)

// load json file (Indonesia / id.json)
var f, _ = linggo.Set("id")

func main() {
    welcome := f.Ts("welcome")
    please := f.Ts("please")

    log.Println(welcome)
    log.Println(please)
}
```
#### Result
```text
selamat datang
silahkan
```