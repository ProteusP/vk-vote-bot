FROM golang:1.24-bookworm

RUN apt-get update && apt-get install -y \
    gcc \
    libssl-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o vk-vote-bot .

CMD ["/app/vk-vote-bot"]
