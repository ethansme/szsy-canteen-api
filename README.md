<p align="center">
    <a href="https://github.com/nczitzk/szsy-canteen-api" target="_blank">
        <img width="300" src="image/szsy-canteen-api-logo.png" alt="szsy-canteen-api-logo">
    </a>
</p>

<p align="center">
    <a href="http://szsy.cn" target="_blank">
        <img src="https://img.shields.io/badge/SZSY-2018-blue.svg?style=flat-square" alt="szsy-2018">
    </a>
    <a href="https://github.com/nczitzk/szsy-canteen-api/blob/master/LICENSE" target="_blank">
        <img src="https://img.shields.io/github/license/nczitzk/szsy-canteen-api.svg?style=flat-square" alt="szsy-canteen-api-license">
    </a>
    <a href="https://github.com/nczitzk/szsy-canteen-api" target="_blank">
        <img src="https://img.shields.io/github/languages/code-size/nczitzk/szsy-canteen-api.svg?style=flat-square" alt="szsy-canteen-api-size">
    </a>
</p>

<p align="center">
    English | 
    <a href="README-zh_CN.md">简体中文</a> | 
    <a href="README-zh_TW.md">正體中文</a>
</p>

This is an API based on the online ordering system owned by [Shenzhen Experimental School](http://www.szsy.cn/). Developers can bypass the originally unnecessary but unavoidable value-passing codes, free themselves from the complicated and incomprehensible logic, and conduct online ordering interaction through the more simplistic interfaces of the packaged __szsy-canteen API__.

__szsy-canteen API__ is written in __Golang__ which can ensure its efficiency and performance. The API is suitable for the secondary development of the online ordering system 💗

## Why use this API

In short, the API is designed to be simple, data-saving and efficient.

For the original intention was aimed at reducing the cost of mobile application development of the school online ordering system, the __logic-optimization__ and __data-saving__ were regard as important indicators in design of the API. Here is an example. The API provides an interface for obtaining a list of "orderable dates" directly, which save the overhead of operations such as "gaining all dates including unnecessary ones" (waste of data traffic) and "judging whether the dates can be ordered" (complex logic).

Of course, some parts of the API is roughly-designed because the developing time is limited and my lack for programming ability. Welcome new contribution, and hope that you can let me know if you made a new implementation of the ordering system and I will update your repo link in the README 😄 

## How to configure

If you are a developer on 64-bit Windows or Linux, you can download the pre-compiled package of your specific system. Double-click to execute it and the API will be enabled on port `2018` on the local machine by default.

If you attempt to compile it and run it yourself, you need to install the developing environment required by __Golang__, do necessary configurations and download the complete source files to compile. You can execute the following command to generate an executable file for your platform.

```
go build main.go
```

## Get started

The __szsy-canteen API__ provides the following interfaces:

- `login` : Login to the online ordering system to get basic information and the key for further operations

- `dates` : Obtain the list of all the dates that can be ordered currently (i.e. the dates before the order deadline)

- `menu` : Get details and order status of the menu for the specific date

- `order` : Submit the order request for the specific date

- `logout` : Log out of the online ordering system

The URL of the API in default configuration should be

```
http://localhost:2018/
```

### Login to online ordering system

Write username and password encrypted with the __MD5__ algorithm to the JSON field. Here is an example.

``` json
{
    "账号":"2152000",
    "密码":"80d9ba95cce518bf747bda3bc98faef8"
}
```

> For security reasons, the encrypting process of the password with MD5 is handed over to the requester on design, that is, you need to encrypted password string yourself and put it into the `密码`(password) item.

Submit to URL with POST method

```
http://localhost:2018/login
```

The API will return a JSON as follows

``` json
{
    "姓名":"XXX",
    "余额":"666.66元",
    "口令":"6E19822908B2FEA56F7BF,bnfh4fsigr5z04ebdeo"
}
```

> In fact, the value of the `口令`(the key) will be much longer than the one in the example. Here, the key has been shortened for typography.

### Obtain orderable dates list

Write the key returned at login to the JSON field

``` json
{
    "口令":"6E19822908B2FEA56F7BF,bnfh4fsigr5z04ebdeo"
}
```

Submit to URL with POST method

```
http://localhost:2018/dates
```

You will get return data as follows

``` json
{
    "可订日期":["2018-06-26","2018-06-27","2018-06-28","2018-06-29",
    "2018-07-02","2018-07-03","2018-07-04","2018-07-05",
    "2018-07-06","2018-07-09","2018-07-10","2018-07-13"]
} 
```

> The date list will contain __all orderable dates__, i.e. the API will also try requesting menus for the next month without missing orderable dates in the next month. The school updates at most two months' menus once. You can leave the job to the __szsy-canteen API__ at ease.

### Obtain the specific date menu

Similarly, write the key to the JSON field.

``` json
{
    "口令":"6E19822908B2FEA56F7BF,bnfh4fsigr5z04ebdeo"
}
```

Submit to the URL with the "specific date" parameter with the POST method

```
http://localhost:2018/menu/?date=2018-07-06
```

You can get return JSON as follows

``` json
{
    "不订餐":null,
    "早餐":
        [
            ["0","套餐","早餐套餐","5.00","0"],
            ["1","牛奶","学生奶","2.04","0"],
            ["2","蛋类","五香鸡蛋","1.40","0"],
            ["3","点心","巧克力面包","1.40","0"],
            ["4","点心","玉米酥","1.40","0"],
            ["5","点心","南瓜米糕","1.40","0"],
            ["6","点心","鲜肉包","1.40","0"],
            ["7","粉面类","水饺","1.60","0"]
        ],
    "午餐":
        [
            ["0","套餐","午餐套餐","12.00","0"],
            ["1","水果","杨桃","1.30","0"],
            ["2","菜肴","蒜茸炒黄芽白","1.80","0"],
            ["3","菜肴","炒小瓜片","1.80","0"],
            ["4","菜肴","剁椒蒸鱼(微辣)","4.40","0"],
            ["5","菜肴","肉末豆角","3.60","0"],
            ["6","菜肴","土豆烧肉","5.60","0"],
            ["7","菜肴","香辣翅根2个","6.70","0"]
        ],
    "晚餐":null
}
```

Each dish includes the following information points

``` json
["编号","类型","菜名","单价","订餐个数"]
```

The item of `不订餐`(not order) should be a `null` value in general, as the example above. Of course, if you choose not to order `早餐`(breakfast) or `午餐`(lunch) , then the value of `不订餐`(not order) will supposed to be

``` json
"不订餐":["早餐","午餐"]
```

> July 10, 2018 is Friday in this example, so the menu for dinner is missing, i.e. the value of `晚餐`(dinner) is `null`

### Submit an order request

Write the key obtained at login and the order information for breakfast, lunch and dinner to the JSON field.

``` json
{
    "口令":"6E19822908B2FEA56F7BF,bnfh4fsigr5z04ebdeo",
    "早餐":"套餐",
    "午餐":"0,0,2,0,0,1,0,1",
    "晚餐":"不订餐"
}
```

There are three formats for order information, and the example has given all the cases.

1. `套餐`(set menu) (i.e. equivalent to "1,0,0,0,0,0,0,0")

2. `不订餐`(not order) (i.e. equivalent to check the "not order" checkbox)

3. Write the number you would like to order of eight dishes with commas as separators

Therefore, in this case, this student chose a "套餐"(set menu) for his breakfast. For lunch, he chose 2 servings of "蒜茸炒黄芽白", 1 serving of "肉末豆角" and 1 serving of "香辣翅根2个", and choose "不订餐"(not order) for dinner.

> On Friday night, i.e. the menu for dinner is missing, you should serve the `晚餐`(dinner) item a `null` value instead of omitting the `晚餐`(dinner) item. Here, in order to show the `不订餐`(not order) field, `晚餐`(dinner) is filled in with it, but it should be `null` 😂

> You may find that the first dish in the eight dishes is `套餐`(set menu). In fact, if one of the dishes following is not `0`, the value of `套餐`(set menu) will be covered by `0` regardless of what the value is.

Then use the POST method to submit to the URL with the "specific date" parameter

```
http://localhost:2018/order/?date=2018-07-06
```

You may get types of return information

1. `提交成功` (the request is submitted successfully, and the order information is written to the system)

2. `账户异常` (this account is frozen or stuck in other abnormal situations, often means in arrear)

3. `超过订餐时间` (the date is over the order deadline, you can avoid this error by checking whether the date in the "orderable date" list before submitting)

4. `提交失败` (failed for the reason that network timeout or school server goes down, etc.)

> The time it takes for the school server to deal with the order information is far more than the time needed normally, so you usually have to wait a few seconds to get feedback (students in SZSY should be impressed) __Do not submit too often__, you might fail, and even worse, get the wrong return information.

### Logout online ordering system

Write the key obtained at login to the JSON field

``` json
{
    "口令":"6E19822908B2FEA56F7BF,bnfh4fsigr5z04ebdeo"
}
```

Submit to URL with POST method

```
http://localhost:2018/logout
```

You will logout after reading `登出成功`.

### Error about key expiration

__szsy-canteen API__ error messages are fairly legible, but there is a special explanation for this.

If you are using a key to request dates list, obtain menu, or submit request with receiving `口令错误或过期` error message, it means your key may have expired. In this case, all you need to do is re-login to get the new key.

### Tips for ordering multiple servings

The school online ordering system limits 3 servings for each dish maximum, which is conducted by the JavaScript script on the website. However, there is no such restriction if using the API. In theory, you can order more than 3 servings.

The practical tests show that the success rate of 5 servings or less is relatively higher, and the submissions will be accepted in most cases. If the number is too exaggerated, the success rate falls to zero basically, and there will be a `提交失败`(submit failed) error returning. The maximum number which successfully submitted I have tried is 20, yes, 20 servings of 卤鸡腿.

Of course, this is not a bug. 20 servings of 卤鸡腿 means you have to pay them as 20 servings (I won't state this here if not so)

This text added here is a piece of advice: As a developer, you can offer an entry that can modify the number of servings in an explicit way when developing the client. After all, adding a dish of "净荤菜" occasionally is also a sort of pleasure.

## Acknowledgement

[@xlfjn](https://github.com/xlfjn)

[@GrakePch](https://github.com/GrakePch)

Thanks to the above two for your technical and moral supports during development 💗