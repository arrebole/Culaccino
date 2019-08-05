export GIN_MODE="release";

(./bin/service.out  > log.txt 2>&1 &);

ps aux |grep service;