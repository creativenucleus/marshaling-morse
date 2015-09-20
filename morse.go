package main;

import(
    "fmt"
    "reflect"
    "bytes"
    "errors"
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
    ' ': "       ",
}


func MorseMarshal(t Test) ([]byte, error) {
    var b bytes.Buffer

    tt := reflect.TypeOf(t)
    vt := reflect.ValueOf(&t).Elem()
    for i := 0; i < tt.NumField(); i++ {
        ttField := tt.Field(i)
        vtField := vt.Field(i)

        parsedTag, _ /*options*/ := parseTag(ttField.Tag.Get("morse"))
        if(parsedTag == "-") {
            continue
        }

//        log.Printf("P: %+v", parsedTag)
//        log.Printf("O %+v", options)

        if(vtField.Kind() != reflect.String) {
            continue
        }

        fmt.Fprintf(&b, "%s: ", ttField.Name)

        for _, c := range vtField.Interface().(string) {
            morse, ok := lookupMorse[c]
            if !ok {
                return []byte{}, errors.New("Unexpected character")
            }

            fmt.Fprintf(&b, "%s ", morse)
        }

        fmt.Fprintf(&b, "\n")
    }

    return b.Bytes(), nil
}
