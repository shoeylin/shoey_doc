# ECMAScript

# 短路求值 真的假的?

## 假的
    0
    -0
    null
    NaN
    undefined
    空白字串("")
    false

## 真的

    不適假的就是真的
    >=1
    <=-1
    任何非空白字串, ex:(" ")
    true

## 無法計算的數學公式,以字串除以數字,會得到NaN(非數字)

    const x='abc'/123;

    console.log(x); //NaN

    console.log(!!x); //false

## 檢查Email格式是否正確,並檢查Email是否重複使用

    const isEmailOK = checkEmailFormat() && checkEmailDuplicate()

    1. checkEmailFormat() == false? return false;

    2. checkEmailFormat() == true? return checkEmailDuplicate();

## 檢查用戶是否有登入, 無論是email登入還是FB登入都可以
    const hasLogin = isLoginByEmail() || isLoginByFacebook()

    1. isLoginByEmail() == true? return true;

    2. isLoginByEmail() == false? return isLoginByFacebook();

## AND運算(&&): 寫法為 a && b, 當a是[假的]時, 回傳a, 否則回傳b

## OR運算(||): 寫法為 a || b, 當a是[真的]時, 回傳a, 否則回傳b

---
    const a1 = true && true //回傳true console.log(!!a1); //true

    const a2 = false && ( 3 == 4); //回傳false console.log(!!a2); // false

    const a3 = 'Cat' && 'Dog'; 
    console.log(a3); //回傳'Dog'
    console.log(!!a3); //true

    const a4 = 'Cat' && false; //回傳false
    console.log(!!a4); //false

    const a5= "" && false; 
    console.log(a5) //回傳空字串""
    console.log(!!a5); //false

    const a6 = 'Cat' && ""; 
    console.log(a6); //回傳空字串""
    console.log(!!a6); //false

---

# 短路求值 : 預設值
//正常寫法

    let value;
    if (!value){
        value = 31;
    }

//短路求值

    let x = value || 31; //31是預設值

---
# 短路求值 : 函式呼叫防呆

//正常寫法

    let email;
    if(user){
        if(user.getEmail){
            email = user.getEmail();
        }
    }

//短路求值

    let email = user && user.getEmail && user.getEmail();

---
# 短路求值 : 減縮程式碼

//正常寫法

    if(!hasLogin){
        checkLogin();
    }

//短路求值

    hasLogin || checkLogin();

---

# 其餘參數

    function sum(a,b,c){
        return a+b+c;
    }

    console.log(sum(1,2,3)); //6

    console.log(sum(1,2)); //NaN

    console.log(sum(1,2,3,4)); //6

---

    function sum(...values){
        let total = 0;
        for(let i=0; i< values.length; i++){
            total += values[i];
        }
        return total;
    }

    funciton (x,y,...valaues){
        //do something
    }

    console.log(sum(1,2,3)); //6

    console.log(sum(1,2)); //3

    console.log(sum(1,2,3,4)); //10

# 陣列展開語法

    const valus = [1,2,3];
    const copiedValues = [...values];

    cost value = [1,2,3];
    let conpiedValues = [];
    for(let i=0; i< values.length; i++){
        copiedValues.push(values[i]);
    }

# 物件

    //空物件
    const emptyObject = {};

    const user = {
        name: 'john Doe'
        id: 1234,
        email: 'johndoe@gmail.com'
    }

    const user = {
        name: 'john Doe',
        id: 1234,
        email: 'johndoe@gamil.com',
        likes:['swimming','hiking','reading'],//陣列
        looks:{'eyecolor':'black','haircolo':'brown'},//物件
        sayHello: function(){console.log('hello!')}//函式
    }

    const userName = user.name;

    function getUserProp(key){
        return user[key];
    }

    let keyToGet = 'email';
    //get email
    consol.log(getUserProp(keyToGet));

    keyToGet ='name'
    //get name
    console.log(getUserProp(keyToGet));

# 類別

## 建構式
    class User{
        constructor(name, email, id){
            this.name = name;
            this.email = email;
            this.id = id;
        }
        sayHello(){
            console.log(this.name + ':Hello!')
        }
    }
    const user = new User('John Doe','johndoe@gmail.com',1234);
    console.log(user.email);
    user.sayHello();

## this
1. 指向全域變數的存放地, 以瀏覽器來說通常式windows物件

2. 在建構式中的this則指向建構出來的新物件

3. 在物件其他函式呼叫, 則會指向這個函式的物件

    class User{
        constructor(name, email, id){
            this.name = name;
            this.email = email;
            this.id = id;
        }
    }

