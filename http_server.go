package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Gelen isteğin IP adresini al
	ip := strings.Split(r.RemoteAddr, ":")[0]

	// Kullanılan işletim sistemini al
	userAgent := r.UserAgent()
	var os string
	if strings.Contains(userAgent, "Windows") {
		os = "Windows"
	} else if strings.Contains(userAgent, "Linux") {
		os = "Linux"
	} else if strings.Contains(userAgent, "Macintosh") {
		os = "MacOS"
	} else {
		os = "Diğer"
	}

	// İstek bilgilerini dosyaya yaz
	writeToLog(fmt.Sprintf("IP: %s, İşletim Sistemi: %s", ip, os))

	// HTML içeriğini oluştur ve gönder
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Merhaba</title>
			<link rel="stylesheet" type="text/css" href="/styles.css">
		</head>
		<body>
			<div class="container">
				<h1>Merhaba, Dünya!</h1>
				<p>IP Adresiniz: %s</p>
				<p>Kullanılan İşletim Sistemi: %s</p>
			</div>
		</body>
		</html>
	`, ip, os)
}

func stylesHandler(w http.ResponseWriter, r *http.Request) {
	css := `
		body {
			font-family: Arial, sans-serif;
			background-color: #f0f0f0;
		}
		.container {
			width: 50%;
			margin: 50px auto;
			text-align: center;
		}
		h1 {
			color: #333;
		}
	`
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprintf(w, css)
}

func writeToLog(message string) {
	file, err := os.OpenFile("demo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Log dosyası açılamadı: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(message)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/styles.css", stylesHandler)

	fmt.Println("Servis çalışıyor...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Sunucu hatası: %s", err)
	}
}
