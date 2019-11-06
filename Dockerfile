FROM  golang:1.13.4 as builder

LABEL maintainer="Gabriel Oliveira <barbosa.olivera1@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Testes
WORKDIR dominio
RUN go test -v
RUN go test -cover


# Fim Testes Inicio BUILD.
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#BUILD WEB
FROM node:10-alpine as builderweb

# Guardar bibliotecas em camadas diferentes evita downloads desnecessarios.
COPY /web/kronos-app/package.json ./
COPY /web/kronos-app/package-lock.json ./
RUN npm ci && mkdir /kronos-app && mv ./node_modules ./kronos-app
WORKDIR /kronos-app
COPY /web/kronos-app .

#Build Web
RUN npm run ng build -- --prod --output-path=dist

######## Start a new stage from scratch #######
FROM alpine:3.10.3
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
WORKDIR /root/static/
COPY  --from=builderweb /kronos-app/dist .
WORKDIR /root/
EXPOSE 80
CMD ["./main"]