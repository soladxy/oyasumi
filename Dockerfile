FROM golang:1.18
LABEL maintainer="sola"
WORKDIR /go/src
COPY . .
EXPOSE 8888
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
# 构建时运行
RUN ["/bin/bash", "/go/src/build.sh"]
# docker run运行
CMD ["/bin/bash", "/go/src/output/bootstrap.sh"]