# E-Commerce Backend

Backend for an e-commerce platform developed in **Golang**, following **Clean Architecture / Hexagonal** principles.

The system supports **product management with variations**, such as size, color, or other configurable options, allowing flexible and scalable product catalogs.

## Categories
<img width="1366" height="630" alt="image" src="https://github.com/user-attachments/assets/f14a3269-83a7-427d-ab59-087348d62517" />

## variations
<img width="1366" height="696" alt="image" src="https://github.com/user-attachments/assets/d415d0c8-3cd1-4c44-bad8-af382d7aecf2" />


## products without variations 
<img width="1366" height="729" alt="image" src="https://github.com/user-attachments/assets/620344b9-0016-49bf-b32f-094d7dbdfe72" />

## Products with variations
<img width="1366" height="694" alt="image" src="https://github.com/user-attachments/assets/0774141c-6267-4952-b790-368bd69835d1" />

<img width="1366" height="693" alt="image" src="https://github.com/user-attachments/assets/dcdcedec-b051-4670-9132-b48896e98b00" />

<img width="1366" height="693" alt="image" src="https://github.com/user-attachments/assets/4ef4f166-9d82-4984-86b3-90cc7249e632" />

<img width="1366" height="694" alt="image" src="https://github.com/user-attachments/assets/cceae67b-039d-412d-a123-e35702ec416a" />


## Db design
<img width="528" height="438" alt="image" src="https://github.com/user-attachments/assets/7af8bf98-7643-40f7-b29c-0d350a755ef6" />

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/amorindev/back-ecomm
    cd back-ecomm
    ```

2. Download dependencies:
    ```bash
    go mod tidy
    ```

3. Set environment variables, add a `.env` file based on `env.example`

4. Start development services (Docker):
    ```bash
    make compose-dev
    ```
    > This runs:
    > ```bash
    > docker-compose -f docker-compose.dev.yml --env-file .env up
    > ```

5. Run the project:
    ```bash
    make run
    ```
    Or:
    ```bash
    go run main.go
    ```

