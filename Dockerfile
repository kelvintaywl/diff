FROM golang:1.9-alpine AS build-env

RUN apk --update add make git

WORKDIR /go/src/github.com/kelvintaywl/diff

ADD . /go/src/github.com/kelvintaywl/diff
RUN make init build


FROM alpine:latest
COPY --from=build-env /go/src/github.com/kelvintaywl/diff/diff /diff

EXPOSE 9999

CMD /diff