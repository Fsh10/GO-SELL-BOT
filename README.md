# GO-SELL Bot

GO-SELL Bot is a production-ready Telegram bot for e-commerce catalog management, designed to provide seamless product browsing, search capabilities, and order processing through an intuitive conversational interface.

## Table of Contents

- [Features](#features)
- [Capabilities](#capabilities)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Architecture](#architecture)
- [Development](#development)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## Features

GO-SELL Bot offers a comprehensive suite of features for managing digital product catalogs:

*   **In-memory caching system** that stores product categories, subcategories, and product details for instant access, significantly reducing API calls and improving response times
*   **Concurrent request processing** where each user update is handled in a separate goroutine, enabling the bot to serve multiple users simultaneously without blocking operations
*   **Thread-safe state management** using `sync.RWMutex` to ensure data consistency when managing chat states and user sessions across concurrent operations
*   **Graceful shutdown mechanism** that properly saves chat state and closes connections when receiving termination signals, ensuring no data loss during deployments
*   **Structured logging** using `zap` logger with context-aware fields, providing detailed insights into bot operations, errors, and performance metrics
*   **Interface-based architecture** that enables easy testing, mocking, and swapping of implementations, particularly for the caching layer
*   **Multi-currency support** with automatic price conversion and display based on user preferences and regional settings
*   **External search API integration** allowing flexible product search capabilities that can be extended with custom search providers

## Capabilities

Users can interact with the bot through various command categories:

*   **Catalog Navigation**: Browse categories, subcategories, and products with hierarchical navigation and pagination support
*   **Product Search**: Search products by name or query with support for external search APIs and fuzzy matching
*   **Order Management**: View product details, pricing information, and initiate purchase flows through integrated payment systems
*   **User Support**: Access help documentation, view customer reviews, and contact support directly through bot commands
*   **Administrative Functions**: Configure bot settings, update product displays, manage currency conversion rates, and monitor active chat sessions (admin-only)
*   **State Persistence**: Automatic saving of chat states to disk every 10 minutes, ensuring user sessions are preserved across bot restarts
*   **Error Handling**: Comprehensive error handling with retry mechanisms for API calls and graceful degradation when external services are unavailable

## Installation

For binary builds, Docker images, and other deployment options, please see the [installation guide](#quick-start) below.

See the [CHANGELOG](CHANGELOG.md) for details on versioning and release information.

## Quick Start

### Prerequisites

*   Go 1.19 or higher
*   Telegram Bot Token (obtained from [@BotFather](https://t.me/botfather))
*   GO-SELL Seller ID

### Local Installation

```bash
git clone https://github.com/Fsh10/GO-SELL-BOT.git
cd GO-SELL-BOT
make build
./bin/bot -token=$BOT_TOKEN -seller-id=$SELLER_ID
```

### Docker Installation

```bash
docker build -t go-sell-bot:latest .
docker run -d \
  --name go-sell-bot \
  --restart unless-stopped \
  -e BOT_TOKEN=$BOT_TOKEN \
  -e SELLER_ID=$SELLER_ID \
  -v $(pwd)/chats.txt:/chats.txt \
  go-sell-bot:latest
```

### Docker Compose

```bash
docker-compose up -d
```

## Configuration

The bot supports configuration via command-line flags or environment variables:

| Flag | Environment Variable | Description | Required |
|------|---------------------|-------------|----------|
| `-token` | `BOT_TOKEN` | Telegram bot token | Yes |
| `-seller-id` | `SELLER_ID` | GO-SELL seller identifier | Yes |
| `-debug` | `DEBUG` | Enable debug-level logging | No |
| `-load-cache` | `LOAD_CACHE` | Load cache from disk on startup (default: true) | No |
| `-chats-file` | `CHATS_FILE` | Path to chat state persistence file (default: chats.txt) | No |
| `-search-url` | `SEARCH_URL` | External search API endpoint URL | No |
| `-uproduct` | `U_PRODUCT` | Universal product ID for search integration | No |
| `-uproductopt` | `U_PRODUCT_OPT` | Universal product option identifier | No |
| `-search-instruction` | `SEARCH_INSTRUCTION` | Display search usage instructions | No |

## Architecture

```
cmd/
  └── main.go          # Application entry point, signal handling, graceful shutdown
bot/
  ├── bot.go           # Core bot logic, update processing, state management
  ├── cache/           # In-memory cache implementation with interface abstraction
  ├── digi/            # GO-SELL API client and integration layer
  ├── search/          # Product search functionality with external API support
  ├── desc/            # Domain entities and data structures
  ├── countries/       # Country and currency handling
  └── commands/        # Command handlers and callback processors
utils/http/            # Generic HTTP client with type-safe request helpers
```

### Design Principles

*   **Dependency Injection**: All dependencies are injected via constructors, enabling testability and flexibility
*   **Interface Segregation**: Core components use interfaces (`inMemoryCache`, `Client`) allowing for easy testing and implementation swapping
*   **Concurrency Safety**: Thread-safe operations using `sync.RWMutex` for read-heavy workloads
*   **Error Propagation**: Structured error handling with context preservation throughout the call stack
*   **Resource Management**: Proper cleanup of file handles, HTTP connections, and goroutines

## Development

### Building from Source

```bash
make build              # Build binary (output: bin/bot)
make test              # Run tests with race detection
make lint              # Run golangci-lint
make fmt               # Format code with gofmt and goimports
make vet               # Run go vet
make all               # Run all checks and build
```

### Tech Stack

*   **Go 1.19+** - Programming language
*   **go-telegram-bot-api/v5** - Telegram Bot API client library
*   **zap** - High-performance structured logging
*   **golangci-lint** - Fast Go linter with multiple analyzers

### Project Structure

The project follows standard Go project layout conventions:
*   Packages organized by functionality and domain boundaries
*   Clear separation between business logic, infrastructure, and utilities
*   Comprehensive error handling and logging throughout

## Deployment

### Production Considerations

*   **Binary Size**: Statically linked binary (~10MB) with no external dependencies
*   **Resource Usage**: Low memory footprint with efficient in-memory caching
*   **Monitoring**: Structured JSON logs compatible with log aggregation systems
*   **State Persistence**: Chat state automatically saved every 10 minutes
*   **Health Checks**: Graceful shutdown on SIGTERM/SIGINT signals

### Monitoring

The bot provides structured logging via `zap` logger:
*   Bot startup and shutdown events
*   API request/response logging with error details
*   Command and callback processing metrics
*   Chat state save operations
*   Performance timing information

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on:
*   Code style and standards
*   Testing requirements
*   Pull request process
*   Issue reporting

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
