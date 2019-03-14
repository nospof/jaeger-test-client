package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/opentracing/opentracing-go/ext"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	interval := flag.Int("interval", 1000, "Interval (ms) at which traces should be generated.")
	flag.Parse()
	fmt.Printf("Generating traces every %d ms\n", *interval)

	ticker := time.NewTicker(time.Millisecond * time.Duration(*interval))

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
	}

	closer, err := cfg.InitGlobalTracer("jaeger-test-client")

	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	for {
		<-ticker.C
		sp := opentracing.StartSpan("dummy-span")
		ext.SamplingPriority.Set(sp, 1)
		sp.Finish()
	}
}
