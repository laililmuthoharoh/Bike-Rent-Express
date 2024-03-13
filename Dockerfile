FROM golang:alpine as build

#build stage
#create folder app
WORKDIR /app

COPY . .

COPY .env .
RUN chmod 777 .env 

RUN go mod download

RUN go build -o bike-rent-express

#final stage
FROM alpine:latest
WORKDIR /app

COPY --from=build /app/bike-rent-express /app/bike-rent-express

ENTRYPOINT ["/app/bike-rent-express"]