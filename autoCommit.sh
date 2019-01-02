#!/usr/bin/env bash
# 自动提交git & 邮件发送
# bash autoCommit.sh "提交注释"

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
gitpull(){
	cd $BASE_PATH && git pull origin develop
}
gitaddpush(){
	git add . && git commit -m "$1" && git push origin develop
}
sendmail(){
	git log -1 | mail -s "git auto commit" $1
}
check(){
	local SUCCESSMSG=$1
	local ERRORMSG=$2
	if [[ $? -ne 0 ]];then
		color_msg red $ERRORMSG
	    exit 1
	fi
	color_msg green $SUCCESSMSG
}

if [[ $# -ne 1 ]];then echo "
1.commit content: $1 
 "
exit 1
fi

BASE_PATH=$(cd "$(dirname "$0")";pwd)

echo $BASE_PATH 
gitpull
check "pull-finish" "pull-failure"
gitaddpush "$1"
check "push-finish" "push-failure"
sendmail "print23@126.com"
check "mail-send-finish" "mail-send-failure"

