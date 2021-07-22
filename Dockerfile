FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY . .

EXPOSE 8000

ENTRYPOINT ["/controller"]
