#!/bin/sh
read -p "Please input FORTIOS_ACCESS_HOSTNAME:" host
read -p "Please input FORTIOS_ACCESS_TOKEN:" token
read -p "Please input FORTIOS_INSECURE:" insecure
read -p "Please input FORTIOS_CA_CABUNDLE:" cabundlefile

p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
export "FORTIOS_ACCESS_HOSTNAME"=$host
export "FORTIOS_ACCESS_TOKEN"=$token
export "FORTIOS_INSECURE"=$insecure
export "FORTIOS_CA_CABUNDLE"=$cabundlefile

echo $FORTIOS_ACCESS_HOSTNAME
echo $FORTIOS_ACCESS_TOKEN
echo $FORTIOS_INSECURE
echo $FORTIOS_CA_CABUNDLE

make -C ../  testacc
