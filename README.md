<p align="center">
  go多模块管理项目
</p>
<p align="center">使用gin框架搭建，实现了admin管理模块的接口。可以作为一个参考项目来做二次开发</p>

---
### 使用
- [gin](https://github.com/gin-gonic/gin)（web框架）
- [redis](https://github.com/gomodule/redigo) 缓存（连接池）
- [gorm](https://github.com/jinzhu/gorm)（数据库操作）
- [ini](https://github.com/go-ini/ini)（配置载入）
- [jwt](https://github.com/dgrijalva/jwt-go)（用户授权）
- [casbin](https://github.com/casbin/casbin)（rbac权限管理）
- [ali-sms](https://github.com/fastgoo/alisms-go)（阿里云发送短信）

### 参考文档
- [gin中文文档](https://learnku.com/docs/gin-gonic/2018/gin-readme/3819)
- [gorm中文文档](http://gorm.book.jasperxu.com/)
- [glide包管理使用教程](https://learnku.com/articles/23503/package-management-tool-glide)
- [casbin权限控制中文文档](https://casbin.org/docs/zh-CN/overview)
- [awesome-go中文文档](https://github.com/jobbole/awesome-go-cn)
- [awesome-go英文文档](https://github.com/avelino/awesome-go)

### 快速开始
```
1、安装golang环境（自行百度）
2、安装 golang.org/x 核心包 ## 传送门：https://www.jianshu.com/p/c77aabdf554c
3、git clone https://github.com/fastgoo/gin-admin-modules.git $GOPATH/src/gin-modules
4、修改配置文件 app.ini.example 为 app.ini
5、修改 app.ini 的数据库、redis配置
6、运行程序：go run main.go -conf app.ini -log=./runtime/log/gin.log
```



### 微信
- huoniaojungege

