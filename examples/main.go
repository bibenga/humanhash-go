package main

import (
	"fmt"

	humanhash "github.com/bibenga/humanhash-go"
)

func main() {
	compressed, err := humanhash.Humanize([]byte{96, 173, 141, 13, 135, 27, 96, 149, 128, 130, 151, 32})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Humanize: compressed=%v\n", compressed)

	uuid, compressed, err := humanhash.NewUuid()
	if err != nil {
		panic(err)
	}
	fmt.Printf("NewUuid: uuid=%v; compressed=%s\n", uuid, compressed)
}
