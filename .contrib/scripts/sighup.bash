#!/usr/local/bin/bash
kill -SIGHUP $(ps aux | grep -i 'goomba server start' | grep -v grep | awk '{print $2}')