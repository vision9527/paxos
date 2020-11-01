# Paxos For Studying

## 一、Paxos Overview

* Paxos( [The Part-Time Parliament](http://lamport.azurewebsites.net/pubs/lamport-paxos.pdf) )共识算法由[Leslie Lamport](https://www.52cs.com/archives/story/acm-sigops%e5%90%8d%e4%ba%ba%e5%a0%82%ef%bc%88%e7%ac%ac%e5%85%ab%e6%9c%9f%ef%bc%89)在1989年首次发布，后来由于大多数人不太能接受他的幽默风趣的介绍方法（其实用比喻的方式介绍长篇的理论，确实让人比较难理解），于是在2001年重新写一篇名叫 [Paxos Made Simple](http://lamport.azurewebsites.net/pubs/paxos-simple.pdf) 论文，相当于原始Paxos算法的简化版，这篇文章与原始文章在讲述Paxos算法上最大的不同就是用的都是计算机术语，看起来也轻松很多

* Paxos算法是分布式系统中的一个共识算法家族，也是第一个被证明正确性的共识算法

* "世界上只有两种分布式共识算法，一种是Paxos算法，另一种是类Paxos算法"

* 现在比较流行的zab和raft算法也是基于Paxos算法设计的

## 二、Basic Paxos

* **Basic Paxos 是在一轮决策中对一个或多个被提议(propose)的值，最终选出一个值达成共识**
![paxos](./images/paxos.png)
<br/>

* **Basic Paxos 可以解决的问题**
    1. 选主
    2. 资源互斥访问
    3. 复制日志的一致性
    4. 其它
<br/>

* **容错模型**
    1. 异步网络，网络是不可靠的，消息可能丢失、重复、延迟、网络分区，但是不包括拜占庭错误，即消息不会被篡改和伪造
    2. 只要大多数（majority）服务器还运行，决议就能继续进行
<br/>

* [FLP](https://groups.csail.mit.edu/tds/papers/Lynch/jacm85.pdf)**定理**
    1. Agreement：所有server必须对同一个值达成共识
    2. Validity：达成共识的值必须是提议的有效值
    3. Termination：最终会对一个值达成共识（Paxos不一定）
<br/>

* [Safety & Liveness](https://lrita.github.io/images/posts/distribution/safety-and-liveness-properties-a-survey.pdf)

    1. Safety（不会有坏事发生）:
    - 只有一个值达成共识
    - 一个server不会知道某个值达成共识，除非它真的已经达成共识

    2. Liveness（好事一定会发生）: 
    - 最终一定会达成共识（在多个proposer的情况下不能保证，即basic paxos是不能保证终止的，为了保证终止就需要选出一个leader，每次的提议只通过一个leader来发起）
    - 如果共识达成最终所有服务器都能知道
<br/>

* **Paxos 角色构成**

    1. Proposer
    - 处理客户端请求，主动发起提议
    2. Acceptor
    - 被动接收来自Proposer的提议消息，并返回投票结果，通知learner
    3. Learner
    - 被动接收来自Acceptor的消息
    4. 实际中一个节点经常扮演三个角色
<br/>

* **Paxos两阶段协议**

1. Phase1

![phase1](./images/phase1.png)

2. Phase2
![phase2](./images/phase2.png)

* **Paxos经典场景**

1. 单个Proposer发起提议情况

![single_proposer](./images/single_proposer.png)

2. 多个Proposer情况

![many_proposer](./images/many_proposer.png)

3. 死锁情况

![live_lock](./images/live_lock.png)

## 三、Multi Paxos

* 当需要决定多个值时就需要连续执行多次Paxos算法，一般执行一次Paxos算法的过程称作A Paxos Run 或者 A Paxos Instance，连续决定多个值则就需要执行多次Paxos

* 为了解决死锁也就是进展性问题，通过选举一个leader，由唯一的leader发起提议

* 两阶段协议效率太低，可以有优化的空间。在单个Leader的情况下，如果前一次已经accept成功，接下来不再需要prepare阶段，直接进行accept

* 多个instance可以并发的进行

![mutil-paxos](./images/mutil_paxos.png)

## 四、Implementing State Machine

1. Basic Paxos -> Multi Paxos -> Replicated State Machine

2. 本人另一篇关于[Raft](https://github.com/vision9527/raft-demo)算法的分享，可以帮助大家更好的理解共识算法的实际应用和实现。

## 五、Basic Paxos Implement With Go

* 运行测试用例：go test -v

## 六、The Part-Time Parliament

* [The The Part-Time Parliament 个人阅读笔记](https://github.com/vision9527/paxos/blob/main/lamport-paxos-vision9527.pdf)

## 七、Reference

1. [The Part-Time Parliament](http://lamport.azurewebsites.net/pubs/lamport-paxos.pdf)

2. [Paxos Made Simple](http://lamport.azurewebsites.net/pubs/paxos-simple.pdf)

3. [FLP定理](https://groups.csail.mit.edu/tds/papers/Lynch/jacm85.pdf)

4. [Safety & Liveness](https://lrita.github.io/images/posts/distribution/safety-and-liveness-properties-a-survey.pdf)

5. [Understanding Paxos](https://www.cs.rutgers.edu/~pxk/417/notes/paxos.html)

6. [Google TechTalks About Paxos](https://www.youtube.com/watch?v=d7nAGI_NZPk)

7. [Paxos lecture (Raft user study)](https://www.youtube.com/watch?v=JEpsBg0AO6o)

8. [raft-demo](https://github.com/vision9527/raft-demo)

9. [Paxos Made Live - Chubby](https://www.cs.utexas.edu/users/lorenzo/corsi/cs380d/papers/paper2-1.pdf)
