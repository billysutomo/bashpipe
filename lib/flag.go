package lib

import (
	"flag"
	"fmt"
)

type Opts struct {
	Port string
}

func Flag() Opts {
	var port = flag.Int("port", 8080, "set custom port, default 8080")
	flag.Parse()
	return Opts{Port: fmt.Sprintf("%d", *port)}
}
