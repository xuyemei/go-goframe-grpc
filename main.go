package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"go-goframe-grpc/internal/cmd"
	_ "go-goframe-grpc/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
