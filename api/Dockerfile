FROM golang:1.23-bullseye

ARG DATABASE_URL
ARG ENV

RUN mkdir -p /app
COPY . /app
WORKDIR /app

# install atlas - db migrations
RUN curl -sSf https://atlasgo.sh | sh

RUN go mod tidy
RUN go mod vendor
RUN go build -o out .

RUN ./scripts/apply-migrations.sh
ENTRYPOINT /app/out
