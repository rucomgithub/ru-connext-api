#=============================================================
#--------------------- build stage ---------------------------
#=============================================================
FROM golang:1.18.2-stretch AS build_stage
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o main .
#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM oracle/instantclient:21 AS final_stage
COPY --from=build_stage /app/main /
COPY environments/config.yaml /environments/
ENTRYPOINT ["/main"]
EXPOSE 8884
