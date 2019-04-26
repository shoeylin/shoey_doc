# Componet 生命週期:掛載

 constructor() 建構式

 componentWillMount() 物件即將載入

 render() 繪製

 componentDidMount() 物件已經載入

 # Component 生命週期:更新

componentWillReceiveProps(nextProps)物件即將接收新屬性

&

this.setState() 元件狀態更新

shouldComponentUpdate(nextPropes, nextState) 物件是否需要更新

true or false

componentWillUpdate(nextProps, nextState) 物件即將更新

render() 繪製

componentDidUpdate(prevProps, prevState) 物件已經更新

# 屬性與狀態

## 屬性:由上層元件以JSX屬性傳入, 可想像成HTML的屬性(attribute), 不可修改

## 狀態:由元件自身擁有, 在建構式中宣告, 存放元件自身的狀態, 可以任意修改

    <img src={logo} className="App-logo" alt="logo" />

    constructor(){
        super();
        this.state = {
            clickCount: 0
        }
    }

    addClickCount(){
        this.setState({
            clickCount: this.state.clickCount++
        )};
    }

# 資料異動

按鈕按下後 呼叫函式新增

該函式取得輸入框中的值

異動待辦事項資料


# 重點整理

## 構造函是如果要使用this必須先super()

## 資料可以透過component的屬性, 傳遞到子component

## function 也可以透過 props傳遞

## 以箭頭函式或用bind來指定function的 this

## 輸入框的操作與防呆

# 以畫面元件驅動開發

    npm instaill --save styled-components

    import styled from 'Styled-components';

    const Clickable = style.span`
        cursor: pointer;
        `;

# 網頁應用程式介紹

## 多頁式網頁應用程式 MPA

### 優點:
    
    技術門檻較低
    對搜尋引擎最友善
    對瀏覽器的功能要求最低, 可觸及的用戶最廣
    首次有效畫面繪製時間成本固定

### 缺點:

    網路頻寬耗用甚鉅
    頁面與頁面之間都需要等待伺服器回應整個頁
    前端與後端技術發點會被綁定無法切割
    開發時間較久且除錯困難
    頁面與頁面之間切換會看到空白
    伺服器負擔較大亦即營運成本較高

## 單頁式網頁應用程式 SPA

### 優點:

    網路頻寬耗用最低,一次性載入所有程式碼再也無須載入
    使用者體驗最佳,頁面與頁面之前的操作僅需等待伺服器回傳最少量資料
    前端與後端技術切割且獨立
    開發時間快速且除錯方便
    較容易針對行動裝置進行優化
    頁面切換不再出現空白
    降低伺服器負擔亦即降低營運成本

### 缺點:

    對搜尋引擎不友善
    首次有畫面繪製時間成本隨應用程式複雜度增加
    對瀏覽器的功能較高,可觸及用戶受限

# React Router 網址與應用程式狀態連結

    npm intall react-router-dom --save

## 核心三元件

### BrowserRouter
### Route
### Link

    import{
        BrowerRouter,
        Route,
        Link
    } from 'react-router-dom'
