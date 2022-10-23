FROM golang:1.19 as build
WORKDIR /server
COPY ["go.mod", "go.sum", "./"]
RUN go mod download
COPY ["utils/*.go", "./utils/"]
COPY ["solver/*.go", "./solver/"]
COPY ["api/*.go", "./api/"]
COPY ["*.go", "./"]
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -buildvcs=false -ldflags="-w -s" -o mezzonic-solver

FROM scratch as final
WORKDIR /app
COPY --from=build /server/mezzonic-solver ./
CMD ["./mezzonic-solver"] 