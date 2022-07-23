FROM golang:1.15.2-alpine


# RUN mkdir /app/backend

WORKDIR /usr/src

CMD ["go", "run", "main.go"]
