#!/bin/sh
# 使用 sh is_vue_env_ok.sh 来运行

# -c: 表示次数，1 为1次 
# -w: 表示deadline, time out的时间，单位为秒，100为100秒。
# -w>=-c
ping -q -c 3 -w 5 vue_admin_env
while [ $? -eq 0 ]  # 成功
do  
    sleep 6s                               
    ping -q -c 3 -w 5 vue_admin_env                  
done 


