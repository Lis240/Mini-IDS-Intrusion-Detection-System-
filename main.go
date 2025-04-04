package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type IPActivity struct {
	count    int
	lastSeen time.Time
}

var ipLog = make(map[string]*IPActivity)
var lock sync.Mutex

const (
	threshold     = 10
	window        = 10 * time.Second
	port          = "2222"
	alertLogFile  = "alerts.log"
	blockEnabled  = true
)

var whitelist = map[string]bool{
	"127.0.0.1": true,
}

func logAlert(message string) {
	log.Println(message)

	f, err := os.OpenFile(alertLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Failed to write alert log:", err)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)
	f.WriteString(logEntry)
}

func blockIP(ip string) {
	if whitelist[ip] {
		log.Println("Whitelist IP detected, not blocking:", ip)
		return
	}

	cmd := exec.Command("iptables", "-A", "INPUT", "-s", ip, "-j", "DROP")
	err := cmd.Run()
	if err != nil {
		logAlert(fmt.Sprintf("❌ Failed to block %s: %v", ip, err))
	} else {
		logAlert(fmt.Sprintf("🚫 Blocked IP via iptables: %s", ip))
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	ip := strings.Split(conn.RemoteAddr().String(), ":")[0]

	if whitelist[ip] {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	now := time.Now()
	if activity, exists := ipLog[ip]; exists {
		if now.Sub(activity.lastSeen) < window {
			activity.count++
		} else {
			activity.count = 1
		}
		activity.lastSeen = now
	} else {
		ipLog[ip] = &IPActivity{count: 1, lastSeen: now}
	}

	if ipLog[ip].count > threshold {
		alert := fmt.Sprintf("[ALERT] High connection rate from %s (%d connections in %s)", ip, ipLog[ip].count, window)
		logAlert(alert)

		if blockEnabled {
			blockIP(ip)
		}

		// Reset counter to avoid spamming
		ipLog[ip].count = 0
	}
}

func main() {
	log.Println("🛡️ Mini IDS started on port", port)
	log.Println("📁 Alerts will be saved in:", alertLogFile)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error starting listener:", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}