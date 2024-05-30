# exoplanet-service

# Overview

The Exoplanet Service is a Go microservice built using the gofr framework and adhering to a three-layered architecture with dependency injection. It provides various functionalities to manage different exoplanets, including adding, listing, getting by ID, updating, and deleting exoplanets. Additionally, it offers fuel estimation for trips to exoplanets based on crew capacity.

# Features
1. Add Exoplanet
Users can add a new exoplanet by providing its name, description, distance from Earth, radius, mass (if applicable), and type (GasGiant or Terrestrial).

2. List Exoplanets
Retrieve a list of all available exoplanets with the support of pagination and filters

3. Get Exoplanet by ID
Retrieve information about a specific exoplanet by its unique ID.

4. Update Exoplanet
Update the details of an existing exoplanet.

5. Delete Exoplanet
Remove an exoplanet by usning exoplanet id.

6. Fuel Estimation
Retrieve an overall fuel cost estimation for a trip to any particular exoplanet for a given crew capacity. Fuel estimation is calculated based on the distance, gravity, and crew capacity.

# Architecture
The Exoplanet Service follows a three-layered architecture consisting of:

1. Http Layer: Handles HTTP requests, responses, and request validation.
2. Service Layer: Contains the core business logic and rules for managing exoplanets.
3. Store Layer: Manages interactions with the MySQL database.

Dependency injection is utilized to manage dependencies between layers and components, allowing for better testability and flexibility.

# Technologies Used
- Go programming language
- gofr framework
- Structs, interfaces, and validations for data modeling and business logic
- MySQL for data storage
- Dependency injection for managing dependencies

# Getting Started
To start using the Exoplanet Service:

1. Clone the repository:
- git clone [https://github.com/your-username/exoplanet-service.git](https://github.com/tejasmd22/exoplanet-service.git)

2. Navigate to the project directory:
- cd exoplanet-service

3. Build and run the service:
- go run main.go
