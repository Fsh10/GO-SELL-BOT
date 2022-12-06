# GO-SELL Bot

Production-ready Telegram bot for e-commerce catalog management with in-memory caching, concurrent request handling, and graceful shutdown.

## Table of Contents

- [Quick Start](#quick-start)
- [Architecture](#architecture)
- [Configuration](#configuration)
- [Development](#development)
- [Deployment](#deployment)
- [License](#license)

## Quick Start

```bash
git clone https://github.com/Fsh10/GO-SELL-BOT.git && cd GO-SELL-BOT
make build
./bin/bot -token=$BOT_TOKEN -seller-id=$SELLER_ID
```

**Docker:**
```bash
docker build -t go-sell-bot .
docker run -d -e BOT_TOKEN=$BOT_TOKEN -e SELLER_ID=$SELLER_ID go-sell-bot
```

## Architecture

```
cmd/
  └── main.go          # Entry point, signal handling, graceful shutdown
bot/
  ├── bot.go           # Core bot logic, update processing
  ├── cache/           # In-memory cache with interface abstraction
  ├── digi/            # GO-SELL API client
  ├── search/          # Product search with external API support
  └── desc/            # Domain entities
utils/http/            # Generic HTTP client with type-safe requests
```

**Key Design Decisions:**
- **Dependency Injection**: All dependencies injected via constructor
- **Interface-based Cache**: `inMemoryCache` interface enables easy testing and swapping implementations
- **Concurrent Processing**: Each update processed in separate goroutine
- **State Management**: Thread-safe chat state with `sync.RWMutex`
- **Structured Logging**: `zap` logger with context-aware fields

## Configuration

| Flag | Env Var | Description |
|------|---------|-------------|
| `-token` | `BOT_TOKEN` | Telegram bot token (required) |
| `-seller-id` | `SELLER_ID` | GO-SELL seller ID (required) |
| `-debug` | `DEBUG` | Enable debug logging |
| `-load-cache` | `LOAD_CACHE` | Load cache on startup (default: true) |
| `-chats-file` | `CHATS_FILE` | Chat state persistence file (default: chats.txt) |
| `-search-url` | `SEARCH_URL` | External search API endpoint |

## Development

```bash
make install-tools    # Install golangci-lint, goimports
make test            # Run tests with race detection
make lint            # Run linter
make fmt             # Format code
make all             # Run all checks and build
```

**Tech Stack:**
- Go 1.19+
- `go-telegram-bot-api/v5` - Telegram Bot API
- `zap` - Structured logging
- `golangci-lint` - Code quality

## Deployment

**Production Build:**
```bash
make build
# Binary: bin/bot (~10MB, statically linked)
```

**Docker:**
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

**Monitoring:**
- Structured JSON logs via `zap`
- Chat state persisted every 10 minutes
- Graceful shutdown on SIGTERM/SIGINT

## License

MIT License - see [LICENSE](LICENSE) file for details.
