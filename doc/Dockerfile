FROM alpine
LABEL maintainer="tchua"
COPY ./ /apps

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /apps

EXPOSE 8080

CMD ["./dingtalk"]

