package main

import (
	"flag"
	"log"
	"time"

	"github.com/opentracing/opentracing-go/ext"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	rate := flag.Int("rate", 1, "Rate at which traces should be generated.")
	// hostname := flag.String("host", "localhost", "Jaeger proxy hostname")
	// port := flag.Int("port", 6831, "Jaeger proxy udp port")

	ticker := time.NewTicker(time.Second * time.Duration(*rate))

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

	println("Generating spans...")
	for {
		<-ticker.C
		sp := opentracing.StartSpan("dummy-span")
		ext.SamplingPriority.Set(sp, 1)
		sp.Finish()
	}
}
