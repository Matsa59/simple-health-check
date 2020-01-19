#
# Builder image
#
FROM golang:alpine as builder
ADD . /src
RUN cd /src && go build -o healthcheck

#
# Runner image
#
FROM alpine
WORKDIR /app
COPY --from=builder /src/healthcheck /app/
EXPOSE 8080
ENTRYPOINT ./healthcheck