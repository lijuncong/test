package main


import (
//    "encoding/json"
    "log"
    "os"
    "time"
    "fmt"
)

type User struct {
       Time time.Time
       Name string
       Age int
    }
var currentTime = time.Now()

func main() {
    file := "./" + time.Now().Format("20180103") + ".log"

    logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
    if nil != err {
	panic(err)
    }

    //创建一个Logger
    //参数1：日志写入目的地
    //参数2：每条日志的前缀
    //参数3：日志属性
    //loger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
    loger := log.New(logFile, "",0 )

    //loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    //loger.SetFlags("")
 
    for {

    //fmt.Println("aaaa----单行danhang----aaaa")
    //loger.Println("aaaa----单行danhang----aaaa")

    //fmt.Printf("%s%s%s","bbbb----多行duohang----\n", "---------------------\n","----多行duohang----bbbb\n")
    //loger.Printf("%s%s%s","bbbb----多行duohang----\n", "---------------------\n","----多行duohang----bbbb")

    fmt.Printf("%s%s%s", "10.10.1.1 :::", currentTime, "::: GET /online/sample HTTP/1.1\n")
    loger.Printf("%s%s%s", "10.10.1.1 :::", currentTime, "::: GET /online/sample HTTP/1.1")


    //user := User{

    //    	Name: "Tab",
    //    	Age: 18,
    //       }
    //       data, err := json.Marshal(user)
    //       if err != nil {
    //    	log.Fatal(err)
    //       }
    //fmt.Printf("%s%s\n", data)
    //loger.Printf("%s",data)
   
    time.Sleep(5 * time.Second)
    }
}
