package main

import (
    "fmt"
    "flag"

    "github.com/wxnacy/wmovie"
)

func main() {
    fmt.Println("Hello World")
    flag.Parse()
    args := flag.Args()
    url := args[0]


    items, _ := wmovie.ParseMMItems(url)
    fmt.Println(items)


    for _, d := range items {
        item_url := wmovie.ParseMMItem(d["url"])
        fmt.Println(item_url)
    }

}
