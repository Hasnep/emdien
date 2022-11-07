FROM emdien AS emdien

FROM tsl0922/ttyd:alpine as ttyd

FROM ghcr.io/charmbracelet/vhs:v0.1.2-devel as vhs

FROM alpine:latest as runner

# Install ttyd
COPY --from=ttyd /usr/bin/ttyd /usr/bin/ttyd

# Install vhs
COPY --from=vhs /usr/bin/vhs /usr/bin/vhs

# Install emdien
COPY --from=emdien /root/.cache/emdien /root/.cache/emdien
COPY --from=emdien /usr/bin/mdn /usr/bin/mdn

# Install Dependencies
RUN apk add --no-cache \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/main \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing \
    ffmpeg \
    chromium \
    bash \
    shadow

# Install fonts
RUN apk add --no-cache \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/main \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing \
    font-inconsolata \
    font-fira-code

# Create user
RUN useradd -u 1976 -U -s /bin/false vhs

WORKDIR /vhs

COPY ./demo/emdien.tape /vhs/emdien.tape

ENTRYPOINT [ "/usr/bin/vhs" ]
