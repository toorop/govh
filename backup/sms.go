package govh

import (
	"fmt"
)

func SmsHandler(cmd Cmd) (err error) {
	fmt.Println(cmd)
	return
}
