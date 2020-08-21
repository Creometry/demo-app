#test comment for webhook - testing deployment
#test new deployment - one more test test tests
FROM golang:alpine
WORKDIR https://github.com/Creometry/demo-app/
COPY . .
RUN CGO_ENABLED=0 go build -v -o app

FROM scratch
COPY --from=0 https://github.com/Creometry/demo-app/app .
EXPOSE 8080
ENTRYPOINT ["/app"]
