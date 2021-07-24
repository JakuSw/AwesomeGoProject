FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY ./controller .

EXPOSE 8000

ENTRYPOINT ["/controller"]
