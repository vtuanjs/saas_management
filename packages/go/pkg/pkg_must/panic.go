package pkgmust

import (
	"fmt"
)

func Check(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("message: %s, error: %v", msg, err))
	}
}
