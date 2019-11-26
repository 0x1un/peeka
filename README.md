这个仓库中放的是一些工作上的小工具

目前的整体目录结构如下

```
    ├── cmd                         // 可执行程序（可直接编译）
    │   ├── attendanceRobot         // 获取企业中指定人的排班
    │   ├── dingtalk                // 钉钉api测试文件 一般我会忽略上传
    │   └── screenshot              // 一个运维小工具，截取zabbix/深信服的图上传七牛并发送到钉钉群聊机器人
    ├── go.mod
    ├── go.sum
    ├── internal                    // 上面的所有程序都是从这里调用核心的功能
    │   ├── chatbot                 // 钉钉机器人的api
    │   ├── component               // 常用的一些组建
    │   ├── dingtalk                // 钉钉的一些api（我用到的）
    │   ├── registry                // windows注册表的一些操作
    │   └── screenshot              // 监控截图工具的核心代码
    ├── pkg
    │   └── common
    └── README.md
```

**\***上面的 attendanceRobot 使用的递归方式，因为写起来比较方便，但资源耗费可能有些多。就这样吧，有时间再调整调整\*\*
