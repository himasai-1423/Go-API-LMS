FROM golang:1.21rc2-bookworm

WORKDIR /app/

COPY . /app/

RUN go get 

CMD [ "./" ]