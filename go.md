# Go

- [A curated list of awesome frameworks, libraries and software](https://github.com/avelino/awesome-go)
- [A curated list of awesome storage projects and libraries](https://github.com/gostor/awesome-go-storage)
- [A curated selection of blog posts](https://github.com/enocom/gopher-reading-list)
- [Build your own web framework](https://www.nicolasmerouze.com/build-web-framework-golang/)
- [Simple HTTP request context example](https://gocodecloud.com/blog/2016/11/15/simple-golang-http-request-context-example/)
- [Best practice for accessing database in HTTP handlers](http://www.alexedwards.net/blog/organising-database-access)
- [Code snippets](snippets.go)
- [Applying the clean architecture](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)
- [Shrink your binaries with this one weird trick](https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/)
- [Stripping dependency bloat](https://medium.com/@valyala/stripping-dependency-bloat-in-victoriametrics-docker-image-983fb5912b0d)
- [How to improve latency by 700% using sync.Pool](http://www.akshaydeo.com/blog/2017/12/23/How-did-I-improve-latency-by-700-percent-using-syncPool/)
- [Packing multiple binaries in a package](https://ieftimov.com/golang-package-multiple-binaries)
- [Supercronic - crontab-compatible job runner designed to run in containers](https://github.com/aptible/supercronic)
- [Error handling](https://blog.golang.org/error-handling-and-go#TOC_3)
- [Best practices for writing high-performance code](https://github.com/dgryski/go-perfbook/blob/master/performance.md)
- [Functional Options - Implementing clean APIs](https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md)
- [Using the Service Object pattern](https://www.calhoun.io/using-the-service-object-pattern-in-go/)
- [Integration with systemd](https://vincent.bernat.im/en/blog/2017-systemd-golang)
- [50 shades of Go: traps, gotchas, and common mistakes](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)
- [Packaging application for OSX](https://medium.com/@mattholt/packaging-a-go-application-for-macos-f7084b00f6b5)
- [How to create an SSH tunnel](http://elliot.land/post/how-to-create-an-ssh-tunnel-in-go)
- [Uber Go style guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)
- [Adopting Go Modules and integrating with Github Actions - Handle different application versions and conditional build in GitHub Actions](https://jcmturner.github.io/go/2020/02/01/adopting-go-modules.html)
- [Thanos Coding Style Guide](https://www.bwplotka.dev/2020/how-thanos-would-program-in-go/)
- [Four Steps to Daemonize Your Go Programs](https://ieftimov.com/post/four-steps-daemonize-your-golang-programs/)
- [Ultimate Go - notes on work with Go](https://github.com/betty200744/ultimate-go)
- [Implementing Clean Arhitecture in Go - DDD](https://dormoshe.io/trending-news/implementing-clean-architecture-in-go-14502)
- [From 0 to Insight with OpenTelemetry in Go](https://www.honeycomb.io/blog/from-0-to-insight-with-opentelemetry-in-go/)
- [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)

## Blogs

- [Dave Cheney](https://dave.cheney.net/)
- [Ivan Daniluk](https://divan.github.io/)
- [Ardan labs](https://www.ardanlabs.com/blog/)
- [Johan Brandhorst](https://jbrandhorst.com/)

## Concurrency

- [Multithreading](https://pragmacoders.com/multithreading-go-tutorial/)
- [Data races and how to fix them](https://www.sohamkamani.com/blog/2018/02/18/golang-data-race-and-how-to-fix-it/)
- [Interesting ways of using channels](http://nomad.uk.net/articles/interesting-ways-of-using-go-channels.html)
- [Concurrency and a mini load-balancer](https://gist.github.com/rushilgupta/228dfdf379121cb9426d5e90d34c5b96)
- [Handling 1 million requests per minute](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/)
- [Common concurrent programming mistakes](https://go101.org/article/concurrent-common-mistakes.html)
- [Channel use cases](https://go101.org/article/channel-use-cases.html)
- [Load balancer](https://medium.com/@owlwalks/load-balancer-at-your-fingertips-golang-ea23d7aaee82)

## Syncing goroutines

- [errgroup](https://godoc.org/golang.org/x/sync/errgroup)
- [tomb v2](https://github.com/go-tomb/tomb/tree/v2)
- [tomb v1](https://github.com/go-tomb/tomb)
- [oklog/run](https://github.com/oklog/run)

## HTTP server

- [Fiber - Express-inspired web framework using Fasthttp](https://gofiber.io/)
- [Goyave - progressive and accessible web application framework](https://github.com/System-Glitch/goyave)
- [chi](https://github.com/go-chi/chi)
- [Gocraft](https://github.com/gocraft/web)
- [Creating the V in MVC](https://www.calhoun.io/intro-to-templates-p4-v-in-mvc/)
- [aah - a secure, flexible, rapid Go web framework](https://aahframework.org/)
- [How I write HTTP services after seven years](https://medium.com/@matryer/how-i-write-go-http-services-after-seven-years-37c208122831)
- [http.Handler wrapper technique](https://medium.com/@matryer/the-http-handler-wrapper-technique-in-golang-updated-bc7fbcffa702)
- [CertMagic - automagic TLS](https://github.com/mholt/certmagic)
- [Serve local files](https://github.com/syntaqx/serve)
- [Authenticate a Go API with JSON web tokens](https://www.thepolyglotdeveloper.com/2017/03/authenticate-a-golang-api-with-json-web-tokens/)
- [Expose Go on the Internet](https://blog.cloudflare.com/exposing-go-on-the-internet/)
- [GO_REUSEPORT](https://github.com/kavu/go_reuseport)
- [API Design Cheat Sheet](https://github.com/RestCheatSheet/api-cheat-sheet)
- [Create Secure Clients and Servers in Golang Using HTTPS](https://medium.com/better-programming/create-secure-clients-and-servers-in-golang-using-https-aa970ba36a13)

## Auth

- [JWT auth in Go](https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-dde432440924)
- [JWT auth in Go Part 2 — Refresh Tokens](https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-part-2-refresh-tokens-d334777ca8a0)
- [An SWT based API for managing users and issuing SWT tokens](https://github.com/netlify/gotrue)
- [Go password validator](https://github.com/lane-c-wagner/go-password-validator)

## Authz

- [Casbin custom functions](https://casbin.org/docs/en/function#how-to-add-a-customized-function)
- [Casbin custom functions](https://github.com/casbin/casbin/issues/33)
- [How to implement RBAC custom function which has to read information from database?](https://github.com/casbin/casbin/issues/326#issuecomment-552875116)

## HTTP client

- [Don’t use default HTTP client in production](https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779)
- [Resty - imple HTTP and REST client library](https://github.com/go-resty/resty/)
- [Transport for http.Client that will cache responses according to the HTTP RFC](https://github.com/gregjones/httpcache)
- [Ferret - web scraper](https://github.com/MontFerret/ferret)
- [Colly - Lightning Fast and Elegant Scraping Framework](https://github.com/gocolly/colly)
- [How to customize Go's HTTP client](https://rafallorenz.com/go/customize-http-client/)
- [Gohalt - Throttler library & client rate limiter](https://github.com/1pkg/gohalt)

## Websocket

- [Minimal and idiomatic WebSocket library](https://github.com/nhooyr/websocket)

## Docker

- [Docker powered development environment for your web app](https://medium.com/developers-writing/docker-powered-development-environment-for-your-go-app-6185d043ea35)
- [Live reloading server in Docker container](https://github.comhttps://blog.golang.org/go1.13-errorsm/how-to-build-microservice-with-mongodb-in-golang/)
- [Mocking Mongo](http://thylong.com/golang/2016/mocking-mongo-in-golang/)
- [Writing integration tests with mongoDB support](https://medium.com/@mvmaasakkers/writing-integration-tests-with-mongodb-support-231580a566cd)

## NATS

- [Introducing NATS](https://medium.com/@shijuvar/introducing-nats-to-go-developers-3cfcb98c21d0)
- [Building distributed systems and microservices NATS Streaming](https://medium.com/@shijuvar/building-distributed-systems-and-microservices-in-go-with-nats-streaming-d8b4baa633a2)

## RabbitMQ

- [Make sure your connection to RabbitMQ stays active](https://ninefinity.org/post/ensuring-rabbitmq-connection-in-golang/)
- [Using AMQP in a single goroutine that publishes and consumes](https://github.com/streadway/amqp/issues/117)
- [HTTP AMQP Proxy](https://github.com/iyidan/http-proxy-amqp)
- [Tips for using RabbitMQ](https://agocs.org/blog/2014/08/19/tips-for-using-rabbitmq-in-go/)
- [Examples](https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go)
- [Make sure your connection to RabbitMQ stays active](https://ninefinity.org/post/ensuring-rabbitmq-connection-in-golang/)

## AWS

- [Code examples](https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code)
- [Creating a Lambda function](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/lambda-go-example-create-function.html)
- [Lambda function handler](https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-handler-types.html)

## Aerospike

- [Driver](https://github.com/aerospike/aerospike-client-go)

## Lambda

- [Shim is a thin layer between API Gateway integration requests via Lambda and the standard library http.Handler interface](https://github.com/iamatypeofwalrus/shim)
- [Porting web applications to AWS Lambda](http://artem.krylysov.com/blog/2018/01/18/porting-go-web-applications-to-aws-lambda/)
- [AWS Lambda proxy - run Lambda function on your machine](https://github.com/mtojek/aws-lambda-go-proxy)
- [Getting started with serverless](https://yos.io/2018/02/08/getting-started-with-serverless-go/)
- [How to build a Serverless API with AWS Lambda](http://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda)

## AI

- [Face Detection using OpenCV and MachineBox](https://www.youtube.com/watch?v=rbZeZNVA-Q4)

## JSON

- [JSON and struct composition](https://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/)
- [JSON decoding](https://attilaolah.eu/2013/11/29/json-decoding-in-go/)
- [Custom JSON marshalling](http://choly.ca/post/go-json-marshalling/)

## SQL

- [SQLX - extensions on top of database/sql package](https://jmoiron.github.io/sqlx/)
- [Grimoire ORM](https://github.com/Fs02/grimoire)
- [kingshard - high-performance proxy for MySQL](https://github.com/flike/kingshard)
- [Vitess - database clustering system for horizontal scaling of MySQL](https://vitess.io/)
- [Nap - library that abstracts access to master-slave physical SQL servers](https://github.com/tsenart/nap)
- [goose - database migration tool](https://github.com/pressly/goose)
- [migrate - database migrations written in Go](https://github.com/golang-migrate/migrate)

## GraphQL

- [Using GraphQL in Golang](https://dev.to/stevensunflash/using-graphql-in-golang-3gg0)
- [gqlgen - Go library for building GraphQL servers](https://github.com/99designs/gqlgen)

## Microservices

- [Distributed tracing in 10 minutes](https://medium.com/opentracing/distributed-tracing-in-10-minutes-51b378ee40f1)
- [Building a microservices application following the CQRS pattern](https://outcrawl.com/go-microservices-cqrs-docker/)
- [Microservices](http://callistaenterprise.se/blogg/teknik/2017/02/17/go-blog-series-part1/)
- [Microservices example](https://github.com/wizelineacademy/GoWorkshop)
- [Automatically set GOMAXPROCS to match Linux container CPU quota](https://github.com/uber-go/automaxprocs)

## CI & CD

- [How to do Continuous Integration like a boss](https://medium.com/pantomath/go-tools-gitlab-how-to-do-continuous-integration-like-a-boss-941a3a9ad0b6)
- [Building a CI system with Jenkins](https://medium.com/@zarkopafilis/building-a-ci-system-for-go-with-jenkins-4ab04d4bacd0)

## Tools

- [GoConvey - automatic web UI for testing](https://github.com/smartystreets/goconvey)
- [Collection of tools and libraries including linters and static analysis](https://github.com/dominikh/go-tools)
- [Concurrently run lint tools and normalize their output](https://github.com/alecthomas/gometalinter)
- [A curated list of awesome linters (more than 60 linters and tools)](https://github.com/golangci/awesome-go-linters)
- [GopherCI helps you maintain high-quality Go projects by checking each GitHub pull request](https://github.com/bradleyfalzon/gopherci)
- [A good Makefile](https://azer.bike/journal/a-good-makefile-for-go/)
- [SonarQube](https://medium.com/red6-es/go-for-sonarqube-ffff5b74f33a)
- [Overview of tooling](https://www.alexedwards.net/blog/an-overview-of-go-tooling)
- [Find outdated dependencies](https://github.com/psampaz/go-mod-outdated)
- [Using Makefiles for Go](https://danishpraka.sh/2019/12/07/using-makefiles-for-go.html)
- [ko - tool for building and deploying Golang applications to Kubernetes](https://github.com/google/ko)
- [profefe continuously collects profiling data](https://github.com/profefe/kube-profefe)
- [vssh - library to handle tens of thousands SSH connections and execute the command(s)](https://github.com/yahoo/vssh)
- [fgprof - profiler that allows you to analyze On-CPU as well as Off-CPU (e.g. I/O) time together](https://github.com/felixge/fgprof)
- [Chronos - static race detector](https://github.com/amit-davidson/Chronos)
- [structslop - checks if struct's fields can be re-arranged to get optimal struct size](https://github.com/orijtech/structslop)

## Raspberry PI

- [Building your own build status indicator](https://scene-si.org/2016/07/19/building-your-own-build-status-indicator-with-golang-and-rpi3/)

## Testing

- [Advanced testing](https://www.youtube.com/watch?v=yszygk1cpEc)
- [In-memory network stack](https://akutz.wordpress.com/2018/04/20/memconn/)
- [Testing HTTP handlers](https://blog.questionable.services/article/testing-http-handlers-go/)
- [go-replayers - tools for recording and replaying RPCs](https://github.com/google/go-replayers)
- [httpexpect - Concise, declarative, and easy to use end-to-end HTTP and REST API testing for Go](https://github.com/gavv/httpexpect)
- [Testing with golden files in Go](https://medium.com/soon-london/testing-with-golden-files-in-go-7fccc71c43d3)
- [copyist - Mocking your SQL database in Go tests has never been easier](https://github.com/cockroachdb/copyist)
- [Testcontainers - friendly API to run Docker container, designed to create runtime environment for automatic tests](https://github.com/testcontainers/testcontainers-go)
- [Go advanced testing tips & tricks](https://povilasv.me/go-advanced-testing-tips-tricks/)
- [aceptadora - boilerplate to orchestrate the containers for an acceptance test](https://github.com/cabify/aceptadora)

## Public API's

- [Client library to access data of public transport companies](https://github.com/michiwend/goefa)

## gRPC

- [Building high performance API's using gRPC and Protocol Buffers](https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b)
- [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [How to develop gRPC microservice - Part 1](https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9)
- [How to develop gRPC microservice - Part 2: HTTP/REST endpoint](https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-af1fff81aeb2)
- [How to develop gRPC microservice - Part 3: Middleware](https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e)
- [protoc-gen-mock - Plugin to generate gRPC mock services from protobuf specifications](https://github.com/carvalhorr/protoc-gen-mock)
- [Awesome gRPC](https://github.com/grpc-ecosystem/awesome-grpc#go)
- [Wombat - Cross platform gRPC client (can be used to test gRPC servers)](https://github.com/rogchap/wombat)

## E-Mail

- [Hermes - Clean, responsive HTML e-mails](https://github.com/matcornic/hermes)

## Config

- [konfig - composable, observable and performant config handling for Go for the distributed processing era](https://github.com/lalamove/konfig)
- [A clean way to pass configs](https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64)

## Command line parsers

- [Kong](https://github.com/alecthomas/kong)

## Miscellaneous

- [File type discovery using magic numbers](https://github.com/h2non/filetype)
- [v8go - Execute JavaScript from Go](https://github.com/rogchap/v8go)
- [x/text - packages related to internationalization (i18n) and localization (l10n)](https://godoc.org/golang.org/x/text)
- [MTProto - Full-native implementation of MTProto (Telegram) protocol](https://github.com/xelaj/mtproto)
- [Mux Go - Official Mux API wrapper for golang projects, supporting both Mux Data and Mux Video](https://github.com/muxinc/mux-go)

## Images

- [caire - Content aware image resize library](https://github.com/esimov/caire)

## Error handling

- [cockroachdb/errors](https://github.com/cockroachdb/errors)

## Desktop and mobile

- [fyne - Cross platform GUI](https://github.com/fyne-io/fyne)

## Vanity import

- [Vanity import paths in Go](https://sagikazarmark.hu/blog/vanity-import-paths-in-go/)
- [Vanity Go Import Paths](https://blog.bramp.net/post/2017/10/02/vanity-go-import-paths/)
- `wget https://mvdan.cc/gofumpt`

## Kubernetes

- [Sample controller](https://github.com/kubernetes/sample-controller)
- [Testing controllers](https://github.com/kubernetes/client-go/blob/master/examples/fake-client/main_test.go)
- [Testing examples](https://github.com/profefe/kube-profefe/blob/master/pkg/kubeutil/kube_test.go)
- [Leader Election Example](https://github.com/kubernetes/kubernetes/tree/323f34858de18b862d43c40b2cced65ad8e24052/staging/src/k8s.io/client-go/examples/leader-election)

## Test tasks

https://play.golang.org/p/n3iH7H11Uoe

```go
package main

const N = 3

func main() {
	m := make(map[int]*int)

	for i := 0; i < N; i++ {
		m[i] = &i //A
	}

	for _, v := range m {
		print(*v)
	}
}
```

https://play.golang.org/p/BW0lY5dW7ZJ

```go
package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}
```

https://play.golang.org/p/zgSTwr035YQ

```go
package main

import (
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}
```

https://play.golang.org/p/6bzWYtdAh9K

```go
package main

func main() {
	s := "123"
	ps := &s
	b := []byte(*ps)
	pb := &b

	s += "4"
	*ps += "5"
	b[1] = '0'

	println(*ps)
	println(string(*pb))
}
```
