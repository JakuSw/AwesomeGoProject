FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY . .

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["/controller"]
