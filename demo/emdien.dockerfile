FROM golang:1.19-alpine AS emdien

COPY ./go.mod /src/
COPY ./go.sum /src/

COPY ./main.go /src/
COPY emdien /src/emdien/

WORKDIR /src/

RUN go get 
RUN go build -o build/mdn
RUN build/mdn --update

RUN cp /src/build/mdn /usr/bin/mdn

ENTRYPOINT [ "/usr/bin/mdn" ]
