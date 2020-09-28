package linggo

import (
	"fmt"
	"log"
	"testing"
)

func Test_Tr(t *testing.T)  {
	id := "halo dunia"
	en := "hello world"

	// tes bahasa indonesia
	message, err := Tr("id", "hello world")
	errorHandler(err)

	if message != id {
		log.Fatal(fmt.Sprintf("hasil yang diharapkan \"halo dunia\", tapi hasil yang didapatkan: %v", message))
	}

	// English test
	message, err = Tr("en", "hello world")
	errorHandler(err)

	if message != en {
		log.Fatal(fmt.Sprintf("the expected result \"hello world\", but the result got: %v", message))
	}
}