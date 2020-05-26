package main

/**
生成UUID
*/

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	uuid, _ := uuid.NewV4()
	fmt.Println(uuid.String() + "cc")
}
