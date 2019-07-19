# react

參考 : 
http://projects.wojtekmaj.pl/react-lifecycle-methods-diagram/
# construct 用法

import React, { Component } from "react";

class Counter extends Component {
state = { count: 0 };
constructor(props) {
super(props);

    // this.state = { count: 0 };
    // this.handleClick = this.handleClick.bind(this);

    //不要用
    //setState
    //props -> state
    //fetch()

}

handleClick = () => {};
render() {
return <div>{this.props.count}</div>;
}
}

export default Counter;

# render 介紹

import React, { Component } from "react";

class Counter extends Component {
//render 主要是把 props or state 轉成視覺畫面
render() {
this.props;
this.state;

    const { count } = this.props;
    const visualCount = count * 2 + 3;

    //不要
    // this.setState()
    // ajax
    // fetch;
    // axios;

    return <div>{this.props.count}</div>;

}
}

export default Counter;

# componentDidMoount

import React, { Component } from "react";

class List extends Component {
state = {
items: []
};
componentDidMount() {
this.fetchList();
}
fetchList = async () => {
const response = await fetch("http://..../api/list");
const data = await response.json();
this.setState({
items: data
});
};
render() {
const { items } = this.state;
return (
<ul>
{items.map(item => (
<li key={item.id}>{item.text}</li>
))}
</ul>
);
}
}

export default List;

## 第二種 componentDidMount

import React, { Component, createRef } from "react";

class Canvas extends Component {
ref = createRef();
componentDidMount() {
const ctx = this.ref.current.getContext("2d");
ctx.fillStyle='red';
ctx.fillRect(10,10,30,30);
}
render() {
return <canvas ref={this.ref} />;
}
}

export default Canvas;

# componentDidUpdatae

import React, { Component } from "react";

class Proflie extends Component {
state = {
userDate: {}
};
componentDidMount() {
this.fetchUser(this.props.userID);
}
componentDidUpdate(prevProps, prevState, snapshot) {
//condition
if (prevProps.userID !== this.props.userID) {
this.fetchUser(this.props.userID);
this.setState({

      })
    }

}
//之後介紹
getSnapshotBeforeUpdate(prevProps,prevState){

}

fetchUser = userID => {
fetch("http://mysite.com/api/user/${userID}")
.then(response => response.json())
.then(data => {
this.setState({
userData: data
});
});
};
render() {
const { userData } = this.state;
return (
<div>
<img src={userData.picture} alt="" />
<label>{userData.name}</label>
</div>
);
}
}

export default Proflie;

# componentWillUnmount

import React, { Component } from "react";

class MyComponent extends Component {
componentDidMount() {
window.addEventListener("scroll", this.onScroll);
}
componentDidUpdate() {
this.timer = setTimeout(() => this.update(), 3000);
  
}
//協助中斷不需要的程序
componentWillUnmount() {
window.removeEventListener("scroll", this.onScroll);
clearTimeout(this.timer)
}
onScroll = () => {
//...
};
render() {
return <div>MyComponent</div>;
}
}

export default MyComponent;

# 無限捲動範例

## App.js
import React, { Component } from "react";
import Course from "./Course";

const api =
  "https://hiskio.com/api/v1/courses/professions?type=all&level=all&sort=latest&profession_id=1";

class App extends Component {
  state = {
    courses: [],
    next: null,
    loading: true
  };
  componentDidMount() {
    this.fetchData(api);
    window.addEventListener("scroll", this.onScroll);
  }
  componentWillUnmount() {
    window.removeEventListener("scroll", this.onScroll);
  }
  onScroll = () => {
    const { next, loading } = this.state;
    if (loading) return;
    if (!next) return;
    if (
      window.scrollY + window.innerHeight >=
      document.body.scrollHeight - 100
    ) {
      // load next
      this.fetchData(next);
    }
  };

  fetchData = url => {
    this.setState({
      loading: true
    });

    fetch(url)
      .then(rs => rs.json())
      .then(data => {
        this.setState({
          courses: [...this.state.courses, ...data.courses],
          next: data.paginate.next_page_url,
          loading: false
        });
      });
  };

  render() {
    const { courses } = this.state;
    return (
      <div>
        {courses.map(course => (
          <Course {...course} />
        ))}
      </div>
    );
  }
}

export default App;



## Course
import React, { Component } from "react";
import style from "./Course.module.css";

class Course extends Component {
  render() {
    const { title, subtitle, cover, teachers } = this.props;
    return (
      <div className={style.Course}>
        <div
          className={style.cover}
          style={{
            backgroundImage: `url('${cover}')`
          }}
        />

        <div className={style.info}>
          <div className={style.title}>{title}</div>
          <label>{teachers[0].username}</label>;
        </div>
      </div>
    );
  }
}

export default Course;


## Course.module.css
.course {
    display: flex;
    margin: 5px 0;
    background: #fed;
    padding: 5px;
}

.cover {
    flex: 0 0 200px;
    height: 100px;
    background: center center no-repeat;
    background-size: cover;
}

.info {
    flex: 1;
    font-family: Mircrosoft JhengHei;
    font-size: 10px;
}


.title {
    font-size: 1.5em;
    color: blue
}


# 非同步處理
// const fs = require("fs");

// fs.readFlie("./file.txt", (err, content) => {
//   console.log(content);
// });

// const content = fs.readFlieSync("./file.txt");

const fs = require("fs");

// 1. callback
const getData = callback => {
  fs.readFlie("./data.json", (err, content) => {
    callback(content);
  });
};
const start = () => {
  const data = getData(data => {
    console.log(data);
  });
};
start();

// 2.promise
const getData = () => {
  return new Promise((resolve, reject) => {
    fs.readFlie("./data.json", (err, content) => {
      if (err) {
        reject(err);
      } else {
        resolve(content);
      }
    });
  });
};

const start = () => {
  getData().then(data => {
    console.log(data);
  });
};
start();

//3 async/await bable/node>=8
const getData = () => {
  return new Promise((resolve, reject) => {
    fs.readFlie("./data.json", (err, content) => {
      if (err) {
        reject(err);
      } else {
        resolve(content);
      }
    });
  });
};

const start = async () => {
  const data = await getData();
  console.log(data);
};
start();

# getDerivedStateFromProps 生命週期函式
import React, { Component } from "react";

class MyComponent extends Component {
  // static getDerivedStateFromProps(props, state) {
  // no async, this, 
  // mounted / props /state
  //   if (props.text !== state.propText) {
  //     return {
  //       text: props.text,
  //       propText: props.text
  //     };
  //   }
  //   return null;
  // }
  state = {
    text: "",
    propText: ""
  };
  constructor(props) {
    super(props);
    this.state = {
      text: props.text
    };
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevProps.text !== this.props.text) {
      this.setState({
        text: this.props.text
      });
    }
  }

  onChange = e => {
    this.setState({
      text: e.target.value
    });
  };
  reset = () => {
    this.setState({
      // text: this.state.propText
      text: this.props.text
    });
  };
  render() {
    const { text } = this.state;
    return (
      <div>
        <input value={text} onChange={this.onChange} />
        <button onClick={this.reset}>Reset</button>
      </div>
    );
  }
}

export default MyComponent;



# shouldComponentUpdate使用時機










---

---

---

# Old

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

## 屬性:由上層元件以 JSX 屬性傳入, 可想像成 HTML 的屬性(attribute), 不可修改

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

## 構造函是如果要使用 this 必須先 super()

## 資料可以透過 component 的屬性, 傳遞到子 component

## function 也可以透過 props 傳遞

## 以箭頭函式或用 bind 來指定 function 的 this

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
