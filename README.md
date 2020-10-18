# paxos for studying

## 一、Paxos共识算法介绍

* Paxos算法是分布式系统中的一个共识算法家族（好像也就这一个？）

* Paxos( [The Part-Time Parliament](http://lamport.azurewebsites.net/pubs/lamport-paxos.pdf) )共识算法由Leslie Lamport首次在1990年发表在TOCS，后来由于大多数人不太能接受他的幽默风趣的介绍方法（其实用比喻的方式介绍长篇的理论，确实让人比较难理解），于是在2001年重新写一篇名叫 [Paxos Made Simple](http://lamport.azurewebsites.net/pubs/paxos-simple.pdf) 论文，相当于原始Paxos算法的简化版，这篇文章与原始文章在讲述Paxos算法上最大的不同就是用的都是计算机术语，看起来也轻松很多

* "世界上只有两种分布式共识算法，一种是Paxos算法，另一种是类Paxos算法"

* 现在比较流行的zab和raft算法也是基于Paxos算法设计的

## 二、Basic paxos（基本的Paxos算法）

目的

假设条件

safety&liveness


## 三、Mutil-paxos

the run of paxos or the instance of paxos

## 四、Implementing state machine

basic paxos -> mutil-paxos -> replicated state machine


## 五、Paxos算法go语言实现

