## [In Progress] UNIQ (Pronounced : UNI-Q)

##### Todo 
- [ ] Ensure high availability of queue using sockets with configurable replicas
- [ ] Handle scenario when replica/master goes down, another one is brought up automatically based on configuration
- [ ] Retry mechanism with configurable retries when offloading work to consumers
- [ ] Support for dlq after 'n' failed retries
- [ ] Expect reponse in 'x' time from consumers, or retry again if eligible
- [ ] Ensure message gets delivered only once(optional)
- [ ] Remove message from queue after acknowledged