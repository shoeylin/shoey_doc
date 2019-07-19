# 重要的KEY

## List.js
import React, { Component } from "react";
import Item from "./Item";

class List extends Component {
  state = {
    list: [
      { label: "name", data: "React Lesson" },
      { label: "price", data: 3200 },
      { label: "videos", data: 60 },
      { label: "teacher", data: "scars" }
    ]
  };
  removeFirst = () => {
    this.setState({
      list: this.state.list.slice(1)
    });
  };
  render() {
    return (
      <div>
      <!-- key 是div內的唯一值 -->
          <div>
        {this.state.list.map(
            ({ label, data }) => (
          <Item key={label} text={`${label}: ${data}`} />
        ),
        )}
        </div>
        <div>
        {this.state.list.map(
            ({ label, data }) => (
          <Item key={label} text={`${label}: ${data}`} />
        ),
        )}
        </div>
        <button onClick={this.removeFirst}>-</button>
      </div>
    );
  }
}

export default List;

## Item.js

import React, { PureComponent } from "react";

class Item extends PureComponent {
  render() {
    console.log("render", this.props.text);
    return <li>{this.props.text}</li>;
  }
}

export default Item;
