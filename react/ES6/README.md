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

## const a1 = true && true //回傳true console.log(!!a1); //true

---

## const a2 = false && ( 3 == 4); //回傳false console.log(!!a2); // false

---

## const a3 = 'Cat' && 'Dog'; 

## console.log(a3); //回傳'Dog'

## console.log(!!a3); //true

---

## const a4 = 'Cat' && false; //回傳false

## console.log(!!a4); //false

---

## const a5= "" && false; 

## console.log(a5) //回傳空字串""

## console.log(!!a5); //false

---

## const a6 = 'Cat' && ""; 

## console.log(a6); //回傳空字串""

## console.log(!!a6); //false

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