#!/bin/sh
p=$(dirname "$PWD");  
export GOPATH=${p%/*/*/*/*}"/"
export "FORTIOS_ACCESS_HOSTNAME"="192.168.52.177"
export "FORTIOS_ACCESS_TOKEN"="tfd351yqtd4Nf719dtHk50Gfgs9Hbs"

echo $FORTIOS_ACCESS_HOSTNAME
echo $FORTIOS_ACCESS_TOKEN

make -C ../  testacc
