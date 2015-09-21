package main

import(
    "log"
    "fmt"
)

type Messages struct {
    FirstString string
    SecondString string `morse:"-"`
    ThirdString string `morse:"shout"`
}

func main() {
    messages := Messages{
        FirstString: "hello world",
        SecondString: "don't encode this!",
        ThirdString: "nice to meet you",
    }

    fmt.Printf("=> Marshalling...\n%+v\n\n", messages)

    marshalled, err := MorseMarshal(messages)
    if err != nil {
        log.Printf("XXX Unable to marshal :(\n")
    } else {
        pretty.Printf("=> Marshalled...\n%+v", string(marshalled[:]))
    }
}
