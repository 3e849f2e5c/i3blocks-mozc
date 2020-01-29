package main

import (
	"fmt"
	"github.com/godbus/dbus"
	"os"
)

var fullText = os.Getenv("MOZC_DISPLAY_MODE") == "1"
var romaji = os.Getenv("MOZC_DISPLAY_ROMAJI") == "1"

func getStatus() string {
	var connection, err = dbus.ConnectSessionBus()
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
	if len(call.Body) >= 2 {
		var body = ""
		if fullText {
			body = call.Body[1].(string)
		} else {
			body = call.Body[0].(string)
		}
		if body != "" {
			return format(body)
		} else {
			return "Error"
		}
	} else {
		if fullText {
			return "ひらがな"
		} else {
			return "あ"
		}
	}
}

func main() {
	var oldMode = getStatus()
	fmt.Println(oldMode)

	var connection, err = dbus.ConnectSessionBus()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer connection.Close()

	const monitor = "org.freedesktop.DBus.Monitoring.BecomeMonitor"
	const rule = "type='signal',sender='org.fcitx.Fcitx',interface='org.fcitx.Fcitx.Status'"

	var call = connection.BusObject().Call(monitor, 0, []string{rule}, uint(0))

	if call.Err != nil {
		fmt.Println(call.Err.Error())
		os.Exit(1)
	}

	var messages = make(chan *dbus.Message)
	connection.Eavesdrop(messages)

	for v := range messages {
		if len(v.Body) >= 3 {
			var body = ""
			if fullText {
				body = v.Body[2].(string)
			} else {
				body = v.Body[1].(string)
			}
			if oldMode != body {
				fmt.Println(format(body))
				oldMode = body
			}
		}
	}
}

func format(input string) string {
	if fullText {
		if romaji {
			return input
		}
		switch input {
		case "Direct":
			return "ローマ字"
		case "Hiragana":
			return "ひらがな"
		case "Full Katakana":
			return "全カタカナ"
		case "Half Katakana":
			return "半カタカナ"
		}
		return input
	} else {
		switch input {
		case "A":
			return "Ａ"
		case "ｱ":
			return "ｱ\u2423"
		default:
			return input
		}
	}
}
