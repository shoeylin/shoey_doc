# React樣式

## SASS範例

    #header{
        #logo{
            display: none;
            color: $main-color;
        }
        #block{
            display: none;
            @include size(20px,40px);
            @extend #logo;
        }
    }

---
## 轉css

    #hear #logo{
        display: none;
        color: red;
    }

    #header #block{
        wideth: 20px;
        height: 40px;
        display: block;
        display: none'
        color: red;
    }

---