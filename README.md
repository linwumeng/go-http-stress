go语言实现的一个http请求性能测试小工具，记录平均响应时间和95%分位响应时间。用goroutine虚拟并发用户，目前不支持proxy。

# 用法
1. 编译源码
编译后运行`go build -o stress.exe`。

1. 运行
`stress.exe --help`
```
NAME:
   stress - http test

USAGE:
   stress.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value, -u value          specify target url (default: "https://www.baidu.com")
   --concurrency value, -c value  specify concurrency level (virutal user) (default: 1)
   --requests value, -n value     specify request number per user (default: 1)
   --interval value, -t value     specify interval(s) of printing a row (default: 1)
   --help, -h                     show help
   --version, -v                  print the version
```

1. 测试www.baidu.com, 单用户
```
stress  -n 100 -t 5 -u https://www.baidu.com 
URL:      https://www.baidu.com
协程数:   1
总请求数: 100
输出间隔: 5s
─────┬────────┬────────┬────────┬────────┬────────┬────────┬────────┬────────
 窗口│ 请求数 │ 成功数 │ 错误率 │中位耗时│最长耗时│最短耗时│95% 耗时│ TPS
─────┼────────┼────────┼────────┼────────┼────────┼────────┼────────┼────────
    1|       6|       6|    0.00|   166ms|   172ms|   138ms|   172ms|    6.20
    2|      38|      38|    0.00|   152ms|   172ms|   135ms|   170ms|    6.49
    3|      70|      70|    0.00|   152ms|   208ms|   135ms|   172ms|    6.48
    4|     100|     100|    0.00|   152ms|   208ms|   133ms|   169ms|    6.49
总耗时:  15.5s
```
1. 测试www.taobao.com，多用户
```
stress  -n 100 -t 5 -c 10 -u https://www.taobao.com
URL:      https://www.taobao.com
协程数:   10
总请求数: 1000
输出间隔: 5s
─────┬────────┬────────┬────────┬────────┬────────┬────────┬────────┬────────
 窗口│ 请求数 │ 成功数 │ 错误率 │中位耗时│最长耗时│最短耗时│95% 耗时│ TPS
─────┼────────┼────────┼────────┼────────┼────────┼────────┼────────┼────────
    1|      86|      86|    0.00|    81ms|   290ms|    67ms|   275ms|   95.01
    2|     591|     591|    0.00|    80ms|   538ms|    65ms|   269ms|  100.46
    3|     952|     952|    0.00|    80ms|   538ms|    65ms|   265ms|  105.63
    4|     952|     952|    0.00|    80ms|   538ms|    65ms|   265ms|  105.63
    5|     952|     952|    0.00|    80ms|   538ms|    65ms|   265ms|  105.63
    6|     952|     952|    0.00|    80ms|   538ms|    65ms|   265ms|  105.63
    7|    1000|    1000|    0.00|    80ms| 21278ms|    65ms|   265ms|   73.46
总耗时:  30.5s
```
