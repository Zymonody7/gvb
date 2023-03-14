# gvb后端
## 新建项目，配置代理
```bash
GOPROXY=https://goproxy.cn,direct
```
## 项目结构
![img.png](img.png)
## 项目配置
### 配置文件
```yaml
// setting.yaml
mysql:
  host: 127.0.0.1
  port: 3306
  db: gvb_db
  user: root
  password: abc123456
  log_level: dev
logger:
  level: info
  prefix: '[gvb]'
  director: log
  show_line: true
  log_in_console: true
system:
  host: '0.0.0.0'
  port: 8080
  env: dev
```
### 配置文件读取
```go
// config/enter.go
type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
// config/conf_system.go
type System struct {
    Host string `yaml:"host"`
    Port string `yaml:"port"`
    Env  string `yaml:"env"`
}
// config/conf_mysql.go
type Mysql struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    DB       string `yaml:"db"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    LogLevel string `yaml:"log_level"`
}
// config/conf_logger.go
package config

type Logger struct {
    Level        string `yaml:"level"`
    Prefix       string `yaml:"prefix"`
    Director     string `yaml:"director"`
    ShowLine     bool   `yaml:"show_line"`
    LogInConsole bool   `yaml:"log_in_console"` // 是否显示打印路径
}
```
安装读取yaml文件的包 `go get gopkg.in/yaml.v2`
```go
// core/conf.go
package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

func InitConf() {
	const ConfigFile = "setting.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf Error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Config Init Unmarshal Error: %v", err)
	}
	log.Println("Config yamlFile load Init Success")
	global.Config = c
}
// 将config放在全局
// global/global.go
package global

import "gvb_server/config"

var (
	Config *config.Config
)
```
在main.go导入后，运行程序。成功打印出配置信息
```go
func main() {
	core.InitConf()
	fmt.Println(global.Config)
}
```
![img_1.png](img_1.png)
### gorm配置
```go
// config/conf_mysql.go
func (m Mysql) Dsn() string {
  return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}

// core/gorm.go
package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	return MysqlConnect()
}

func MysqlConnect() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("未配置MySQL，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", err))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间
	return db
}

```
