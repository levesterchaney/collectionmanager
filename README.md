# Collection Manager

A full-stack application for managing personal collections with item tracking, valuations, and organization features.

## 🚀 Quick Start

### Prerequisites
- Docker and Docker Compose
- Go 1.21+ (for backend development)
- Node.js 20+ (for frontend development)

### Running with Docker Compose

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd collectionmanager
   ```

2. **Start all services**
   ```bash
   docker-compose up -d
   ```

3. **Access the application**
   - **Frontend**: http://localhost:3000
   - **Backend API**: http://localhost:8080
   - **MinIO Console**: http://localhost:9001 (admin/password)
   - **Database**: PostgreSQL on localhost:5432

### Health Check
Test the services are running:
```bash
# Backend API
curl http://localhost:8080/api/v1/health

# Frontend
curl http://localhost:3000
```

## 🏗️ Architecture

### Frontend (Vue 3)
- **Framework**: Vue 3 with Composition API
- **Language**: TypeScript with strict mode
- **Build Tool**: Vite
- **Styling**: Tailwind CSS
- **Routing**: Vue Router 4
- **State Management**: Pinia
- **Testing**: Vitest + Vue Test Utils
- **Code Quality**: ESLint + Prettier

### Backend (Go)
- **Framework**: Gin HTTP router
- **Database**: PostgreSQL with migrations
- **Storage**: MinIO for file storage
- **Authentication**: JWT tokens
- **Docker**: Multi-stage builds for optimization

### Infrastructure
- **Database**: PostgreSQL 15 Alpine
- **File Storage**: MinIO (S3-compatible)
- **Web Server**: Nginx (for frontend)
- **Containerization**: Docker & Docker Compose
- **Networking**: Internal Docker network with service discovery

## 📁 Project Structure

```
collectionmanager/
├── backend/                 # Go backend service
│   ├── cmd/api/            # Application entry point
│   ├── internal/           # Internal packages
│   │   ├── auth/          # Authentication logic
│   │   ├── collections/   # Collection management
│   │   ├── items/         # Item management
│   │   ├── users/         # User management
│   │   ├── valuations/    # Valuation tracking
│   │   ├── db/           # Database connections
│   │   ├── config/       # Configuration
│   │   └── storage/      # File storage
│   ├── migrations/        # Database migrations
│   ├── go.mod            # Go dependencies
│   └── Dockerfile        # Backend container
├── frontend/              # Vue 3 frontend
│   ├── src/
│   │   ├── components/   # Reusable Vue components
│   │   ├── views/        # Page-level components
│   │   ├── router/       # Vue Router configuration
│   │   ├── stores/       # Pinia state stores
│   │   ├── composables/  # Vue composables
│   │   └── assets/       # Static assets
│   ├── public/           # Public assets
│   ├── package.json      # Node dependencies
│   ├── vite.config.ts    # Vite configuration
│   ├── tailwind.config.js # Tailwind CSS config
│   ├── nginx.conf        # Production Nginx config
│   └── Dockerfile        # Frontend container
├── docker-compose.yml     # Multi-service orchestration
└── README.md             # This file
```

## 🔧 Development

### Frontend Development

1. **Navigate to frontend directory**
   ```bash
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Run development server**
   ```bash
   npm run dev
   ```

4. **Available scripts**
   ```bash
   npm run dev          # Start dev server
   npm run build        # Build for production
   npm run test         # Run tests
   npm run lint         # Lint code
   npm run format       # Format code
   npm run type-check   # TypeScript checking
   ```

### Backend Development

1. **Setup environment**
   ```bash
   cd backend
   cp .env.example .env
   # Edit .env with your configuration
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run locally**
   ```bash
   # Start database with Docker
   docker-compose up -d db
   
   # Run backend
   make run
   ```

### Database Setup

The application uses PostgreSQL with automatic migrations. Tables include:
- `users` - User accounts and authentication
- `collections` - Collection metadata and organization
- `items` - Individual items within collections
- `valuations` - Historical value tracking

## 🌐 API Endpoints

### Health Check
- `GET /api/v1/health` - Service health status

### Authentication (Planned)
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Token refresh

### Collections (Planned)
- `GET /api/v1/collections` - List user collections
- `POST /api/v1/collections` - Create collection
- `GET /api/v1/collections/{id}` - Get collection details
- `PUT /api/v1/collections/{id}` - Update collection
- `DELETE /api/v1/collections/{id}` - Delete collection

### Items (Planned)
- `GET /api/v1/collections/{id}/items` - List collection items
- `POST /api/v1/collections/{id}/items` - Add item
- `GET /api/v1/items/{id}` - Get item details
- `PUT /api/v1/items/{id}` - Update item
- `DELETE /api/v1/items/{id}` - Delete item

## 🔌 Services & Ports

| Service | Port | Description |
|---------|------|-------------|
| Frontend | 3000 | Vue.js application |
| Backend | 8080 | Go API server |
| Database | 5432 | PostgreSQL database |
| MinIO API | 9000 | S3-compatible storage |
| MinIO Console | 9001 | Storage web interface |

## 🔑 Environment Variables

### Backend
| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `DATABASE_URL` | PostgreSQL connection string | `postgres://cmuser:cmdev1@db:5432/collectionmanager` |
| `JWT_SECRET` | JWT signing secret | `your-development-secret-key` |
| `ENVIRONMENT` | Runtime environment | `development` |

### Frontend
Environment variables are processed at build time and embedded in the application.

## 🎯 Features

### ✅ Completed
- Full Docker containerization with multi-service orchestration
- Vue 3 frontend with TypeScript, Vite, and Tailwind CSS
- Go backend API with Gin framework
- PostgreSQL database setup
- MinIO S3-compatible storage
- Development tooling (ESLint, Prettier, testing frameworks)
- Production-ready Nginx configuration
- Service health monitoring

### 🔄 In Progress
- User authentication and authorization
- Collection and item management APIs
- Frontend UI components and routing
- Database migrations and models

### 📋 Planned
- Collection creation and management UI
- Item cataloging with photos and metadata
- Valuation tracking and history
- Search and filtering capabilities
- File upload and storage integration
- Export functionality
- Automated valuation suggestions
- Analytics and insights

## 🧪 Testing

### Frontend Testing
```bash
cd frontend
npm run test           # Run unit tests
npm run test:ui        # Run tests with UI
npm run test:coverage  # Generate coverage report
```

### Backend Testing
```bash
cd backend
go test ./...          # Run all tests
make test-coverage     # Generate coverage report
```

## 🚀 Deployment

The application is containerized and ready for deployment:

```bash
# Build all services
docker-compose build

# Deploy to production
docker-compose -f docker-compose.prod.yml up -d
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and linting
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Submit a pull request

## 📄 License

[Add your license here]

---

**Status**: 🚧 Active Development - Full-stack foundation complete, building core features!