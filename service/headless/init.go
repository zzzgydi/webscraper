package headless

import (
	"context"
	"fmt"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/zzzgydi/webscraper/common/initializer"
)

var (
	allocCtx      context.Context
	readabilityJS string
)

func initHeadless() error {
	// options := []chromedp.ExecAllocatorOption{
	// 	chromedp.Flag("headless", true),
	// 	chromedp.ExecPath("/headless-shell/headless-shell"),
	// 	chromedp.Flag("disable-gpu", true),
	// 	chromedp.Flag("blink-settings", "imagesEnabled=false"),
	// 	// chromedp.Flag("disable-web-security", true),
	// }
	// options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	// // set http proxy
	// if config.AppConf.HttpProxy != "" {
	// 	_, err := url.Parse(config.AppConf.HttpProxy)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	options = append(options, chromedp.ProxyServer(config.AppConf.HttpProxy))
	// }

	// ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)

	ctx, _ := chromedp.NewRemoteAllocator(context.Background(), "ws://chromedp:9222/")
	allocCtx = ctx

	js, err := os.ReadFile("assets/Readability.js")
	if err != nil {
		return fmt.Errorf("failed to read JavaScript file: %w", err)
	}

	readabilityJS = string(js)
	return nil
}

func init() {
	initializer.Register("headless", initHeadless)
}
