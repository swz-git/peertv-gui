package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"path"

	"github.com/skratchdot/open-golang/open"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type SearchResult struct {
	URL        string  `json:"url"`
	Title      string  `json:"title"`
	Seed       string  `json:"seed"`
	Leech      string  `json:"leech"`
	MagnetLink string  `json:"magnetlink,omitempty"`
	FileSize   int     `json:"filesize"`
	Engine     string  `json:"engine"`
	Score      float64 `json:"score"`
}

func (a *App) Search(term string) []SearchResult {
	cmd := exec.Command(path.Join(GetGopath(), "bin/peertv"), "search", term, "-i", "https://search.swz.works", "--json")
	out, e := cmd.Output()
	if e != nil {
		return []SearchResult{}
	}

	var result []SearchResult
	json.Unmarshal(out, &result)
	return result
}

func (a *App) OpenLink(link string) {
	open.Run(link)
}

func (a *App) PlayMagnet(magnet string) {
	cmd := exec.Command(path.Join(GetGopath(), "bin/peertv"), "play", magnet)
	fmt.Println("Starting to play: ", magnet)

	// stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	e := cmd.Start()

	if e != nil {
		panic(e)
	}

	runtime.EventsOnce(a.ctx, "quitplayer", func(optionalData ...interface{}) {
		cmd.Process.Kill()
	})

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text() + "\n"
		fmt.Print(m)
		runtime.EventsEmit(a.ctx, "log", m)
	}

	cmd.Wait()
	runtime.EventsEmit(a.ctx, "quitted")
}

// // Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }
