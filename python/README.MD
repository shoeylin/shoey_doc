# 莫煩 python3

## 循環

    condition =1
    while condition < 10:
        print(condition)
        condition = condition+1


    while True:
        print("I'm True")


    example_list = [1,2,3,4,5,6,7,12,543,876,12,3,2,5]
    for i in example_list:
    print(i)


    for i in range(1,10,1):
    print(i)

    for i in range(1,10,3):
    print(i)

## 判斷

    x = 1
    y = 2
    z = 3

    if x>y>z:
        print('x is greater than y')
    else:
        print('x is less than y')

    if x!=y:
    if x==y:
    if x>=y:
    if x<=y:

-----

    x = -2
    y = 2
    z = 3

    if x>1:
        print('x > 1')
    elif x<-1:
        print('x < -1')
    elif x<1:
        print('x < 1')    
    else:
        print('x = 1')

## 函數def

    def function():
        print('This is a function')
        a=1+2
        print(a)

    function()

    def fun(a, b):
    c = a-b
    print('the c is',c)

    fun(a=2,b=10)

    def sale_car(price, brand, colour='red', is_second_hand=True):
    print('price:', price,
        'colour:',colour,
        'brand:',brand,
        'is_second_hand:',is_second_hand)

    sale_car(131000,'BMW',colour='blue')

## 文件寫入

    text='This is my first test.\nThis is next line.\nThis is last line.'

    print(text)

    my_file = open('my_file.txt','w')
    my_file.write(text)
    my_file.close()

---
    重複寫入文件後方

    append_text='\nThis is a appended file.'

    my_file = open('my_file.txt','a') #applend
    my_file.write(append_text)
    my_file.close()
---
    讀取文件
    file = open('my_file.txt','r') #read
    content = file.read() #file.readline() file.readlines()
    print(content)

## class 類

    class Calculator:
        name = 'Good calculator'
        price = 18
        def add(self, x, y):
            print(self.name)
            result = x + y
            print(result)
        def minus(self, x, y):
            result = x - y
            print(result)
        def times(self, x, y):
            print(x*y)
        def divide(self, x, y):
            print(x/y)

    calcul = Calculator()
    print(calcul.name)
    print(calcul.price)
    print(calcul.add(10,11))
    print(calcul.minus(10,11))
    print(calcul.times(10,11))
    print(calcul.divide(13,2))

---

### init

    class Calculator:
        name = 'Good calculator'
        price = 18
        def __init__(self, name, price, hight=18, width=20, weight=19):
            self.name=name
            self.price=price
            self.h=hight
            self.wi=width
            self.we=weight
            self.add(1,2) # 生成時就執行
        def add(self, x, y):
            result = x + y
            print(result)
        def minus(self, x, y):
            result = x - y
            print(result)
        def times(self, x, y):
            print(x*y)
        def divide(self, x, y):
            print(x/y)

    calcul = Calculator('Bad calculator',12)
    print(calcul.name)
    print(calcul.price)
    print(calcul.add(10,11))
    print(calcul.minus(10,11))
    print(calcul.times(10,11))
    print(calcul.divide(13,2))
    print(calcul.price)
    print(calcul.h)
    print(calcul.wi)
    print(calcul.we)

## input

    a_input = int(input('Please give a number: ')) ## return a string
    if a_input == '1':
        print('This is a good one')
    elif a_input == 2:
        print('See you next time')
    else :
        print('Good luck', a_input)

## 元組列表 tuple list

# tuple list
    a_tuple = (12,3,5,15,6)
    another_truple = 2,4,6,7,8

    a_list = [12,3,67,7,82]
            #[ 0,1, 2,3, 4]
    for content in a_tuple:
        print(content)

    for content in another_truple:
        print(content)

    for content in a_list:
        print(content)

    for index in range(len(a_list)):
        print('index',index,'number in list=',a_list[index])
        a_list[1]

