package main

import (
	"fmt"
)

func main() {
	b := []byte("11AB398765UJ1A050200N23")
	mapa, err := separaDatos(b)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(mapa)
	}
}
