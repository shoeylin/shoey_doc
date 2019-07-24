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
        ctx.fillStyle = "red";
        ctx.fillRect(10, 10, 30, 30);
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
          this.setState({});
        }
      }
      //之後介紹
      getSnapshotBeforeUpdate(prevProps, prevState) {}

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
        clearTimeout(this.timer);
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

# shouldComponentUpdate 使用時機

    import React, { Component } from "react";

    class MyComponent extends Component {
      shouldComponentUpdate(nextProps, nextState) {
        // if (
        //   parseInt(nextProps.ms / 100, 10) !== parseInt(this.props.ms / 100, 10)
        // ) {
        //   return true;
        // }
        // return false;

        return (
          parseInt(nextProps.ms / 100, 10) !== parseInt(this.props.ms / 100, 10)
        );
      }

      render() {
        const { ms } = this.props;
        return <div>{this.props(ms / 1000, 10)}</div>;
      }
    }

    //ms 100 200 300

    export default MyComponent;

## shouldComponentUpdate (2) PureComponent 的內建內容

    import React, { PureComponent } from "react";

    class MyComponent extends PureComponent {
      //PureComponent內建
      shouldComponentUpdate(nextProps, nextState) {
        Object.keys(nextState).some(key => {
          this.props[key] !== nextProps[key];
          //shallow compare
        });
      }

      render() {
        const { ms } = this.props;
        return <div>{this.props(ms / 1000, 10)}</div>;
      }
    }

    //ms 100 200 300

    export default MyComponent;

# getSanpshotBeforeUpdate

    EX:log

    import React, { Component, createRef } from "react";

    class Log extends Component {
      ref = createRef();
      getSnapshotBeforeUpdate(prevProps, prevState) {
        if (prevProps.items.length !== this.props.items.length) {
          const list = this.ref.current;
          return list.scrollHeight - list.scrollTop;
        }
        return null;
      }
      componentDidUpdate(prevProps, prevState, snapshot) {
        if (snapshot) {
          const list = this.ref.current;
          list.scrollTop = list.scrollHeight - shapshot;
        }
      }

      render() {
        const { items } = this.props;
        return <div ref={this.ref}>{items.map}</div>;
      }
    }

    export default Log;

# rander (react 16.2) 更靈活

    import React, { Component,Fragment } from 'react';

    class MyComponent extends Component {
      render() {
        return //reacte elements, bool, int, null ,Array, fragments
        (

          <Fragment>
          <h1></h1>
          <div></div>
            <button></button>
          </Fragment>
        );
      }
    }


    export default MyComponent;

# Frangment

    import React, { Component, Fragment } from "react";

    // babel
    class MyComponent extends Component {
      render() {
        return (
          <Router>
            <Fragment>
              <h1 />
              <div />
              <button />
            </Fragment>
          </Router>
        );
      }
    }

    export default MyComponent;

# 錯誤處理

## MyComponent

    import React, { Component } from "react";

    class MyComponent extends Component {
      componentDidMount() {
        JSON.props("asdf");
        //Promise錯誤不會被抓到 throw Error
        return Promise.reject("test");
      }
      render() {
        return <div>MyComponent</div>;
      }
    }

    export default MyComponent;

## App

    import React, { Component,axios } from "react";
    import MyComponent from "./MyComponent";

    class App extends Component {
      state = {
        hasError: false
      };
      static getDerivedStateFromError(err) {
        return { hasError: true };
      }
      //有error被呼叫
      componentDidCatch(error, info) {
        axios.post('/api/logger', {info})
      }
      render() {
        const hasError = this.state;
        if (hasError) {
          return <h1>Error!</h1>;
        }
        return (
          <div>
            <h1>App</h1>
            <MyComponent />
          </div>
        );
      }
    }

    export default App;

# Portal

## LessonModal.js

    import React, { Component } from "react";
    import { createPortal } from "react-dom";

    class LessonModal extends Component {
      render() {
        return createPortal(<div> </div>, document.getElementById("modal"))();
      }
    }

    export default LessonModal;

