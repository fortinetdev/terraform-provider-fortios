#!/bin/sh
p=$(dirname "$PWD");  
export GOPATH=${p%/*/*/*/*}"/"
export "FORTIOS_ACCESS_HOSTNAME"="50.18.197.179"
export "FORTIOS_ACCESS_TOKEN"="6mQybp1cpxy4y778q4gc44s8jHf7mn"

echo $FORTIOS_ACCESS_HOSTNAME
echo $FORTIOS_ACCESS_TOKEN

make -C ../  cleantestacc
make -C ../  testacc
