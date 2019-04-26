# Redux 絕對真理資料流 

    UI = f(store);

# Reducer (歸納)

## 鐵律(之一)

    任何異動都必須要回傳新的state, 沒有異動過的資料回傳原本的state

# 物件比較 

 javascript物件為記憶體位置比較

## immutable Data

    var oldString = 'Hello';
    newString = oldString.concat('World');
    console.log(oldString);
    console.log(newString);

    1+1=2

## 鐵律(之二)

    Reducer函式必須要式 "純" 函式

    函式的產出結果, 必然與傳入的參數相關

### 純函式

    //加法
    function add(a, b){
        return a+b;
    }

### 非純函式

    //加法 加上目前幾號
    function addDate(a, b){
        return a + b + new Date().getDate();
    }

## 資料可預測性


# Redux三部曲之三:建立Store與資料流

## 使用Redux函式庫建立Store

    import{ createStore } from 'redux' // 從redux函式庫中匯入createStore函式
    import todoReducer from './reduces' // 匯入我們自己定義的reducers
    let store = createStore(todoApp) // 以reducers建立Store

## 透過Store發動Action

    import{ createStore } from 'redux' // 從redux函式庫匯入createStore函式
    import todoReducer from './reduces' // 匯入我們自己定義的reducers
    import{ addTodo, removeTodo } from './actions'; //匯入我們自己定義的actions
    let store = createStore(todoApp) // 以reducers建立Store

    store.dispatch(addTodo('買雞蛋'));

## Redux名詞總覽

    State->定義 UI ->觸發 Action -> 送出 Reducer -> 更新 Store -> 包含 State

# Redux只是概念 函式庫

    npm install --save react-redux

    import { Provider, connect } from 'react-redux';


## Connect語法(比喻)

    <Connect>
        <TodoList
            todos=(store.state)
            remove={(id) => store.dispatch(removeTodo(id))}
        />
    </Connect>

## 何謂HOC(High Order Component)

    是用來對 '既有元件' 增加 '額外功能' 的函式
    該函式接收一個既有元件, 然後 '回傳' 另一個包覆既有元件的新元件
    如果HOC的資料有所異動, 函式會重新執行, 並更新元件
    包含可重複使用的常用邏輯, 使元件架構更佳,更簡潔
    HOC架構可以不慣的包覆元件, 無限次數也不限HOC數量

    例以withRouter以及connect HOC函式包覆HomeView元件:
    export default withRouter(
        connect(mapStateToProps, mapActionToProps,null,{
            withRef: true
        })(HomeView)
    );

## 以上就是Decorator Pattern(修飾模式)

    //original SimpleWindow
    SimpleWindow sw = new SimpleWindow();
    printInfo(sw);
    //HorizontalScrollBar mixed Window
    HorizontalScrollBar hbw = new HorizontalScrolBar(sw);
    printInfo(hbw);
    //VerticalScrollBar mixed Window
    VerticalScrollBar vbw = new VerticalScrollBar(hbw);
    printInfo(vbw)

    discription: simple window
    discription: simple window,including horizontal scrollbars
    discription: simple window,including horizontal scrollbars,including vertical scrollbars

# 為什麼

    為什麼需要Redux
    Store, Reducer, Action 它們是什麼
    組件如何取得Store的資料
    怎麼改變Store的資料
    掌握資料流向的小工具
    手把手讓React Todolist變成React Redux Todolist

    安裝npm install redux
    安裝npm install react-redux


# google chrome商店安裝 redux devtools 

https://github.com/zalmoxisus/redux-devtools-extension

把
    
    window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()

複製到 App.js 內

ex: 

    let store = createStore(todoReducer,
     window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
     )

chrome F12 可以看到store狀態

# 重點整理

    Store 其實是由Reducer產生
    Action定義行為, 而行為須透過Dispatch觸發
    Store的State,必須是Immutable
    透過conncet讓組件取得Store裡的State
    Redux的好處

# React與後端API呼叫 fetch()


## XHR 

    const xhr = new XMLHttpRequest();
    xhr.onload = function(){
        const data = JSON.parse(this.responseText);
        console.log(data)
    }
    xhr.onerror = function(err){
        console.log(err);
    }
    xhr.open('get','/api/todo',true);
    xhr.send();

