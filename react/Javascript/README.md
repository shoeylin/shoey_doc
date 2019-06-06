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

