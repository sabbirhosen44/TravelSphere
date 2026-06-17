# TravelSphere

TravelSphere is a lightweight, responsive travel-planning web application built using **Go** and the **Beego** framework. It integrates external APIs to display country statistics, weather details, and local tourist attractions, alongside a JSON-based wishlist planner.

## Features

- **Country Exploration:** Search and browse countries using the RestCountries API.
- **Tourist Attractions:** View points of interest powered by the OpenTripMap API.
- **Wishlist Management:** Add, update, and remove travel destinations with custom notes and status.
- **Mock Authentication:** Session-based simulator to handle restricted access pages.
- **Local Storage:** Zero-configuration local database using `wishlist.json`.

---

## Getting Started

### Prerequisites

- **Go** (Version 1.26 or higher recommended)
- **Beego CLI (optional, for hot-reloading development)**:
  ```bash
  go install github.com/beego/bee/v2@latest
  ```

### Installation & Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/sabbirhosen44/TravelSphere.git
   cd TravelSphere
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configure environment settings**:
   Copy the example configuration to build the main configuration file:
   ```bash
   cp conf/app.conf.example conf/app.conf
   ```
   Open `conf/app.conf` and configure the environment variables/keys as necessary (e.g., `countries_api_key`).

---

## Running the Application

### Development (Live Reload)
To run with live-reloading during development, use Beego's `bee` tool:
```bash
bee run
```
The application will be running at `http://localhost:8080`.

### Standard Mode / Production
To build the application and run the compiled binary:
```bash
go build -o travelsphere.exe
./travelsphere.exe
```
Or run the file directly without compiling:
```bash
go run main.go
```

---

## Running Tests

To run the unit tests:
```bash
go test ./tests/...
```

---

## Project Structure

```text
├── conf/            # App configurations (app.conf)
├── controllers/     # MVC controller logic and middleware filters
├── models/          # Data structures and representations
├── routers/         # Endpoint and page routing definitions
├── services/        # Service layers for external API requests
├── static/          # Static assets (CSS, JS, images)
├── views/           # Beego HTML template views (.tpl)
├── wishlist.json    # Local file database for wishlist storage
└── main.go          # Application entry point
```
