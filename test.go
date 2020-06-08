package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

//import (
//    log "github.com/sirupsen/logrus"
//)
//
//func main() {
//    log.WithFields(log.Fields{
//	"animal": "walrus",
//    }).Info("A walrus appears")
//}

var currentTime = time.Now()
type User struct {
    Time time.Time
    Name string
    Age int
}

func main() {
    for {
	standard_danhang()
	fmt.Printf("\n")
	standard_duohang()
	fmt.Printf("\n")
	standard_json()
	fmt.Printf("\n")
	standard_fengefu()
        fmt.Printf("\n")
	time.Sleep(5 * time.Second)
    }

}
//func ran(a int)int {
//    rand.Seed(time.Now().UnixNano())
//    x := rand.Intn(a)
//    return x
//}

func standard_danhang() {

    fmt.Printf("%s%s",currentTime,"----------danhang----------aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
}

func standard_duohang() {

    fmt.Printf("%s%s",currentTime,"\n", "-----------duohang---------")
}

func standard_json() {
    user := User{
        Time: currentTime,
	Name: "Tab",
	Age: 18,
    }
    data, err := json.Marshal(user)
    if err != nil {
	log.Fatal(err)
    }
    fmt.Println(string(data))
}

func standard_fengefu() {
    fmt.Printf("%s%s%s", "10.10.1.1 ::: ",currentTime, "::: GET /online/sample HTTP/1.1")
    
}
