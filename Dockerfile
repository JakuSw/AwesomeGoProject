FROM ubuntu

WORKDIR /

COPY . .

EXPOSE 8000

ENTRYPOINT ["/controller"]
