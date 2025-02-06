package server

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/flamego/cors"
	"github.com/flamego/flamego"
	"github.com/spf13/cobra"
	"go-backend-template/config"
	"go-backend-template/core/kernel"
	"go-backend-template/core/logx"
	"go-backend-template/core/store/mysql"
	"go-backend-template/internal/app/appInitialize"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	configYml string
	engine    *kernel.Engine
	StartCmd  = &cobra.Command{
		Use:     "server",
		Short:   "Set Application config info",
		Example: "main server -c config/config.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setUp()
			loadStore()
			loadApp()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config.yaml",
		"Start server with provided configuration file")
}

// 初始化配置和日志
func setUp() {
	// 初始化全局 ctx
	ctx, cancel := context.WithCancel(context.Background())

	// 初始化资源管理器
	engine = &kernel.Engine{Ctx: ctx, Cancel: cancel}
	kernel.Kernel = engine

	// 加载配置
	config.LoadConfig(configYml)

	engine.Fg = flamego.New()
	engine.Fg.Use(flamego.Recovery(), flamego.Renderer(), cors.CORS(cors.Options{
		AllowCredentials: true,
		Methods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
	}))
}

func loadStore() {
	engine.Mysql = mysql.MustNewMysqlOrm(config.GetConfig().Mysql)
}

// 加载应用，包含多个生命周期
func loadApp() {
	apps := appInitialize.GetApps()
	for _, app := range apps {
		_err := app.Init(engine)
		if _err != nil {
			logx.Errorw("failed to init app", zap.Field{Key: "err", Type: zapcore.StringType, String: _err.Error()})
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.Load(engine)
		if _err != nil {
			logx.Errorw("failed to load app", zap.Field{Key: "err", Type: zapcore.StringType, String: _err.Error()})
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.Start(engine)
		if _err != nil {
			logx.Errorw("failed to start app", zap.Field{Key: "err", Type: zapcore.StringType, String: _err.Error()})
			os.Exit(1)
		}
	}
}

func run() {
	port, _ := strconv.Atoi(config.GetConfig().Port)
	engine.Fg.Run(port)
}
