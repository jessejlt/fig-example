FROM golang:1.3.1-onbuild

EXPOSE 8080

ENTRYPOINT ["/go/bin/app"]