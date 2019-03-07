#!/bin/bash

. /etc/profile

BASE_PATH=$(cd "$(dirname "$0")";pwd)

# common
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

# user process start
start(){
    stop

    . /etc/profile
    PHPLIB=`which php71`

	if [ -d "${BASE_PATH}/bin" ] ; then
	    for startup in `ls ${BASE_PATH}/bin`; do
	        chmod +x ${BASE_PATH}/bin/${startup}
	        $PHPLIB ${BASE_PATH}/bin/${startup} start &
	        # wait 2s
	        sleep 2
	    done
	    color_msg green "start sucess"
	fi
}

# user process stop
stop(){
    count=`ps -ef | grep php-kafka-dc.php | grep start | wc -l` 
    if [ 0 == $count ]; then
        color_msg green "nothing to stop"
    else
        kill `ps -ef | grep php-kafka-dc.php | grep start | awk '{print $2}'`
        if [[ $? -ne 0 ]];then
            color_msg red "stop failure"
            exit 1
        else
            color_msg green "stop success"
        fi
    fi
}

# check
checkStatus(){
	count=`ps -ef | grep bin/cmal | grep start | wc -l`
    if [ 0 == $count ]; then
        color_msg green "stoped ..."
    else
        color_msg green "running ..."
    fi
}

testFunction(){
    # PHP_SYMLINK=$(readlink -f /usr/local/bin/php71)
    # PHP_BASEDIR=${PHP_SYMLINK:-/usr/local/php71}
    # echo $PHP_SYMLINK
    # echo $PHP_BASEDIR
    . /etc/profile
    PHPLIB=`which php71`
    echo $PHPLIB

}



case $1 in
	start)
	start
	;;
    stop)
    stop
    ;;
    check)
    checkStatus
    ;;
    test)
    testFunction
    ;;
	*)
	echo "Usage: $0 {start| stop| check}"
	exit 1
	;;

esac
exit 0