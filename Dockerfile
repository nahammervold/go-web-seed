FROM alpine:latest

ARG GOBIN
ENV GOBIN ${GOBIN:-.}
ENV APP_NAME go-web-seed

# Remember to build with GOOS=linux GOARCH=amd64 go build
COPY $GOBIN/$APP_NAME /app/$APP_NAME
RUN chmod +w /app/$APP_NAME

ENV PORT 3000
EXPOSE 3000

ENTRYPOINT /app/$APP_NAME