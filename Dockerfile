FROM --platform=$BUILDPLATFORM node:22-alpine AS build
WORKDIR /app
RUN corepack enable && corepack prepare pnpm@9.15.0 --activate
COPY --chown=node:node package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
COPY --chown=node:node . .
RUN pnpm run build

FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS server
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY server/go.mod server/main.go ./
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /server .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=server --chown=nonroot:nonroot /server /server
COPY --from=build --chown=nonroot:nonroot /app/dist /srv
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/server", "-healthcheck"]
ENTRYPOINT ["/server"]
