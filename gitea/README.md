# gitea 二進制執行檔 建置流程

使用centOS7 mysql

##CentOS7 設定

###開通網路
vim /etc/sysconfig/network-scripts/ifcfg-eth0

ONBOOT=yes

###開通防火牆 ex: 3000port
查詢
firewall-cmd --zone=public --list-all
對外開放3000port
firewall-cmd --zone=public --add-port=3000/tcp --permanent
重新讀取設定黨
firewall-cmd --reload
再次查詢
firewall-cmd --zone=public --list-all
通了


## 安裝mysql 

###安裝前檢查有無自帶mysql
rpm -qa | grep mysql

###有安裝的話先移除
rpm -e mysql　　// 普通删除模式
rpm -e --nodeps mysql　　// 强力删除模式

###下載mysql
可以去http://repo.mysql.com/看一下版本

wget http://repo.mysql.com/mysql80-community-release-fc29.rpm
https://link.juejin.im/?target=http%3A%2F%2Fdev.mysql.com%2Fget%2Fmysql-community-release-el7-5.noarch.rpm

rpm -ivh mysql80-community-release-fc29.rpm
yum update
yum install mysql-server

###設置權限
chown mysql:mysql -R /var/lib/mysql

###初始化mysql
mysqld --initialize

###啟動mysql
systemctl start mysqld

/etc/init.d/mysqld start 
/etc/init.d/mysqld stop 
/etc/init.d/mysqld restart 
或 
service mysqld start 
service mysqld stop 
service mysqld restart 
或 
service mysql start 
service mysql stop 
service mysql restart

###查看mysql
systemctl status mysqld

###MariaDB管理mysql 沒試過
yum install mariadb-server mariadb 

systemctl start mariadb  #啟動MariaDB
systemctl stop mariadb  #停止MariaDB
systemctl restart mariadb  #重啟MariaDB
systemctl enable mariadb  #設置開機啟動
systemctl status  mariadb #檢查服務運行狀態

###驗證mysql
mysqladmin --version

###登入mysql
mysql -uroot mysql -p

###建立一個資料庫
create database giteadb;

###更改使用者密碼
UPDATE mysql.user SET password=PASSWORD("your_new_password") WHERE user="root";

###新增一個使用者帳號
create user 'testuser'@'localhost' identified by 'password';

ex:
create user 'gitea'@'%' identified by 'gitea';

###設定 testuser 這個帳號可以使用 testdb 這個資料庫
grant all on testdb.* to 'testuser'@'localhost';

ex:
grant all on giteadb.* to 'gitea'@'%';

mariadb查詢帳號全縣
show grants for gitea@'%'; 

##安裝git 

###移除舊的git
sudo yum remove git

###安裝 相依套件
sudo yum install epel-release

###安裝 IUS repository
sudo yum install https://centos7.iuscommunity.org/ius-release.rpm

###安裝git2u
sudo yum install git2u


##安裝gitea

###創建帳號
useradd git
登入
su - git
改密碼
passwd
換帳號這邊會影響gitea裡ssh的連線開頭

###下載gitea
可以先到https://dl.gitea.io/gitea/ 找版本
wget -O gitea https://dl.gitea.io/gitea/master/gitea-master-linux-amd64

###增加執行權限
chmod +x gitea

###執行
./gitea web

##進入安裝gitea頁面
localhost:3000
###創建資料庫相關設定
![windows icon](gitea-install-01.jpg)
![windows icon](gitea-install-02.jpg)
可以改port

###創建第一個帳號為管理員
之後關閉創建帳號後 就用這個管理員新增帳號

###設定檔
可到custom/conf/app.ini更改設定檔
設定gitea安全性
[repository]
DEFAULT_PRIVATE：private：創建新存儲庫時默認為private。[最後，私人，公眾]
DISABLE_HTTP_GIT：true：禁用通過HTTP協議與存儲庫交互的功能。
USE_COMPAT_SSH_URI：true：當使用默認SSH端口時，強制ssh：//克隆url而不是scp-style uri。

[service]
DISABLE_REGISTRATION：true：禁用註冊，之後只有管理員可以為用戶創建帳戶。
REQUIRE_SIGNIN_VIEW：true：啟用此選項以強制用戶登錄以查看任何頁面。

[openid]
ENABLE_OPENID_SIGNIN：false：允許通過OpenID進行身份驗證。
ENABLE_OPENID_SIGNUP：false！DISABLE_REGISTRATION：允許通過OpenID註冊。

參考配置
https://docs.gitea.io/zh-tw/config-cheat-sheet/


#####參考資料
http://www.runoob.com/mysql/mysql-install.html
https://blog.gtwang.org/linux/centos-7-install-mariadb-mysql-server-tutorial/
https://medium.com/verybuy-dev/%E5%9C%A8-centos-7-3-%E4%B8%8A%E5%AE%89%E8%A3%9D-git-2-8052587dd1fd
