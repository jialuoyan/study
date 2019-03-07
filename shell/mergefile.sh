#!/bin/bash
# 删除掉部分不需要的tar文件,交由镜像提交
# 此处作为shell 模版

CURRENT_PATH=$(cd "$(dirname "$0")"; pwd)

#######useage#######
usage()
{
	echo -e "\n\tUsage:"
	echo -e "\t-----------------------------------------"
        echo -e "\t| $0 < rmtarfile | test >"
	echo -e "\t-----------------------------------------"
	cat<<EOF
            | rmtarfile   -- delete tar tmp file ;
            | mergefile   -- merge sql file 
            | test        -- test ;
EOF
	echo -e "\t-----------------------------------------"
    echo -e "\n\n"
}

#Common Function
color_msg(){
    local COLOR=$1
    local MSG=$2
    local OFFSET=$3

    NORMAL="\033[0m"
    case $OFFSET in
        left)
            OFFSET=""
            ;;
        right)
            OFFSET="\033[60G"
            ;;
        *)
            OFFSET=""
            ;;
    esac
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
checkEmpty() {
    if [ -z $1 ]; then
        color_msg red "$0 cannot be empty"
        exit 1
    else
        echo "cat start file name:"
        color_msg green $1 right
    fi
}

testfucntion(){
    # echo 'test'
}

rmtarfile(){
    targetfile=`git status | grep "target"`
    for file in $targetfile
    do   
    	rm -rf $CURRENT_PATH'/'$file
    	echo rm -rf $CURRENT_PATH'/'$file
    done
    echo 'rm done'
}


mergefile(){
	checkEmpty $1

    file_start=$1
    savetmpfile=$CURRENT_PATH'/savetmpfile.sql'

    cd 'bootstrap/src/main/resources/db/migration'

    if [ -f "$savetmpfile" ];then
        echo "delete file:"
        color_msg green $savetmpfile right
        rm $savetmpfile
    fi

    mergefile=`ls -l | awk '{print $9}'`

    flag=false
    for file in $mergefile
    do   
        # echo $file_start" : "$file
        if [[ $flag == "true" ]]; then
            echo $file
            cat $file >> $savetmpfile && echo ';' >> $savetmpfile
            endfilename=$file
        else
            if [ "$file" = "$file_start" ]; then
                startfilename=$file
                flag=true
                color_msg green "###########start############" left
            fi
        fi
    done
    color_msg green "###########finish############" left
    if [ ! -f "$savetmpfile" ]; then
        color_msg red " $1 file is not exist, please check" left
        echo "example : rmjar.sh mergefile V20181107.5__t_driver_settlement.sql"
    else
        color_msg green "${startfilename}_${endfilename}" right
        mv $savetmpfile $CURRENT_PATH"/${startfilename}_${endfilename}"
    fi
}


case $1 in
	rmtarfile)
	rmtarfile
	;;
    test)
    testfucntion
    ;;
    mergefile)
    mergefile $2
    ;;
	*)
	usage
	exit 1
	;;

esac
exit 0