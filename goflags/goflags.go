
package main


import (
    "os"
    "fmt"
    "flag"
)

func main() {

    fmt.Println("arguments ", os.Args)

    fmt.Printf("arguments %d\n", len(os.Args))


    //declare all flags
    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")


    // parse the command line args after declaring them
    flag.Parse()

    // need to dereference the points with e.g. `*wordPtr`
    // to get the actual option values.
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("positional args:", flag.Args())
}
