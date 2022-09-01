#! /bin/bash

# Build web UI
cd ~/golang_streaming/video_server/web
go install
cp ~/golang/bin/web ~/golang/bin/video_server_web_ui/web
cp -R ~/golang_streaming/video_server/templates ~/golang/bin/video_server_web_ui/web