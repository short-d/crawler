# Crawler
Explore the web in parallel on thousands of machines

## TODO

- [ ] Productionize master & worker with [app](https://api.short-d.com/r/fw) framework
- [ ] Configure continuous delivery to bootstrap productivity
- [ ] Create CLI to trigger crawling for a certain site
- [ ] Health check for workers & master
- [ ] Abstract out worker scheduling to support custom algorithms
- [ ] Assign job to new worker when previously assigned worker failed
- [ ] Support checkpoint for master
- [ ] Auto recover system when the master is back online
- [ ] Support pluggable worker script
- [ ] Load worker script from network file system
- [ ] TLS
- [ ] Support Docker Swam
- [ ] Support k8s

## Getting Started

```bash
cd bin
```

### Start Master

```bash
go run master.go 8080
```

#### Output

```
Master started at 8080
```

### Start Workers

Run the following command at different terminals:

```bash
go run worker.go 8081 localhost 8080
go run worker.go 8082 localhost 8080
go run worker.go 8083 localhost 8080
```

#### Output
At the master side:

```
Worker registed: ID(0) IP(localhost) PORT(8081) SECRET(encrypted)
Worker registed: ID(1) IP(localhost) PORT(8082) SECRET(encrypted)
Worker registed: ID(2) IP(localhost) PORT(8083) SECRET(encrypted)
```

At the work side:

```
Registered with master: id(0)
Registered with master: id(1)
Registered with master: id(2)
```

### Try It Out

Send the following gRPC calls to the master:

```grpc
{
 "url": "https://leetcode.com"
}
```

#### Output

On the master side:

```
Start exploring https://leetcode.com
https://leetcode.com
/support/
/jobs/
/bugbounty/
/terms/
/privacy/
/region/
mailto:billing@leetcode.com?subject=Billing%20Issue&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
mailto:support@leetcode.com?subject=General%20Support&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
mailto:feedback@leetcode.com?subject=Other%20Inquiries&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
Finish exploring https://leetcode.com

```

On the worker side:

```
// Worker 0

Start extracting links from https://leetcode.com
/support/
/jobs/
/bugbounty/
/terms/
/privacy/
/region/
mailto:billing@leetcode.com?subject=Billing%20Issue&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
mailto:support@leetcode.com?subject=General%20Support&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
mailto:feedback@leetcode.com?subject=Other%20Inquiries&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
Start extracting links from /support/
Start extracting links from /bugbounty/
Start extracting links from /region/
Start extracting links from /privacy/
```

```
// Worker 1
Start extracting links from /jobs/
Start extracting links from mailto:support@leetcode.com?subject=General%20Support&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
```

```
// Worker 2
Start extracting links from mailto:feedback@leetcode.com?subject=Other%20Inquiries&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
Start extracting links from mailto:billing@leetcode.com?subject=Billing%20Issue&body=Name:%0D%0A%0D%0AUsername:%0D%0A%0D%0AMessage:%0D%0A%0D%0A
Start extracting links from /terms/
```

## Author

- **Yang Liu** - *Initial work* - [byliuyang](https://github.com/byliuyang)
- **Vinod Krishnan** - *Incremental improvements* - [vtkrishn](https://github.com/vtkrishn)

## License
This project is maintained under MIT license.
