#### Usage

在这里对此程序做一个详细的使用说明

在运行程序之前，你必须确认该计算机已经安装了 Chrome/Chromium 浏览器。

若在执行程序后超过 5 秒没有任何响应，也许你开了 VPN，请你将它暂时关闭再重试。

执行`screenshot.exe -h`会得出类似以下的参数说明，值得注意的是`-t`这个参数，它指定的是毫秒(ms)，我建议设置 1000ms, 也许速度会慢很多。

如果在程序执行完成之后发现图片不完整，没有截取到流量图的部分，大概率的原因是你的电脑打开浏览器缓慢导致的。其次就是要访问的这个服务器出现了问题，导致流量图无法加载。

执行该程序必须指定`-p`参数，因为不确定密码是什么，默认的密码是`zabbix`。不过一般情况都会对网站的密码进行修改，视情况而定。

在程序执行完成之前，我不建议你直接关掉运行窗口，因为我没有对 KILL 信号进行处理，懒得加了，就这样吧。使用`Ctrl-C`安全退出该程序！

```
MonitorCrawler version: MonitorCrawler/0.0.1
Usage: screenshot [-h]

Options:
  -c config
    	配置文件地址 config (default "config.json")
  -h	help..
  -p password
    	zabbix的用户密码 password (default "zabbix")
  -s serverName
    	zabbix服务器地址 serverName (default "140.246.36.89:8096")
  -t waitTime
    	网页加载等待时间, (t) == int && (t) >= 1, 单位(ms) waitTime (default 1000)
  -u username
    	zabbix的用户名 username (default "Admin")
```

#### 关于源代码部分

这个小程序是用 Golang 写的，至于为什么不用 Python，因为我写腻了。再者就是 Golang 写的程序部署非常容易，一次写好全平台无环境依赖运行！

代码内容非常简单，其中用到`Chromedp`这个 Golang 的 chrome 驱动，感谢该作者!

代码的注释很少，可以说是几乎没有，后续有闲时我会再次添加一些注释。
