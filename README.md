# Ex0r


/ctfx-backend/
├── cmd/                        # Application entry points
│   └── api/                    # Main API server
│       └── main.go             # Entry point for the API server
├── internal/                   # Private application code
│   ├── api/                    # API handlers and routes
│   │   ├── middleware/         # Request middleware
│   │   ├── handlers/           # Request handlers
│   │   └── router.go           # Route definitions
│   ├── auth/                   # Authentication logic
│   │   ├── jwt.go
│   │   └── password.go
│   ├── challenge/              # Challenge management
│   │   ├── service.go          # Business logic
│   │   ├── repository.go       # Data access
│   │   └── model.go            # Challenge data models
│   ├── team/                   # Team management
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   ├── user/                   # User management
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   ├── submission/             # Flag submission
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   ├── scoreboard/             # Scoreboard calculations
│   │   └── service.go
│   ├── file/                   # File storage
│   │   └── service.go
│   └── websocket/              # Real-time updates
│       └── hub.go
├── pkg/                        # Public library code
│   ├── config/                 # Configuration handling
│   │   └── config.go
│   ├── database/               # Database connection
│   │   └── postgres.go
│   ├── cache/                  # Redis operations
│   │   └── redis.go
│   ├── validator/              # Input validation
│   │   └── validator.go
│   └── logger/                 # Logging utilities
│       └── logger.go
├── migrations/                 # Database migrations
│   ├── 000001_create_users.up.sql
│   ├── 000001_create_users.down.sql
│   └── ...
├── config/                     # Configuration files
│   ├── config.dev.yaml
│   └── config.prod.yaml
├── scripts/                    # Utility scripts
│   ├── seed.go                 # Database seeding
│   └── migrate.sh              # Migration runner
├── Dockerfile                  # Container definition
├── docker-compose.yml          # Local development setup
├── go.mod                      # Go module definition
└── go.sum                      # Go module checksums