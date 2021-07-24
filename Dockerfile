FROM ubuntu

WORKDIR /

COPY ./controller .

EXPOSE 8000

ENTRYPOINT ["/bin/bash"]
