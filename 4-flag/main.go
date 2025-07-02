package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put ur DB host")
	var port *int = flag.Int("port", 0, "Put ur DB port")
	var username *string = flag.String("username", "root", "Put ur DB username")
	var password *string = flag.String("password", "root", "Put ur DB username")
	flag.Parse()

	fmt.Println(*host, *port, *username, *password)
}