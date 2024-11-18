package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"looklook/app/usercenter/cmd/rpc/internal/config"
	"looklook/app/usercenter/model"
	"os"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
	DB            *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化sqlx连接
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	//sqlDB, _ := sqlConn.RawDB() // This returns the *sql.DB instance used by sqlx.

	// 初始化 GORM 连接
	myLogger := logger.New(
		//设置Logger
		//NewMyWriter(),

		//输出在控制台，方便debug
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Silent, //仅仅在控制台输出指定Debug的语句
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)
	//dsn := "root:PXDN93VRKUm8TeE7@tcp(mysql:3306)/looklook_usercenter?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	gormDB, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{
		Logger: myLogger,
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	//sqlDB.SetMaxOpenConns(100)
	//sqlDB.SetMaxIdleConns(100)
	//sqlDB.SetConnMaxLifetime(time.Hour)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),

		UserAuthModel: model.NewUserAuthModel(sqlConn, c.Cache),
		UserModel:     model.NewUserModel(sqlConn, c.Cache),
		DB:            gormDB, // 设置 GORM 连接对象
	}
}
