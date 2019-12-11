FROM golang:latest as builder
# Add Maintainer Info
LABEL maintainer="Rajendra Gosavi <raje.g.995@gmail.com>"
RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 9595
CMD ["./main"] 
