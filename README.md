# 🛡️ Mini IDS (Intrusion Detection System) in Go

A simple network-based IDS that detects abnormal connection activity in real-time. The IDS listens on a TCP port and logs and blocks suspicious IPs based on the number of connection attempts in a short time window.

## 🚀 Features

- **High connection rate detection**: Detects IPs making too many connections within a specified time window.
- **Whitelisted IPs**: Allows you to configure trusted IPs that will not be blocked.
- **Auto-blocking**: Automatically blocks IPs exceeding the connection threshold using `iptables` (Linux only).
- **Alert Logging**: Logs suspicious activities and blocks to an `alerts.log` file.
- **Configurable Settings**: Customize thresholds, time windows, and port.

## ⚙️ How It Works

- The IDS listens on a specified TCP port (default is `2222`).
- It tracks the number of connection attempts for each IP address.
- If an IP exceeds a threshold of connections within a defined time window, an alert is logged and the IP is automatically blocked via `iptables`.
- The `alerts.log` file records all alerts, including blocked IPs, timestamps, and connection counts.

## 🧪 Usage

1. Clone the repository:

```bash
git clone https://github.com/Lis240/Mini-IDS-Intrusion-Detection-System-
cd mini-ids
