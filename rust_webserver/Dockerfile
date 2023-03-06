FROM rust:alpine as builder

WORKDIR /usr/src/myapp
COPY . .
RUN cargo install --path .


FROM alpine

RUN apk update && rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/local/cargo/bin/rust_webserver /usr/local/bin/rust_webserver

EXPOSE 8080

RUN adduser --disabled-password --no-create-home john-doe

ENTRYPOINT ["rust_webserver"]

USER john-doe
