#FROM golang:1.15

FROM ubuntu
RUN  sed -i "s@/archive.ubuntu.com/@/mirrors.163.com/@g" /etc/apt/sources.list \
        && rm -rf /var/lib/apt/lists/* \
        && apt-get update --fix-missing -o Acquire::http::No-Cache=True \
    && apt-get install -y gcc libc6-dev make wget

ENV COMPOSE_CONVERT_WINDOWS_PATHS=true