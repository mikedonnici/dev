build:
	protoc -I greet/greetpb/ greet/greetpb/greet.proto --go_out=plugins=grpc:greet/greetpb
	protoc -I calculator/calculatorpb/ calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:calculator/calculatorpb
