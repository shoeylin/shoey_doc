# jsx

    1. must close
    2. seft close  /  一定要有關閉
    3. className / htmlFor
    4. onClick 駝峰式命名
    5. {}

# Component

    import React, { Component } from "react";

    //static property
    // static method
    class Item extends Component {
    static propTypes ={

    }

    static getDerivedStateFromProps(){

    }

    bioname

    age = 0

    state={
        x:1,
    }
    constructor() {
        this.state={
        x:1,
        }
    }
    componentDidMount() {}

    componentDidUpdate() {}

    onChange = () => {

    }

    getData(){

    }

    render() {
        return <li> hello world</li>;
    }
    }


    Item.propTypes={}
    export default Item;

# state

    import React, { Component } from "react";

    class Message extends Component {
    state = {
        title:'Massage',
        text: 'Hello',
    }
    seyHi=()=>{
        this.setState({
            text:'Hi',
        });
    }
    render() {
        return (
            <div>
                <h3>{this.state.title}</h3>
                <h3>{this.state.text}</h3>
                <button onClick={this.seyHi}>Say hi</button>
            </div>
        );
    }
    }

    export default Message;

# this and bind 說明

    import React, { Component } from "react";

    class Message extends Component {
    state = {
        title: 'Massage',
        text: 'Hello 2',
    }
    constructor(props){
        super(props);

        this.sayHi=this.sayHi.bind(this);
    }
    // seyHi=()=>{
    //     this.setState({
    //         text:'Hi',
    //     });
    // }
    sayHi(){
        this.setState()
    }
    render() {
        return (
            <div>
                <h3>{this.state.title}</h3>
                <h3>{this.state.text}</h3>
                <button onClick={this.seyHi}>Say hi</button>
            </div>
        );
    }
    }

    export default Message;

# props

    從上面傳下來的會存在props
    List
    <Item text="Learn JavaScript" price={100}/>
    ---
    Item
    {this.props.text} ({this.props.price +1})

## children

    放在中間可以從children取用
    List
    <Item> Learn JavaScript </Item>
    ---
    Item
    {this.props.children}

## counter

    import React, { Component } from "react";

    class Counter extends Component {
    state = {
        count: 0,
        step: 1
    };
    addCount = () => {
        const { count, step } = this.state;
        this.setState({
        count: count + step,
        step: step + 1
        });
    };
    render() {
        const { count, step } = this.state;
        return (
        <div>
            <h1>{count}</h1>
            <button onClick={this.addCount}>+{step}</button>
            counter
        </div>
        );
    }
    }

    export default Counter;

# props

    可以使用傳來的props來進行初始化
    且可以使用PropTypes檢查類型
    import React, { Component } from "react";
    import PropTypes from "prop-types";

    class Counter extends Component {
        // static propTypes = {
        //     initCount: PropTypes.number
        //   };
    constructor(props) {
        super(props);
        this.state = {
        count: props.initCount,
        step: 1
        };
    }
    addCount = () => {
        const { count, step } = this.state;
        this.setState({
        count: count + step,
        step: step + 1
        });
    };
    render() {
        const { count, step } = this.state;
        return (
        <div>
            <h1>{count}</h1>
            <button onClick={this.addCount}>+{step}</button>
            counter
        </div>
        );
    }
    }

    Counter.defaultProps = {
    initCount: 0
    };

    Counter.propTypes = {
    initCount: PropTypes.number
    };

    export default Counter;

# state

    import React, { Component } from "react";

    class Counter extends Component {
    state = {
        count: 0,
        step: 1
    };

    addCount = () => {
        this.setState({
        count: this.state.count + 1
        },()=>{
            this.sendCount();
        });

    };
    sendCount = () => {
        fetch(`/api/count?value=${this.state.count}`);
    };

    render() {
        const { count, step } = this.state;
        return (
        <div>
            <h1>{count}</h1>
            <button onClick={this.addCount}>+{step}</button>
            counter
        </div>
        );
    }
    }

    export default Counter;

# ref focus()

    import React, { Component, createRef } from "react";

    class Input extends Component {

    myInput = createRef();

    componentDidMount() {
        this.myInput.current.focus();
    }

    render() {
        return (
        <div>
            <h3>Enter your name</h3>
            <input type="text" ref={this.myInput} />
        </div>
        );
    }
    }

    export default Input;

# 父子溝通

## 父

    import React, { Component, createRef } from "react";
    import Child from "./Child";

    class Parent extends Component {
    childRef = createRef();
    state = {
        count: 0
    };
    addCount = () => {
        this.setState({
        count: this.state.count + 1
        });
    };
    addChildCount=()=>{
        this.childRef.current.addCount();
    }
    render() {
        return (
        <div>
            <h1>Parent:{this.state.count}</h1>
            <button onClick={this.addCount}>+Parent</button>
            <button onClick={this.addChildCount}>+Child</button>

            <Child ref={this.childRef} addParentCount={this.addCount} />
        </div>
        );
    }
    }

    export default Parent;

## 子

    import React, { Component } from "react";

    class Child extends Component {
    state = {
        count: 0
    };
    addCount = () => {
        this.setState({ count: this.state.count + 1 });
    };
    render() {
        return (
        <div>
            <h1>Child:{this.state.count}</h1>
            <button onClick={this.props.addParentCount}>+Parent</button>
            <button onClick={this.addCount}>+Child</button>
        </div>
        );
    }
    }

    export default Child;

---

## 另一種 比較常用

## 父

    import React, { Component, createRef } from "react";
    import Child from "./Child";

    class Parent extends Component {
    childRef = createRef();
    state = {
        count: 0,
        childCount:0,
    };
    addCount = () => {
        this.setState({
        count: this.state.count + 1
        });
    };
    addChildCount=()=>{
        this.setState({
        childCount: this.state.childCount + 1
        });
    }
    render() {
        return (
        <div>
            <h1>Parent:{this.state.count}</h1>
            <button onClick={this.addCount}>+Parent</button>
            <button onClick={this.addChildCount}>+Child</button>

            <Child
            count={this.state.childCount}
            addParentCount={this.addCount}
            addChildCount={this.addChildCount} />
        </div>
        );
    }
    }

    export default Parent;

## 子

    import React, { Component } from "react";

    class Child extends Component {
    render() {
        const { count, addParentCount, addChildCount } = this.props;
        return (
        <div>
            <h1>Child:{count}</h1>
            <button onClick={addParentCount}>+Parent</button>
            <button onClick={addChildCount}>+Child</button>
        </div>
        );
    }
    }

    export default Child;
