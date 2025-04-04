# 🛡️ Mini IDS (Intrusion Detection System) in Go

A simple network-based IDS that detects abnormal connection activity in real-time.

## 🚀 Features

- Detects IPs making too many connections in a short time
- Logs suspicious IPs with timestamps
- Easy to configure and extend

## ⚙️ How It Works

- Listens on a TCP port (default `2222`)
- Tracks connection attempts per IP
- Triggers alert if an IP exceeds N attempts in T seconds

## 🧪 Usage

```bash
go run main.go

## 🔐 Advanced Features

- Logs alerts to `alerts.log`
- Whitelist IPs: add trusted IPs to the `whitelist` map in `main.go`
- Auto-blocks abusive IPs using `iptables` (Linux only)

> ⚠️ Auto-blocking requires **root privileges**.
> Run with `sudo` if needed:

```bash
sudo go run main.go
