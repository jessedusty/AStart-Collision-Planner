#!/bin/bash

set -e 
ssh ubuntu "cd messages && source /opt/ros/melodic/setup.sh && PATH=\$PATH:/home/jesse/go/bin ./gen_msgs.sh"
rm -rf ../msgs/
scp -r ubuntu:~/messages/msgs ..
(cd .. && go run tools/fix_msgs_imports.go)