## html

    <!DOCTYPE html>
    <html lang="en">

    <head>
        <meta charset="utf-8" />
        <title>React App</title>
    </head>

    <body>
        <div id="app"></div>
        <div id="model"></div>
        <div id="messenger"></div>
    </body>

    </html>

##Lesson

    import React, { Component } from "react";
    import LessonModal from "./LessonModal"

    class Lesson extends Component {
      state = {
        detailVisible: false
      };
      showMore = () => {
        this.setState({
          detailVisible: true
        });
      };
      render() {
        const { detailVisible } = this.state;
        return (
          <div>
            <button onClick={this.showMore}>Show More</button>
            {detailVisible && <LessonModal />}
          </div>
        );
      }
    }

    export default Lesson;

# Context API

## 資料傳遞與視覺結構綁定

## Product.js

import React from "react";

const Product = ({ id, name, addOrder }) => (

<li>
<label>{name}</label>
<button onClick={() => addOrder(id)}>+</button>
</li>
);

export default Product;

## ProductList.js

    import React from "react";
    import Product from "./Product";

    const products = [
      { id: 1, name: "牛肉鍋" },
      { id: 2, name: "豬肉鍋" },
      { id: 3, name: "羊肉鍋" },
      { id: 4, name: "雞肉鍋" },
      { id: 5, name: "好肉鍋" }
    ];

    const ProductList = ({ addOrder }) => (
      <ul>
        {products.map(product => (
          <Product
            {...product}
            addOrder={addOrder}
            key={product.id}
          />
        ))}
      </ul>
    );
    export default ProductList;

## Header.js

    import React, { Component } from "react";

    const Header = ({ orders = [] }) => (
      <div>
        <span>{`購物車 (${orders.length})`}</span>
        <hr />
      </div>
    );
    export default Header;

## App.js

    import React, { Component } from "react";

    import Header from "./Header";
    import ProductList from "./ProductList";

    class App extends Component {
      state = {
        orders: []
      };
      addOrder = order => {
        this.setState({
          orders: [...this.state.orders, order]
        });
      };
      render() {
        const { orders } = this.state;
        return (
          <div>
            <Header orders={orders} />
            <ProductList addOrder={this.addOrder} />
          </div>
        );
      }
    }

    export default App;

# 使用 createContext

## OrderContext

    import { createContext } from "react";

    export const { Provider, Consumer } = createContext({
      orders: [],
      addOrder: () => {}
    });

## Header.js

    import React, { Component } from "react";
    import {Consumer} from "./OrderContext";

    const Header = ({ orders = [] }) => (
      <div>
        <span>
          <Consumer>
            {value => `購物車 (${value.orders.length})`}
          </Consumer>
        </span>

        <hr />
      </div>
    );
    export default Header;

## Product.js

    import React from "react";

    import {Consumer} from "./OrderContext";

    const Product = ({ id, name }) => (
      <li>
        <label>{name}</label>
        <Consumer>
          {value => <button onClick={() => value.addOrder(id)}>+</button>}
        </Consumer>
      </li>
    );

    export default Product;

## ProductList.js

    import React from "react";
    import Product from "./Product";

    const products = [
      { id: 1, name: "牛肉鍋" },
      { id: 2, name: "豬肉鍋" },
      { id: 3, name: "羊肉鍋" },
      { id: 4, name: "雞肉鍋" },
      { id: 5, name: "好肉鍋" }
    ];

    const ProductList = () => (
      <ul>
        {products.map(product => (
          <Product
            {...product}

            key={product.id}
          />
        ))}
      </ul>
    );
    export default ProductList;

## App.js

    import React, { Component } from "react";

    import Header from "./Header";
    import ProductList from "./ProductList";

    import {Provider} from "./OrderContext";

    class App extends Component {
      state = {
        orders: []
      };
      addOrder = order => {
        this.setState({
          orders: [...this.state.orders, order]
        });
      };
      render() {
        const { orders } = this.state;
        const contextValue = {
          orders,
          addOrder: this.addOrder
        };
        return (
          <div>
            <Provider value={contextValue}>
              <Header />
              <ProductList />
            </Provider>
          </div>
        );
      }
    }

    export default App;

# 列表過濾

