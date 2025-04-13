package main

import (
	"html/template"
	"net/http"
)

func handleLogs(w http.ResponseWriter, r *http.Request) {
	logsMutex.Lock()
	defer logsMutex.Unlock()

	tmpl := template.Must(template.ParseFiles("templates/logs.html"))

	data := struct {
		Logs        []string
		StopRefresh bool
	}{
		Logs:        logs,
		StopRefresh: stopRefresh,
	}

	tmpl.Execute(w, data)
}

func handleClearLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		logsMutex.Lock()
		defer logsMutex.Unlock()
		logs = []string{}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleStopRefresh(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		stopRefresh = true
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
