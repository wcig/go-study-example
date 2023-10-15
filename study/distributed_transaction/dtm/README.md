# dtm 示例

## 1.环境准备
**1.数据库准备**
MySQL 执行 `init.sql`. 初始化的 3 个数据库分别对应:
* dtm: tm server 连接使用
* dtm_barrier: rm server 连接使用
* dtm_busi: rm server 连接使用

**2.运行 dtm 服务**
```shell
dtm -c config.yml
```

## 2.运行测试
```shell
go run xa/main.go
```