## App.js 存前端

    import React, { Component } from "react";

    import lessons from "./lessons.json";

    class App extends Component {
      state = {
        filter: "",
        lessons
      };
      onChangeFilter = e => {
        this.setState({
          filter: e.target.value
        });
      };
      render() {
        const { filter, lessons } = this.state;
        const visibleLessons = lessons.filter(lesson =>
          lesson.title.toLowerCase().includes(filter.toLocaleLowerCase())
        );
        return (
          <div>
            <input value={filter} onChange={this.onChangeFilter} />
            <ul>
              {visibleLessons.map(lesson => (
                <li key={lesson.id}>{lesson.title}</li>
              ))}
            </ul>
          </div>
        );
      }
    }

    export default App;

# 列表過濾 有後端 server

## server.js

const express = require("express");
const lessons = require("./src/lessons.json");

const app = express();
const PORT = 3001;
app.listen(PORT, () => console.log(`server started at ${PORT}`));

app.get("/api/lessons", (req, res) => {
const { filter = "" } = req.query;

    if (!filter) {
      res.json(lessons);
    } else {
      res.json(
        lessons.filter(({ title }) =>
          title.toLowerCase().includes(filter.toLowerCase())
        )
      );
    }

});

## App.js

    import React, { Component } from "react";

    class App extends Component {
      state = {
        filter: "",
        lessons: []
      };
      componentDidMount() {
        this.fetchLessons();
      }
      onChangeFilter = e => {
        const filter = e.target.value;
        this.setState({ filter }, this.fetchLessons);
      };
      fetchLessons = () => {
        clearTimeout(this.timer);
        this.timer = setTimeout(() => {
          const { filter } = this.state;
          fetch(`/api/lessons?filter=${filter}`)
            .then(rs => rs.json())
            .then(lessons => this.setState({ lessons }));
        }, 300); //debounce
      };

      render() {
        const { filter, lessons } = this.state;
        return (
          <div>
            <input value={filter} onChange={this.onChangeFilter} />
            <ul>
              {lessons.map(lesson => (
                <li key={lesson.id}>{lesson.title}</li>
              ))}
            </ul>
          </div>
        );
      }
    }

    export default App;

# 井字遊戲

## Game.js

    import React, { Component } from "react";

    import style from "./Game.module.css";

    const toSymbol = n => {
      switch (n) {
        case 0:
          return "";
        case 1:
          return "O";
        case -1:
          return "X";
      }
    };

    // 0 1 2
    // 3 4 5
    // 6 7 8

    const lines = [
      [0, 1, 2],
      [3, 4, 5],
      [6, 7, 8],
      [0, 3, 6],
      [1, 4, 7],
      [2, 5, 8],
      [0, 4, 8],
      [2, 4, 6]
    ];

    class Game extends Component {
      state = {
        grids: [0, 0, 0, 0, 0, 0, 0, 0, 0],
        player: 1,
        winner: 0
      };
      componentDidUpdate(prevProps, prevState) {
        if (prevState.grids !== this.state.grids) {
          this.setState({
            winner: this.getWinner()
          });
        }
      }

      handleClick = idx => {
        if (this.state.winner !== 0) return;
        const grids = [...this.state.grids];
        if (grids[idx] !== 0) return;
        grids[idx] = this.state.player;
        this.setState({
          grids,
          player: -this.state.player
        });
      };
      reset = () => {
        this.setState({
          grids: [0, 0, 0, 0, 0, 0, 0, 0, 0],
          player: 1
        });
      };
      getWinner = () => {
        const { grids } = this.state;
        for (const line of lines) {
          const [i, j, k] = line;
          if (grids[i] === grids[j] && grids[j] === grids[k]) {
            return grids[i];
          }
        }
        return 0;
      };
      render() {
        const { grids, player, winner } = this.state;
        return (
          <div>
            <div className={style.board}>
              {grids.map((elm, idx) => (
                <div className={style.grid} onClick={() => this.handleClick(idx)}>
                  <span>{toSymbol(elm)}</span>
                </div>
              ))}
            </div>
            <div>Player: {toSymbol(player)}</div>
            <div>Winner: {toSymbol(winner)}</div>
            <button onClick={this.reset}>Reset</button>
          </div>
        );
      }
    }

    export default Game;

