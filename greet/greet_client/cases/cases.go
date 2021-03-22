package cases

import "github.com/dipjyotimetia-anz/gogrpc/greet/greetpb"

type Suite []Case

type Case struct {
	Name     string
	Request  *greetpb.GreetRequest
	Response *greetpb.GreetResponse
}

var TestCases = Suite{
	{
		Name: "Test Dip",
		Request: &greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Dipjyoti",
				LastName:  "Metia",
			}},
		Response: &greetpb.GreetResponse{
			Result: "Dipjyoti",
		},
	},
	{
		Name: "Test Smile",
		Request: &greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: ":-)",
				LastName:  ":-):-)",
			}},
		Response: &greetpb.GreetResponse{
			Result: ":-)",
		},
	},
	{
		Name: "Test Emoji",
		Request: &greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "⛄🙄",
				LastName:  "\U0001F976",
			}},
		Response: &greetpb.GreetResponse{
			Result: "⛄🙄",
		},
	},
}
