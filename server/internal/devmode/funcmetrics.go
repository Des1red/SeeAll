package devmode

import (
	"SeeAll/internal/model"
	"fmt"
	"log"
	"time"
)

func ShowFuncMetrics(t string, posts []model.Post, total time.Duration) {
	if !isDev() {
		return
	}
	if isPrewarming.Load() {
		return
	}
	fmt.Println("NEWS REQUEST")
	fmt.Println("Type:", t)
	fmt.Println("Posts:", len(posts))
	fmt.Printf("Total: %dms\n", total.Milliseconds())
	fmt.Println("RSS:", model.Usage.RSS)
	fmt.Println("JSONRSS:", model.Usage.JSONRSS)
	fmt.Println("ATOM:", model.Usage.Atom)
	fmt.Println("-------------------")
	model.Usage = model.FuncUsage{}
	for name, m := range getSourceMetrics() {
		log.Printf("  %-20s %dms  %d posts", name, m.Duration.Milliseconds(), m.Posts)
	}
	resetSourceMetrics()
}
