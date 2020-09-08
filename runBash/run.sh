#!/bin/bash
go run ./main.go run /bin/bash
#running [/bin/bash]

ps
#  PID TTY           TIME CMD
# 8929 ttys000    0:00.04 /Applications/iTerm.app/Contents/MacOS/iTerm2 --server login -fp didi
# 8931 ttys000    0:00.28 -bash
#17531 ttys000    0:00.39 /usr/local/go/bin/go run ./main.go run /bin/bash
#17544 ttys000    0:00.01 /var/folders/r9/35q9g3d56_d9g0v59w9x2l9w0000gn/T/go-build770253821/b001/exe/main
#17545 ttys000    0:00.05 /bin/bash
#78768 ttys002    0:00.05 /Applications/iTerm.app/Contents/MacOS/iTerm2 --server login -fp didi
#78770 ttys002    0:00.32 -bash

exit
#exit