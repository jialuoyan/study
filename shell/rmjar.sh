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
testfucntion(){
    echo 'test'
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
	tmpfile=$CURRENT_PATH'/newfile.sql'
	mergefile=`find bootstrap/src/main/resources/db/migration/* `
	for file in $mergefile
    do   
    	cat $CURRENT_PATH'/'$file >> $tmpfile && echo ';' >> $tmpfile
    done
}


case $1 in
	rmtarfile)
	rmtarfile
	;;
    test)
    testfucntion
    ;;
    mergefile)
    mergefile
    ;;
	*)
	usage
	exit 1
	;;

esac
exit 0