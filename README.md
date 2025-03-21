# Weather-api
This project demonstrates how to build a Weather API in Golang that fetches weather data from a 3rd party API (such as Visual Crossing API) and implements Redis caching to optimize API calls and reduce response time.

[.](https://roadmap.sh/projects/weather-api-wrapper-service)
🚀 Features
✅ Fetch real-time weather data using 3rd party API
✅ Store results in Redis with an expiration period
✅ Return cached results if data is available
✅ Efficient error handling with appropriate HTTP status codes
✅ Environment variables for API keys and Redis config

🛠️ Tech Stack
Golang - Backend API
Redis - In-memory caching
Gin Framework - RESTful API framework
Resty - HTTP client for 3rd party API requests

⚙️ Installation
1. Clone the Repository
git clone https://github.com/radhikabhut/weather-api.git
cd weather-api

2. Docker
REDIS_HOST=localhost
REDIS_PORT=6379
API_KEY=your_visual_crossing_api_key

3. Install Dependencies
go mod tidy

4. Run the Application
go run ain.go
