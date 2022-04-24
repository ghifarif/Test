#!/bin/bash

# pull github & update golang service
cd /home/go/src/Test/ && git pull https://github.com/ghifarif/Test
go build bmi.go

# update daemon service
systemctl restart bmi