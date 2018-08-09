		
		Dep + plugin build programs cal
		Golang
			
		/src/project/main.go + /src/project/plugin.go 
		$ dep init
		$ dep ensure			
		$ go build -buildmode=plugin plugin.go
		$ rm plugin.go
		$ go build 
		
