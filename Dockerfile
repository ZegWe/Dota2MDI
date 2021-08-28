FROM golang:alpine AS builder

RUN mkdir /build
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 go build -o app

FROM scratch AS final
COPY --from=builder /build/app  .
CMD [ "/app" ]