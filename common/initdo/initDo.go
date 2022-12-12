package initdo

import (
	"fmt"
	"go_forum/common"
	"go_forum/common/snowflake"
	"time"

	"go.uber.org/zap"

	"go_forum/common/setUp/config"
	"go_forum/common/setUp/logger"
	"go_forum/dao/mysql"
	"go_forum/dao/redis"
)

func InitDO(ch chan int) {
	//初始化logger
	if err := logger.InitLogger(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("初始化logger失败, err:%v\n", err)
		panic(err)
	}

	//Timer定时写入
	go func() {
		<-time.NewTimer(1 * time.Hour).C
		zap.L().Sync()
	}()
	defer zap.L().Sync() //写入磁盘

	//初始化mysql连接
	if err := mysql.InitMysql(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("初始化mysql失败, err:%v\n", err)
		panic(err)
	}
	//defer mysql.DBClose()

	//初始化redis连接
	if err := redis.InitRedis(config.Conf.RedisConfig); err != nil {
		fmt.Printf("初始化redis失败, err:%v\n", err)
		panic(err)
	}
	//defer redis.RDBClose()

	//雪花算法生成分布式uid
	if err := snowflake.InitSnowFlake(config.Conf.StartTime, config.Conf.MachineID); err != nil {
		fmt.Printf("雪花算法生成uid失败, err:%v\n", err)
		panic(err)
	}

	//初始化gin内置支持的校验器(validator)的翻译器(en zh)
	if err := common.InitTrans("zh"); err != nil {
		fmt.Printf("初始化validator翻译器失败, err:%v\n", err)
		return
	}

	<-ch
	return
}
