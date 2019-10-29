TypeScript 課程
參考
https://www.youtube.com/watch?v=_wltEKw3BI0&list=PLliocbKHJNwtCfLQu5U3LF_AS-wuP6M7w

https://www.typescriptlang.org/

https://github.com/microsoft/typescript

# nvm 安裝 此為 npm 的管理套件

https://github.com/coreybutler/nvm-windows/releases

## 使用 nvm 安裝 node

    $ nvm install <version>

## 檢視目前安裝的 node 版本

    $ nvm ls

## 切換 node 版本

    $ nvm use <version>

## 檢視 node 版本

    $ node -v

## 安裝最新 node

    nvm install 10.16.2

# tsc 安裝

```bash
$ npm install -g typescript
$ tsc -v
$ tsc --help
```

# npm bin -g 此為 npm 全局目錄

    C:\Program Files\nodejs

## 安裝 ESlint vscode 開發 js 格式化

    設定-> Text Editor -> Formatting -> Format On Save

## Playground 可以看 tx 轉成 js 的

    官網

## tsc 命令 編譯指定位置

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

- number:數值類型
- string:字符串
- boolean:布林類型
- symbol:符號類型標示唯一對象
- any:任意類型
- object:對象類型 (數組 元祖 類 接口 函數)

```js
let myname: string = "koma";
console.log("My name is " + myname);
myname = 100;

let myage: number = 25;
console.log("My age is " + myage);
myage = "aabbb";
```

## any 類型

    不推薦

```js
let myage: any = "aaa";
console.log("My age is " + myage);
myage = 100;
console.log("My age is " + myage);
```

# var let 區別

```js
if (true){
    let myname:string ="koma";
}

console.log(`helo ${myname}.`)
----
let myname:string = "koma";
console.log(`myname is ${myname}.`)

let myname:string = "xiaoma";
console.log(`myname is ${myname}.`)
```

# 常量

```js
const name: type = initial_value;

const dbstr: string = "ip=192.168.1.10";
console.log(dbstr);

dbstr = "new";

//不可改
const dbstr: number[] = [100, 200, 300];
console.log(dbstr);

dbstr = [123, 123, 123];

//可改
const data: number[] = [100, 200, 300];
console.log(data);

data[0] = 10;
data[1] = 20;
data[2] = 30;
console.log(data);

//增加
data[3] = 40;
console.log(data);
```

# 數組使用

```js
//定義
//let name:type[] = initial;
let data: string[] = ["java", "ruby", "kotlin", "dart"];
console.log(data);

console.log(data[0]);

console.log(data[3]);
data[4] = "php";
console.log(data[4]);
data[10] = "Go";
console.log(data);
```

# 多為數組的使用

```js
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
```

# 枚舉類型

```js
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

switch (+member_sex) {
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


console.log(checkSex(member_sex));

function checkSex(sex: Sex) {
  let result: string = "";
  switch (sex) {
    case Sex.MALE:
      result = "男";
      break;
    case Sex.FEMALE:
      result = "女";
      break;
    case Sex.UNKNOWN:
      result = "不明";
      break;
  }

  return result;
}

```

# 聯合類型

可以包含多種類型的變量類型, 不常使用

```js
定義
let data: type1 | type2 | type3 | ....;

let mydata: string | boolean;

mydata = "Helo TS!";
console.log(mydata);

mydata = true;
console.log(mydata);

mydata = 100; //錯誤
```

# null 空檢查

tsc --init
產生 tsconfig.json
"strictNullChecks": true,

```js
let mydata1: string = undefined;
let mydata2: string = null;
let mydata3: string = "";
mydata3 = null;
```

# 第三方庫的使用

## mode-request 安裝

```bash
npm init
#產生package.json

npm install --save request
#node_modules 第三方庫
```

##利用 node-request 庫

```js
const request = require("request");

request("http://192.168.0.1:6688/api/announce?token=abcde", function(
  error,
  response,
  body
) {
  if (error) console.error(error);
  else {
    // console.log(body);
    var body = JSON.parse(body);
    console.log("11111", body.result);
    for (var i = 0; i < body.data.length; i++) {
      var item = body.data[i];
      console.log(item.title);
    }
  }
});
```

