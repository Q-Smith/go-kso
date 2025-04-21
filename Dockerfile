# Base image
FROM flant/shell-operator:latest

# Image labels
LABEL maintainer="Q-Smith"

# Build arguments
ARG SHELL_OPTS=${SHELL_OPTS:--eux}

# Environment variables
ENV LANG=en_US.UTF-8 \
    LANGUAGE=en_US:en;

# Prepare directories
RUN set ${SHELL_OPTS}; \
    mkdir -p /hooks;

# Install software
# curl -sSLO https://github.com/mike-engel/jwt-cli/releases/download/6.2.0/jwt-linux.tar.gz;
# tar -xf jwt-linux.tar.gz -C /usr/bin/;
RUN apk add tree wget curl;

# Hooks
COPY ./bin/kso-linux-amd64 /usr/bin/kso
COPY ./hooks/ /hooks/
