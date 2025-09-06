# Oxygen Blog

A modern blog platform built with Go, Fiber, GORM, PostgreSQL, Handlebars, and Bootstrap.

## Features

- **Modern Web Framework**: Built with Go Fiber for high performance
- **Database**: PostgreSQL with GORM ORM for data persistence
- **Templates**: Handlebars templating engine for dynamic content
- **UI Framework**: Bootstrap for responsive design
- **Admin Panel**: Dedicated admin interface for content management
- **User Management**: User authentication and profile management
- **Post Management**: Create, edit, and manage blog posts
- **Tag System**: Categorize posts with tags
- **Responsive Design**: Mobile-friendly interface

## Tech Stack

- **Backend**: Go 1.25+
- **Web Framework**: Fiber v2
- **Database**: PostgreSQL
- **ORM**: GORM
- **Templates**: Handlebars
- **Frontend**: Bootstrap
- **Configuration**: Environment-based configuration

## Project Structure

```
oxygenBlog/
├── config/          # Configuration files
├── domain/          # Domain models and business logic
├── public/          # Static assets and templates
│   ├── static/      # CSS, JS, images
│   └── views/       # Handlebars templates
├── route/           # HTTP route handlers
├── main.go          # Application entry point
├── go.mod           # Go module dependencies
└── go.sum           # Go module checksums
```

## Prerequisites

- Go 1.25 or higher
- PostgreSQL 12 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd oxygenBlog
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
# Database configuration
export OXYGEN_DB_USER=your_db_user
export OXYGEN_DB_PASS=your_db_password
export OXYGEN_DB_NAME=your_db_name
export DB_HOST=localhost
export DB_PORT=5432

# Initial admin user
export INIT_NAME="Admin User"
export INIT_EMAIL="admin@example.com"
export INIT_USER=admin
export INIT_PASS=admin
```

4. Create PostgreSQL database:
```sql
CREATE DATABASE your_db_name;
```

5. Run the application:
```bash
go run main.go
```

The application will be available at `http://localhost:8080`

## Configuration

The application uses environment variables for configuration:

### Database Configuration
- `OXYGEN_DB_USER`: Database username (required)
- `OXYGEN_DB_PASS`: Database password (required)
- `OXYGEN_DB_NAME`: Database name (required)
- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)

### Initial User Configuration
- `INIT_NAME`: Initial admin user name (required)
- `INIT_EMAIL`: Initial admin user email (required)
- `INIT_USER`: Initial admin username (default: admin)
- `INIT_PASS`: Initial admin password (default: admin)

## Development

### Running in Development Mode
```bash
go run main.go
```

The application runs with hot reload enabled for templates.

### Building for Production
```bash
go build -o oxygenBlog main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Go Fiber](https://gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Handlebars](https://handlebarsjs.com/) - Templating engine
- [Bootstrap](https://getbootstrap.com/) - CSS framework
