# LEVEL 1: NORMAL BUILD => Có thể big image lớn vì chứa những thứ không cần thiết
# Command để build image là docker build . -t <image-name>
# . ở đây là thư mục hiện tại, và docker sẽ tìm Dockerfile trong thư mục này và các file cần thiết khác để build image
# -t flag để đặt tên cho image
FROM golang:alpine 
# This is an default image "golang:alpine" from Docker Hub for golang and you use it as a base image to build your app image

WORKDIR /build
# This is the working directory in your container where you will be working on your app => mọi command về docker sẽ execute tại đây

COPY . .
# Copy all the files from your project "." to the working directory "."in the container

RUN go mod download
# Download all the dependencies of the app such as packages, libraries, etc.

RUN go build -o main ./cmd/server/
# tạo ra binary file (exe file) có tên "main" từ file main.go trong thư mục cmd/server và lưu vào thư mục /build trong container vì chung ta đang ở WORKDIR /build
# -o flag is used to specify the name of the output file => ở đây file binary được tạo ra có tên là "main"

WORKDIR /dist
# Change the working directory to /dist

RUN cp /build/main .
# Copy the binary file from /build to /dist
# ở đây file binary tên main sẽ được copy từ thư mục /build sang thư mục /dist

COPY ./config /dist/config
# Copy the config directory from the local host machine to the /dist/config directory in the container


EXPOSE 8002
# Expose the port 8002 => mở cổng 8002
# Chúng ta sẽ sử dụng cổng này để truy cập vào ứng dụng của mình locahost:8002 và container khác cũng có thể interact với container này thông qua cổng 8008

CMD ["/dist/main"]
# Run the binary file => khi container được start thì file binary main sẽ được chạy



# LEVEL 2: Dùng Multi-stage builds (specify "AS builder") để giảm dung lượng của image => PRODUCTION nhé 
# Có nghĩa là đầu tiên nó sẽ tạo một container để build app, sau đó copy binary file từ container này sang container khác để chạy app
# => Điều này giúp giảm dung lượng của image vì container build app chỉ chứa những thứ cần thiết để build app, không chứa những thứ không cần thiết để chạy app

# FROM golang:alpine AS builder

# WORKDIR /build
# # Đây là working directory trong builder stage và chúng ta sẽ làm việc với app ở đây

# COPY . .

# RUN go mod download
# # Download all the dependencies of the app such as packages, libraries, etc.

# RUN go build -o crm.shopdev.com ./cmd/server/
# # Tạo ra binary file (exe file) có tên "crm.shopdev.com" từ file main.go trong thư mục cmd/server và lưu vào thư mục /build trong container vì chung ta đang ở WORKDIR /build

# FROM scratch 
# # This is an empty image => không chứa bất kỳ thứ gì

# COPY ./config /config
# # Copy the configs folder from the local host machine to the container at /configs

# COPY --from=builder /build/crm.shopdev.com /
# # Copy the binary file from the builder stage to the scratch image
# # "/build/crm.shopdev.com" là đường dẫn của file binary tên crm.shopdev.com trong builder stage
# # "/"" là đường dẫn của thư mục root trong container scratch

# ENTRYPOINT ["/crm.shopdev.com", "config/local.yaml"]
# # Run the binary file => khi container được start thì file binary tên "crm.shopdev.com" sẽ được chạy và truyền vào file "config/local.yaml" làm argument