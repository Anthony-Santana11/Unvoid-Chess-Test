FROM golang:1.24.2-alpine

WORKDIR /app
COPY . .

WORKDIR /app/cmd

CMD ["go", "run", "."]