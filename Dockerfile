FROM iron/go
WORKDIR /run
COPY . /run
EXPOSE 8080
CMD /run/goforbeer
