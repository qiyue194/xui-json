# 使用教程

在json文件中配置你xui面板配置的订阅地址


```
https://aa.demo.com:20304/[自定义的订阅地址]/
```

这里注意需要  **`/`**  结尾

### 你在正常复制xui面板订阅地址时应该是

##### ~~你不知道如何使用xui的订阅？  打开google.com  搜搜【xui订阅设置】~~

https://aa.demo.com:20304/[自定义的订阅地址]/[用户Subscription]

之后你便可以使用
http://ip:30303/xuijson?token=[用户Subscription]

如果配置了ssl证书就是

https://ip:30303/xuijson?token=[用户Subscription]

即可获取多个面板的订阅组合信息

注意这里的【用户Subscription】需要你的多个xui面板的【用户Subscription】设置一样才可以从token=【用户Subscription】时组合获取订阅地址

执行方法

上传文件`xui-json`和`config.json`文件到服务器

```shell
chmod +x xui-json
#在线执行
./xui-json
#后台执行
nohup ./xui-json >/dev/null 2>& 1&
```



