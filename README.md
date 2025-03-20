# **Fiber URL Shortener**
A high-performance URL shortener built using **Golang**, **Fiber**, **Redis**, and **Nginx Load Balancer**.

---

## 🛠 **Tech Stack**
- **Backend:** Golang + Fiber
- **Database:** Redis (for fast lookups)
- **Load Balancer:** Nginx
- **Containerization:** Docker + Docker Compose
- **Testing:** Go Testing + Testify
- **API Documentation:** Swagger

---

## 🔥 **Setup & Run the Project**
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/abhaybraja/distributed-url-shortener.git
cd distributed-url-shortener
```

### **2️⃣ Install Dependencies**
```sh
go mod tidy
```

### **3️⃣ Run Locally**
```sh
go run cmd/main.go
```

### **4️⃣ Run with Docker**
```sh
docker-compose up --build -d
```

### **5️⃣ Access Swagger API Documentation**
- Open: [`http://localhost:3000/swagger/index.html`](http://localhost:3000/swagger/index.html)

---

## 📌 **API Endpoints**
| Method | Endpoint | Description |
|--------|---------|-------------|
| **POST** | `/api/shorten` | Shorten a long URL |
| **GET** | `/{shortcode}` | Redirect to the original URL |
| **GET** | `/api/analytics/{shortcode}` | Get click analytics for a short URL |
| **GET** | `/swagger/*` | OpenAPI Swagger documentation |

---

## 🚀 **API Usage Examples**
### **1️⃣ Shorten a URL**
**Request**
```sh
curl -X POST http://localhost:3000/api/shorten \
     -H "Content-Type: application/json" \
     -d '{"original_url":"https://google.com"}'
```

**Response**
```json
{
  "short_url": "http://localhost:3000/G1xYz8",
  "short_code": "G1xYz8"
}
```

---

### **2️⃣ Redirect to Original URL**
**Request**
```sh
curl -v http://localhost:3000/G1xYz8
```

**Response**
- Redirects to `https://google.com`

---

### **3️⃣ Get Click Analytics**
**Request**
```sh
curl -X GET http://localhost:3000/api/analytics/G1xYz8
```

**Response**
```json
{
  "short_code": "G1xYz8",
  "clicks": 42,
  "last_clicked": "2025-03-20 14:30:00"
}
```

---

## 📦 **Docker Deployment**
### **Run Containers**
```sh
docker-compose up --build -d
```

### **Stop Containers**
```sh
docker-compose down
```

---

## 📜 **Swagger API Documentation**
Swagger UI is available at:
```
http://localhost:3000/swagger/index.html
```

To regenerate Swagger docs:
```sh
swag init -g cmd/main.go
```

---
