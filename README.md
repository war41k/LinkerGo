		
		Dep + plugin build programs call
		Golang
			
		/src/project/main.go + /src/project/plugin.go 
		$ dep init
		$ dep ensure			
		$ go build -buildmode=plugin plugin.go
		$ rm plugin.go
		$ go build main.go
		
