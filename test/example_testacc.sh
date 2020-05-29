#!/bin/sh
p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"

# Need super_admin for the token
export "FORTIOS_ACCESS_HOSTNAME"="192.168.52.166"
export "FORTIOS_INSECURE"="true"
export "FORTIOS_ACCESS_TOKEN"="xxxxxxxxxxxxxxxxxxxxxxxxxxx"

echo $FORTIOS_ACCESS_HOSTNAME
echo $FORTIOS_ACCESS_TOKEN


# Need FGT name for the FMG is "myfirewall"
# Don't use the same FGT with previous
# Must keep FORTIOS_FMG_TESTACC not empty
# show login list on FMG: diagnose system admin-session list(25 minutes timeout, max limit: 250)
# unset FORTIOS_FMG_TESTACC when un-TESTACC

export "FORTIOS_FMG_HOSTNAME"="192.168.52.178"
export "FORTIOS_FMG_USERNAME"="admin"
export "FORTIOS_FMG_PASSWORD"="xxxxxxxxxxxxx"
export "FORTIOS_FMG_INSECURE"="true"
export "FORTIOS_FMG_TESTACC"="true"

echo $FORTIOS_FMG_HOSTNAME
echo $FORTIOS_FMG_USERNAME
echo $FORTIOS_FMG_PASSWORD
echo $FORTIOS_FMG_INSECURE
echo $FORTIOS_FMG_TESTACC


make -C ../  cleantestacc
make -C ../  testacc

