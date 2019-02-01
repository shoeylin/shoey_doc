# elasticsearch筆記

## 安裝jdk
java -version 檢查版本 太舊移除

去官網下載java
cp jdk-8u201-linux-x64.tar.gz /usr/java/
cd /usr/java/
tar zxvf jdk-8u201-linux-x64.tar.gz

## 安裝elasticsearch

官網下載tar
解壓縮

但有些要求要處理 處理完再執行比較好 當然也可以邊跑邊查邊改
過程都可以去logs裡面elasticsearch.log查看

### 問題一 會有不能用root的問題 
要創建一個使用者 useradd xxx

執行後
./bin/elasticsear -d 
如果你先用root執行一遍 會有一個很靠杯的問題
執行會自動創建文件夾
你root創建的文件都要改成新的使用者可以用才行

不然就要換其他使用者路徑

### 問題二 jvm問題
max virtual memory areas vm.max_map_count [65530] is too low,
由於elasticsearch6.0默認分配jvm空間大小1g，需要改小一點
可能因為用虛擬機 沒給這麼多記憶體
-Xms1g -> -Xms512m
-Xmx1g -> -Xms512m

### 問題三 連線問題 
訪問問題 修改config/elasticsearch.yml中的network.host:0.0.0.0
讓外網任何ip都可以來訪問

#### 使用者資源權限使用問題
vi /etc/security/limits.conf
內加上
* soft nofile 300000
* hard nofile 300000
* soft nproc 102400
* soft memlock unlimited
* hard memlock unlimited

# End of file

limits.conf是 pam_limits.so的 配置文件
可參考
https://www.itread01.com/content/1501938140.html

#### 如果外面還是連不到可能是防火牆沒開
CentOS7開通防火牆9200port
查詢
firewall-cmd --zone=public --list-all
對外開放3000port
firewall-cmd --zone=public --add-port=9200/tcp --permanent
重新讀取設定黨
firewall-cmd --reload
再次查詢
firewall-cmd --zone=public --list-all
通了

### 問題四  如果還是不行 再改
[1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
更換root帳號 sysctl -w vm.max_map_count=262144

持久性的做法是在 /etc/sysctl.conf 文件中修改 vm.max_map_count 參數
echo "vm.max_map_count=262144" > /etc/sysctl.conf
sysctl -p

現在執行應該可以了
./bin/elasticsear -d 

以上教學參考
https://www.jianshu.com/p/04f4d7b4a1d3


### 用docker安裝
docker run -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.5.4


## 路徑說明
/bin        運行elasticsearch 實例和管理插件的一些腳本
/config     配置文件路徑 包含elasticsearch.yml
/data       在節點上每個索引/碎片的數據文件的位置 可以有多個目錄
/lib        elasticsearch使用的庫
/logs       日誌文件夾
/pulgings   已經安裝的插件的存放位置


##插件介紹
head 集群管理工具
可視化插件
https://github.com/mobz/elasticsearch-head

git clone git://github.com/mobz/elasticsearch-head.git

