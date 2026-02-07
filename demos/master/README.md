# Omnidraw Master

A premium, high-performance Lucky Draw system demonstration. This module showcases a clean Go architecture, real-time player management, and a sophisticated CLI interface.

---

## Features

-   **Sophisticated CLI**: An interactive terminal UI built with `promptui` and `fatih/color`, featuring a curated color palette and smooth navigation.
-   **Smart Service Logic**:
    -   **Deduplication**: Automatically filters out duplicate names and whitespace during import.
    -   **O(1) Performance**: Uses an optimized "swap-and-shrink" algorithm for drawing, ensuring constant-time performance regardless of the pool size.
-   **Clean Architecture**: Separation of concerns across `cmd`, `handler`, `service`, and `router` layers.
-   **Thread-Safe**: Fully protected by `sync.Mutex` for concurrent API access.

---

## Project Structure

```text
demos/master/
├── cmd/
│   ├── cli/            # Interactive CLI application
│   └── server/         # RESTful API server (Gin-based)
├── internal/
│   ├── handler/        # HTTP controllers & request binding
│   ├── router/         # Route definitions & Gin configuration
│   └── service/        # Core business logic & state management
├── test_players.txt    # Standardized test dataset (50+ players)
├── go.mod              # Dependencies (Gin, promptui, color)
└── README.md           # This documentation
```

---

## Getting Started

### 1. Start the Engine
The API server handles the player pool and lottery logic.
```bash
go run cmd/server/main.go
```
*The server listens on `:8080` by default.*

### 2. Launch the Experience
In a separate terminal, start the interactive CLI:
```bash
go run cmd/cli/main.go
```

---

## CLI Interaction

The CLI is designed for a seamless, interactive experience:

1.  **Import List**: Select this to load your candidate file (e.g., `test_players.txt`).
    - Success output: `✓ Loaded 57 players`
2.  **Draw Winner**: Trigger the randomized draw with a subtle animation.
    - Winner found: `★ Wendy` followed by `56 left in pool`
    - Pool empty: `✗ no players remaining for the draw`
3.  **Exit**: Safely terminate the session.
    - Feedback: `✓ Session ended. Goodbye.`

---

## API Reference

### Import Players
`POST /master/players`

| Field | Type | Description |
| :--- | :--- | :--- |
| `players` | `[]string` | List of player names to load into the pool. |

**Example Response (200 OK):**
```json
{ "message": "success" }
```

### Draw Winner
`GET /master/draw`

**Example Response (200 OK) - Winner Found:**
```json
{
  "message": "success",
  "winner": "Wendy",
  "remaining": 56
}
```

**Example Response (200 OK) - No Players Left:**
```json
{
  "message": "no players remaining for the draw"
}
```
*Note: `remaining` is returned as a JSON number.*

---

## Advanced

You can connect the CLI to a custom server address:
```bash
go run cmd/cli/main.go -server http://your-api-host:8080/master
```
