FROM golang:1.16 as builder

WORKDIR /apps

COPY ./ /apps
RUN export GOPROXY=https://goproxy.cn \
    && go build  -ldflags "-s -w" -o dingtalk  main.go \
    && chmod +x dingtalk

# 第二阶段构建
FROM alpine
LABEL maintainer="tchua"
COPY --from=builder /apps/  /apps
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN echo -e  "http://mirrors.aliyun.com/alpine/v3.4/main\nhttp://mirrors.aliyun.com/alpine/v3.4/community" >  /etc/apk/repositories \
&& apk update && apk add tzdata \
&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Shanghai/Asia" > /etc/timezone \
&& apk del tzdata

WORKDIR /apps

EXPOSE 18080

CMD ["./dingtalk","-f","etc/app.toml"]

