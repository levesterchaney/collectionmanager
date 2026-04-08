# Collection Manager

A full-stack application for managing personal collections with item tracking, valuations, and organization features.

## 🚀 Quick Start

### Prerequisites
- Docker and Docker Compose
- Go 1.21+ (for local development)
- Node.js 18+ (for frontend development)

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
   - Backend API: http://localhost:8080
   - MinIO Console: http://localhost:9001 (admin/password)
   - Frontend: http://localhost:3000 (when available)

### Health Check
Test the backend is running:
```bash
curl http://localhost:8080/api/v1/health
```

## 🏗️ Architecture

### Backend (Go)
- **Framework**: Gin HTTP router
- **Database**: PostgreSQL with migrations
- **Storage**: MinIO for file storage
- **Authentication**: JWT tokens

### Frontend (React - Coming Soon)
- **Framework**: React with TypeScript
- **UI**: Modern component library
- **State Management**: Context API or Redux

### Infrastructure
- **Database**: PostgreSQL 16
- **File Storage**: MinIO (S3-compatible)
- **Containerization**: Docker & Docker Compose

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
├── frontend/              # React frontend (coming soon)
├── docker-compose.yml     # Multi-service setup
└── README.md             # This file
```

## 🔧 Development

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

### API Endpoints (Planned)

- **Authentication**
  - `POST /api/v1/auth/register` - User registration
  - `POST /api/v1/auth/login` - User login
  - `POST /api/v1/auth/refresh` - Token refresh

- **Collections**
  - `GET /api/v1/collections` - List user collections
  - `POST /api/v1/collections` - Create collection
  - `GET /api/v1/collections/{id}` - Get collection details
  - `PUT /api/v1/collections/{id}` - Update collection
  - `DELETE /api/v1/collections/{id}` - Delete collection

- **Items**
  - `GET /api/v1/collections/{id}/items` - List collection items
  - `POST /api/v1/collections/{id}/items` - Add item
  - `GET /api/v1/items/{id}` - Get item details
  - `PUT /api/v1/items/{id}` - Update item
  - `DELETE /api/v1/items/{id}` - Delete item

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `DATABASE_URL` | PostgreSQL connection string | `postgres://localhost/collectionmanager` |
| `JWT_SECRET` | JWT signing secret | `your-secret-key` |
| `ENVIRONMENT` | Runtime environment | `development` |

## 🎯 Features (Planned)

### Core Features
- ✅ User authentication and authorization
- ✅ Collection creation and management
- ✅ Item cataloging with photos and metadata
- ✅ Valuation tracking and history
- 🔄 Search and filtering capabilities
- 🔄 File upload and storage
- 🔄 Export functionality

### Advanced Features
- 🔄 Automated valuation suggestions
- 🔄 Collection sharing and privacy controls
- 🔄 Bulk import/export
- 🔄 Analytics and insights
- 🔄 Mobile-responsive design

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

[Add your license here]

---

**Status**: 🚧 Early Development - Backend API scaffolded, frontend coming soon!
