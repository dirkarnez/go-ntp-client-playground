// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time)
}
