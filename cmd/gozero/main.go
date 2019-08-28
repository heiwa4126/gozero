package main

import (
	"fmt"
	"github.com/heiwa4126/gozero"
)

func main() {
	fmt.Printf("version=%s\n", gozero.Version)
	fmt.Println("世界の皆さん、こんにちは!")

	w1 := gozero.Word{English: "Apple", Japanese: "りんご"}
	fmt.Printf("% +v\n", w1)

	fmt.Println(gozero.Add(29, 13)) //-> 42
}