---

    const user = new User('john Doe', 'johndoe@gmail.com', 1234);

    user.name
    user.email
    user.id

    sayHello(){
        console.log(this.name + ':Hello!')
    }

    user.sayHello()

## 物件封裝

    class User{
        constructor(name, email, id){
            this.name = name;
            this.email = email;
            this.id = id;
            this._youdontseeme = true;
        }
    }
    const user = new User('John Doe', 'johndoe@email.com', 1234);
    console.log(user._youdontseeme);//true

##　靜態屬性

    const userSays = User.says;

    // https://jsbin.com/nomexe/edit?js,console

    class User {
        constructor(name, email, id) {
            this.name = name;
            this.email = email;
            this.id = id;
        }
    
        sayHello() {
        console.log(this.name + ': ' + User.says); // 使用靜態屬性      
        }
    }

    User.says = 'Hello!';

    const user = new User('John Doe', 'johndoe@gmail.com', 1234);

    console.log(user.says); // undefined，靜態屬性不是物件的
    console.log(User.says); // 'Hello!'，靜態屬性是類別的

    user.sayHello();

    const jane = new User('Jane Doe', 'janedoe@gmail.com', 1234);
    jane.sayHello(); 

    // 修改類別靜態變數，此時無論是user還是jane，sayHello的結果都會跟著改變
    User.says = 'Hello World!'; 
    user.sayHello(); // Hello World!
    jane.sayHello();  // Hello World!

## ES7靜態屬性寫法(大多數瀏覽器未採用)

    class User{
        static says = 'Hello!';
        constructor(name, email, id){
            this.name = name;
            this.email = email;
            this.id = id;
        }
        sayHello(){
            console.log(this.name + ': ' + User.says); //使用靜態屬性
        }
    }

## 繼承

    // https://jsbin.com/yufiho/edit?js,console

    class User {
        constructor(name, email, id) {
        this.name = name;
        this.email = email;
        this.id = id;
        } 
        sayHello() {
        console.log(this.name + ': Hello!')      
        }
    }

    class SuperUser extends User {
        constructor(name, email, id, address) {
        super(name, email, id);
        this.address = address;
        } 
        sayHello() {
        console.log(this.name + ': Hello, I am super user!');
        }
    }


    const user = new SuperUser('John Doe', 'johndoe@gmail.com', 1234, 'Taipei');

    console.log(user.address);
    user.sayHello(); // John Doe: Hello, I am super user!

## 物件複製

    // https://jsbin.com/foqosup/edit?js,console

    const user = {
        name: 'John Doe',
        id: 1234,
        email: 'johndoe@gmail.com'    
    };

    const user2 = {
        name: 'John Doe',
        id: 1234,
        email: 'johndoe@gmail.com'    
    };

    console.log(user === user2); // false

    let user3 = {
        name: 'John Doe',
        id: 1234,
        email: 'johndoe@gmail.com'    
    };

    let user4 = user3;

    user3.id = 4567; 

    console.log(user4.id); // 4567
    // 物件複製語法如下：
    // const newObj = Object.assign(target, ...sources);

    let user5 = Object.assign({}, user);

    user5.id = 4567; 

    console.log(user.id); // 1234

    const o1 = { a: 1 };
    const o2 = { b: 2 } ;
    const o3 = { c: 3 };

    const obj = Object.assign({}, o1, o2, o3);
    console.log(obj);  // { a: 1, b: 2, c: 3};

    let o4 = { a: 1, b:1, c: 1 };
    let o5 = { b: 2, c: 2 } ;
    let o6 = { c: 3 };

    const obj2 = Object.assign(o4, o5, o6);
    console.log(obj2);  // { a: 1, b: 2, c: 3};
    console.log(o4);  // { a: 1, b: 2, c: 3}; 

    let user6 = {
        name: 'John Doe',
        id: 1234,
        email: 'johndoe@gmail.com' ,
        looks: { 'eyecolor': 'black', 'haircolo': 'brown'}, // 物件
    };

    let user7 = Object.assign({}, user6);

    console.log(user6 === user7); // false

    user6.looks.eyecolor = 'blue';

    console.log(user7.looks.eyecolor); // blue

    let user8 = {
        name: 'John Doe',
        id: 1234,
        email: 'johndoe@gmail.com' ,
        looks: { 'eyecolor': 'black', 'haircolo': 'brown'}, // 物件
        sayHello: function(){ console.log('hello!') } // 函式
    };

    // 將user物件轉成JSON，再解析JSON成為新的物件
    let user9 = JSON.parse(JSON.stringify(user8)); 

    console.log(user8 === user9); // false

    user8.looks.eyecolor = 'blue';

    console.log(user9.looks.eyecolor); // black

    user9.sayHello() // 錯誤，自JSON轉換回來的物件會喪失函式

