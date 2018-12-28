#!/usr/bin/env bash
# 自动提交git & 邮件发送

BASE_PATH=$(cd "$(dirname "$0")";cd ../;pwd)
echo $BASE_PATH 
cd $BASE_PATH && git pull origin develop
if [[ $? -ne 0 ]];then
    exit 1
fi
echo "git pull finish"
git add . && git commit -m "git auto commit" && git push origin develop
if [[ $? -ne 0 ]];then
    exit 1
fi
echo "git add & push finish"
git log -1 | mail -s "git auto commit successfully!" print23@126.com
if [[ $? -ne 0 ]];then
    exit 1
fi
echo "send mail finish"
