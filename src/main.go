package main

import (
	"fmt"
	"github.com/godbus/dbus"
	"os"
)

func getStatus () string {
	var connection, err= dbus.ConnectSessionBus()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer connection.Close()

	obj := connection.Object("org.fcitx.Fcitx", "/Status")
	call := obj.Call("org.fcitx.Fcitx.Status.Get", 0, "mozc-composition-mode")
	if call.Err != nil {
		fmt.Println(call.Err)
		os.Exit(1)
	}

	if len(call.Body) >= 1 && call.Body[0] != "" {
		return formatLetter(call.Body[0].(string))
	} else {
		return "あ"
	}
}

func main() {
	var connection, err= dbus.ConnectSessionBus()
	if err != nil {
		fmt.Println(err.Error())
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

	var oldMode = getStatus()

	fmt.Println(oldMode)
	for v := range messages {
		if len(v.Body) >= 3 {
			var newMode = v.Body[1].(string)
			if oldMode != newMode {
				var mode = formatLetter(newMode)
				fmt.Println(mode)
				oldMode = mode
			}
		}
	}
}

func formatLetter (letter string) string {
	switch letter {
	case "A":
		return "Ａ"
	case "ｱ":
		return "ｱ\u2423"
	default:
		return letter
	}
}