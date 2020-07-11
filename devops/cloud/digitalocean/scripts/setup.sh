#!/usr/bin/env bash

DO_DROPLET_IP=`terraform output plaything_ip`

scp scripts/install.sh root@$DO_DROPLET_IP:/root
echo $DIGITALOCEAN_TOKEN | pbcopy
ssh root@$DO_DROPLET_IP -- echo /root/install.sh
tssh root@$DO_DROPLET_IP
