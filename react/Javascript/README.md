var let const

scope
block

1. function scope / global scope
2. scope out > in

var: scope
let/const: block

---

# main

    import math, { PI as P, ROOT_2 as R} from './math';

    console.log( P, R);

---

# math

    const math = {
        double : x => x * 2,
        square : x => x * x,
        area : (w, h) => w * h
    };

    // named export

    export const PI = 3.1415;

    export const ROOT_2 = 1.414;

    export default math;

# class

    class Dog extends Animal{
        // constructor(){
        //     this.age = 0;
        // }
        age = 0;
        // bark(){
        //     console.log('woof');
        // }
        bark = () => console.log('woof');
    }

    const spot = new Dog();

    method (function)
    property (variable)

    babel class-properties

# 解構

    const point = [1, 2, 3];

    const x = point[0];

    const y = point[1];

    const z = point[2];

    const [x, y, z, w = 0] = point;
    //rest [2,3]

    const point = { x: 1, y: 2, z: 3 };

    const x = point.x;
    const y = point.y;
    const z = point.z;

    const { x, ...rest } = point;
    //rest{y:2,z:3}

    const { w = 0 } = point;

    const { w: px = 0 } = point;

    console.log(pw);

    const draw = ({ x, y, z }) => {
    //canvas
    };

# 字串模板

    const message = "I am " + age + " years old";

    const message = `
        I am ${age * 2} years old
        I have ${cars} cars and ${house} houses.
        Would you marry me ? (Y/N)
        `;

# async / await

    callback
    callSomeAPI(function(result) {
    console.log(result);
    });

    //promise
    callSomeAPI().then(result => {
    console.log(result);
    });

    //async/await
    const start = async () => {
    const result = await callSomeAPI();
    };
    start();

#　列表渲染 map array(ren 成元素)

import React, { Component } from "react";
import Item from "./Item";

const steps = ["Learn JavaScript", "Learn React", "Make Money", "Buy a House"];

class List extends Component {
render() {
return (

<div>
{steps.map((text, idx, array) => (
<Item>{text({ idx })}</Item>
))}
</div>
);
}
}

export default List;

## 物件(ren 成元素)

import React, { Component } from "react";
import Item from "./Item";

// 如果要排序宣告成陣列 物件會不排序
const info = {
name: "React Lesson",
price: 3200,
videos: 60,
teacher: "scars"
};
class List extends Component {
render() {
return (

<div>
{Object.keys(info).map(key => {
const value = info[key];
return (
<Item>
{key}:{value}
</Item>
);
})}
</div>
);
}
}

export default List;

## 陣列排序

import React, { Component } from "react";
import Item from "./Item";

// 如果要排序宣告成陣列 物件會不排序
const info = [
{ label: "name", data: "React Lesson" },
{ label: "price", data: 3200 },
{ label: "videos", data: 60 },
{ label: "teacher", data: "scars" }
];
class List extends Component {
render() {
return (

<div>
{info.map(({ label, data }) => (
<Item>
{label}: {data}
</Item>
))}
</div>
);
}
}

export default List;

# rander json 檔案樹

## App.js

    import React, { Component } from "react";
    import files from "./files.json";
    import './style.css'

    import FilesNode from "./FilesNode";

    class App extends Component {
    render() {
        return (
        <div>
            <FilesNode name={files.name} files={files.files} />
            <FilesNode {...files} />
        </div>
        );
    }
    }

    export default App;

## FilesNode.js

    import React, { Component } from "react";

    class FilesNode extends Component {
    state = {
        open: false
    };
    toggle = () => {
        this.setState({
        open: !this.state.opne
        });
    };
    render() {
        const { name, files } = this.props;
        const { open } = this.state;
        if (!files) {
        return <li>{name}</li>;
        }
        return (
        <div>
            <div className={`folder ${open ? "open" : ''}`} onClick={this.toggle}>
            {name}
            </div>
            {!open ? null : (
            <ul>
                {files.map(file => (
                <FilesNode {...file} />
                ))}
            </ul>
            )}
        </div>
        );
    }
    }

    export default FilesNode;

## styly.css

    .folder {
        color: blue;
        cursor: pointer;
    }

    .folder:before {
        content: '📁';
    }

    .folder.open:before{
        content: '📂';
    }

# Array 介紹

    //map
    const array = [1, 2, 3, 4, 5, 6];

    const result = array.map(function(elemm, idx, arr) {
    return <li>{elem * 2}</li>;
    });

    const result = array.map((elemm, idx, arr) => <li>{elem * 2}</li>);

    //filter
    const result = array.filter((elem, idx, arr) => {
    return elem % 2 === 0;
    });

    // reduce 總和
    // initValue
    //             0  1  3  6 10 15
    const array = [1, 2, 3, 4, 5, 6];
    //             1  3  6 10 15 21
    const result = array.reduce((accumulator, elem, idx, arr) => {
    return accumulator + elem;
    }, 0);

    const result = array.reduce((acc, value) => acc + value, 0);

    //以上三個不會去改變array的值

    //以下會改變array的值 mutate
    const array = [1, 2, 3, 4];
    array.pop(); //後面彈 4
    array.push(5); //後面塞 5
    array.shift(); //前面彈 1
    array.unshift(0); //前面塞 0

    array.reverse(); //倒過來
    const newArr = array.slice(); //複製一個新的array
    newArr.reverse();

    const newArr=[...array] //ES6這樣也是複製

    array.sort()
    array.splice()

    //mutate 很少用到 通常會先複製

