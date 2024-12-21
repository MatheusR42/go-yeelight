# Go Yeelight

A GO package that implements Yeelight WiFi Light Inter-Operation Specification. 

This code will find Yeelight lamp connected in LAN using Simple Service Discovery Protocol (SSDP), and send commands via Telnet Protocol.

## How to use

 1. Plug in the yeelight smart bulb
 2. Turn on the light switch
 3. Run `go run main.go`

By default it should turn on the bulb and change the collor to white. 

You can change the commands in main.go file following the specifications of the section `4.1 COMMAND message` of the [Yeelight_Inter-Operation_Spec.pdf](./Yeelight_Inter-Operation_Spec.pdf)


### Credits

This package is heavily based on [goupnp SSDP implementation](https://github.com/huin/goupnp). Thanks guys :)

### References:

- https://github.com/huin/goupnp
- https://github.com/4thel00z/upnpctl/blob/8eb0d27fa37067c2a600ecab9d1d145f157bb2ff/cmd/scan.go
- https://ops.tips/blog/udp-client-and-server-in-go/
- https://github.com/kbinani/screenshot
- https://stackoverflow.com/questions/58735783/how-to-get-the-dominant-color-in-an-image-with-go-and-imagemagick
- https://github.com/RobCherry/vibrant
- https://www.yeelight.com/en_US/developer
- https://www.yeelight.com/download/Yeelight_Inter-Operation_Spec.pdf
- https://stackoverflow.com/a/49236233/5786900
