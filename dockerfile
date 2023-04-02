FROM golang:1.20-alpine  as build

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . /app

##buat executeable
RUN go build -o BeliKuy .
 .

##jalankan executeable
CMD ["/app/BeliKuy"]
