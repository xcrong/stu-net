## 一、STU-NET 程序说明

```sh
(base) PS C:\Users\zlz_g> stu-net
stu-net 是一款可以在命令行中运行的校园网工具。
可以用于无头环境下校园网的登录、登出、状态查询、流量查询。

Usage:
    stu-net [flags]
    stu-net [command]

Available Commands:
    check       查看当前是否登录
    completion  Generate the autocompletion script for the specified shell
    flux        查询剩余流量
    help        Help about any command
    login       通过账号和密码登录到校园网
    logout      登出校园网
    restore     重置配罝文件

Flags:
    -h, --help          help for stu-net
    -i, --interactive   使用交互模式

Use "stu-net [command] --help" for more information about a command.

Examples:
    stu-net -i # 进入交互模式
    stu-net login \
    -u 20xxxxx \
    -p my_pwd \
    -s # 指定用户名和密码登录，同时将之保存到配置
    stu-net flux # 查询流量

``

## 二、校园网接口说明

整个校园网的登录、登出、流量查询、状态查询都基于 `POST` 请求。

- 登录、登出、状态查询的 URL 都是： `https://a.stu.edu.cn:444/ac_portal/login.php`
- 流量查询的 URL 是： `https://a.stu.edu.cn:444/ac_portal/userflux`

它们 POST 的 payload 不同，下面详述。

### 1. 登录

登录成功后，除了会返回一个 表示成功的JSON；在响应头中，还会设置 cookie，这个 cookie 在流量查询时需要用到.

响应头节选： `Set-Cookie:[AUTHSESSID=cd726d747d9e; HttpOnly;Secure;]`

```py
payload = {
    "opr": "pwdLogin",
    "userName": Username, # 20abcdef
    "pwd": Password, # mypassword
    "ipv4or6": "",
    "rememberPwd": "0"
}
```

### 2. 登出

```py
payload = {
    "opr": "logout",
    "ipv4or6": ""
}
```

### 3. 流量查询

进行流量查询时，需要用到登录时获得的 cookie， 将 cookie 添加到请求头中，然后直接 POST 流量查询的 URL 即可，不需要参数。


### 4. 状态查询

```py
payload = {
    "opr": "online_check"
}
```
