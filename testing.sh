#!/bin/bash

# test cases
t1=$(curl -sS http://<API-IP>:10485/快鲜?height=165&weight=52); o1=$(echo $t1 | jq '.label')
t2=$(curl -sS http://<API-IP>:10485/快鲜?height=105&weight=52); o2=$(echo $t2 | jq '.label')
t3=$(curl -sS http://<API-IP>:10485/快鲜?height=205&weight=52); o3=$(echo $t3 | jq '.label')
t4=$(curl -sS http://<API-IP>:10485/快鲜?height=0&weight=52); o4=$(echo $t4 | jq '.text')
t5=$(curl -sS http://<API-IP>:10485/快鲜?height=abc&weight=52); o5=$(echo $t5 | jq '.text')

# report
if [[ ${o1//\"/} = "normal" ]]; then t1=pass; else t1=NG; fi
if [[ ${o2//\"/} = "overweight" ]]; then t2=pass; else t2=NG; fi
if [[ ${o3//\"/} = "underweight" ]]; then t3=pass; else t3=NG; fi
if [[ ${o4//\"/} = "invalid input" ]]; then t4=pass; else t4=NG; fi
if [[ ${o5//\"/} = "invalid input" ]]; then t5=pass; else t5=NG; fi
printf "normal:$t1\nover:$t2\nunder:$t3\nzero:$t4\nalpha:$t5" > /var/www/html/bmitest.txt
