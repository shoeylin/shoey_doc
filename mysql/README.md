mysql 雙主架構

https://blog.csdn.net/xiaoyi23000/article/details/80525625

使用 vmware 安裝兩台 linux 裝圖形化介面 vm tools 才好安裝 A

## 安裝 mysql

### 安裝前檢查有無自帶 mysql

rpm -qa | grep mysql

### 有安裝的話先移除

rpm -e mysql 　　// 普通删除模式

rpm -e --nodeps mysql 　　// 强力删除模式

### MariaDB 管理 mysql

#### 使用 yum 安裝 mariadb

將以下資料輸入在一份文件中 (MariaDB.repo)

---

    [mariadb]
    name = MariaDB
    baseurl = http://yum.mariadb.org/10.3/centos7-amd64
    gpgkey = https://yum.mariadb.org/RPM-GPG-KEY-MariaDB
    gpgcheck = 1

---

    複制並粘貼到/etc/yum.repos.d/下的文件中 建議命名文件MariaDB.repo

#### 添加 Repo 之後，運行:(這應該強制重新加載存儲庫）

    sudo yum clean all

#### 然後運行:(這將列出 MariaDB 的可用存儲庫）。

    sudo yum list --showduplicates MariaDB-server

#### 再進行安裝

    sudo yum install MariaDB-server MariaDB-client

#### mariadb 簡單指令

    systemctl start mariadb  #啟動MariaDB

    systemctl stop mariadb  #停止MariaDB

    systemctl restart mariadb  #重啟MariaDB

    systemctl enable mariadb  #設置開機啟動

    systemctl status  mariadb #檢查服務運行狀態

### 驗證 mysql

    mysqladmin --version

### 登入 mysql

    mysql -uroot mysql -p
    預設無密碼

### 建立一個資料庫

    create database giteadb;

### 更改使用者權限

    grant all privileges on *.* to 'root'@'localhost';
    grant all privileges on *.* to 'root'@'%';

    DROP USER IF EXISTS 'user'@'localhost';
    CREATE USER 'user'@'localhost' IDENTIFIED BY 'pass';
    GRANT ALL PRIVILEGES ON database_name.* TO 'user'@'localhost';

    DROP USER IF EXISTS 'shoey'@'%';
    CREATE USER 'shoey'@'%' IDENTIFIED BY 'abcd1234';
    GRANT ALL PRIVILEGES ON *.* TO 'shoey'@'%';

    grant replication slave on *.* to zhang@'192.168.0.104';

### 更改使用者密碼

    UPDATE mysql.user SET password=PASSWORD("your_new_password") WHERE user="root";

    UPDATE mysql.user SET password=PASSWORD("abcd1234") WHERE user="root";

#### 要記得執行(commit)

    flush privileges;

### 開通網路

    vim /etc/sysconfig/network-scripts/ifcfg-ens33

    ONBOOT=yes

    重新啟動網路卡：
    ifdown ens33
    ifup ens33

    ifconfig 查看

### 開通防火牆 ex: 3000port

    查詢
    firewall-cmd --zone=public --list-all
    對外開放3000port
    firewall-cmd --zone=public --add-port=3000/tcp --permanent
    重新讀取設定黨
    firewall-cmd --reload
    再次查詢
    firewall-cmd --zone=public --list-all
    通了

#-------------------------------------------------

# 開始

# master(主庫從庫)

    vim /etc/my.cnf

    [mysqld]
    #GTID:
    server-id=129
    auto-increment-increment=2
    auto-increment-offset=1

    #binlog
    log_bin=/var/log/mysql/mysql-bin.log

# slave(主庫從庫)

    vim /etc/my.cnf

    [mysqld]
    #GTID:
    server-id=131
    auto-increment-increment=2
    auto-increment-offset=2

    #binlog
    log_bin=/var/log/mysql/mysql-bin.log
    relay_log=/var/log/mysql/mysql-relay-bin.log

    #read_only=1 只讀

## 建立 log 位置(主庫從庫)

    mkdir /var/log/mysql

## 修改 mysql 可存取(主庫從庫)

    chown mysql:mysql /var/log/mysql

## mysql 重啟

    systemctl stop mariadb
    systemctl start mariadb

## 登入

    mysql -uroot -p

## 創建同步使用者(主庫從庫)

    create user 'shoey2'@'192.168.253.%'identified by 'abcd1234';

## 同步使用者權限設定(主庫從庫)

    grant replication slave, replication client on *.* to 'shoey2'@'192.168.253.%' identified by 'abcd1234';

## 查詢主庫運行狀態(主庫): log 的名稱,log 從第幾行開始: master_log_file & master_log_pos

    show master status;

## 設定同步機制(從庫) 使用主庫查到的資料填入 master_log_file, master_log_pos

    CHANGE MASTER TO MASTER_HOST='192.168.253.129', MASTER_USER='shoey2',MASTER_PASSWORD='abcd1234', MASTER_LOG_FILE='mysql-bin.000001', MASTER_LOG_POS=328;

## 開啟同步機制(從庫)

    start slave;

### 如果失敗 重設 slave(從庫)

    stop slave;
    reset slave;

    主庫ip 帳號 密碼 主庫log名稱行數
    CHANGE MASTER TO MASTER_HOST='192.168.253.129', MASTER_USER='shoey2',MASTER_PASSWORD='abcd1234', MASTER_LOG_FILE='mysql-bin.000001', MASTER_LOG_POS=328;

    執行
    start slave;

## 查詢 slave 設定是否成功

    show slave status\G

## 查詢從庫列表(主庫)

    show slave hosts;

## 查詢 binlog 文件列表

    show binary logs;

# show slave status\G 錯誤 Table 'mysql.gtid_slave_pos' doesn't exist，

    至mysql 新增一個

    use mysql;

    CREATE TABLE `gtid_slave_pos` (
        `domain_id` int(10) unsigned NOT NULL,
        `sub_id` bigint(20) unsigned NOT NULL,
        `server_id` int(10) unsigned NOT NULL,
        `seq_no` bigint(20) unsigned NOT NULL,
        PRIMARY KEY (`domain_id`,`sub_id`)
        ) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='Replication slave GTID state';

    stop slave;
    start slave;
