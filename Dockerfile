#golang yang digunakan
FROM golang:1.20-alpine

#directory didalam image
WORKDIR /app

#copy file gomod dan gosum serta semua packet kedalam directory
COPY go.mod ./
COPY go.sum ./
RUN go mod download

#copy semuanya kedalam directory
COPY . .

#build aplikasi dan define dimana output itu berada, untuk case ini didalam tugas
RUN go build -o /tugas

#define port aplikasi
EXPOSE 8000

CMD ["/tugas"]