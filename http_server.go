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
	ipNumber := strings.Split(r.RemoteAddr, ":")

	// Port bilgisi
    var port string
    if len(ipNumber) > 1 {
        port = ipNumber[1]
    } else {
        port = "Belirtilmedi"
    }


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
	writeToLog(fmt.Sprintf("IP: %s,Port: %s,Sistem: %s", ipNumber[0], port, os))

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
				<h1>Golang, for hackers!</h1>
				<p>IP Adresi: %s</p>
				<p>İşletim Sistemi: %s</p>
				<p>Port Numarası: %s</p>
			</div>
		</body>
		</html>
	`, ipNumber[0], os,port)
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
