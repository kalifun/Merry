package main

import (
	"fmt"
	"github.com/kalifun/common"
)

func main()  {
	fmt.Sprintf(common.Logo, "\x1b[35m", "\x1b[1m",
		"\x1b[0m", "\x1b[32m", "\x1b[1m", "\x1b[0m")
}