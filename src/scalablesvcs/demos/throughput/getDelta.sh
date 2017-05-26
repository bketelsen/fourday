#!/bin/bash
variance=`awk '/%/ { print $4 }' $1`
sign=`echo $variance | grep -Eo '([+-])'`
plus=+
if [ $sign == "+" ]
then
	percent=`echo $variance | grep -Eo '(100|0|(^0(\.([0-9]{1,})?)|([1-9]{1})?[0-9](\.([0-9]{1,})?)?))'`
	exit -1
fi

