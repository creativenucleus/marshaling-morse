package main;

import(
    "log"
    "fmt"
)

type Test struct {
    Mystr string ``
    Mystr2 string `morse:"-"`
}

func main() {
    test := Test{
        Mystr: "hello world",
        Mystr2: "dwedlloe3h",
    }

    log.Printf("Marshalling: %+v", test)

    marshalled, err := MorseMarshal(test)
    if err != nil {
        log.Printf("Unable to marshal :(")
    } else {
        fmt.Printf("Marshalled: %+v", string(marshalled[:]))
    }
}
