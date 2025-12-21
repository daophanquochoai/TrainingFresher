# Hướng dẫn sử dụng Docker và Docker Compose

## Yêu cầu
- Docker Engine 20.10+
- Docker Compose 2.0+

## Cấu trúc

Dự án bao gồm:
- **PostgreSQL**: Database server
- **Redis**: Cache server
- **productservice-http**: Product service HTTP API (port 8080)
- **productservice-grpc**: Product service gRPC server (port 9090)
- **categoryservice-http**: Category service HTTP API (port 8081)

## Cách sử dụng

### 1. Tạo file .env (tùy chọn)

Copy file `.env.example` thành `.env` và chỉnh sửa các biến môi trường nếu cần:

```bash
cp .env.example .env
```

### 2. Build và chạy tất cả services

```bash
docker-compose up -d
```

### 3. Xem logs

```bash
# Xem logs tất cả services
docker-compose logs -f

# Xem logs của một service cụ thể
docker-compose logs -f productservice-http
docker-compose logs -f categoryservice-http
```

### 4. Dừng services

```bash
docker-compose down
```

### 5. Dừng và xóa volumes (xóa dữ liệu)

```bash
docker-compose down -v
```

## Cấu hình biến môi trường

Tất cả các biến môi trường có thể được cấu hình trong file `.env` hoặc trực tiếp trong `docker-compose.yml`. Các biến môi trường chính:

### PostgreSQL
- `POSTGRES_USER`: Tên user (mặc định: admin)
- `POSTGRES_PASSWORD`: Mật khẩu (mặc định: admin)
- `POSTGRES_DB`: Tên database (mặc định: db)
- `POSTGRES_PORT`: Port (mặc định: 5432)

### Redis
- `REDIS_HOST`: Hostname (mặc định: redis)
- `REDIS_PORT`: Port (mặc định: 6379)
- `REDIS_PASSWORD`: Mật khẩu (mặc định: trống)
- `REDIS_DB`: Database number (mặc định: 0)

### Product Service
- `PRODUCT_APP_NAME`: Tên app (mặc định: product)
- `PRODUCT_HTTP_PORT`: HTTP port (mặc định: 8080)
- `PRODUCT_GRPC_PORT`: gRPC port (mặc định: 9090)

### Category Service
- `CATEGORY_APP_NAME`: Tên app (mặc định: category)
- `CATEGORY_HTTP_PORT`: HTTP port (mặc định: 8081)

## Kiểm tra services

Sau khi chạy, bạn có thể kiểm tra:

- Product HTTP API: http://localhost:8080
- Category HTTP API: http://localhost:8081
- Product gRPC: localhost:9090
- PostgreSQL: localhost:5432
- Redis: localhost:6379

## Rebuild images

Nếu bạn thay đổi code và cần rebuild:

```bash
docker-compose build
docker-compose up -d
```

Hoặc rebuild và restart:

```bash
docker-compose up -d --build
```

