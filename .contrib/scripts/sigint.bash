#!/usr/local/bin/bash
kill -SIGINT $(ps aux | grep -i 'goomba server start' | grep -v grep | awk '{print $2}')