
# redis 可以用來存儲 string hash list 集合常用來提供數據結構服務

redis可以做存儲(storge) memcached是用來做緩存(cache)

redis可以存儲字符串,鍊表,hash結構 集合 有續集合

## 介紹
redis-benchmark 性能測試

redis-check-aof 檢查aof日誌工具

redis-check-dump 檢查rdb日誌工具

redis-cli 連接用的客戶端

redis-server redis服務進程

## conf配置文件

daemonize yes 可以不中斷運行


docker run redis

當前所有keys pattern 查詢相應的key
keys *

# 插入值
set age 29

set site www.shoey.com

get age 

get site

# 模糊查詢
## 通配符 [] ? *
keys sit[ey] 

keys si?e

keys s* 

## 返回隨機key
randomkey

# key類型
type site

# key存不存在
exists site

# 刪除key
del age

# 改key名
rename site sss

# 檢查新key有沒有重名 重名不改
renamenx site search

# 選擇數據庫 預設為16個 0-15
select 1

# 移動數據庫的key到別的數據庫
move site 1

# 設定有效期限
## 查詢有效期 ttl key
ttl search
-1 永久有效
-2 不存在
返回秒數

## expire key 整形值(秒數)
expire search 10

## pexpire key 整形值(毫秒數)
pexpire search 9000

### 查詢毫秒
pttl search

## persist key 變為永久有效
persist search


# list
lpush character a

rpush character b

rpush character c

lpush character 0

## 查看列表
左0  右-1
lrange character 0 -1

## 彈出列表 lpop key  or  rpop key
lpop character

## 刪除 lrem key count 
lrem answer 1 b 重前面刪

lrem anser -2 a 重後面刪

## 剪接


