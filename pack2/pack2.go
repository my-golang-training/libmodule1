package pack2

import(
	"log/slog"
)

func Nop(){
	slog.Info("Nop() of libmodule1/pack2/pack2.go")
}

func init(){
	slog.Info("init() of libmodule1/pack2/pack2.go")
}
