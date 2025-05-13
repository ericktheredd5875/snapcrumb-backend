# 🚀 SnapCrumb

> **Shorten in a snap. Share your crumb trail.**  

SnapCrumb is a fast and lightweight URL shortening service designed for quick redirection and easy sharing.  
Built with modern backend practices, Dockerized for easy deployment, and ready for cloud scaling.

---

[![codecov](https://codecov.io/gh/ericktheredd5875/snapcrumb-backend/graph/badge.svg?token=T0T34K27RD)](https://codecov.io/gh/ericktheredd5875/snapcrumb-backend)

---

## 🛠️ Features

- Shorten long URLs into small, easy-to-share links
- Redirect users instantly via custom shortcodes
- Track creation dates and (optionally) click counts
- Designed for containerization (Docker)
- Ready for cloud deployment (AWS, GCP, Azure)
- Future-proof structure for scaling and monitoring

---

## 📚 API Endpoints

### 1. Create a Shortened URL

**Request**

POST /shorten Content-Type: application/json { "url": "https://www.example.com/some/very/long/path" }

**Success Response**

201 Created { "shortcode": "aBcD123", "shortened_url": "https://yourdomain.com/aBcD123" }

---

### 2. Redirect to Original URL

**Request**

GET /{shortcode}

**Behavior**
- 302 Redirect to original URL if shortcode exists
- 404 Not Found if shortcode is missing

---

## 🗄️ Database Schema

```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    shortcode VARCHAR(10) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP,
    visit_count INTEGER DEFAULT 0
);
```

## 📦 Project Structure

```bash
snapcrumb/
├── cmd/server/          # App entrypoint
├── internal/api/        # HTTP Handlers
├── internal/db/         # Database logic
├── internal/model/      # Data models
├── pkg/utils/           # Helpers
├── scripts/             # Deployment scripts
├── Dockerfile           # Container config
├── docker-compose.yml   # Local dev setup
└── README.md            # This file
```

## 🐳 Quickstart (Local Dev) -- NOT COMPLETED YET

1. Clone the repo

```bash
git clone https://github.com/yourusername/snapcrumb.git
cd snapcrumb
```

2. Spin up using Docker

```bash
docker-compose up --build
```

3. Visit your local server

```bash
http://localhost:8080
```

## 🚀 Development

To start the server with auto-reloading:

- **On macOS/Linux**:
```bash
  ./dev.sh
```
- **On Windows (PowerShell)**:
```powershell
./dev.ps1
```
**Requires air (https://github.com/air-verse/air) to be installed**


## 📈 Future Improvements (Roadmap)

* URL expiration feature
* Authentication for link management
* Admin dashboard (metrics, stats)
* Analytics dashboard (link clicks, source locations)
* Fully automated CI/CD with GitHub Actions
* Kubernetes deployment (advanced)

## ✨ Credits

Built by Eric Harris as part of a backend engineering and DevOps skill upgrade journey 🚀.
