package headless

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/zzzgydi/webscraper/common/config"
	"github.com/zzzgydi/webscraper/common/initializer"
)

var (
	allocCtx      context.Context
	readabilityJS string
)

func initHeadless() error {
	cs := config.AppConf.Chrome
	if cs.RemoteUrl != "" {
		slog.Info("use remote chrome", slog.String("url", cs.RemoteUrl))

		ctx, _ := chromedp.NewRemoteAllocator(context.Background(), cs.RemoteUrl)
		allocCtx = ctx
	} else {

		options := []chromedp.ExecAllocatorOption{
			chromedp.Flag("headless", true),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("blink-settings", "imagesEnabled=false"),
		}
		if cs.ExecPath != "" {
			slog.Info("use local chrome", slog.String("path", cs.ExecPath))
			options = append(options, chromedp.ExecPath(cs.ExecPath))
		} else {
			slog.Info("use default chrome")
		}

		options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

		// set http proxy
		if config.AppConf.HttpProxy != "" {
			_, err := url.Parse(config.AppConf.HttpProxy)
			if err != nil {
				return err
			}

			slog.Info("use http proxy", slog.String("proxy", config.AppConf.HttpProxy))
			options = append(options, chromedp.ProxyServer(config.AppConf.HttpProxy))
		}

		ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)
		allocCtx = ctx
	}

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
