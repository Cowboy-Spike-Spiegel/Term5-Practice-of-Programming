# Notation description

记法说明



1. 脚本记法文件格式选择

  鉴于程序采用的是go语言，对于yaml，json等有很好的支持，所以这里采用json进行script的存储。这种键值对的存储很适合脚本逻辑的存储。

  

2. 脚本整体结构设计

  脚本整体设计为一个大的json格式：

  ```json
  {
    "Steps": {
      "Start" : {
      },
      "Welcome" : {
      },
      "Silence" : {
      },
      "Complain" : {
      },
      "Reply" : {
      },
      "Bill" : {
      },
      "Mishandled" : {
      },
      "Exit" : {
      }
    }
  }
  ```

  这里我们存了一个对象，key为steps，value为各个step的名称。

  

3. 脚本单元结构设计

  对于每个step，我们有详细的结构设计：

  ```json
  "Start" : {
   "speak": "您好，$name，我在。",
   "listen": 0,
   "branch": {},
   "silence": "",
   "default": "Welcome"
  },
  "Welcome" : {
   "speak": "您好，$name，请问有什么可以帮您？（投诉、账单、退出）",
   "listen": 20,
   "branch": {
       "投诉": "Complain",
       "账单": "Bill",
       "退出": "Exit"
   },
   "silence": "Silence",
   "default": "Mishandled"
  },
  "Silence" : {
   "speak": "还在嘛？请问有什么可以帮您的嘛？（投诉、账单、退出）",
   "listen": 20,
   "branch": {
       "投诉": "Complain",
       "账单": "Bill",
       "退出": "Exit"
   },
   "silence": "Silence",
   "default": "Mishandled"
  },
  "Complain" : {
   "speak": "请提出您宝贵的建议：",
   "listen": 15,
   "branch": {},
   "silence": "Complain",
   "default": "Reply"
  },
  "Reply" : {
   "speak": "我们收到了，感谢！",
   "listen": 0,
   "branch": {},
   "silence": "",
   "default": "Welcome"
  },
  "Bill" : {
   "speak": "您好，您本月的账单是$amount，账户余额是$balance。",
   "listen": 0,
   "branch": {},
   "silence": "",
   "default": "Welcome"
  },
  "Mishandled" : {
   "speak": "我没听懂呜呜呜。",
   "listen": 0,
   "branch": {},
   "silence": "",
   "default": "Welcome"
  },
  "Exit" : {
   "speak": "感谢您的来电，再见",
   "listen": 0,
   "branch": {},
   "silence": "",
   "default": "End"
  }
  ```

  我们规定每个step都有：speak，listen，branch，silence，default。

  - speak：string，代表这个步骤要说的话

  - listen：int，代表这个步骤要等待的时间

  - branch：map[string]string，代表根据用户的输入来匹配下一个步骤名称

  - silence：string，代表用户回应超时时因该跳转到的步骤名称

  - default：string，代表用户回应或者等待时间为0时需要跳转到的步骤名称

  

4. 脚本特殊属性值设计

	step：名称必须包含**Start, Exit**，且Exit的default必定是**End**

	- speak：有$name, $amount, $balance, 这些值都要在生成回复信息时通过数据库来查找对应的数据进行替换

	- listen：>= 0，0代表不需要等待，大于0代表需要等待的时间。

	- branch：下一个步骤名称必须是script中的步骤名称

	- silence：string，代表用户回应超时时因该跳转到的步骤名称

	- default：string，代表用户回应或者等待时间为0时需要跳转到的步骤名称



 



