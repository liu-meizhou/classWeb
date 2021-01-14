#!/bin/bash
# 检测的文件
go_mod=go.mod
go_sum=go.sum
# 记录 md5值的文件
go_mod_md5=go.mod.md5
go_sum_md5=go.sum.md5
# 创建新的md5信息
go_mod_md5_new=$(md5sum -b $go_mod | awk '{print $1}'|sed 's/ //g')
go_sum_md5_new=$(md5sum -b $go_sum | awk '{print $1}'|sed 's/ //g')

# 读取旧的md5信息
go_mod_md5_old=$(cat $go_mod_md5|sed 's/ //g')
go_sum_md5_old=$(cat $go_sum_md5|sed 's/ //g')

# 判断文件是否存在
if [ ! -f $go_mod_md5 ] ; then
        echo "go_mod_md5 is not exsit,create md5file......."
        echo $go_mod_md5_new > $go_mod_md5
fi
if [ ! -f $go_sum_md5 ] ; then
        echo "go_sum_md5 is not exsit,create md5file......."
        echo $go_sum_md5_new > $go_sum_md5
fi

echo "go mod md5 old/new"
echo $go_mod_md5_old
echo $go_mod_md5_new

echo "go sum md5 old/new"
echo $go_sum_md5_old
echo $go_sum_md5_new

# 对象对比判断
if [ "$go_mod_md5_new" != "$go_mod_md5_old" ] || [ "$go_sum_md5_new" != "$go_sum_md5_old" ]
then
        echo "md5 is changed"
        echo $go_mod_md5_new > $go_mod_md5
        echo $go_sum_md5_new > $go_sum_md5
        go env -w GO111MODULE=on
        go env -w GOPROXY="https://goproxy.io,direct"
        go mod tidy
        go mod vendor
else
        echo "md5 is not changed"
fi