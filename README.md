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
* /v1/hongid/register/smscode/atk[post]: 注册第一步完成
```
    http status:

```
* /v1/hongid/register[post]: 绑定昵称&密码


[![Build Status](https://travis-ci.org/ogstation/gas-station.svg)](https://travis-ci.org/ogstation/gas-station)

gas-station
===========

How to build ?
======
mvn clean install

How to run ?
======
mvn jetty:run

API
======
* /api/error/400: handle 400 error
* /api/error/401: handle 401 error
* /api/error/403: handle 403 error
* /api/error/404: handle 404 error
* /api/error/500: handle 500 error



* /api/stations(GET): retrieve a list of gas stations by paging
```
Default size is 10, customize by /api/stations?page=2&size=20
```

* /api/stations/{id}(GET): retrieve a specific gas
```
Error 404 if not found
```

* /api/stations(POST): create gas station
```
Error if mandatory fields not filled.
{
    "fieldErrors": [
        {
            "field": "provinceCode",
            "message": "province should not be empty"
        },
        {
            "field": "countryCode",
            "message": "country should not be empty"
        },
        {
            "field": "contact",
            "message": "contact should not be empty"
        },
        {
            "field": "addressDetails",
            "message": "address details should not be empty"
        },
        {
            "field": "cityCode",
            "message": "city should not be empty"
        },
        {
            "field": "name",
            "message": "name should not be empty"
        }
    ]
}
```
* /api/stations/{id}(PUT): update gas station
```
Error 404 if not found
```
* /api/stations/{id}(DELETE): delete gas station
```
Error 404 if not found
```
* /api/stations/search(POST): search gas station by gas station