## typescript 的使用方法

1.找到 tsd 文件 2.導入 tsd 文件 3.編寫代碼 4.編譯運行

### 找到 tsd 文件(TypeSearch-微軟提供)

https://microsoft.github.io/TypeSearch/

### 導入 tsd 文件

```bash
npm install --save @types/request
```

###　編寫 ts

```js
import request = require("request");

request("http://192.168.0.1:6688/api/announce?token=abcde", function(
  error,
  response,
  body
) {
  if (error) console.error(error);
  else {
    // console.log(body);
    var body = JSON.parse(body);
    console.log("11111", body.result);
    for (var i = 0; i < body.data.length; i++) {
      var item = body.data[i];
      console.log(item.title);
    }
  }
});
```

# 函數的使用

## 定義

```js
function fname(param1:type, ...) : return_type{
  //函數內容
}
```

## 範例

```js
function add(x: number, y: number): number {
  return x + y;
}

console.log(add(10, 5));

function multiply(x: number, y: number): number {
  return x * y;
}

console.log(multiply(10, 5));

let func_add = add;
let func_mul = multiply;
console.log(func_add(10, 5));
console.log(func_mul(10, 5));
```

# 箭頭函數

## 定義

```js
(param:type, ...):return_type => { 函數內容 }
```

## 使用

```js
//原
function add(x: number, y: number): number {
  return x + y;
}
console.log(add(10, 5));
//箭頭
let func_add = (x: number, y: number): number => {
  return x + y;
};

console.log(func_add(100, 5));
```

# 可省略參數

## 定義

```js
function sayHelo(name: string): string {
  if (name === undefined) {
    return "Helo Shoey.";
  } else {
    return "Helo " + name + ".";
  }
}

console.log(sayHelo("trump"));
```

## 定義符號 ?

```js
function sayHelo(name?: string): string {
  if (name === undefined) {
    return "Helo Shoey.";
  } else {
    return "Helo " + name + ".";
  }
}

//   console.log(sayHelo("trump"))
console.log(sayHelo());
```

## 函數內先定義

```js
function sayHelo(name: string = "Shoey"): string {
  return "Helo " + name + ".";
}

console.log(sayHelo("trump"));
console.log(sayHelo());
```

# 參數 72 變

## 參數個數不定

```js
// 普通函數定義
function add1(vals: number[]): number {
  let result = 0;
  for (let val of vals) {
    result += val;
  }
  return result;
}

console.log(add1([1, 2, 3]));

// 參數可變長函數定義
function add2(...vals: number[]): number {
  let result = 0;
  for (let val of vals) {
    result += val;
  }
  return result;
}

console.log(add2(1, 2, 3));
```

# 類的定義與使用

## 定義

```js
class classname{
    property_name1: type;
    property_name2: type;
    //...

    //構造函數
    constructor(param1:type, param2:type, ...){
      //構造函數內容
    }

    //類方法
    method(param1:type, param2:type, ...):return_type{
      //函數內容
    }
    //method2, method3, ...

}
```

## 使用

```js
class Person {
  name: string;
  sex: number; //0:女 1:男

  constructor(name: string, sex: number) {
    this.name = name;
    this.sex = sex;
  }

  sayHelo() {
    console.log(`${this.name},妳好!`);
  }
}

let koma = new Person("小馬", 1);
koma.sayHelo();

console.log(koma.name);
console.log(koma.sex);
```

# 類的訪問限制

## 知識點

- 類的訪問修飾符
  - public: 共有訪問
  - protected: 本類和子類訪問
  - private: 本類訪問

## 實戰演習

```js
class Database {
  dbname: string;

  constructor(dbname: string) {
    this.dbname = dbname;
  }

  // 外部調用
  public showDB(){
    console.log(`數據庫: ${this.dbname}`);
  }

  // 本類和子類訪問
  protected connect(){
    console.log(`${this.dbname}`,"連接中...");
  }

  // 本類調用
  private disconnect(){
    console.log(`${this.dbname}`,"關閉.");
  }
}

let oracle = new Database("Oracle");
oracle.showDB();
// oracle.connect();
// oracle.disconnect();

class PostgreSql extends Database {
  public doIt() {
    super.connect();
    // super.disconnect();
  }
}

let postgres = new PostgreSql("PostgreSql");
postgres.showDB();
postgres.doIt();

```