## 模組系統

### ES6的模組以檔案為單位, 一個檔案亦即一個模組

### ES6的模組透過export以及import語法來輸出與輸入語法

    //user-modules.js
    export const userSay = 'Hello';

    export function sayHello(name){
        console.log(name + ': ' + userSays);
    }

    export const user = {
        name: 'John Doe',
        id: 1234,
        email: 'johndoe@gmail.com'
    };

    export class User {
        constructor(name, email, id){
            this.name = name;
            this.email = email;
            this.id = id;
        }
        sayHello(){
            console.log(this.name + ': Hello!')
        }
    }

    // index.js
    import { userSays, sayHello, user, User} from './user-module';
    //使用模組中輸入資料 方法 物件以及類別
    console.log(userSays);
    sayHello(user.name);
    console.log(user);
    const john = new User('Jone Doe', 'johndoe@gmail.com.', 1234);
    john.sayHello();

---

### default

    //user-modules.js
    export const userSay = 'Hello';

    export function sayHello(name){
        console.log(name + ': ' + userSays);
    }

    export const user = {
        name: 'John Doe',
        id: 1234,
        email: 'johndoe@gmail.com'
    };

    export default class User {
        constructor(name, email, id){
            this.name = name;
            this.email = email;
            this.id = id;
        }
        sayHello(){
            console.log(this.name + ': Hello!')
        }
    }

    // index.js
    import DefaultUser from './user-module';
    import * as userModule from './user-module';

    //使用模組中輸入資料 方法 物件以及類別
    console.log(userSays);
    sayHello(user.name);
    console.log(user);
    const john = new User('Jone Doe', 'johndoe@gmail.com.', 1234);
    john.sayHello();

# 解構赋值

    const user = {
        name: 'john Doe',
        id: 1234,
        email: 'johndoe@gmail.com'
    };

    const name = user.name;
    const id = user.id;
    const email = user.email;

---

    const {name, id, emial} = user; //自user物件提取name, id與 email
    console.log(name);

## 指派式

    const a = 1;

    const b = a;

    const { key: newName} = user;

## 完整語法

    const user = {
        name: 'john Doe',
        id: 1234,
        email: 'johndoe@gmail.com'
    };

    const {name: userName } = user;

    console.log(userName);

---
### 同 (記憶體位置不一樣)

    const user = {
        name: 'john Doe',
        id: 1234,
        email: 'johndoe@gmail.com'
    };

    const { name } = user;

    console.log(name);

## 陣列解構

    const arr = [4, 5, 6, 7, 8];

    const [, ,thirdNum, ...restNums] = arr;

    console.log(thirdNum); //6

    console.log(restNumr); //[7,8]


# 預設值

    function sayHello(name){
        cosole.log(name + ': Hello!');
    }

    sayHello();

    undefined: Hello!   

---

    function getFirstItem(items){
        console.log(items[0]);
    }

    getFirstItem(); //Error!

## 短路求值

    function getFirstItem(items){
        let _items = items || [];
        console.log(_items[0]);
    }

    getFirstItem(); //undefined

## ES6預設值語法

    function getFirstItem(items = []){
        console.log(items[0]);
    }

    getFirstItem(); //undefined

---

    class User{
        constructor(name, email, id, address='no where'){
            this.name = name;
            this.email = email;
            this.id = id;
            this.address = address;
        }
    }
    const user = new User('John Doe', 'johndoe@gmail.com', 1234);

## 解構赋值

    const user = {
        name: 'John Doe',
        id: 1234,
        email: 'johndoe@gmail.com'
    };
    //自user物件提取name, id與email以及地址
    const {name, id, email, address = 'no where'} = user
    console.log(address); // no where

# 箭頭函式

    const saySomething = (name, says) =>
    {console.log(name + ': '+says)};

    saySomething('John Doe', 'Hello world!');

## 同上

    function saySomething(name, says){
        return console.log(name + ': ' + says);
    }

    saySomething('John Doe', 'Hello world!');

