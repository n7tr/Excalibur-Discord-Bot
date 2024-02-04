# For deployment on railway.app
FROM golang:latest

WORKDIR /Inferno

COPY . .

RUN go build Inferno

CMD [ "./Inferno" ]