## Game.moduls.css

    .board {
        width: 240px;
        height: 240px;
        display: flex;
        flex-flow: row wrap;
        align-content: flex-start;
    }

    .grid {
        flex-basis: 33%;
        height: 33%;
        border: 1px solid #ccc;
        margin: -1px 0 0 -1px;
        box-sizing: border-box;
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 60px;
        color: #666;
        cursor: pointer;
    }

# context api 跳窗實作

## App.js

    import React, { Component } from "react";
    import Header from "./Header";
    import Modal from "./Modal";
    import "./App.css";
    import { Provider } from "./context";

    class App extends Component {
      state = {
        modal: null,
        login: false
      };

      render() {
        const contextValue = {
          state: this.state,
          setState: this.setState.bind(this)
        };
        return (
          <Provider value={contextValue}>
            <div className="app">
              <Header />
              <Modal />
            </div>
          </Provider>
        );
      }
    }

    export default App;

## App.css

    html,
    body,
    .app {
        margin: 0;
        padding: 0;
    }

## context.js

    import { createContext } from "react";

    export const { Provider, Consumer } = createContext({});

## Header.css

    .header {
        background: #333;
        color: #ccc;
        text-align: right;
    }

    .header>span {
        display: inline-block;
        padding: 0.5em 1em;
    }

    .header a {
        color: white;
    }

## Header.js

    import React, { Component } from "react";
    import "./Header.css";
    import { Consumer } from "./context";

    const Header = () => (
      <Consumer>
        {({ setState }) => (
          <div className="header">
            <span>
              <a
                href="#"
                onClick={e => {
                  e.preventDefault();
                  setState({ modal: "login" });
                }}
              >
                Login
              </a>
            </span>
          </div>
        )}
      </Consumer>
    );

    export default Header;

## Modal.css

    .modal {
        position: fixed;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
        background: rgba(0, 0, 0, 0.8);
        color: white;
    }

    .modal-content {
        position: absolute;
        left: 30px;
        top: 30px;
        bottom: 30px;
        right: 30px;
        border-radius: 10px;
        background: white;
    }

## Modal.js

    import React, { Component } from "react";
    import { Consumer } from "./context";
    import ModalLogin from "./ModalLogin";
    import "./Modal.css";

    class Modal extends Component {
      renderContent = (modal, setState) => {
        switch (modal) {
          case "login":
            return <ModalLogin setState={setState} />;
          default:
            return null;
        }
      };

      render() {
        return (
          <Consumer>
            {({ state, setState }) => {
              const { modal } = state;
              if (!modal) return null;

              return (
                <div className="modal">
                  <div className="modal-content">
                    {this.renderContent(modal, setState)}
                  </div>
                </div>
              );
            }}
          </Consumer>
        );
      }
    }

    export default Modal;

## ModalLogin.js

    import React, { Component } from "react";
    import axios from "axios";

    class ModalLogin extends Component {
      state = {
        username: "",
        password: ""
      };
      onChangeUsername = e => {
        this.setState({
          username: e.target.value
        });
      };
      onChangePassword = e => {
        this.setState({
          password: e.target.value
        });
      };
      onSubmit = e => {
        e.preventDefalt();
        const { username, password } = this.state;
        axios.post("/api/login", { username, password }).then(rs => {
          if (rs.data.success) {
            this.props.setState({
              modal: null,
              login: true
            });
          }
        });
      };
      render() {
        const { username, password } = this.state;
        return (
          <form className="modal-login" onSubmit={this.onSubmit}>
            <input type="text" value={username} onChange={this.onChangeUsername} />
            <br />
            <input
              type="password"
              value={password}
              onChnage={this.onChangePassword}
            />
            <br />
            <button type="submit">Submit</button>
          </form>
        );
      }
    }

    export default ModalLogin;

# hooks

    import React, { useState } from "react";

    const App = () => {
      const [count, setCount] = useState(0);
      const addCount = () => setCount(count + 1);
      return (
        <div>
          <h1>{count}</h1>
          <button onClick={addCount}>Add</button>
        </div>
      );
    };
    export default App;

