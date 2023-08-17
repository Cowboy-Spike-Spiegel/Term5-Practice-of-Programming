BUPT Program Design
===========================
BUPT程序设计实践——DSL



##### 环境依赖
go 1.19.3
node.js v16.14.2
mysql Ver 8.0.28 for Win64 on x86_64



##### 目录结构描述

├── README.md
├── bin
│   ├── release.exe         // 发行版
│   ├── release_1.exe       // 女仆版
│   ├── release_2.exe       // 强尼银手版
│   └── script_test.exe     // 脚本本地测试版
├── src
│   ├── router              // 路由层
│   │   └── router.go
│   ├── controller          // 控制层
│   │   └── controller.go
│   ├── service             // 服务层
│   │   └── service.go
│   ├── dao                 // 数据访问层
│   │   ├── dao.go
│   │   ├── db.go
│   │   └── script.go
│   ├── output              // 输出
│   │   ├── print.go
│   │   └── log.go
│   └── test                // 脚本本地测试
│       └── test.go
├── main.go                 // 源程序入口
├── go.mod
├── go.sum 
├── scripts                 // 脚本相关
│   ├── script.json
│   ├── script_1.json
│   ├── script_2.json
│   └── test                // 测试数据
│       ├── input           // 输入数据
│       │   ├── input_script.txt
│       │   ├── input_script_1.txt
│       │   └── input_script_2.txt
│       └── reference       // 对比数据
│           ├── input_script.txt
│           ├── input_script_1.txt
│           └── input_script_2.txt
├── static
│   ├── css
│   │   ├── login.css
│   │   └── service.css
│   ├── js
│   │   └── jquery-3.6.1.min.js
│   └── pics
│       ├── Spike.jpg
│       └── 02.jpg
├── templates
│   ├── login.html
│   └── service.html
└── docs
     ├── Notation description.md
     ├── design.jpg
     └── Lab Report.pdf



##### 设计模式

![design.jpg](/docs/design.jpg)