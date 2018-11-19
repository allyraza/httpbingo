FROM golang:1.11.2
WORKDIR /go/src/github.com/allyraza/httpbingo/
COPY . .
RUN make linux

FROM scratch
COPY --from=0 /go/src/github.com/allyraza/httpbingo/httpbin-server .
COPY --from=0 /go/src/github.com/allyraza/httpbingo/config.json .
EXPOSE 80
ENTRYPOINT ["/httpbin-server"]