# hooks todolist

## App.js

    import React, { useState } from "react";
    import TodoLists from "./TodoLists";

    const App = () => <TodoLists/>

    export default App;

## TodoList.js

    import React, { useState } from "react";
    import TodoInput from "./TodoInput";

    const TodoList = () => {
      const [items, setItems] = useState([]);
      const addItem = text => {
        setItems([...items, { id: Date.now(), text }]);
      };
      const removeItem = id => {
        setItems(items.filter(item => item.id !== id));
      };
      return (
        <div>
          <TodoInput addItem={addItem} />
          <ul>
            {items.map(item => (
              <li key={item.id} onClick={() => removeItem(item.id)}>
                {item.text}
              </li>
            ))}
          </ul>
        </div>
      );
    };
    export default TodoList;

## TodoInput.js

    import React, { useState, useEffect, useRef } from "react";

    const TodoInput = ({ addItem }) => {
      const [text, setText] = useState("");
      const onChangeText = e => {
        setText(e.target.value);
      };
      const onSubmit = e => {
        e.preventDefault();
        addItem(text);
        setText("");
      };

      const ref = useRef();
      useEffect(() => {
        ref.current.focus();
      });

      return (
        <form onSubmit={onSubmit}>
          <input type="text" value={text} onChange={onChangeText} ref={ref} />
          <button type="submit">Submit</button>
        </form>
      );
    };

    export default TodoInput;

# Cookie Clicker 遊戲

## App.js

    import React, { Component } from "react";

    const upgrades = [
      {
        name: "Grandma",
        Price: 10,
        auto: 1
      },
      {
        name: "Farm",
        Price: 100,
        auto: 8
      },
      {
        name: "Mine",
        Price: 1000,
        auto: 60
      }
    ];

    class App extends Component {
      state = {
        cookies: 1000,
        auto: 0,
        upgrades
      };
      componentDidMount() {
        this.time = Date.now();
        this.updateCookie();
      }
      //   autoStep = () => {
      //     this.setState(({ cookies, auto }) => ({
      //       cookies: cookies + auto
      //     }));
      //     setTimeout(this.autoStep, 1000);
      //   };

      updateCookie = () => {
        const { cookies, auto } = this.state;
        const time = Date.now();
        const stap = (auto / 1000) * (time - this.time);
        this.time = time;

        this.setState(
          {
            cookies: cookies + stap
          },
          () => {
            requestAnimationFrame(this.updateCookie);
          }
        );
      };

      addCookie = () => {
        const { cookies } = this.state;
        this.setState({
          cookies: cookies + 1
        });
      };
      upgrades = idx => {
        const { cookies, auto, upgrades } = this.state;
        const u = upgrades[idx];
        if (cookies < u.Price) return;

        this.setState({
          cookies: cookies - u.Price,
          auto: auto + u.auto,
          upgrades: upgrades.map((v, i) =>
            i !== idx ? v : { ...v, Price: parseInt(v.Price * 1.15, 10) }
          )
        });
      };
      //   buyFarm = () => {
      //     const { cookies, auto, priceFram: price } = this.state;
      //     if (cookies < price) return;
      //     this.setState({
      //       cookies: cookies - price,
      //       auto: auto + 8,
      //       priceFram: parseInt(price * 1.5, 10)
      //     });
      //   };
      render() {
        const { cookies, upgrades } = this.state;
        return (
          <div>
            <h3>Cookies</h3>
            <h1>{parseInt(cookies,10)}</h1>
            <hr />
            <button onClick={this.addCookie}>Click</button>
            <br />
            {upgrades.map(({ name, Price }, idx) => [
              <button onClick={() => this.upgrades(idx)}>{name}</button>,
              <span>{Price}</span>,
              <br />
            ])}
          </div>
        );
      }
    }

    export default App;

