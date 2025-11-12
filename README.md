# BANK STATEMENT VIEWER - Go + Next.js
#### Created by  Adrian Milano
## Goal

Build a small full-stack app that lets users upload a bank statement file, view insights, and
inspect transaction issues.

## 🛠️ Setup Instructions

### Option 1: Automated Setup (Recommended)

```bash
# Clone and setup everything automatically
git clone <repository-url>
cd bank-statement-viewer
make setup         # Install all dependencies
make dev           # Start both backend and frontend
make test-coverage # Unit test both backend and frontend
```

### Option 2: Manual Setup

#### 1. Clone the Repository

```bash
git clone <repository-url>
cd bank-statement-viewer
```

#### 2. Backend Setup

```bash
cd backend
# Download and install Go dependencies
go mod download
go mod tidy
# Copy .env
cp .env.example .env
# Start the server
go run ./cmd/main.go
# OR use Make
make run
```
The backend will start on `http://localhost:8080`

#### 3. Frontend Setup (New Terminal)

```bash
cd frontend
# Install Node.js dependencies
npm install
# Start development server
npm run dev
# OR use Make
make dev
```


## 🎯 Architecture Decisions

```
┌──────────────────────────────┐
│          Frontend            │
│  Next.js + TypeScript + CSS  │
│  - Pages (Next.js routes)    │
│  - Components (UI logic)     │
│  - Services (API calls)      │
│  - State management (hooks)  │
└──────────────┬───────────────┘
               │  REST API calls
               ▼
┌──────────────────────────────┐
│           Backend            │
│            GoLang            │
│  - main.go (entry point)     │
│  - handler/ (HTTP handlers)  │
│  - service/ (business logic) │
│  - repository/ (data layer)  │
│  - model/ (data types)       │
│  - utils/ (helpers, respond) │
└──────────────┬───────────────┘
               │
               ▼
┌──────────────────────────────┐
│       Storage / CSV File     │
│  Transactions are persisted  │
│  to CSV for simplicity.      │
│  Repo reads/writes parsed    │
│  CSV records to memory.      │
└──────────────────────────────┘
```

### Layer Responsibilities
| Layer                  | Description                                                                                                                  |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| **Frontend (Next.js)** | Responsible for UI rendering, user interactions, pagination, sorting (delegated to backend), and calling REST APIs.          |
| **Backend (Go)**       | Exposes REST API endpoints for uploading, parsing, and querying transaction data. Implements clean layering for testability. |
| **Service Layer**      | Contains domain logic like filtering, sorting, and pagination. Keeps business logic separate from transport (HTTP).          |
| **Repository Layer**   | Manages data persistence (in-memory + CSV). Acts as a data abstraction.                                                      |
| **Handler Layer**      | Maps HTTP requests to service calls, formats responses with `respondJSON()` and `respondErr()`.                              |
| **Model Layer**        | Defines `Transaction`, `Meta`, and other core domain models.                                                                 |

### Data Flow Overview
1. User uploads a CSV file via the frontend.
2. Frontend sends it via `POST /api/upload` to the backend.
3. Backend handler delegates to service → repository → CSV parsing & store.
4. User requests `/api/issues?page=1&limit=10&sort_order=desc&sort_by=name` → backend service filters + sorts data → returns paginated results.
5. Frontend displays issues in a dynamic table with sortable columns.

### Technologies Used
#### Frontend 
- Next.js
- TypeScript
- Jest + React Testing Library for testing
- Native CSS 
#### Backend
- Go 
- net/http for lightweight routing
- encoding/csv + json for I/O
- Testify for unit testing
#### CI/CD
GitHub Actions pipeline runs:
- `go test` for backend
- `npm test && npm run build` for frontend
- Linting and type checking before merge

### Design Decisions
- **Separation of Concerns** - clear distinction between HTTP transport, business logic, and persistence
- **Stateless Services** - backend services can scale horizontally since data is read from file/memory.
- **Testability** - service and repo layers are independently testable using mocks.
- **Extensibility** - backend can be easily extended to support a database later (e.g., MySQL, PostgreSQL).