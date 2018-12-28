#!/usr/bin/env bash
if [[ $# -ne 3 ]];then echo "
1.redoDate: $1 
2.conf file: $2   
3.sqlTag: $3
 "
exit 1
fi

echo $1
echo $2
echo $3