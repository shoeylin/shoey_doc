# windows 開發 編輯成 linux 執行檔

    GOOS=linux GOARCH=amd64 go build

# linux 背景執行

    nohup ./TheProject  > /dev/null 2>&1 &

# linux crontab 設定(linux 排程設定)

    crontab -e
    */1 * * * * /home/user/TheProject/run.sh

    Linux crontab 不重覆執行, 前一次若尚在執行則跳過
    */1 * * * * flock -xn /home/user/TheProject/run.lock -c '/home/user/TheProject/run.sh'

    crontab -l 查詢

    注意 sudo crontab 與 crontab 內容會看到不一樣的

## run.sh 內容

    cd /home/user/TheProject && ./TheProject

## 給予 run.sh 執行權限

    chmod 755 run.sh


# Windows 開發 建立 windows 執行檔 build.sh

    #!/bin/sh

    GOOS=windows GOARCH=386 go build -o bin/WRA1Batch_win32.exe *.go
    GOOS=windows GOARCH=amd64 go build -o bin/WRA1Batch_win64.exe *.go

    GOOS=linux GOARCH=386 go build -o bin/WRA1Batch_linux32 *.go
    GOOS=linux GOARCH=amd64 go build -o bin/WRA1Batch_linux64 *.go

## Windows 排程 有自己內建的工作排程器
