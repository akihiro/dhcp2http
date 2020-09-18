package main

import (
	dhcp "github.com/krolaw/dhcp4"
	"log"
)

type Handler struct {
}

func (h *Handler) ServeDHCP(req dhcp.Packet, msgType dhcp.MessageType, options dhcp.Options) dhcp.Packet {
	log.Println(req, msgType, options)
	return dhcp.Packet{}
}
