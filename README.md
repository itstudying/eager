# eager
一个golang实现支持热更新的配置解析库。通常程序更改配置需要停机服务，但是停机会造成短时间的服务终止，如果能配置更改后，服务自动刷新，并加载相应配置多好。`写来练手的，慎用`

> 暂时支持以下功能
1. 解析toml文件
2. 解析配置至结构体
3. 解析配置至内存，可使用get、getE等等方法获取值，
4. 配置热更新，可自定义监听时间
5. 易扩展解析不同格式配置文件

> 暂不支持，待支持
1. 指定配置修改调用指定回调，比如db链接更改，自动加载新配置并重连db
2. 更多配置格式，ini、conf、yaml.....
3. 统一尽量多格式的配置文件获取值方法

## 快速入门

```
go get github.com/itstudying/eager
```

运行 `example/main.go`，修改config.toml配置文件，发现值发生变化
