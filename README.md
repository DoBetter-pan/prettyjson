# prettyjson

A little tool to pretty json 

# Usage:

## Dependencies:
- Golang http://golang.org/

## Install and run:

    mkdir ~/go_codes/src
    cd ~/go_codes/src
    git clone https://github.com/DoBetter-pan/prettyjson.git
    cd prettyjson
    source setenv.sh
    go build
    go install
    add your GOPATH/bin into PATH

    vim 1.json
    :%!prettyjson
    You will see the pretty json in vim.

Note:
You must setup your golang environment first and will change your GOPATH when executing "source setenv.sh".
I tested it in Ubuntu. If you are using other OS, it is the same way as in Ubuntu. Please try.

# Support and contact

If you have any question and advice, please tell me.
