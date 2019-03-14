## jaeger-test-client

Simple test client for Jaeger that will emit traces every N seconds.

The client will connect to localhost on port 6831.

You may specify the interval at which traces are generated and sent.

## Example:
Use -interval 2000 to generate a trace every 2 seconds.

`go run main.go -interval 2000`
