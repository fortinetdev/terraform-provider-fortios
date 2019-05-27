#!/bin/sh
read -p "Please input FORTIOS_ACCESS_HOSTNAME:" host
read -p "Please input FORTIOS_ACCESS_TOKEN:" token

p=$(dirname "$PWD");  
export GOPATH=${p%/*/*/*/*}"/"
export "FORTIOS_ACCESS_HOSTNAME"=$host
export "FORTIOS_ACCESS_TOKEN"=$token

echo $FORTIOS_ACCESS_HOSTNAME
echo $FORTIOS_ACCESS_TOKEN

make -C ../  testacc
