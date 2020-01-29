package main

import (
	"fmt"
	"github.com/godbus/dbus"
	"os"
)

func main() {
	var connection, err = dbus.ConnectSessionBus()
	if err != nil {
		fmt.Println("Bruh " + err.Error())
		os.Exit(1)
	}

	defer connection.Close()

	const monitor = "org.freedesktop.DBus.Monitoring.BecomeMonitor"
	const rule = "type='signal',sender='org.fcitx.Fcitx',interface='org.fcitx.Fcitx.Status'"

	var call = connection.BusObject().Call(monitor, 0, []string{rule}, uint(0))

	if call.Err != nil {
		fmt.Println("Bruh " + call.Err.Error())
		os.Exit(1)
	}

	var messages = make(chan *dbus.Message)
	connection.Eavesdrop(messages)

	var oldMode = "あ"

	fmt.Println(oldMode)
	for v := range messages {
		if len(v.Body) >= 3 {
			var newMode = v.Body[1].(string)
			if oldMode != newMode {
				switch newMode {
					case "A":
						fmt.Println("Ａ")
						break
					case "ｱ":
						fmt.Println("ｱ\u2423")
						break
					default:
						fmt.Println(newMode)
				}
				oldMode = newMode
			}
		}
	}
}