## fetch()

    fetch('/api/todo' , {method:'get'})
    .then(function(response){
        //處理response
    }).catch(function(err){
        console.log(err);
    });

# fetch四大天王

## 1.Request物件

### 可重複使用的Request物件

    const getReq = new Request(API_URL, {method: 'GET'});

    fetch(getReq).then(function(response){
        //處理response
    }).catch(function(err){
        console.log(err)
    })

### 可當作範本的Request物件

    const postReq = new Request(getReq, {method: 'POST'});

### Request物件參數一覽

    url: 要求網址
    method: HTTP Request Method, 支援GET, POST, PUT, DELET, HEAD
    headers: 詳見Header物件解釋
    mode: 跨網欲存取資源相關設定, 常用的參數如下:
        cors: 允許跨網域資料存取
        same-origin: 只限定與網頁應用程式相同網域的資源存取

    credentials: 是否傳送cookie至後端, 常用的參數如下:
        omit: 不傳送任何cookie
        same-origin: 只限定與網頁應用程式相同網域的後端呼叫時,傳送cookie
        include: 無論如何都傳送cookie

    redirect: 當收到伺服器回傳HTTP狀態301或302等轉只代碼時,是否要跟隨轉址:
        follow: 自動跟隨轉址
        error: 不跟隨轉址,視作出錯
        manual: 多數瀏覽器的預測值,開發者必須自行決定是否要跟隨

    cache: 快取機制該採用何者, 可用的選項有:
        default: 由瀏覽器原本的快取機制決定是否快取,也就是:
            1.檢查快取中的結果是否依然有效,若有效則回傳快取中的結果
            2.若已經無效,向後端伺服器發送檢查用的要求,若伺服器說資源沒有異動,則回傳快取中的值,否則該資源會重新從伺服器下載,並更新快取中的值
            3.若快取無值,就直接從伺服器下載該資源

        no-store: 瀏覽器將不檢查自身快取是否已經友值,收到結果後'不會'更新自身快取中的值

        reload: 瀏覽器將不檢查自身快取是否已經有值,收到結果後'會'更新自身快取中的值

        no-cache: 
            1.無論快取中的值是否有效,都會向伺服器檢查資源有效性,若伺服器說資源沒有異動,則回傳快取中的值,否則該資源會重新從伺服器下載,並更新快取中的值
            2.若快取中無值,就直接從伺服器下載該資源

        force-cache:
            1.無論快取中的值是否有效,都回傳快取中的值
            2.若快取中無值,就直接從伺服器下載該資源

        only-if-cached:
            1.無論快取中的值是否有效,都回傳快取中的值
            2.若快取中無值,回傳錯誤
            3.只能與same-origin mode並用

    body:請看body物件詳解

## 2.Header物件

    const headerKeyValues = { 'Content-Type':'image/jpeg','Accept-Charset':'utf-8'};

    const headerObj = new Headers(headerKeyValues);

    const reqWithMyHeader = new Request(API_URL,{headers: headerObj});

## 3.Body物件

    const formData = { name: 'John Doe', sex: 'Male'};

    fetch(API_URL,{
        method: 'post',
        body: JSOB.stringify(formData);
    });

## 4.Response物件

### JSON Response範例

    fetch(API_URL).then(function(response){

        if (response.ok){

            //將Response物件直接轉乘JSON格式
            return response.json();

        } else{
            //錯誤處理
        }
        
    
    }).then(function(jsonData){

        console.log(jsonData); //jsonData為JSON格式的物件

    }).catch(fuction(err){

        console.log(err);

    })

### Response物件屬性

    type: basic或cors
    url: 回應網址
    redirected: true代表本次回應有經過轉址
    status: HTTP狀態碼
    ok: true代表本次回應的HTTP狀態碼介於200-299之間,視作成功
    headers: 本次回應相關的Header物件

### 文字格式 Response範例

    fetch(API_URL).then(function(response){

        if (response.ok){

            //將Response物件直接轉乘純文字格式
            return response.text();

        } else{
            //錯誤處理
        }
        
    
    }).then(function(text){

        console.log(text); //text為純文字格式的物件

    }).catch(fuction(err){

        console.log(err);

    })


# Redux進階

Fat action, slim reducer.

