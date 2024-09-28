# api-elevator
Simple API for Elevator Logic Use Queue

## Install And Running
### Required

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Step By Step

1. **Clone Repository:**
   ```bash
   git clone https://github.com/SyaifulDistroid/api-elevator.git
   cd api-elevator
   ```

2. **Build And Run Dockercompose**
   ```bash
    cd docker
    docker-compose up --build
   ```

3. **API**
    ```bash
    Request Floor :
    curl -X POST http://localhost/request-floor -d "floor=5"
    ```
    ```bash
    Status :
    curl http://localhost/elevator-status
    ```

### Core Library
- **Golang:** Backend logic
- **RabbitMQ:** Message broker
- **Gin:** HTTP framework untuk REST API
- **Docker:** Containerization
- **Nginx:** Load balancer
- **Testify:** Unit testing