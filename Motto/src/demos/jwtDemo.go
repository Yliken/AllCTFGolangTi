package main

import (
	"SummerVactionSQL/function"
	"fmt"
)

func main() {
	jwt, err := function.EnJwt("admin", "Yliken")
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
	deJwt, err := function.DeJwt(jwt)
	if err != nil {
		panic(err)
	}

	fmt.Println((*deJwt)["username"].(string))
}
