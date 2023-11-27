FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /doc_app
COPY go.mod .
COPY  go.sum .
RUN go mod download
COPY . .
RUN go build -o server cmd/job-portal-apis/main.go
#CMD [ "./server" ]


FROM  scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
#RUN apk install ssh  if application needs install we use this command
#RUN sshKey
COPY --from=builder /doc_app/server .
COPY --from=builder /doc_app/private.pem .
COPY --from=builder /doc_app/pubkey.pem .

CMD [ "./server" ]



