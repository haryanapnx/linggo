# Linggo v1.1.0

Ini adalah package bahasa pemrograman go sederhana untuk mengkonfigurasikan tampilan 
pesan ke dalam multi bahasa.

### func Tr()
```go
func Tr(file string, key string) (message interface{}, err error)
```
### func Set()
```go
func Set(file string) (jsonFile JFile)
```
### func Ts()
```go
func (f JFile) Ts(key string) (message interface{})
```

### Instalasi
`go get github.com/michaelwp/linggo` atau
`go get github.com/michaelwp/linggo@v1.1.0`

### Contoh penggunaan
A. Buat Json file untuk menampung bahasa yang digunakan
- id.json (bahasa Indonesia)
```json 
{
  "hello world": "halo dunia",
  "welcome": "selamat datang",
  "please": "silahkan"
}
```
- en.json (bahasa Inggris)
```json 
{
  "hello world": "hello world"
}
```
- fr.json (bahasa Prancis)
```json 
{
  "hello world": "Bonjour le monde"
}
```
- en.json (bahasa China)
```json 
{
  "hello world": "你好，世界"
}
```
B. memanggil file json setiap kali proses translate

```go 
package main

import (
    "github.com/michaelwp/linggo"  // import linggo package
    "log"
)

func main(){
    // contoh translate ke dalam bahasa China
    msg, err := linggo.Tr("ch", "hello world")
    // handle error (disarankan untuk menghandle error dengan cara yang lebih baik)
    if err != nil {
        log.Fatal(err)
    }
    // menampilkan hasil ke layar
    log.Println(msg)
}
```
#### Hasil
```text
你好，世界
```

C. memanggil file json hanya sekali diawal.

```go 
package main

import (
    "github.com/michaelwp/linggo"  // import linggo package
    "log"
)

// load json file (bahasa indonesia/ id.json)
var f = linggo.Set("id")

func main() {
    welcome := f.Ts("welcome")
    please := f.Ts("please")

    log.Println(welcome)
    log.Println(please)
}
```
#### Hasil
```text
selamat datang
silahkan
```