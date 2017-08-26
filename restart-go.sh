#!/bin/sh

screen -S goserv -X kill
chmod +x snuggle-world
screen -S goserv -dm ./snuggle-world -t $SUNBOT_KEY