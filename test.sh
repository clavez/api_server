#!/bin/bash
mkdir /home/csl/backup/$1
useradd -d /home/csl/backup/$1 $1
chown $1:$1 /home/csl/backup/$1

setquota -u $1 $2 0 0 0 /home/csl/backup/
