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
        type: 1(注册验证码),2(找回密码）int
        tel: 18610888575 string
    status:
        200 OK [0:ok, -202:参数错误, -302:验证码发送失败]
```
* /v1/hongid/register/smschk [post]: 注册第一步完成
```
    body:
        tel: 18609484759 string
        code: 123456 string
    status:
        200: OK [0:ok, -202:参数错误, -303:验证码错误, -301: 手机号已经被注册, -400:内部错误]
```
* /v1/hongid/register/info [post]: 绑定昵称&密码
```
    body:
        tel:18610889275 string
        nickname: mint string
        password: 123456 string

    status:
        200: OK [0:ok, -202:参数错误, -400:内部错误]
```
* /v1/hongid/login [post]: 登录
```
    body:
        username: 18610889275 string
        password: 123456    string
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
* /v1/store/brush [post]: 筛选商铺
```
    body:
        gps: 40.047669,116.313082 string
        ordertype: 1(综合), 2(销量), 3(人气) int
        orderorder: 1(升序), 2(降序) int
        productcount: 3 int
        storepage: 1 int
        storesize: 10 int
        industryid: 0(全部) int
        distance: 0(附近), 1(1公里), 2(3公里), 3(5公里), 4(10公里) int
        keywordsearch: 辣条 string
    status
        200: OK [0: ok, -202:参数错误, -400:内部错误, ]
        "body": {
            "stores":{
                [
                    {
                        "storeid": 1, int
                        "storename": 辣条馆, string
                        "industryname": 餐饮, string
                        "distance": 549, int
                        "avatar": http://pic.baidu.com/jiejal.jpg, string
                        "products":{
                            [
                                {
                                    "productid": 1, int
                                    "productname": 辣条, string
                                    "productavatar": http://pic.baidu.com/latiao.jpg, string
                                    "productprice": 100, int
                                    "productunit": 袋 string
                                },
                                {}
                            ]
                        }
                    },
                    {}
                ]
            },
            "filters":{
                "industryfilters":{
                    [
                        {
                            "itemid":0, int
                            "itemname":"全部" string
                        },
                        {"itemid":1,"itemname":"餐饮"}
                    ]
                },
                "distancefilters":{
                    [
                        {"itemid":0,"itemname":"附近"},
                        {"itemid":1,"itemname":"1公里"}
                    ]
                }
            }
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