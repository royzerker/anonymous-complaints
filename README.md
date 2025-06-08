# 🛡️ Anonymous and Traceable Citizen Complaints System

Secure platform to register complaints either **anonymously or with an alias**, with verifiable traceability through **unique hashes**, identity encryption, and advanced search by type, location, and date.

---

## 🧩 Features

- 📄 Register complaints anonymously or with alias
- 🔒 Identity encryption using AES or JOSE
- 🧾 Generate SHA-256 hash per complaint (traceability)
- 🔍 Search by type, location, and dates (Elasticsearch)
- 🧠 JWT authentication with optional alias
- 📎 Attach files to complaints
- 🛠️ Backend in Go + RESTful API

---

## ⚙️ System Requirements

### 🔧 Main Technologies

- **Go (Golang)** — REST API, business logic, authentication, encryption
- **MongoDB** — Storage for complaints, users, and attachments
- **Elasticsearch** — Advanced search (type, location, date)
- **Docker & Docker Compose** — Unified dev environment
- **Nextjs (optional)** — Frontend for complaint submission and viewing
- **JWT** — Anonymous authentication with optional alias
- **AES/JOSE** — Identity encryption

---

## 🗃️ Project Structure

