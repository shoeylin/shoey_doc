# Windows安裝React環境

## 安裝nvm 此為nodo.js與npm管理工具

### 先移除現有node.js 
1. 檢查有沒有安裝過node 

    node -v

2. 執行 Run npm cache clean --force

3. 安裝移除程序進行刪除 node.js  後重新啟動

4. 尋找以下路徑刪除資料(路徑可能有資料 可能沒有資料)

    C:\Program Files (x86)\Nodejs
    C:\Program Files\Nodejs
    C:\Users\{User}\AppData\Roaming\npm (or %appdata%\npm)
    C:\Users\{User}\AppData\Roaming\npm-cache (or %appdata%\npm-cache)
    C:\Users\{User}\.npmrc (and possibly check for that without the . prefix too)
    C:\Users\{User}\AppData\Local\Temp\npm-*

參考:

https://stackoverflow.com/questions/20711240/how-to-completely-remove-node-js-from-windows

---
## 安裝nvm

下載位置

https://github.com/coreybutler/nvm-windows/releases

下載nvm-setup.zip安裝即可

## 安裝node.js

    nvm install 9.5.0

### 查詢已有版本

    nvm list

### 使用並查詢node.js 與 npm版本

    use 9.5.0

    node -v

    npm -v

---

## 欲升級npm至6.9.0

安裝完成後 node v9.5.0可能只有 npm 5.6.0  

查詢npm -v

將以下幾個檔案進行檔名變更 且執行node指令

    cd %APPDATA%\nvm\v9.5.0           # or whatever version you're using
    mv npm npm-old
    mv npm.cmd npm-old.cmd
    cd node_modules\
    mv npm npm-old
    cd npm-old\bin
    node npm-cli.js i -g npm@latest

### 再次查詢npm -v 

變更完成

## 專案還是跑不動 還需要再更新 安裝更新工具

參考:

https://www.npmjs.com/package/npm

    npm i npm-windows-upgrade

系統管理員身分開啟cmd

    Set-ExecutionPolicy Unrestricted -Scope CurrentUser -Force

    npm-windows-upgrade


專案node_modules刪除重建 

    rm -rf node_modules/ && npm -i

本地啟動專案 

    npm devforwindows


# 創建新專案

安裝create-react-app

    npm install create-react-app -g

create-react-app 專案名稱

