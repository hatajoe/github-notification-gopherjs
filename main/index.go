package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/gopherjs/gopherjs/js"
)

func readConfig(filePath string) error {
	fs := js.Global.Call("require", "fs")
	s := fs.Call("readFileSync", filePath, "utf-8").String()
	m := map[string]interface{}{}
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		return err
	}
	js.Global.Set("interval", m["interval"].(float64))
	js.Global.Set("token", m["token"].(string))
	js.Global.Set("apihostname", m["apihostname"].(string))
	js.Global.Set("githostname", m["githostname"].(string))
	return nil
}

func main() {
	js.Global.Call("require", "crash-reporter").Call("start")

	app := js.Global.Call("require", "app")
	tray := js.Global.Call("require", "tray")
	menu := js.Global.Call("require", "menu")

	_, filename, _, _ := runtime.Caller(1)
	dirName := path.Dir(filename)
	if err := readConfig(path.Join(app.Call("getPath", "home").String(), ".github-notification-gopherjs")); err != nil {
		log.Println(err)
		app.Call("quit")
	}

	app.Get("dock").Call("hide")

	app.Call("on", "window-all-closed", func() {
		if js.Global.Get("process").Get("platform").String() != "darwin" {
			log.Println("window all closed")
			app.Call("quit")
		}
	})

	app.Call("on", "ready", func() {
		browserWindow := js.Global.Call("require", "browser-window")
		mainWindow := browserWindow.New(map[string]interface{}{
			"show": false,
		})
		mainWindow.Call("loadUrl", fmt.Sprintf("file://%s/%s", dirName, "index.html"))
		appIcon := tray.New(path.Join(dirName, "icon.png"))
		appIcon.Call("setToolTip", "GitHub Notification GopherJS")
		appIcon.Call("setContextMenu", menu.Call("buildFromTemplate", []map[string]interface{}{
			map[string]interface{}{
				"label":       "Quit",
				"accelerator": "Command+Q",
				"click": func() {
					app.Call("quit")
				},
			},
		}))
		mainWindow.Call("on", "closed", func() {
			mainWindow = nil
		})
	})
}
