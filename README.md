# SnapCrumb – A Scalable URL Shortener in Go

SnapCrumb is a lightweight, modular URL shortener built with Go and PostgreSQL. Designed for speed, simplicity, and scalability, it provides redirect tracking, stats logging, and a clean RESTful API—all built without external frameworks.

## 🚀 Features

- ✅ Built entirely in Go (Golang)
- 🗄️ PostgreSQL-based schema for link storage and tracking
- 🔒 IP and click tracking (with future Redis caching planned)
- 🧩 Middleware architecture using Chi router
- 🧪 Unit-tested core packages
- ♻️ Designed for extensibility (analytics, auth, etc.)

## 📐 Architecture Overview

- **Router**: Powered by [Chi](https://github.com/go-chi/chi), uses idiomatic middleware
- **Handlers**: Each route has its own lightweight handler
- **Storage**: Repository layer abstracts PostgreSQL
- **Models**: Clean data model separation
- **Stats**: Tracks hits, IPs, and redirect count (pluggable)

```
📦 cmd/
├── server.go      # App entry point
📚 internal/
├── handler/       # Route handlers
├── middleware/    # Custom middleware (logging, etc)
├── model/         # URL model + validation
├── repo/          # DB repository interface & PostgreSQL impl
🧪 tests/
├── handler_test.go, repo_test.go
```

## 🛠️ Getting Started

```bash
# Clone the repo
git clone https://github.com/ericktheredd5875/snapcrumb.git && cd snapcrumb

# Set up your environment
cp .env.example .env

# Run the app
go run cmd/server.go
```

Make sure PostgreSQL is running and matches the credentials in `.env`.

## 🧪 Running Tests

```bash
go test ./...
```

## 🔮 Roadmap

- [ ] Redis-based analytics caching
- [ ] URL expiration
- [ ] User accounts + auth
- [ ] Admin dashboard

## 🙋‍♂️ About Me

Hi, I'm [Eric Harris](https://github.com/ericktheredd5875) — a backend engineer passionate about building scalable systems. I built SnapCrumb to level up my Go/PostgreSQL skills and contribute a clean, practical project to my portfolio.

---

## 📬 Contact

Questions? Want to collaborate?  
📧 ericktheredd5875@gmail.com  
🔗 [LinkedIn](https://www.linkedin.com/in/eric-harris-20579232/)
