# Triglav Agent

Triglav Agent is a tool for [Triglav](https://github.com/kentaro/triglav). Which provides some functionalities as below:

  1. System information collector
  2. Message proxy between Triglav server and agents in each hosts
  3. Command executor

## 1. System information Collector

Triglav is an inventry management system. Triglav Agent is installed in each hosts to work as a one-shot command or agent daemon to collect system information.

```
          <----- [WAN] <----- [Agent]
[Triglav] <----- [WAN] <----- [Agent]
          <------------------ [Agent]
```

## 2. Message proxy

Not all hosts can access to a Triglav server if the server can be reached on the Internet. In case that, Triglav Agent can work also message bypassing proxy between a Triglav server and agents.

```
          <----- [WAN] <-------------------- [Agent]
[Triglav] <----- [WAN] <----- [Proxy] <----- [Agent]
          <--------------------------------- [Agent]
```

## 3. Command executor

We're thinking of that once agents in hosts and a Triglav server are set, the system can provide a feature that is like [MCollective](http://puppetlabs.com/mcollective), even though it might become sort of fragile. It sounds, however, practical at least for us.

```
          -----> [WAN] --------------------> [Agent]
[Triglav] -----> [WAN] -----> [Proxy] -----> [Agent]
          ---------------------------------> [Agent]
```

## See Also

  * [Triglav](https://github.com/kentaro/triglav)

## Author

  * [Kentaro Kuribayashi](http://kentarok.org/)

## License

  * MIT [http://kentaro.mit-license.org/](http://kentaro.mit-license.org/)