# useState

    import React, { useState } from "react";

    const App = () => {
      const [{ count1, count2 }, setState] = useState({ count1: 0, count2: 0 });

      const addCount1 = () => {
        setState(state => ({ ...state, count1: state.count1 + 1 }));
      };

      const addCount2 = () => {
        setState(state => ({ ...state, count2: state.count2 + 1 }));
      };
      return (
        <div>
          <h1>count1: {count1}</h1>
          <button onClick={addCount1}>add</button>
          <h1>count2: {count2}</h1>
          <button onClick={addCount2}>add</button>
        </div>
      );
    };

    export default App;

# useEffect

## 原本

    import React, { Component } from "react";

    class FetchUseEffect extends Component {
      state = {
        email: "",
        picture: ""
      };
      componentDidMount() {
        this.fetchUser();
      }
      fetchUser = () => {
        fetch("https://randomuser.me/api")
          .then(rs => rs.json())
          .then(data => {
            const [user] = data.results;
            this.setState({
              email: user.email,
              picture: user.picture.medium
            });
          });
      };
      render() {
        const { email, picture } = this.state;

        return (
          <div>
            <img src={picture} />
            <div>{email}</div>
          </div>
        );
      }
    }

    export default FetchUseEffect;

## useEffect

    import React, { useState, useEffect } from "react";

    const FetchUseEffect = () => {
      const [state, setState] = useState({
        email: "",
        picture: ""
      });

      useEffect(() => {
        fetch("https://randomuser.me/api")
          .then(rs => rs.json())
          .then(data => {
            const [user] = data.results;
            setState({
              email: user.email,
              picture: user.picture.medium
            });
          });
      }, []); //[]等於 componentDidMount ,沒填就是componenetDidUpdate

      const { email, picture } = state;

      return (
        <div>
          <img src={picture} />
          <div>{email}</div>
        </div>
      );
    };

    export default FetchUseEffect;

## useEffect 代替 componentWillUnmount

    import React, { useState, useEffect } from "react";

    const UseEffectFetch = () => {


      useEffect(() => {
        const onScroll=()=>{};
        window.addEventListener('scroll',onScroll)
        return ()=>{
          window.removeEventListener('scroll',onScroll)
        }
      }, []);
      //1. 檢查[]是否重複 ,[]等於 componentDidMount ,沒填就是componenetDidUpdate
      //2. 檢查return
      //3. 執行內容
      //4. 存下return


      return (
        <div>

        </div>
      );
    };

    export default UseEffectFetch;

# hooks context

## context.js

    import { createContext } from "react";
    const context = createContext();
    export const { Provider, Consumer } = context;
    export default context;

## UseContextOpen.js

    import React, { useState } from "react";
    import UseContextOpenButton from "./UseContextOpenButton";
    import { Provider } from "./context";

    const UseContextOpen = () => {
      const [open, setOpen] = useState(false);
      const toggle = () => {
        setOpen(!open);
      };
      const contextValue = {
        open,
        toggle
      };

      return (
        <Provider value={contextValue}>
          <div>
            <UseContextOpenButton />
            {open && <div>Some Content</div>}
          </div>
        </Provider>
      );
    };

    export default UseContextOpen;

## UseContextOpenButten.js

    import React, { useContext } from "react";
    import context from "./context";

    const UseContextOpenButton = () => {
      const { open, toggle } = useContext(context);

      return <button onClick={toggle}>{open ? "Close" : "Open"}</button>;
    };

    export default UseContextOpenButton;

# hooks UseRef

## useRef.js

    import React, { useEffect, useRef } from "react";

    const UseRefFocus = () => {
      const ref = useRef();
      useEffect(() => {
        ref.current.focus();
      }, []);

      return (
        <div>
          <input ref={ref} />
        </div>
      );
    };

    export default UseRefFocus;

## UseRefTimer.js

    import React, { useState, useRef } from "react";

    const UseRefCounter = () => {
      const [count, setCount] = useState(0);
      const ref = useRef();
      const startCounter = () => {

        ref.current = setInterval(() => setCount(c => c + 1), 100);
      };
      const stopCounter = () => {
        clearInterval(ref.current);
      };
      return (
        <div>
          <h1>{count}</h1>
          <button onClick={startCounter}>Start</button>
          <button onClick={stopCounter}>Stop</button>
        </div>
      );
    };

    export default UseRefCounter;

---

---

---

# Old react 15 以前

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
