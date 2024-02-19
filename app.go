package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SetWindowTitle(pid int, title string) {

}

func (a *App) BeginFetchingWindows() {
	go func() {
		for {
			w := GetWindows()
			runtime.EventsEmit(a.ctx, "window-update", w)
			time.Sleep(500 * time.Millisecond)
		}
	}()
}
