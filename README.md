# 初始化ZK数据

config.yml 配置ZK和需要创建的路径值

config.yml 配置文件和执行程序在同级目录

格式如下

``` yaml
# zk服务器地址
server: 127.0.0.1:2181
# 需要导入ZK的数据，每个子对象相当于下级目录
data: 
  vrv: 
    test: 
      value1: test value 1
      value2: test value 2
    test2: 
      value1: test2 value 1
      value2: test2 value 2
```

```shell
# 执行测试
go run init-zk-data

# 打包
go build
```
