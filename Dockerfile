FROM golang:1.19

ENV APP_NAME stocktickers
ENV PORT 8080

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}


#RUN go get ./
RUN go build -o ./bin/${APP_NAME} ./src/cmd/server

CMD ./bin/${APP_NAME}

EXPOSE ${PORT}