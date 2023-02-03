#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc calculator/calcpb/calc.proto --go_out=plugins=grpc:.

protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.

protoc -I calculator/calcpb -I calculator/calcpb/ --go_out=plugins=grpc:. calculator/calcpb/calc.proto

protoc -I calculator/calcpb -I calculator/calcpb/ --grpc-gateway_out=logtostderr=true:. calculator/calcpb/calc.proto

protoc -I calculator/calcpb -I calculator/calcpb/ --swagger_out ./gen/swagger --swagger_opt logtostderr=true calculator/calcpb/calc.proto

protoc -I blog/blogpb -I blog/blogpb/ --swagger_out ./gen/swagger --swagger_opt logtostderr=true blog/blogpb/blog.proto


