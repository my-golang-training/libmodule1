package pack1

import(
	"log/slog"
)

func Nop(){
	slog.Info("Nop() of libmodule1/pack1/pack1.go")
	slog.Info("Release", slog.String("version","1.0.0"))
	return
}

func init(){
	slog.Info("Init() of libmodule1/pack1/pack1.go")
}
