#app-svr

##API
* /v1/hongid/register/smscode [post]: 获取手机验证码
```
    http body:
    {
        tel: "18610889275"
    }

    http status:
        400: 参数错误
        {
            code: -201,
            message: "sjdfljojfow",
            errors: [
                {
                    field: "sss",
                    message: "sldjfls"
                },
                {
                    field: "sss",
                    message: "sldjfls"
                }
            ]
        }
        200:
        {
            code: 0,
            message: "sfsafasf",
            description: "safewfsdf",
        }

        自定义返回code：
         -202：参数错误
         -302：向短信平台发送验证码失败
         0：一切ok

```
* /v1/hongid/register/smscode/atk [post]: 注册第一步完成
```
    http body:
    {
        tel: "18610889275",
        code: "748573"
    }
    http status:
        400: 参数错误
                {
                    code: -201,
                    message: "sjdfljojfow",
                    errors: [
                        {
                            field: "sss",
                            message: "sldjfls"
                        },
                        {
                            field: "sss",
                            message: "sldjfls"
                        }
                    ]
                }
        200:
                {
                    code: 0,
                    message: "sfsafasf",
                    description: "safewfsdf",
                }

        自定义返回code：
         -202：参数错误
         -303：验证码错误
         -400：系统内部错误
         -301：手机号已经注册
         0：一切ok

```
* /v1/hongid/register[post]: 绑定昵称&密码
```
    http body:
    {
        tel: "18610889275",
        nickname: "sss",
        password: "lsjdfoiwejfoj"
    }
    400: 参数错误
                    {
                        code: -201,
                        message: "sjdfljojfow",
                        errors: [
                            {
                                field: "sss",
                                message: "sldjfls"
                            },
                            {
                                field: "sss",
                                message: "sldjfls"
                            }
                        ]
                    }
    500: 服务器内部错误
    200:
                    {
                        code: 0,
                        message: "sfsafasf",
                        description: "safewfsdf",
                    }
                    自定义返回code：
                             -202：参数错误
                             -400：系统内部错误
                             0：一切ok
```
* /v1/hongid/login [post]: 登录
```
    http body:
    {
        username: "18610889275",
        password: "lsjdfoiwejfoj"
    }
    400: 参数错误
                    {
                        code: -201,
                        message: "sjdfljojfow",
                        errors: [
                            {
                                field: "sss",
                                message: "sldjfls"
                            },
                            {
                                field: "sss",
                                message: "sldjfls"
                            }
                        ]
                    }
    500: 服务器内部错误
    403:
     {
                            code: 0,
                            message: "sfsafasf",
                            description: "safewfsdf",
                        }
    200:
                    {
                        id: 3434,
                        uuid: "asdfsdfsaf",
                        hongid: "343434",
                        nickname: "sdfasdf"
                    }
                    自定义返回code：
                             -202：参数错误
                             -304: 用户名错误
                             -305: 密码错误
                             -400：系统内部错误
                             0：一切ok
```
