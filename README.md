# kube-nurse
<p align="center"><img width="200px" src="https://kube-nurse.oss-cn-beijing.aliyuncs.com/kube-nurse.png"/></p>
kube-nurse is a kubernetes system diagnostic tool. Command <b>kubectl cluster-info dump</b> can dump many useful information about the cluster.But the content isn't very friendly for human reading. kube-nurse can split the whole content into seperate files and analysis key words about the core component of kubernetes. 


## usage 
```shell
kubectl cluster-info dump > cluster.dump 
kube-nurse diagnose cluster.dump 
```
## *License*
This software is released under the Apache 2.0 license.
