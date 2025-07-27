# 🌡 temp – Raspberry Pi Temperature Monitoring

[![Go Test](https://github.com/myuser/temp/actions/workflows/test.yml/badge.svg)](https://github.com/myuser/temp/actions/workflows/test.yml)

> A lightweight client/server system to monitor and visualize temperature data from Raspberry Pi sensors.

---

## 📦 Features

- Raspberry Pi client pushes data via HTTP
- Server stores to SQLite (or pluggable backends)
- Static HTML+JS frontend (Chart.js) with:
  - Auto-refresh
  - Live point markers
  - Manual override + multi-client support
- API key authentication
- Modular storage & queue implementations
- 90%+ test coverage

---

## 🧪 Run Tests

```bash
go test ./... -cover

