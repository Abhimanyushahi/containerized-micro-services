FROM golang AS builder

WORKDIR /app  
COPY main.go .
RUN go env -w GO111MODULE=off
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /abhi
COPY --from=builder /app .
RUN ls -la
EXPOSE 81
CMD ["./main"] 
