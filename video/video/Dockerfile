FROM ubuntu:22.04

# 安装必要依赖
RUN apt-get update && \
    apt-get install -y ca-certificates mysql-server && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 复制二进制和静态资源
COPY files/main /app/main
COPY static /app/static
COPY sql /app/sql

COPY files/start.sh /app/start.sh
RUN chmod +x /app/start.sh
RUN chmod +x /app/main

ENV DASFLAG DASCTF{test12345}

EXPOSE 9090

CMD ["/app/start.sh"]