#app-svr

ps -ef | grep bee | awk '{print $2}' | xargs kill -9
ps -ef | grep hongId | awk '{print $2}' | xargs kill -9
ps -ef | grep app-svr | awk '{print $2}' | xargs kill -9

##API
```
通用返回结构
    {
        "status":{
            code: 0,
            message: "sfsafasf"
        }
        "body":{
            xxxxxxx
        }
    }

状态码
    0：everything is ok.
    -202：参数错误.
    -301：手机号已被注册.
    -302：短信网关错误
    -303：验证码错误
    -304：用户名错误
    -305：密码错误
    -306：会员卡不正确
    -307：会员卡已被其他用户绑定
    -400：系统内部错误
```
* /v1/hongid/register/smscode [post]: 获取手机验证码
```
    json body:
    {
        tel: "18610889275"
    }

    status:
        200 OK [0:ok, -202:参数错误, -302:验证码发送失败]
```
* /v1/hongid/register/smscode/atk [post]: 注册第一步完成
```
    json body:
    {
        tel: "18610889275",
        code: "748573"
    }
    status:
        200: OK [0:ok, -202:参数错误, -303:验证码错误, -301: 手机号已经被注册, -400:内部错误]
```
* /v1/hongid/register[post]: 绑定昵称&密码
```
    json body:
    {
        tel: "18610889275",
        nickname: "sss",
        password: "lsjdfoiwejfoj"
    }
    200: OK [0:ok, -202:参数错误, -400:内部错误]
```
* /v1/hongid/login [post]: 登录
```
    json body:
    {
        username: "18610889275",
        password: "lsjdfoiwejfoj"
    }
    200: OK [0: ok, -202:参数错误, -304: 用户名错误, -305: 密码错误, -400:内部错误]
          "body": {
             "id": 1,
             "hongid": "86046",
             "tel": "18610889275",
             "email": "",
             "nickname": "mint"
           }

```
* /v1/memberCards/card/:card/bind/:owner [post] :绑定会员卡
```
    200: OK [0: ok, -202:参数错误， -306:会员卡不正确, -307:会员卡已被其他会员绑定，-400:内部错误]
```
* /v1/memberCards/card/:card/unbind [post] :解除绑定会员卡
```
    200: OK [0: ok, -202:参数错误， -306:会员卡不正确， -400:内部错误]
```