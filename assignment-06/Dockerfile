FROM golang:1.17.2 as gobuild
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /wallet


FROM alpine
RUN mkdir app
RUN mkdir app/.config
COPY ./.config ./app/.config/
COPY --from=gobuild /wallet ./app/wall
ENV APP_ENV=test
RUN echo ls -la

CMD ["./app/wall"]