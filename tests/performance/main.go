package main

import (
	"fmt"
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	blogpb "github.com/dipjyotimetia-anz/gogrpc/blog/blogPb"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
)

var (
	endpoint = map[string]string{"CreateBlog": "blog.BlogService.CreateBlog", "ReadBlog": "blog.BlogService.ReadBlog"}
)

const (
	Host        = "localhost:50051"
	Concurrency = 5
	TotalReq    = 10
	Insecure    = true
)

func main() {
	report, err := runner.Run(
		endpoint["CreateBlog"],
		Host,
		// runner.WithProtoFile("blog/blogPb/blog.proto", []string{}), //TODO: gRPC Reflection api is configured
		runner.WithData(&blogpb.CreateBlogRequest{
			Blog: &blogpb.Blog{
				Id:       strconv.Itoa(int(uuid.New().ID())),
				AuthorId: "Dip",
				Title:    "Dip gRPC Mongo perf",
				Content:  "mongo service perf",
			}}),
		runner.WithConcurrency(Concurrency),
		runner.WithTotalRequests(TotalReq),
		runner.WithInsecure(Insecure),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	reportPrinter := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}

	ts := make(map[int]int64)
	ts[50] = 10
	ts[75] = 12
	ts[90] = 15
	ts[95] = 16
	ts[99] = 25

	ValidateLatency(ts, reportPrinter.Report)
	_ = reportPrinter.Print("influx-details")
}

// ValidateLatency validate latency
func ValidateLatency(latencyTime map[int]int64, report *runner.Report) {
	for _, details := range report.LatencyDistribution {
		switch details.Percentage {
		case 50:
			if details.Latency.Milliseconds() > latencyTime[50] {
				log.Fatalf("P90: latency differene :%d", latencyTime[50]-details.Latency.Milliseconds())
			}
		case 75:
			if details.Latency.Milliseconds() > latencyTime[75] {
				log.Fatalf("P75: latency differene :%d", latencyTime[75]-details.Latency.Milliseconds())
			}
		case 90:
			if details.Latency.Milliseconds() > latencyTime[90] {
				log.Fatalf("P90: latency differene :%d", latencyTime[90]-details.Latency.Milliseconds())
			}
		case 95:
			if details.Latency.Milliseconds() > latencyTime[95] {
				log.Fatalf("P95: latency differene :%d", latencyTime[95]-details.Latency.Milliseconds())
			}
		case 99:
			if details.Latency.Milliseconds() > latencyTime[99] {
				log.Fatalf("P99: latency differene :%d", latencyTime[99]-details.Latency.Milliseconds())
			}
		}
	}
}
