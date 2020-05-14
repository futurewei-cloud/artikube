# artikube - Golang implementation of an artifact repository manager

Artikube is an open-source aritfact repository manager written in Go (Golang), with the intention to support various cloud storage backends and be deployed in containers orchestrated by Kubernetes.   

One object of this project is explore the simple and powerful Golang and its libraries implemented and supported by the community, by comparing to the tranditional Java impelmentation and JVM runtime.  We would also like to test some of tooling in Golang such as gofmt, go mod, go doc, go test, etc.  

## Golang Framework for Web Development 
Surprisingly, there are so many web framework already released by the community.  I spend time reading a few of them including Gin, Mux, Beego and Go standard library net/http.  I decide to go with Gin mainly because of its performance, adoption rate and the acivitity in the Github repo. 

## Logging 
This seems to be an easy decision to choose Zap package published by Uber. Lots of projects I follow in Github are using this framework 