# 表單控制

    import React, { Component } from "react";

    class App extends Component {
    state = {
        text: "abc"
    };
    onChange = e => {
        this.setState({
        text: e.target.value
        });
    };
    render() {
        const { text } = this.state;
        return (
        <div>
            <input type="text" value={text} onChange={this.onChange} />
            <h2>{text}</h2>
        </div>
        );
    }
    }

    export default App;

## 表單 數字處理

    import React, { Component } from "react";

    class App extends Component {
    state = {
        count: 10
    };
    onChange = e => {
        this.setState({
        count: parseInt(e.target.value, 10)
        });
    };
    render() {
        const { count } = this.state;
        return (
        <div>
            <input type="number" value={count} onChange={this.onChange} />
            <h2>next:{count + 1}</h2>
        </div>
        );
    }
    }

    export default App;

## 下拉選單 select

    import React, { Component } from "react";

    const relations = [
    { label: "父", value: 0 },
    { label: "母", value: 1 },
    { label: "子", value: 2 },
    { label: "女", value: 3 },
    { label: "妻", value: 4 },
    { label: "友", value: 5 }
    ];
    class App extends Component {
    state = {
        rel: `${relations[0].value}`
    };
    onChange = e => {
        this.setState({
        rel: e.target.value
        });
    };
    render() {
        const { rel } = this.state;
        return (
        <div>
            <select value={rel} onChange={this.onChange}>
            {relations.map(relations => (
                <option value={relations.value} key={relations.label}>
                {relations.label}
                </option>
            ))}
            </select>
            <h1>{relations.find(r => `${r.value}` === rel).label}</h1>
        </div>
        );
    }
    }

    export default App;

## 單選與多選

    import React, { Component } from "react";

    class App extends Component {
    state = {
        gender: "male",
        like: {
        male: false,
        female: false
        }
    };
    onChangeGender = e => {
        this.setState({
        gender: e.target.value
        });
    };
    onChangeLike = e => {
        const key = e.target.value;
        this.setState(state => ({
        like: {
            ...state.like,
            [key]: !state.like[key]
        }
        }));
    };
    render() {
        const { gender, like } = this.state;
        return (
        <div>
            <div>
            Your gender:
            <input
                type="radio"
                value="male"
                onChange={this.onChangeGender}
                checked={gender === "male"}
            />
            <label>Male</label>
            <input
                type="radio"
                value="female"
                onChange={this.onChangeGender}
                checked={gender === "female"}
            />
            <label>Female</label>
            </div>
            <div>
            You Like:
            <input
                type="checkbox"
                value="male"
                onChange={this.onChangeLike}
                checked={like.male}
            />
            <label>Male</label>
            <input
                type="checkbox"
                value="female"
                onChange={this.onChangeLike}
                checked={like.female}
            />
            <label>Female</label>
            </div>
            <div>
            <pre> {JSON.stringify(this.state, null, 2)}</pre>
            </div>
        </div>
        );
    }
    }

    export default App;

# 檔案上傳與圖片預覽

    import React, { Component } from "react";

    class App extends Component {
    state = {
        file: null,
        img: ""
    };
    onChange = e => {
        const file = e.target.files.item(0);
        const fr = new FileReader();
        fr.addEventListener("load", this.fileLoad);
        fr.readAsDataURL(file);
        this.setState({
        file
        });
    };
    fileLoad = e => {
        this.setState({
        img: e.target.result
        });
    };
    submit = () => {
        //json base64
        // axios.post("/img", { img: this.state.img });


        // multipart
        // const form = new FormData();
        // form.append(this.state.file);
        // axios.post("/img", { form });
    };
    render() {
        return (
        <div>
            <input type="file" onChange={this.onChange} />
            <img width="100%" src={this.state.img} />
            <button onClick={this.submit}>Submit</button>
        </div>
        );
    }
    }

    export default App;

# 計算薪水範例

import React, { Component } from "react";

class Salary extends Component {
state = {
month: 0,
startTime: Date.now(),
currentTime: Date.now()
};
componentDidMount() {
setInterval(() => {
this.setState({
currentTime: Date.now()
});
}, 100);
}
onChangeMonth = e => {
this.setState({
month: parseInt(e.target.value, 10)
});
};
render() {
const { month, startTime, currentTime } = this.state;
const time = (currentTime - startTime) / 1000;
return (
<div>
<label>月薪</label>
<input type="number" value={month} onChange={this.onChangeMonth} />
<br />
<label>時薪</label>
<input type="number" value={month / 240} />
<br />
<label>分薪</label>
<input type="number" value={month / 240 / 60} />
<br />
<label>秒薪</label>
<input type="number" value={month / 240 / 60 / 60} />
<br />
<label>經過時間</label>
<spen>{time.toFixed(1)}</spen>
<br />
<label>偷取薪水</label>
<spen>{((time \* month) / 240 / 60 / 60).toFixed(2)}</spen>
</div>
);
}
}

export default Salary;

# 單人聊天室

    import React, { Component } from "react";
    import stlye from "./ChatRoom.module.css";

    class ChatRoom extends Component {
    state = {
        text: "",
        content: []
    };
    onChangeText = e => {
        this.setState({
        text: e.target.value
        });
    };
    submit = e => {
        e.preventDefault();
        const { text, content } = this.state;
        this.setState({
        text: "",
        content: [{ id: Date.now(), text }, ...content]
        });
    };
    render() {
        const { text, content } = this.state;
        return (
        <div>
            <form onSubmit={this.submit}>
            <input tpye="text" value={text} onChange={this.onChangeText} />
            <button type="submit" onClick={this.submit}>
                Submit
            </button>
            </form>
            <ul>
            {content.map(item => (
                <li key={item.id}>{item.text}</li>
            ))}
            </ul>
        </div>
        );
    }
    }

    export default ChatRoom;
