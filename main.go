package main

import (
	"flag"
	"log"
	"net"
	"net/url"
	"os"
	"strings"

	dhcp "github.com/krolaw/dhcp4"
)

var (
	listenAddr = flag.String("listen", ":67", "Listen address")
	backend    = flag.String("backend", "http://localhost/", "Backend URL")
)

func main() {
	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
			if err := f.Value.Set(s); err != nil {
				log.Fatal(err)
			}
		}
	})
	flag.Parse()

	backendURL, err := url.Parse(*backend)
	if err != nil {
		log.Fatalf("backend is invalid: %s", err)
	}
	switch backendURL.Scheme {
	case "http":
	case "https":
	default:
		log.Fatalf("backend schmea is invalid: %s", backendURL.Scheme)
	}
	if backendURL.Host == "" {
		log.Fatalln("backend host not found")
	}

	l, err := net.ListenPacket("udp4", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	h := &Handler{}
	log.Fatalln(dhcp.Serve(l, h))
}
