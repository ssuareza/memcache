#!/bin/bash
# add key flags exptime bytes [noreply]
# value
# key: name of the key
# flags: 
# exptime: expiration time in seconds
# bytes:  length of the data that needs to be stored 

KEY=$1
VALUE=$2
LENGHT=$(echo -n $VALUE|wc -c |awk '{print $NF}')
EXP=900

printf "add $KEY 0 $EXP $LENGHT\r\n$VALUE\r\n" | ncat localhost 11211
