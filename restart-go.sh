#!/bin/sh

screen -S goserv -X kill
screen -S goserv -dm go run ./snuggle.go