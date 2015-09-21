package main

import(
    "fmt"
    "reflect"
    "bytes"
    "errors"
    "strings"
)

var lookupMorse = map[rune]string {
    'a': ".-",
    'b': "-...",
    'c': "-.-.",
    'd': "-..",
    'e': ".",
    'f': "..-.",
    'g': "--.",
    'h': "....",
    'i': "..",
    'j': ".---",
    'k': "-.-",
    'l': ".-..",
    'm': "--",
    'n': "-.",
    'o': "---",
    'p': ".--.",
    'q': "--.-",
    'r': ".-.",
    's': "...",
    't': "-",
    'u': "..-",
    'v': "...-",
    'w': ".--",
    'x': "-..-",
    'y': "-.--",
    'z': "--..",
    ' ': "    ",
}


func MorseMarshal(m Messages) ([]byte, error) {
    var b bytes.Buffer

    mType := reflect.TypeOf(m)
    mValue := reflect.ValueOf(&m).Elem()
    for i := 0; i < mType.NumField(); i++ {
        mTypeField := mType.Field(i)
        mValueField := mValue.Field(i)

        tags := DecodeTag(mTypeField.Tag, "morse")

        _, omit := tags["-"]
        if(omit) {
            continue
        }

        if(mValueField.Kind() != reflect.String) {
            continue
        }

        fmt.Fprintf(&b, "%s: ", mTypeField.Name)

        for _, c := range mValueField.String() {
            morse, ok := lookupMorse[c]
            if !ok {
                return []byte{}, errors.New("Unexpected character")
            }

            _, shout := tags["shout"]
            if(shout) {
                morse = strings.Replace(morse, ".", "•", -1)
                morse = strings.Replace(morse, "-", "=", -1)
            }

            fmt.Fprintf(&b, "%s ", morse)
        }

        fmt.Fprintf(&b, "\n\n")
    }

    return b.Bytes(), nil
}
