package main

import (
	"fmt"

	"github.com/NatthanonPPP/Golang/phim"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello, world!2")
	fmt.Println("id :", id)
	phim.HelloPhim()
}
