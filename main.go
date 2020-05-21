package main

import (
	"fmt"
	"github.com/egawata/mockery_test/service"
)

func main() {
	s := service.NecoService{}

	h := s.GetNeco(123)
	fmt.Printf("Neco = %#v\n", h)
}
