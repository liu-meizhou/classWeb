#!/bin/sh
# -c: 表示次数，1 为1次 
# -w: 表示deadline, time out的时间，单位为秒，100为100秒。
# -w>=-c

ping -c 3 -w 5 vue_admin_env
echo "ping 了一次前端环境"
while [[ $? == 0 ]]  # 成功
do     
    echo " ping success "
    sleep 1                               
    ping -c 3 -w 5 vue_admin_env                                  
done 

echo " ping fail "

