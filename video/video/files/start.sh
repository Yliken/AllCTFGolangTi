#!/bin/bash
set -e

# 启动 MySQL 服务
service mysql start

# 初始化数据库和导入数据
mysql -uroot -e "ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root'; FLUSH PRIVILEGES;"
mysql -uroot -proot -e "CREATE DATABASE IF NOT EXISTS video;"

for f in /app/sql/*.sql; do
  mysql -uroot -proot video < "$f"
done

# 启动 Go 程序
/app/main
