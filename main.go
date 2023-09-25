package main

import (
	_ "goframe-argo-workflow/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-argo-workflow/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
