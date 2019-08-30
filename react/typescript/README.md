TypeScript課程
參考
https://www.youtube.com/watch?v=_wltEKw3BI0&list=PLliocbKHJNwtCfLQu5U3LF_AS-wuP6M7w

https://www.typescriptlang.org/

https://github.com/microsoft/typescript

# nvm安裝 此為npm的管理套件

https://github.com/coreybutler/nvm-windows/releases

## 使用 nvm 安裝 node
   
    $ nvm install <version>

## 檢視目前安裝的 node 版本
    
    $ nvm ls

## 切換 node 版本

    $ nvm use <version>

## 檢視 node 版本

    $ node -v

## 安裝最新node

    nvm install 10.16.2

# tsc安裝

~~~bash
$ npm install -g typescript
$ tsc -v
$ tsc --help
~~~

# npm bin -g 此為npm全局目錄

    C:\Program Files\nodejs

## 安裝ESlint  vscode 開發 js格式化

    設定-> Text Editor -> Formatting -> Format On Save

## Playground 可以看tx轉成js的

    官網

## tsc命令 編譯指定位置

    tsc --outDir dist ./src/helloworld.ts

## tsconfig.json

    tsc --init

#@ 自動編譯

    tsc -w

# 變量聲明要注意

## 變量使用前要定義

    // let/var 變量名:變量類型 = 默認值
    let myname:string = "koma";
    console.log("My name is " + myname);

    let myage:number = 25;
    console.log("My age is " + myage);

## 變量類型

+ number:數值類型
+ string:字符串
+ boolean:布林類型
+ symbol:符號類型標示唯一對象
+ any:任意類型
+ object:對象類型 (數組 元祖 類 接口 函數)

~~~js
let myname:string ="koma"
console.log("My name is " + myname);
myname=100;

let myage:number = 25;
console.log("My age is " + myage);
myage="aabbb"
~~~

## any類型
    不推薦

~~~js
let myage:any = "aaa";
console.log("My age is " + myage);
myage=100
console.log("My age is " + myage);
~~~

# var let 區別
~~~js
if (true){
    let myname:string ="koma";
}

console.log(`helo ${myname}.`)
----
let myname:string = "koma";
console.log(`myname is ${myname}.`)

let myname:string = "xiaoma";
console.log(`myname is ${myname}.`)
~~~

# 常量
~~~js
const name:type =initial_value;

const dbstr:string = "ip=192.168.1.10"
console.log(dbstr)

dbstr="new"

//不可改
const dbstr:number[] = [100,200,300]
console.log(dbstr)

dbstr=[123,123,123]

//可改
const data:number[] = [100,200,300]
console.log(data)

data[0]=10;
data[1]=20;
data[2]=30;
console.log(data)

//增加
data[3]=40;
console.log(data)
~~~

# 數組使用

~~~js
//定義
//let name:type[] = initial;
let data:string[] = ['java','ruby','kotlin','dart'];
console.log(data)

console.log(data[0]);

console.log(data[3]);
data[4]='php'
console.log(data[4]);
data[10] = 'Go';
console.log(data);
~~~

# 多為數組的使用

~~~js
// let name:type[][]= [[],[],[],...];

let data: number[][] = [
  [1, 2, 3, 4, 5],
  [10, 20, 30, 40, 50],
  [100, 200, 300, 400, 500]
];
console.log(data);
console.log(data[0]);
console.log(data[1]);
console.log(data[2]);

console.log(data[0][0]);
console.log(data[0][1]);
~~~

# 枚舉類型


~~~js
//定義方法
// enum name {name1,name2,name3, ...}

enum Sex {
  MALE,
  FEMALE,
  UNKNOWN
}

let member_sex: Sex = Sex.FEMALE;
console.log(member_sex);
console.log(Sex[member_sex]);
//console.log(checkSex[member_sex]);

switch (member_sex) {
  case Sex.MALE:
    console.log("男");
    break;
  case Sex.FEMALE:
    console.log("女");
    break;
  case Sex.UNKNOWN:
    console.log("不明");
    break;
}
~~~