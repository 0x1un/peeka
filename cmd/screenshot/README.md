screenshot 是一个日常运维的小工具

## Usage:

`Windows*`: screenshot.exe -h
`Linux`: screenshot -h

相关的帮助使用`-h`就足够了。

### config.json 配置格式

对于抓取线路的网址，其格式配置如下

```json
{
  "成都-线路": [
    {
      "电信": "http://127.0.0.1:8888"
    },
    {
      "联通": "http://127.0.0.1:8887"
    }
  ]
}
```

**\*** 只支持 zabbix4.2.x 版本及以上，因为深信服的垃圾 web 页面，暂时只写了首页的抓取。\*\*

### 使用.env 文件作为配置文件

```
# 高朋运维群机器人
ROBOT_TOKEN=36e452cedd99e028151d3ce6f3b90b9a3994d9ab8a62811c49b7f9da7fd9a
# 七牛云空间密钥
ACCESSKEY=GSDdabqaln9Vgou3Hhhu79ADFAF979FDA9F9A9AA
SECRETKEY=TnTWGN_KG_2UG890F8DASFA0A0A0AD8F0AS0FA0hbdvV0
BUCKET=test

# zabbix服务器地址与账户密码设置
ZABBIX_SERVER=127.0.0.1:1100
ZABBIX_USERNAME=Admin
ZABBIX_PASSWORD=zabbix

# 深信服账户密码
SANGFOR_USERNAME=admin
SANGFOR_PASSWORD=xxxx3

# 深信服抓取等待时间
SANGFOR_LOGIN_TIME=3
SANGFOR_PAGE_TIME=20

# zabbix抓取等待时间(ms)
TIMEOUT=1000
# 该程序超时时间(s)
TOTALTIMEOUT=120
# 图片质量
QUALITY=100
# zabbix抓图的时间段(1h,3h,6h,12h,24h,15m,30m)
TIMERANGE=24h
```

就这些吧，如有其他需求自行更改源码进行自定义
