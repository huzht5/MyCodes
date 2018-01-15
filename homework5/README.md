# 程序说明

## 实现的功能有：

​	①将mysql数据库用xorm包装
    ②实现delete功能。

---

## 实验结果：

ab测试结果如下：

```bash
$ ab -n 1000 -c 100 -T 'application/x-www-form-urlencoded' -p post.txt http://localhost:8080/service/insertThis is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/insert
Document Length:        111 bytes

Concurrency Level:      100
Time taken for tests:   1.179 seconds
Complete requests:      1000
Failed requests:        977
   (Connect: 0, Receive: 0, Length: 977, Exceptions: 0)
Total transferred:      236785 bytes
Total body sent:        192000
HTML transferred:       112785 bytes
Requests per second:    848.23 [#/sec] (mean)
Time per request:       117.892 [ms] (mean)
Time per request:       1.179 [ms] (mean, across all concurrent requests)
Transfer rate:          196.14 [Kbytes/sec] received
                        159.04 kb/s sent
                        355.18 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   2.9      0      19
Processing:    34  115  68.7     86     327
Waiting:       33  114  67.1     86     308
Total:         34  117  68.7     88     329

Percentage of the requests served within a certain time (ms)
  50%     88
  66%    115
  75%    137
  80%    151
  90%    246
  95%    277
  98%    316
  99%    328
 100%    329 (longest request)
```

得出的结论：

- xorm  ①实现了dao的自动化  ②提高了编程效率，少了一层架构  ③使用反射，因此损失了一点性能

---

## 使用方法:

1. 根据如下步骤自己打开默认数据库。

   ```bash
   $ sudo docker run -p 2048:3306 --name mysql2 -e MYSQL_ROOT_PASSWORD=root -d mysql
   $ sudo docker run -it --net host mysql "sh"
   # mysql -h127.0.0.1 -P2048 -uroot -proot
   mysql: [Warning] Using a password on the command line interface can be insecure.
   Welcome to the MySQL monitor.  Commands end with ; or \g.
   Your MySQL connection id is 3
   Server version: 5.7.20 MySQL Community Server (GPL)

   Copyright (c) 2000, 2017, Oracle and/or its affiliates. All rights reserved.

   Oracle is a registered trademark of Oracle Corporation and/or its
   affiliates. Other names may be trademarks of their respective
   owners.

   Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

   mysql> create database test;
   Query OK, 1 row affected (0.00 sec)
   ```

   此时输入./app即可。默认属性为：

   ​	./app root root 2048 test

2. 输入./app 或 ./app name password port databasename 开启服务器。

   其中./app 等同于 ./app root root 2048 test

3. 服务器开始运行时，用curl测试：

   ```bash
   插入：
   $ curl -d "username=XXX&departname=abc" http://localhost:8080/service/insert
   查找所有：
   $ curl http://localhost:8080/service/find?userid=
   通过Id查找：
   $ curl http://localhost:8080/service/find?userid=1
   删除：
   $ curl http://localhost:8080/service/delete?userid=1
   ```