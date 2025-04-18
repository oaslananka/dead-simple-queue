# Dead Simple Queue

Dead Simple Queue is a minimal, file-based task queue microservice written in Go. It accepts tasks via an HTTP API, processes them asynchronously using a worker, and logs them to disk. It comes with Docker support, unit tests, and CI/CD integration using GitHub Actions.

---

## ✨ Features

* 📅 Task enqueueing via HTTP API
* 🧠 Background worker for task processing
* 📃 Logs all tasks to a persistent `tasks.log` file
* 🧪 Simple unit test coverage
* 🛣️ Docker support
* ⚙️ GitHub Actions CI/CD pipelines

---

## 📦 API Usage

### Enqueue Task

```http
POST /enqueue
Content-Type: application/json

{
  "id": 1,
  "payload": "hello world"
}
```

The worker processes and logs the task automatically after the request.

---

## 🛠️ Running Locally

### 1. Requirements

* Go 1.20+
* Git
* Docker (optional)

### 2. Run the Server

```bash
go run main.go
```

The server listens on port `8080` and starts the background worker.

---

## 🧪 Running Tests

```bash
go test ./...
```

All tests are located in the `tests/` directory.

---

## 🚀 Docker Support

```bash
docker build -t dead-simple-queue .
docker run -p 8080:8080 dead-simple-queue
```

---

## 🔄 CI/CD

CI/CD is handled via GitHub Actions:

* `ci.yml` runs unit tests on every push
* `cd.yml` builds the Docker image on push to `main`

---

## 📁 Project Structure

```
dead-simple-queue/
├── api/             # HTTP API logic
├── queue/           # Task definition and worker
├── storage/         # File output logic
├── tests/           # Unit tests
├── .github/workflows
├── Dockerfile
├── go.mod
├── main.go
└── README.md
```

---

## 📄 License

MIT License

(c) 2025, Osman Aslan
