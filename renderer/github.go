package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	github := js.Global.Call("require", "octonode")
	remote := js.Global.Call("require", "remote")
	app := remote.Call("require", "app")
	shell := remote.Call("require", "shell")

	interval := time.Duration(remote.Call("getGlobal", "interval").Float())
	token := remote.Call("getGlobal", "token").String()

	desktopNotification := js.Global.Get("Notification")
	client := github.Call("client", token)
	ghme := client.Call("me")
	noticed := map[string]bool{}
	for {
		ghme.Call("notifications", map[string]interface{}{}, func(e *js.Object, d []map[string]interface{}, h *js.Object) {
			if e != nil {
				log.Println("error: ", e)
				app.Call("quit")
			}
			if len(d) > 0 {
				for _, v := range d {
					id := v["id"].(string)
					if _, ok := noticed[id]; !ok {
						noticed[id] = true
						subject := v["subject"].(map[string]interface{})
						n := desktopNotification.New(v["repository"].(map[string]interface{})["name"].(string), map[string]interface{}{
							"tag":  id,
							"body": fmt.Sprintf("%s\n%s", subject["type"].(string),subject["title"].(string)),
						})
						n.Set("onclick", func() {
							shell.Call("openExternal", "https://github.com/notifications")
						})
					}
				}
			}
		})
		time.Sleep(interval * time.Second)
	}
}

/*
var remote = require('remote');
var app = remote.require('app');
var shell = remote.require('shell');
var path = require('path');
var fs = require('fs');
var github = require('octonode');

if (Notification.permission != 'granted') {
	Notification.requestPermission(function (permission) {
  	if (permission != "granted") {
			app.quit();
		}
	})
}

var dataFilePath = path.join(app.getPath('home'), '.github-electron');
var data;
if (!fs.existsSync(dataFilePath)) {
	data = {};
}
data = JSON.parse(fs.readFileSync(dataFilePath, 'utf-8'));
var token = null
if ('token' in data) {
	token = data['token'];
}

var client, ghme;
if (token != null)  {
  client = github.client(token);
  ghme = client.me();
}

var notifications = []
var timer = setInterval(function() {
	if (ghme != undefined) {
 	  ghme.notifications({}, function(err, data, headers) {
 	 		if (data.length > 0) {
				data.forEach(function(e,i,a) {
					if (notifications.indexOf(e.id) >= 0) {
						return;
					}
					notifications.push(e.id);

					var notification = new Notification(e.subject.type, { tag: e.id,  body: e.subject.title });
					notification.onclick = function () {
						shell.openExternal("https://github.com/notifications")
					};
				})
 	   	}
		});
	}
}, 3000);
*/