---
    list處理

    a=[1,2,3,4,5,6,7,8,-1]

    a.append(0) # 添加在後面 0
    a.insert(1,0) # 指定位置+0
    a.remove(3) # 移除指定位置元素
    print(a)
    print(a[0]) # 第一位
    print(a[-1]) # 最後一位
    print(a[0:3]) # 第0到3位
    print(a[-3:]) # 最後三個數字

    print(a.index(8)) # 第一個出現的數字為哪個索引
    print(a.count(0)) # 出現的數字的次數
    a.sort()
    print(a)
    a.sort(reverse=True) # 倒序 預設false
    print(a)

## 多維列表

    a = [1,2,3,4,5]
    multi_dim_a = [[1,2,3],
                    [2,3,4],
                    [3,4,5]]
                
    print(a[1])
    print(multi_dim_a[2][2])

## 字典 d = {}

    a_list = [1,2,3,5,4,5,4]
    d = {'apple':[1,2,3,4],'pear':2,'orange':{1:3,3:'a'}}
    d2 = {1:'a','c':'b'}

    print(d['apple'])
    print(a_list[0])

    del d['pear']
    print(d)

    d['b'] = 20
    print(d)

    print(d['orange'][3])

## 載入模塊 import

    import time as t
    import time
    # from time import time,localtime
    from time import *

    # print(time.localtime())
    # print(t.localtime())

    print(localtime())
    print(time())
    print(gmtime())

## 自己的模塊

    m1.py----

    def printdata(data):
    print('I am m1')
    print(data)

    hello.py----

    import m1

    m1.printdata('I am python1')

### site-packages
    下載的模塊都在這
    C:\Users\shoey\AppData\Local\Programs\Python\Python37\Lib\site-packages

## continue & break

    while True:
        b = input('type something ')
        if b == '1':
            break
            # continue
        else:
            pass
        print('still in while')

    print('finish run')

## 錯誤處理 try

    try:
        file = open('eeee','r+')
    except Exception as e:
        print('there is no file named as eeee')
        response = input('do you want to create a new file (y/n) ? ')
        if response == 'y':
            file = open('eeee','w')
        else:
            pass
    else:
        file.write('ssss')
    file.close()

    try有錯誤就執行except 
    try沒錯誤就執行else

## zip lambda map

    a = [1,2,3]
    b = [4,5,6]

    # zip
    print((zip(a,b)))
    print(list(zip(a,b)))

    for i,j in zip(a,b):
        print(i/2,j*2)

    print(list(zip(a,a,b)))

    def fun1(x,y):
        return (x+y)

    print(fun1(2,3))

    # lambda
    fun2 = lambda x,y:x+y
    print(fun2(2,3))

    # map
    print(map(fun1,[1],[2]))
    print(list(map(fun1,[1],[2])))
    print(list(map(fun1,[1,3],[2,5])))

## copy & deepcopy

    import copy
    a = [1,2,3]
    b = a
    print(id(a))
    print(id(b))

    b[0]=11
    print(a)

    a[1]=22
    print(b)

    print(id(a)==id(b)) #true
    c = copy.copy(a)
    print(id(a)==id(c)) #false

    c[1]=222222
    print(c)
    print(a)

    a=[1,2,[3,4]]
    d=copy.copy(a)
    print(id(a)==id(d)) #false
    print(id(a[2])==id(d[2])) #true

    a[0]=11
    print(a)
    print(d)

    a[2][0]=333
    print(a)
    print(d)

    e=copy.deepcopy(a)
    print(id(e[2])==id(a[2])) #false

## 多線程 多進程 tkinter窗口

## pickle

    import pickle

    # a_dict={'da':111,2:[23,1,4],'23':{1:2,'d':'sad'}}

    # file= open('pickle_example.pickle','wb')
    # pickle.dump(a_dict,file)
    # file.close()

    # file = open('pickle_example.pickle','rb')
    # a_dict = pickle.load(file)
    # file.close()
    # print(a_dict)

    # with可不用關閉file
    with open('pickle_example.pickle','rb') as file:
        a_dict1=pickle.load(file)

    print(a_dict1)

