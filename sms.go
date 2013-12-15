package govhBck

import (
	"fmt"
)

func SmsHandler(cmd Cmd) (err error) {
	fmt.Println(cmd)
	return
}
