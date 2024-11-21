# E-Commerce Product Listing System in Go

This repository demonstrates a basic e-commerce product listing system using Golang. It includes:
- Database schema design for product catalog management.
- Backend APIs for CRUD operations.
- Database replication with separate read replicas.
- Caching implementation to improve read performance.
- Cache invalidation on catalog updates.

---

## Features
1. **Database Schema Design**: Robust schema for storing product catalog.
2. **API Service**: Expose RESTful APIs for managing and viewing products.
3. **DB Replication**: Separate write and read replicas for improved scalability.
4. **Cache Integration**: Read-heavy API optimized with caching.
5. **Cache Invalidation**: Cache is invalidated when product catalog is updated.

---

## How to Run
1. Clone the repository: `git clone <>`
2. Install dependencies: `go mod tidy`
3. Set up the database and configure environment variables in `.env`.
4. Run the server: `go run main.go`

---

## API Endpoints
### Product Endpoints:
- **GET /products**: List all products (uses cache for reads).
- **GET /products/{id}**: Get a product by ID (uses cache for reads).
- **POST /products**: Add a new product.
- **POST /products/buy/{id}**: Buy a product (invalidates cache).
- **PUT /products/{id}**: Update a product (invalidates cache).
- **DELETE /products/{id}**: Delete a product.

---

## Environment Variables
| Variable         | Description                          |
|-------------------|--------------------------------------|
| `DB_HOST`        | Primary database hostname            |
| `DB_READ_REPLICA`| Read replica database hostname       |
| `CACHE_HOST`     | Redis or cache server hostname       |

---

## Contributing
Feel free to fork the repository, raise issues, and submit pull requests to improve the project.

---