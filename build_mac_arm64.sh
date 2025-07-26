#!/bin/bash

function log_info()
{
    datestring=`date "+%Y-%m-%d"`
    timestring=`date "+%H:%M:%S"`
    echo -e "\033[32m[Info ][$datestring $timestring]" "$1 \033[0m"
}

function log_warn()
{
    datestring=`date "+%Y-%m-%d"`
    timestring=`date "+%H:%M:%S"`
    echo -e "\033[33m[Warn ][$datestring $timestring]" "$1 \033[0m"
}

function log_error()
{
    datestring=`date "+%Y-%m-%d"`
    timestring=`date "+%H:%M:%S"`
    echo -e "\033[31m[Error][$datestring $timestring]" "$1 \033[0m"
}

# params:
#   $1 : message
function check_error() {
    if [ $? -eq 0 ];then
        log_info "[OK] $1"
    else
        log_error "[Failed] $1"
        exit 1
    fi
}



# Main
TargetBin="ip-notify"

# configs
BUILD_TIME=`date "+%F_%H:%M:%S"`
COMMIT_SHA1=`git rev-parse HEAD`

log_info "build start"
rm -rf ${TargetBin}
rm -rf ${TargetBin}.tar.gz

log_info "go build ..."

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags \
"-s -w -X github.com/shenhailuanma/local-ip-notify/version.gBuildTime=${BUILD_TIME} \
-X github.com/shenhailuanma/local-ip-notify/version.gCommitID=${COMMIT_SHA1}" \
-o ${TargetBin} main.go

log_info "build end"