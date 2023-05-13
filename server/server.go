package server

import (
	"fmt"

	"github.com/tobbstr-examples/client-server-monorepo/shared"
)

func Endpoint(p shared.Person) {
	fmt.Printf("Person is %s and is %d years old\n", p.Name, p.Age)
}
