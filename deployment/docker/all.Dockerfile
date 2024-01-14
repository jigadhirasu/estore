FROM alpine:3.10
RUN apk add --no-cache ca-certificates && \ 
apk add tzdata && \ 
cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime && \ 
rm -rf /var/cache/apk/*

WORKDIR /ding4pro
ENV PATH /ding4pro:$PATH

ADD fs /ding4pro/fs
# ADD loginservice /ding4pro/loginservice
# ADD ownservice /ding4pro/ownservice
# ADD resourceservice /ding4pro/resourceservice

CMD [ "echo", "'ding4pro microservice'" ]
