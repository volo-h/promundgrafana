#!/bin/bash

for ((i=1;i<=100;i++)); do curl "http://localhost:8080/api/v1/balances"; done

#for i in $(seq 5)
#do
#  ping -c 1 google.com
#done
