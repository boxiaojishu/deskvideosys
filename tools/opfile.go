package tools
import (
    "log"
    "os"
)

func CreateEmtpyfile(filename string) {
    file, err := os.Create(filename)
    if err != nil {
        log.Fatal("create file err", err)
    }
    log.Println(file)
    file.Close()
}

func Removefile(filename string){
    err := os.Remove(filename)
    if err != nil {
        log.Fatal(err)
    }
}

