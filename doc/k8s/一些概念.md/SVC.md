# 功能
Kubernetes 之所以需要 Service，一方面是因为 Pod 的 IP 不是固定的，另一方面则是因为一组 Pod 实例之间总会有负载均衡的需求。
1. 服务发现，固定服务访问入口
2. 实现网络四层上的负载均衡，IP+Port 传输层

# 理解
service本质上是一些iptables规则,即kube-proxy通过service进行配置iptables，实现流量在集群内的转发

# 使用细节
1. NodePort方式的svc可以对机群外暴露端口访问，curl每个节点都可以。
2. 对外暴露的端口号可以随机，也可以指定 []ports.nodePort
   