---

    // https://jsbin.com/texofel/edit?js,console

    // 基本語法
    const saySomething = (name, says) => {console.log(name + ': ' + says)};
    saySomething('John Doe', 'Hello world from saySomething!');

    // 單一參數時無須括號
    const sayHello = name => { console.log(name + ': Hello from sayHello!')};
    sayHello('John Doe');

    // 單一參數但要給預設值時須要括號
    const sayDefaultHello = (name='John Doe') => { console.log(name + ': Hello from sayDefaultHello!')};
    sayDefaultHello();

    // 無參數時一定要給括號
    const johnDoeSayHello = () => { console.log('John Doe: Hello there!')};
    johnDoeSayHello();

    // 當然也可以return值
    const getUserSays = (user) => { 
        const userSays = user.name + ': ' + user.says;
        return userSays;
    };

    const userSays = getUserSays({
        name: 'John Doe',
        says: 'Hello from getUserSays!'
    });

    console.log(userSays);
    
    // 如果函式內容只有一行且直接return，可以省略大括號

    const getUserId = user => user.id;

    const userId = getUserId({
        name: 'John Doe',
        id: 1234
    });

    console.log('User ID = ' + userId);

---
## this的自動綁定

    // https://jsbin.com/xidicen/edit?js,console

    function Person() {
        // 這個 Person() 建構式定義 `this` 為它自己
        this.age = 0;

        setInterval(function growUp() {
            // growUp() 函式的this會是執行growUp函式的呼叫者
            // 在setInterval的環境下，瀏覽器的window物件會作為呼叫者
            this.age++;
            // 然而，window物件中並沒有age屬性，視作對undefined進行加法，
            // 結果會是NaN
            console.log('New age = ' + this.age);
        }, 1000);
    }

    var p = new Person();

---

### 箭頭函式解決 this的自動綁定問題

    //https://jsbin.com/hoviyabuse/edit?js,console

    function Person() {
        // 這個 Person() 建構式定義 `this` 為它自己
        this.age = 0;
        
        setInterval(() => {
            // 這邊的this指向的就是Person物件的this,也就Person物件的this
            this.age++;
            console.log(this.age);
        }, 100);
    }
    var p = new Person();

# ES6異步運算

## 同步運算

    console.log('Hello world');
    console.log(1+1);

## 異步運算

    setTimeout(()=> console.log('Hello world'), 100);
    console.log(1+1);

    setTimeout(()=> console.log('Hello world'), Math.random()*500);
    setTimeout(()=> console.log(1+1), Math.random()*500);

## Promise規格的崛起

    //Promise類別語法

    const promise = new Promise((resolve, reject) =>{
        if (true){
            //成功時, 呼叫resolve, 並帶入回傳值
            resolve(good);
        }
        //失敗時, 呼叫reject, 並帶入失敗原因
        reject('not good!')
    });

    //結果處理

    promise.then((result) => {
        console.log('Promise success, result = ' + result);
    });

## Promise三態

### 未解(pending)
### 已解, 完成(settled, fullfilled)
### 已解, 失敗(settled, rejected)

    const promise = new Promise((resolve, reject) =>{
        if (true){
            //成功時, 呼叫resolve, 並帶入回傳值
            resolve(good);
        }
        //失敗時, 呼叫reject, 並帶入失敗原因
        reject('not good!')
    });

    //結果處理

    promise.then((result) => {
        console.log('Promise success, result = ' + result);
    } , (reason) => {
        console.log('Promise failed, reason = ' + reason);
    });

### 錯誤處理

    promise.catch((error) => {
        console.log(error);
    });

#### 同

    promise.then(undefined, (error) => {
        console.log(error);
    });

### 連鎖處理

    promise.then((result) => {
        console.log('Promise sucess, result = ' + result);
    }).catch((error) => {
        console.log(error);
    });

---

    promise
        .then((result) => { console.log('Promise sucess, result = ' + result);})
        .catch((error) => { console.log(error);});

# ES6太新 瀏覽器支援

## 填充技術 Polyfill

    const arr = [1, 2, 3, 4, 5];
    const hasTwo = arr.includes(2);
    console.log(hasTwo); //true

---

    function includes(arr, element){
        return arr.indexOf(element)!== -1;
    }

    const arr = [1, 2, 3, 4, 5];
    const hasTwo = includes(arr, 2);
    console.log(hasTwo); //true

## Babel

    const user = {
        name: 'John Doe',
        id: 1234;
        email: 'johndoe@gmail.com'
    };
    //自user物件提取name, id 與 email
    const { name, id, email } = user;
    console.log(name);
    console.log(id);

### 轉

    var user = {
        name: 'John Doe',
        id: 1234;
        email: 'johndoe@gmail.com'
    };
    //自user物件提取name, id 與 email
    var name = user.name,
        id = user.id,
        email = user.email;
    console.log(name);
    console.log(id);

## Webpack工具

打包工具