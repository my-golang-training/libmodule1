package pack1

import(
	"log/slog"
)

func Nop(){
	slog.Info("Nop() of libmodule1/pack1/pack1.go")
	return
}

func init(){
	slog.Info("Init() of libmodule1/pack1/pack1.go")
}
