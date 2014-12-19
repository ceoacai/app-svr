#app-svr

##API
```
通用返回结构
    参数错误：
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
    逻辑返回:
    {
        code: 0,
        message: "sfsafasf",
        description: "safewfsdf",
    }
```
* /v1/hongid/register/smscode [post]: 获取手机验证码
```
    json body:
    {
        tel: "18610889275"
    }

    http status:
        400 请求参数错误 [-202:参数错误]
        200 OK [0:ok -302:验证码发送失败]
```
* /v1/hongid/register/smscode/atk [post]: 注册第一步完成
```
    json body:
    {
        tel: "18610889275",
        code: "748573"
    }
    http status:
        400: 请求参数错误 [-202:参数错误]
        500: 服务器内部错误 [-400:内部错误]
        200: OK [0:ok, -303:验证码错误, -301: 手机号已经被注册]
```
* /v1/hongid/register[post]: 绑定昵称&密码
```
    json body:
    {
        tel: "18610889275",
        nickname: "sss",
        password: "lsjdfoiwejfoj"
    }
    400: 请求参数错误 [-202:参数错误]
    500: 服务器内部错误 [-400:内部错误]
    200: OK [0:ok]
```
* /v1/hongid/login [post]: 登录
```
    json body:
    {
        username: "18610889275",
        password: "lsjdfoiwejfoj"
    }
    400: 请求参数错误 [-202:参数错误]
    500: 服务器内部错误 [-400:内部错误]
    403: 禁止请求 [-304: 用户名错误, -305: 密码错误]
    200: OK [0: ok]
        json body
        {
            id: 3434,
            uuid: "asdfsdfsaf",
            hongid: "343434",
            nickname: "sdfasdf"
        }
```
* /v1/memberCards/card/:card/bind/:owner [post] :绑定会员卡
```
    400: 请求参数错误 [-202:参数错误]
    500: 服务器内部错误 [-400:内部错误]
    200: OK [0: ok, -306:会员卡不正确, -307:会员卡已被其他会员绑定]
```
* /v1/memberCards/card/:card/unbind [post] :解除绑定会员卡
```
    400: 请求参数错误 [-202:参数错误]
    500: 服务器内部错误 [-400:内部错误]
    200: OK [0: ok, -306:会员卡不正确]
```