package main

import (
    "github.com/ajones239/gkademlia"
    "fmt"
)

func main() {
    var t node.Node
    t.PeerCount = 6
    fmt.Println(t.PeerCount)
}
