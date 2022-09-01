#! /bin/bash

# Build web UI
cd /d/golang_streaming/video_server/web
go install
cd /d/golang_streaming/video_server/api
go install
cd /d/golang_streaming/video_server/scheduler
go install
cd /d/golang_streaming/video_server/streamserver
go install
cp /d/golang/bin/web /d/golang/bin/video_server_web_ui/web
cp -R /d/golang_streaming/video_server/templates /d/golang/bin/video_server_web_ui/