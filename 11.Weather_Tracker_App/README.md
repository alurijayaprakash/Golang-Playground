# Weather Tracker Application in Golang

This is a Weather Tracker Application implemented in the Go programming language. It provides a simple way to fetch and display weather information for different locations.

# Features
1. Exposes a RESTful API for retrieving weather data.
2. Fetches weather data from a weather API.
3. Returns current weather conditions including temperature, humidity, wind speed, and more in JSON format.
4. Support for multiple locations.
5. Independent backend implementation.

### Installation

Clone the repository: 

```
https://github.com/alurijayaprakash/Golang-Playground.git
```

Navigate to the project directory:
```
cd Weather_Tracker_App
```

### Part-A : RUN BACKEND APP
1. Navigate to the backend project directory:
```
cd Backend/cmd
```
2. Run the application:

```
go run main.go
```

### Part-A : RUN FRONTEND APP
1. Navigate to the frontend project directory:
```
cd Frontend
```
2. Run the application:

open index.html in any browser

#### Usage
The Weather Tracker Application supports the following:

This App Fetches and displays the weather information for the major location.


#### Configuration:
The application requires an API key to fetch weather data. To set your API key
API key Provider : [openweathermap.com](http://api.openweathermap.org)

navigate to Weather_Tracker_App/Backend/internal/router/router.go and set the API key   
```
const APIKEY = "b3740c800de5629c909s6b5f206w8d74b5"
```
Make sure to replace your-api-key with your actual API key.

#### Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

#### License
This project is licensed under the MIT License.

