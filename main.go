package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/matheusr42/go-yeelight/httpu"
	"github.com/matheusr42/go-yeelight/ssdp"
	"github.com/reiver/go-telnet"
)

const (
	maxWaitSeconds int = 2
	numSends       int = 1
)

type Device struct {
	Stn      string
	Usn      string
	Location *url.URL
}

func main() {
	client, err := httpu.NewHTTPUClient()
	if err != nil {
		fmt.Print("Error client")
		fmt.Errorf(err.Error())
		return
	}
	defer client.Close()

	responses, errSsdp := ssdp.SSDPRawSearch(client, "wifi_bulb", maxWaitSeconds, numSends)

	if errSsdp != nil {
		fmt.Print("Error ssdp")
		fmt.Errorf(errSsdp.Error())
		return
	}

	devices := make([]Device, 0, len(responses))
	for _, resp := range responses {
		// loc := resp.Header.Get("Location")
		loc, err := resp.Location()

		if err != nil {
			fmt.Print("Error Location")
			fmt.Errorf(errSsdp.Error())
			return
		}

		device := Device{
			Location: loc,
		}
		devices = append(devices, device)
	}

	if len(devices) > 0 {
		conn, _ := telnet.DialTo(devices[0].Location.Host)
		r := 255
		g := 0
		b := 0

		rgb := (r * 65536) + (g * 256) + b

		command := "{ \"id\": 1, \"method\": \"set_power\", \"params\":[\"on\", \"smooth\", 500]}"
		conn.Write([]byte(command))
		conn.Write([]byte("\r\n"))

		time.Sleep(3 * time.Second)

		command = fmt.Sprintf("{\"id\":1,\"method\":\"set_rgb\",\"params\":[%d, \"smooth\", 500]}", rgb)
		fmt.Println(command)
		conn.Write([]byte(command))
		conn.Write([]byte("\r\n"))
	}
}
