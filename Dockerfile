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

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASS=password123
ENV DB_NAME=db_bike_rent_express
ENV MAX_IDLE=1
ENV MAX_CONN=2
ENV MAX_LIFE_TIME=1h

ENV PORT=8080
ENV LOG_MODE=1

ENTRYPOINT ["/app/bike-rent-express"]