package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "%s\n", "helloworld! = unbuffered")

	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf, "%s\n", "helloworld! = buffered")
	buf.Flush()
}