# 設計同步Action Creator

    1. 開始呼叫後端資料API
        1.1. 通知使用者應用程式用再擷取後端資料, 呈現載入中狀態
        1.2. 驗證參數是否正常,是否有必要取得網路資料
            1.2.1. 參數錯誤, 或沒有必要抓取後端資料,通知使用者,並取消載入狀態
        1.3. 組裝呼叫後端資料API所需參數呼叫,fetch函式抓取後端資料
    2. 收到後端資料API結果
        2.1. 網路問題無法呼叫,通知使用者,並取消載入狀態
        2.2. 伺服器回應錯誤代碼(4xx,5xx),通知使用者,並取消載入狀態
        2.3. 伺服器回傳內容格式有誤,通知使用者,並取消載入狀態
        2.4. 伺服器回傳正確格式資料,消化資料後,通知使用者,並取消載入狀態

# 同步Action

    // 開始抓取後端資料, 並設定isFetchingTodoList為true來呈現載入中狀態
    {tpye:'beginFetchTodoList',payload:{isFetchingTodoList:true}}
    
    // 後端資料抓取結束,並設定isFetchingTodoList為false來取消載入中狀態
    // 若有任何錯誤,帶上錯誤內容(參數錯誤,網路錯誤,錯誤代碼,格式錯誤)
    {tpye:'finishFetchTodoList',payload:{isFetchingTodoList:false,error:error}}

    // 成功接收後端資料
    {type:'receFetchTodoListResult',payload:{todosLtodos}};

# Middleware

    State (呼叫)-> Action Creator (建構)-> Action (送交)-> Reducer (產生)->State (更新)->State

## Middleware架構(ES5)

    export default function(store){
        return function(next){
            return function(action){
                //Middleware邏輯
                //最後以retrun next(action)或是store.dispatch(action)結束
            }
        }
    }

## Middleware架構(ES6)

    export default store => next => action =>{
        // Middleware邏輯
        //最後以retrun next(action)或是store.dispatch(action)結束
    }

    React (呼叫)-> Action Creator (建構)-> Action (送交)-> Middleware (轉交)-> Reducer (產生)->State (更新)->React

## redux-thunk

    return({dispatch,getState}) => next => action =>{
        if(typeof action === 'function'){
            return action(dispatch,getState,extraArgument);
        }
    
        return next(action);
    };

    安裝
    npm instaill redux-thunk --save

# 設計非同步Action

# 建置

    npm run build

## 發布前的測試

    npm install -g serve
    serve -s build

    http://localhost:5000

# React的效能陷阱

    shouldComponentUpdate預設回傳true

    shouldComponentUpdate(nextPorps, nextState){
        return nextProps !== this.props || nextState !== this.state;
    }

## PureComponent

    import React,{ PureComponent } from 'react';

### 不建議使用時機

    元電的更新與否無法單存以淺層比對屬性或狀態決定, 必須深層比對

    該元件的下層元件因上述原因無法使用PureComponent


# Google Progressive Web Application

Google PWA

## Web
   
    PWA是一個網頁
    可以使用網址的形式分享
    不需要經過'上架'的程序,上傳至網頁伺服器可以提供用戶使用
    內容可以被搜尋引擎分析
    使用HTTPS技術提供安全防護

## Application

    PWA也是一個行動裝置應用程式
    正如同多數的行動裝置應用程式一樣,不需連網也能使用
    可以下載並且安裝置行動裝置,不需要透過App Store或Market機制即可安裝
    沒有網址列,可全螢幕操作,並依據行動裝置螢幕大小給予適切的畫面設計
    可使用推播Push機制通知

## Progressive

    漸進式載入與安裝
    能顯示安裝進度條
    可在網路連線品質不佳的情況下使用
    可快速啟動
    可在背景下載更新檔或其他資源

# 網頁效能兩個重要數字

    1. First MeaningfuL Paint, FMP, 對用戶來說第一次有意義的繪製
    2. Time to Interactive, TTI, 用戶需要等待多久才能開始操作

# ES7動態輸入語法

    import('./moduleA')
        .then(({moduleA}) => {
            //自此模A可供使用
        })

## 檢查PWA是否正常運作

    npm run build

    serve -s build

# PWA的陷阱

    會先抓取catch

# 列表效能優化

    列表對於效能的影響
    不同的撰寫方式,對效能會有多大影響?
    如何優化列表效能
    在業界怎麼解決常列表對於效能的傷害?

# 虛擬化列表


## react-virtualized

    npm install react-virtualized