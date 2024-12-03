from golang:1.22 as base

WORKDIR /app

copy go.mod .

RUN go mod download

copy . .

RUN go build -o main .

#distroless image

from gcr.io/distroless/base

copy --from=base /app/main .

copy --from=base /app/static ./static

expose 8080

cmd ["./main"]





