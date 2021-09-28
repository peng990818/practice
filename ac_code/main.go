package main

import (
    "fmt"
    "practice/ac_code/ac"
)

type Address struct {
    name  string
    state string
}
type Person struct {
    name string
    age  int
    Address
}

func main() {
/*outer:
    for i := 0; i < 3; i++ {
        for j := 1; j < 4; j++ {
            fmt.Printf("i = %d , j = %d\n", i, j)
            if i == j {
                break outer
            }
        }

    }*/

    /*welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)*/
    w := []string{"asd","abc","a","bd"}
    trie := ac.NewTrie(false)
    ac := ac.NewAC(trie)
    ac.Trie.Add(w...)
    ac.BuildFailurePath()
    res := ac.Find("asdaaaabdaaabc")
    fmt.Println(res)
}

func change(s ...string) {
    s[0] = "Go"
    s = append(s, "playground")
    fmt.Println(s)
}