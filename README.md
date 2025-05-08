# Telegram Bot in Go

A simple Telegram bot built using Go and the telebot.v3 library.

## Prerequisites

- Go 1.18 or later
- Telegram Bot Token (get it from [@BotFather](https://t.me/BotFather))

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/telegram-bot.git
   cd telegram-bot
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Configure the bot:
   - Edit the `.env` file and replace `your_telegram_bot_token_here` with your actual Telegram bot token.

## Running the Bot

```bash
go run cmd/bot/main.go
```

Or build and run:

```bash
go build -o bin/tgbot cmd/bot/main.go
./bin/tgbot
```

## Project Structure

```
.
├── cmd/
│   └── bot/
│       └── main.go       # Entry point
├── config/
│   └── config.go         # Configuration management
├── internal/
│   └── handlers/
│       └── handlers.go   # Bot command handlers
├── .env                  # Environment variables
├── .gitignore
├── go.mod
├── go.sum
├── main.go               # Root level main file (optional)
└── README.md
```

## Adding New Commands

To add new commands, edit the `internal/handlers/handlers.go` file and register new handlers in the `RegisterHandlers` function.

## License

MIT