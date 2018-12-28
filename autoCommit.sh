#!/usr/bin/env bash
# 自动提交git & 邮件发送

#color Function
color_msg(){
    local COLOR=$1
    local MSG=$2
    OFFSET="\033[60G"
    NORMAL="\033[0m"
    case $COLOR in
        red)
            COLOR="\033[1;40;31m"
            ;;
        green)
            COLOR="\033[1;40;32m"
            ;;
        yellow)
            COLOR="\033[1;40;33m"
            ;;
        *)
            COLOR="\033[0m"
            ;;
    esac
    echo -en "$OFFSET [$COLOR $MSG $NORMAL"
    echo     "]"
}

BASE_PATH=$(cd "$(dirname "$0")";pwd)
echo $BASE_PATH 
cd $BASE_PATH && git pull origin develop
if [[ $? -ne 0 ]];then
	color_msg red "git pull failure"
    exit 1
fi
color_msg green "git pull finish"
git add . && git commit -m "git auto commit" && git push origin develop
if [[ $? -ne 0 ]];then
	color_msg green "git push failure"
    exit 1
fi
color_msg green "git push finish"
git log -1 | mail -s "git auto commit successfully!" print23@126.com
if [[ $? -ne 0 ]];then
	color_msg green "send mail failure"
    exit 1
fi
color_msg green "send mail finish"
