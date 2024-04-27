FROM golang:alpine3.19
LABEL authors="miguoliang"

WORKDIR /code

COPY . /code

RUN go build -o ./build/arch-go ./cmd && \
    chmod +x ./build/arch-go && \
    mv ./build/arch-go /usr/local/bin/arch-go && \
    rm -rf /code

EXPOSE 8080

CMD ["arch-go"]
