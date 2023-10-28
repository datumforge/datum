FROM golang:1.21 as builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/datum -a -ldflags '-linkmode external -extldflags "-static"' .

FROM gcr.io/distroless/static:nonroot

# `nonroot` coming from distroless
USER 65532:65532

# Copy the binary that goreleaser built
COPY --from=builder /go/bin/datum /bin/datum

# Run the web service on container startup.
ENTRYPOINT [ "/bin/datum" ]
CMD ["serve","--debug","--pretty","--dev"]