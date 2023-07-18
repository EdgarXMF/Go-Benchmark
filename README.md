# GO-Benchmark

This repository contains the resolution to a series of proposed tasks ([jig/bench](https://github.com/jig/bench)) on programming, documentation, and management of a repository.

## Task 1: Performance Testing with Apache's ab

For the first task we have used as web server a [load balancer](https://github.com/EdgarXMF/Go-Benchmark/tree/main/nginxLoadBalancer) made with nginx and 5 services deployed with docker compose.
The measurements have been performed using 2 computers with 4 cores each, connected in the same network, and both using virtualization. 

Running the `apache benchmark tool`, the average latency values obtained are different using the same concurrency configuration and number of messages sent, because one of the computers is connected to a wifi network which is unstable, but the values remain between 200 and 600 ms.

When checking the cpu load of the computer that contains the server, we can see that it increases quite a lot when multiplying by 10 the concurrency, reaching 95% of the cpu used, thus requiring that if the web service is expected to be accessed by more than 1000 users concurrently, it will be necessary to replicate the services in different machines, since in the same one they still share CPU.


## Task 2: Custom Implementation of goab

For the second task, a reduced version of the ab tool has been implemented with the `-c` `-k` `-n` flags in `Go` and tested with the load balancer mentioned above. This program has also been `dockerized`. 

This [goab](https://github.com/EdgarXavier/GO-Benchmark/tree/main/goab) makes requests with HTTP/1.1 since `Go` does not allow the use of HTTP/1.0 as `ab` does.
 
In this case some of the results comparing `ab` and `goab` would be:

- HTTP Server: Nginx load balancer
  - `ab` Results:
    - Concurrency Level: 100
      - TPS: 961
      - Average Latency: 104 ms
      - Errored Responses: 0 (0%)
    - Concurrency Level: 1000
      - TPS: 2048
      - Average Latency: 488 ms
      - Errored Responses: 0 (0%)
  - `goab` Results:
    - Concurrency Level: 100
      - TPS: 786
      - Average Latency: 5 ms
      - Errored Responses: 0 (0%)
    - Concurrency Level: 1000
      - TPS: 1499
      - Average Latency: 667 ms
      - Errored Responses: 0 (0%)

 
We can observe that as the concurrency and requests increase the latency grows quite a lot and errors start to appear, as the connections timed out.

In comparison with the ab tool, we can see that using the same values in the `-n` and `-k` flag, the TPS and average latency are worse than using `ab`, but not too much.


## Task 3: HTTP Server Benchmarking

A [server](https://github.com/EdgarXavier/GO-Benchmark/tree/main/goserver) has been implemented in go that allows access with https and http (port 80) that redirects to https (port 443).

This server is `dockerized` and can also be deployed with `kubernetes`.

The results after using apache's ab tool in this new http server are:

- `goserver` Results:
  - Concurrency Level: 1000
    - TPS: 1502
    - Average Latency: 665 ms
    - Errored Responses: 0 (0%)

The values obtained are quite similar to those of the nginx load balancer, but the values of the server cpu load change quite a lot since the load is at most 2/3 of what it was with the nginx load balancer. 
