# Indonesian Time Format
Library ini dapat digunakan untuk memudahkan pemanggilan Format `time` 
khusus untuk indonesia. Jika ada yang kurang silahkan tambahkan untuk membantu kontribusi disini.
## How to Use ?
```go
    package main

    import (
        indotime "github.com/vaksi/indonesiandate"
    )
    
    func main() {
        // siapkan variabel time format
        t := time.Now()
        // case for DefaultFormatIDDateFull
        dateString := indonesiandate.New(t).Format(indonesiandate.FormatIDDateWithTimeFull)
        ....
    }   
```

