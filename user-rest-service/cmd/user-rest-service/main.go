package main

import (
	"fmt"
	"os"

	"github.com/paypay3/kakeibo-rest-api-ddd/user-rest-service/infrastructure/router"
)

func main() {
	if err := router.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
