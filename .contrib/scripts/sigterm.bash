#!/usr/local/bin/bash
kill -SIGTERM $(ps aux | grep -i 'goomba server start' | grep -v grep | awk '{print $2}')