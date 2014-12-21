#app-svr

ps -ef | grep bee | awk '{print $2}' | xargs kill -9
ps -ef | grep hongId | awk '{print $2}' | xargs kill -9
ps -ef | grep app-svr | awk '{print $2}' | xargs kill -9

##API
```
通用返回结构
    {
        "status":{
            "code": 0,  int
            "msg": "sfsafasf"  string
        }
        "body":{
            "xxx":"xxx"
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
* /v1/public/smscode [post]: 获取手机验证码
```
    body:
        tel=18610889275&type=1     string,int   注册验证码
        tel=18610889275&type=2     string,int   找回密码验证码
    status:
        200 OK [0:ok, -202:参数错误, -302:验证码发送失败]
```
* /v1/hongid/register/smschk [post]: 注册第一步完成
```
    body:
        tel=18610889275&code=123455     string,string
    status:
        200: OK [0:ok, -202:参数错误, -303:验证码错误, -301: 手机号已经被注册, -400:内部错误]
```
* /v1/hongid/register/info [post]: 绑定昵称&密码
```
    body:
        tel=18610889275&nickname=mint&password=123456   string,string,string
    status:
        200: OK [0:ok, -202:参数错误, -400:内部错误]
```
* /v1/hongid/login [post]: 登录
```
    body:
        username=18610889275&password=123456    string,string
    status
        200: OK [0: ok, -202:参数错误, -304: 用户名错误, -305: 密码错误, -400:内部错误]
          "body": {
             "id": 1,   int64
             "hongid": "86046",     string
             "tel": "18610889275",  string
             "email": "",   string
             "nickname": "mint"     string
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