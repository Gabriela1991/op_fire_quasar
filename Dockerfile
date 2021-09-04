FROM registrysecaas.azurecr.io/secaas/golang-dev:1.12-latest

WORKDIR /app
USER root

COPY . .
RUN pwd
RUN go build /app/src/cmd/main.go
RUN ls -la /root

# ENTRYPOINT ["go", "run", "/go/src/main.go"]
#ENTRYPOINT ["/app/main"]
CMD ["tail -f /dev/null"]
