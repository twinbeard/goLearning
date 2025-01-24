# Setting Up Prometheus in Golang

## Introduction

This guide will help you set up Prometheus monitoring for a Golang application.

## Step 1: Install Prometheus Client Library

First, you need to install the Prometheus client library for Go.

1. Prometheus-go: [LINK🥰](https://github.com/prometheus/client_golang).

2. Prometheus-go guide: [LINK🥰](https://github.com/prometheus/client_golang).

```zsh
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp

```

## Step 2: Set up Graphana and Prometheus in Docker

### Prometheus

```docker
prometheus:
    image: prom/prometheus:latest
    container_name: pre-event-prometheus
    restart: unless-stopped
    volumes:
      - ./prometheus/prometheus.yml:/etc/prometheus/prometheus.yml

# ./prometheus/prometheus.yml -> Đây là location của file prometheus.yml để configure
# ghi đúng file là yml or yaml sao cho cũng khớp với file tương ứng trong working directory luôn nhé.Otherwise 0, thì nó không chạy đâu

      - ./data/prometheus_data:/prometheus
# Khi docker run thì data trong docker sẽ được map về địa chỉ này ./data/prometheus_data

    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/etc/prometheus/console_libraries'
      - '--web.console.templates=/etc/prometheus/consoles'
      - '--web.enable-lifecycle'
    extra_hosts:
      - host.docker.internal:host-gateway
    ports:
      - "9092:9090"
    networks:
      - pre-go-local-network
```

### Graphana

```docker
grafana:
    image: grafana/grafana
    container_name: pre-event-grafana
    hostname: grafana
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
      - GF_USERS_ALLOW_SIGN_UP=false
      - GF_SERVER_DOMAIN=localhost
      #enable logger
      - GF_LOG_MODE=console file
      - GF_LOG_FILTERS=alerting.notifier.slack:debug alermanager:debug ngalert:debug
    volumes:
      - ./grafana-storage:/var/lib/grafana
    ports:
      - "3002:3000"
    networks:
      - pre-go-local-network # nhớ set up custom netwrok
```

## Step 3: Set up prometheus.yml file

Configure the target where the prometheuse is supposed to scrape for monitoring

```yml
global: # Cấu hình global cho prometheus
  scrape_interval: 15s # Định kỳ 15s scrape data từ các target mỗi

## host.docker.internal là một địa chỉ đặc biệt của docker, nó cho phép các container trong cùng một network có thể truy cập vào nhau thông qua địa chỉ này

scrape_configs:
  - job_name: "shopdev go prometheus" # Prometheus sẽ lấy metrics của bản thân nó
    scrape_interval: 5s # Định kỳ 5s scrape data từ các target
    static_configs:
      - targets: ["host.docker.internal:9092"] # Đây là địa chỉ của prometheus trong docker-compose.yml
    metrics_path: "/metrics" # Đây là api endPoint mà prometheus sẽ lấy metrics từ đó

  - job_name: "shopdev go api" # Prometheus sẽ lấy metrics của golang app
    scrape_interval: 5s
    static_configs:
      - targets: ["host.docker.internal:8002"] # Đây là địa chỉ của golang app trong docker-compose.yml during production or r.run(":8002") during development

  - job_name: "ecommerce node exporter" # Prometheus sẽ lấy metrics của CPU
    scrape_interval: 5s
    static_configs:
      - targets: ["host.docker.internal:9100"] # Cái này không cần metric path nhé
     #! Chưa add 2 cái này vào docker-compose-dev.yml
  - job_name: "ecommerce mysql exporter" # Prometheus sẽ lấy metrics của mysql
     scrape_interval: 5s
     static_configs:
     - targets: ["host.docker.internal:9104"]
     metrics_path: "/metrics"

  - job_name: "ecommerce redis exporter" # Prometheus sẽ lấy metrics của golang app
     scrape_interval: 5s
     static_configs:
     - targets: ["host.docker.internal:9121"]
     metrics_path: "/metrics"
```

## Step 4: Wrap prometheus to Gin Gonic

To use Prometheus with the Gin Gonic framework, follow these steps:

1. **Import the necessary packages:**

   ```go
   import (
        "github.com/gin-gonic/gin"
        "github.com/prometheus/client_golang/prometheus/promhttp"
   )
   ```

2. **Set up the Gin router and apply the middleware:**

   ```go
   func main() {
   	r := initialize.Run()
   	prometheus.MustRegister(pingCounter) // TEST ONLY - Nhớ xoá: Đăng ký metric
   	r.GET("/ping/200", ping)             // TEST ONLY - Nhớ Xoá
   	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
   	r.Run(":8002")
   }
   ```

With these steps, you have integrated Prometheus with the Gin Gonic framework. You can now monitor your application's metrics by accessing the `/metrics` endpoint.

## Step 5: Run Your Application

Run your Go application.

```Makefile
make dev or go run cmd/server/main.go
```

Your application will be available at `http://localhost:8002` and metrics will be exposed at `http://localhost:8002/metrics`.

## Step 6: Set up graphana to monitor

1. Go to graphana dashboard: **[http://localhost:3002](http://localhost:3002)**

- 3002: _the port you set up in docker-compose-dev.yaml_

2. Set up **DATASOURCE**:

   ![alt text](<Screenshot 2025-01-21 at 16.26.28.png>)

   ![alt text](<Screenshot 2025-01-21 at 16.27.41.png>)

   ![alt text](<Screenshot 2025-01-21 at 16.32.13.png>)

   ![alt text](<Screenshot 2025-01-21 at 16.32.53.png>)

3. Scroll down => Click save test + Click build dashboard

   - Here is dashboard ID for golang: **14061**
     ![alt text](<Screenshot 2025-01-21 at 16.41.02.png>)

4. Import dashboard ID to graphana:

   ![alt text](<Screenshot 2025-01-21 at 16.42.41.png>)

   ![alt text](<Screenshot 2025-01-21 at 16.43.42.png>)

   ![alt text](<Screenshot 2025-01-21 at 16.44.13.png>)

5. Select the right datasrouce and the job that you specify prometheus to scrape which compatible with DASHBOARD ID

   ![alt text](<Screenshot 2025-01-21 at 16.45.40.png>)

_**DASHBOARD** chỉ là là những promql có sẵn mà người khác đã làm trước. Vì thế mình nên chọn dash board phù hợp với các dạng job được specify trong prometheus.yml_

_**Chú ý**: yaml và yml là như nhau nhưng nếu ghi yaml hoặc yml không thông nhất thì coi chừng không chạy được nhé_
