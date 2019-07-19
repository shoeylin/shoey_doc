# 方案一

ex: css

## MyButton.module.css

.btn {
background: #cde;
border-radius: 20px;
font-size: 2em;
width: 250px;
text-align: center;
padding: 0.5em 0;
margin: 0.5em;
}

.btnText{
color: red;
}

## MyButton.js

import React, { Component } from "react";

import style from "./MyButton.module.css";
const MyButton = ({ children }) => (
<button className={style.btn}>
<span className={style.btnText}>{children}</span>
</button>
);

export default MyButton;

---

# 方案二

## yarn add styled-components

## MyButton.js

import React, { Component } from "react";

import styled from "styled-components";

// import style from "./MyButton.module.css";

const Button = styled.button`background: #cde; border-radius: 20px; font-size: 2em; width: 250px; text-align: center; padding: 0.5em 0; margin: 0.5em;`;

const ButtonText = styled.span`color: red;`;

const MyButton = ({ children }) => (
<Button>
<ButtonText>{children}</ButtonText>
</Button>
);

export default MyButton;


# 方案三 sass
yarn add node-sass

## MyButton.module.sass
 .btn {
    background: #ecc;
    border-radius: 20px;
    font-size: 2em;
    width: 250px;
    text-align: center;
    padding: 0.5em 0;
    margin: 0.5em;
    &:after{
        content: '!!!!!'
    }
}


## MyButton.js
import React, { Component } from "react";

import style from "./YourButton.module.scss";

const YourButton = ({ children }) => (
  <button className={style.btn}>{children}</button>
);
export default YourButton;
