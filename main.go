package main

import (
	"fmt"
	"log"
	"net/http"
)

func getGreeting() string {
	return "Welcome to EFD Go project!"
}

func getPage() string {
	greeting := getGreeting()
	return "<!DOCTYPE html>" +
		"<html lang=\"en\">" +
		"<head><meta charset=\"utf-8\" /><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" /><title>EFD Go Welcome</title>" +
		"<style>:root{--bg:#1f0d0d;--panel:#ffffff;--accent:#ef4444;--text:#0f172a;--muted:#64748b;}*{box-sizing:border-box;}body{margin:0;font-family:'Segoe UI',Arial,sans-serif;background:linear-gradient(135deg,var(--bg),#3b1111);color:var(--text);min-height:100vh;display:grid;place-items:center;padding:24px;}" +
		".card{background:var(--panel);border-radius:20px;padding:2.5rem 3rem;box-shadow:0 20px 50px rgba(0,0,0,0.25);text-align:center;max-width:560px;width:100%;}" +
		".badge{display:inline-block;padding:0.4rem 0.8rem;border-radius:999px;background:rgba(239,68,68,0.12);color:var(--accent);font-weight:700;text-transform:uppercase;letter-spacing:0.08em;font-size:0.8rem;margin-bottom:1rem;}" +
		"h1{margin:0 0 0.75rem;font-size:2rem;color:var(--accent);}p{margin:0;color:var(--muted);line-height:1.6;}</style></head>" +
		"<body><div class=\"card\"><div class=\"badge\">EFD • Go</div><h1>" + greeting + "</h1><p>This polished welcome page is served from the Go project and is ready to be viewed in a browser.</p></div></body></html>"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, getPage())
	})

	fmt.Println("Server started at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
