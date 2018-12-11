FROM golang:1.11.2
WORKDIR /go/src/github.com/allyraza/httpbingo/
COPY . .
RUN make linux

FROM scratch
COPY --from=0 /go/src/github.com/allyraza/httpbingo/httpbingo-server .
COPY --from=0 /go/src/github.com/allyraza/httpbingo/config.json .
EXPOSE 8080
ENTRYPOINT ["/httpbingo-server"]
