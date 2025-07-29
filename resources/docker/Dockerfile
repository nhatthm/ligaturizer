ARG GO_VERSION=1.24
ARG ALPINE_VERSION=3.19
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

ARG VERSION=dev
ARG REVISION=""
ARG USER=""

WORKDIR /go/src/go.nhat.io/ligaturizer
COPY ../.. .

RUN set -ex \
    && apk --no-cache add ca-certificates bash gcc musl-dev git pkgconfig make python3-dev

RUN VERSION=$VERSION REVISION=$REVISION USER=$USER make build

FROM ghcr.io/nhatthm/ligaturizer:base

ARG VERSION=dev
ARG REVISION=""

COPY --from=builder /go/src/go.nhat.io/ligaturizer/out/ligaturizer /usr/bin/

ENV LIGATURE_FONT_DIR=/opt/fonts/FiraCode3/otf
ENV LIGATURIZER_VERSION=$VERSION

RUN set -ex \
    && apk --no-cache add bash-completion \
    && ligaturizer completion bash > /usr/share/bash-completion/completions/ligaturizer \
    && echo "source /etc/bash/bash_completion.sh" >> ~/.bashrc

WORKDIR /opt/fonts

ENTRYPOINT ["/usr/bin/ligaturizer"]
