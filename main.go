package main

import (
	"fmt"
	"net"

	"github.com/goakshit/uniq/config"
)

func main() {

	fmt.Println("Initialising...")
	conf := config.New()

	fmt.Printf("Booting up %s server on %s:%d\n", conf.ConnConfig.ConnType, conf.ConnConfig.ConnHost, conf.ConnConfig.ConnPort)
	l, err := net.Listen(conf.ConnConfig.ConnType,
		fmt.Sprintf("%s:%d", conf.ConnConfig.ConnHost, conf.ConnConfig.ConnPort))
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		panic(err)
	}
	defer l.Close()

	fmt.Println("Waiting for connections...")
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
	}
}
