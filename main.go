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
		fmt.Println("Error client")
		fmt.Errorf(err.Error())
		return
	}
	defer client.Close()

	responses, errSsdp := ssdp.SSDPRawSearch(client, "wifi_bulb", maxWaitSeconds, numSends)
	if errSsdp != nil {
		fmt.Println("Error ssdp")
		fmt.Errorf(errSsdp.Error())
		return
	}

	devices := make([]Device, 0, len(responses))
	for _, resp := range responses {
		loc, err := resp.Location()
		if err != nil {
			fmt.Println("Error Location")
			fmt.Errorf(errSsdp.Error())
			return
		}

		device := Device{
			Location: loc,
		}
		devices = append(devices, device)
	}

	commands := []string{}

	if len(devices) > 0 {
		host := devices[0].Location.Host
		fmt.Printf("connecting to %s\n", host)

		conn, err := telnet.DialTo(host)
		if err != nil {
			fmt.Println("Error connecting to device")
			fmt.Println(err.Error())
			return
		}
		defer conn.Close()
		r := 255
		g := 255
		b := 255

		rgb := (r * 65536) + (g * 256) + b

		command := "{\"id\": 1, \"method\": \"set_power\", \"params\":[\"on\", \"smooth\", 500]}"
		commands = append(commands, command)

		command = fmt.Sprintf("{\"id\": 2,\"method\": \"set_rgb\",\"params\" :[%d, \"smooth\", 500]}", rgb)
		commands = append(commands, command)

		for _, command := range commands {
			time.Sleep(3 * time.Second)
			_, err := conn.Write([]byte(command + "\r\n"))
			if err != nil {
				fmt.Println("Error sending command")
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("Sent command: %s\n", command)
		}
	}
}
