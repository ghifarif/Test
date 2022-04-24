#!/bin/bash

# prepare base working directory
mkdir -p /home/go/src/
chmod -R 755 /home/go/src/

# pull github & setup golang service
cd /home/go/src/ && git clone https://github.com/ghifarif/Test
go build bmi.go

# setup daemon service
cat <<'EOF' >> /etc/systemd/system/bmi.service
[Unit]
Description=bmi
After=network.target

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/home/go/src/Test
ExecStart=/home/go/src/Test/bmi
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload && systemctl start bmi && systemctl enable bmi
