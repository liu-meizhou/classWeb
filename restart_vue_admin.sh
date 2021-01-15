#!/bin/bash
# 该脚本是给Jenkinsfile运行的, 要运行服务请查看README或者看看对应文件有没有注释
docker restart vue_admin_env
docker restart vue_admin_nginx