# 看到不該看的東西

如何看到 class 中定義的 private 內容

## 知識點

- getter
- setter

```js
class Database {
  private dbname: string;
  constructor(dbname: string) {
    this.dbname = dbname;
  }

  get name(): string {
    return this.dbname;
  }

  set name(val: string) {
    this.dbname = val;
  }
}

let db = new Database("oracle is good");
console.log(db.name);

db.name = "mysql is good too."
console.log(db.name);

```

tsc .\test.ts --target es5

# 靜態成員

不用類實例化就能夠訪問使用的屬性和方法

## 知識點

- 靜態成員定義方法
- 關鍵字: static

```js
class Database {
  public static dbname: string;

  constructor() {}

  public static connect() {
    if (this.dbname) console.log(this.dbname + "->連接中...");
    else console.log("數據庫未指定");
  }

  public showDB() {
    console.log("Hi my database.");
  }
}

// let db =new Database(...);

Database.dbname = "dbtype = Oracle;ip=192.168.1.1;uid=admin;pwd=12345678;";
console.log(Database.dbname);
Database.connect();

// Database.showDB();
```

# 名空間

把類分家

## 知識點

- 定義
- 關鍵字: namespace

```js
//定義方法
namespace ns_name{
  export class class_name {}
  export function func_name {}
  export namespace ns_name {}
}
```

## 使用

```js
namespace com.komavideo{
  //導出類
  export class KVUser{
    constructor(){}

    sayHelo(){
      console.log("hi, my dear.")
    }
  }

  //導出函數
  export function showVer(){
    console.log("version is 1.0);
  }

  // 嵌套子名空間
  export namespace util{
    export class MyDatabase{
      private dbname: string;
      constructor(dbname: string){
        this.dbname =dbname;
      }
      showMe(){
        console.log(this.dbname + "is good.");
      }
    }
  }
}

// let user =new com.komavideo.KVUser();
// user.sayHelo();

// com.komavideo.showVer()

let db = new com.komavideo.MyDatabase("oracle");
```

上面怪怪的 XDD

# 類繼承

## 設計

```js
class child_class extends parent_class{
  ...
}
```

## 實例

```js
class Database {
  protected name: string;

  constructor(name: string) {
    this.name = name;
  }

  showInfo() {
    console.log(`您使用的數據庫是, ${this.name}`);
  }
}

class MySql extends Database{}

let mydb = new MySql("mysql");
mydb.showInfo();

```

# 接口

## 定義

```js
interface name {
  //接口定義內容
}
```

## 使用

```js
interface IDatabase {
  connent(): void;
  close(): void;
  exeSql(sql: string): number;
}

class Oracle implements IDatabase {
  connect() {
    console.log("[Oracle]數據庫連接");
  }
  close() {
    console.log("[Oracle]數據庫關閉");
  }
  exeSql(sql: string) {
    console.log("[Oracle]執行成功sql:", sql);
    return 0;
  }
}

let mydb: IDatabase = new Oracle();
mydb.connect();
mydb.exeSql("select * from table1");
mydb.close();

class PostgreSql implements IDatabase {
  connect() {
    console.log("[PostgreSql]數據庫連接");
  }
  close() {
    console.log("[PostgreSql]數據庫關閉");
  }
  exeSql(sql: string) {
    console.log("[PostgreSql]執行成功sql:", sql);
    return 0;
  }
}

mydb = new PostgreSql();
mydb.connect();
mydb.exeSql("select * from table1");
mydb.close();

//接口參數
function doIt(db: IDatabase) {
  db.connect();
  db.exeSql("");
  db.close();
}
let oracle: IDatabase = new Oracle();
let pgsql: IDatabase = new PostgreSql();
doIt(oracle);
doIt(pgsql);
```
