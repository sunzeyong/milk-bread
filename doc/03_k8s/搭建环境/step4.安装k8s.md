> 本文适合安装1.23版本及其以前版本，1.24版本后不再使用docker，默认使用的运行时变了

# 检查
1. 关闭swap
   
```
sudo swapoff -a
sudo vi /etc/fstab
# remove the line with swap keyword
```

2. 为了让Kubernetes能够检查、转发网络流量，你需要修改iptables的配置，启用“br_netfilter”模块：
   
```
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward=1 # better than modify /etc/sysctl.conf
EOF

sudo sysctl --system
```

3. 修改hostname确保不重复
```
sudo vi /etc/hostname
```

4. docker cgroup的驱动程序改成 systemd
见docker安装


# 安装kubeadm kubectl kubelet

## 配置国内源 下载更快
```
sudo apt install -y apt-transport-https ca-certificates curl

curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -

cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

sudo apt update
```

## 下载
```
sudo apt install -y kubeadm=1.23.3-00 kubelet=1.23.3-00 kubectl=1.23.3-00
```
会输出如下信息
The following additional packages will be installed:
  conntrack cri-tools ebtables kubernetes-cni socat
The following NEW packages will be installed:
  conntrack cri-tools ebtables kubeadm kubectl kubelet kubernetes-cni socat
0 upgraded, 8 newly installed, 0 to remove and 33 not upgraded.
Need to get 75.5 MB of archives.
After this operation, 336 MB of additional disk space will be used.


# 使用kubeadm下载k8s所需要的镜像
查询需要的镜像
kubeadm config images list --kubernetes-version v1.23.3

下载镜像，指定国内的镜像源
kubeadm config images pull --image-repository registry.aliyuncs.com/google_containers

# 初始化master节点
初始化master，即使下载好镜像也需要指定镜像源，不指定就会去下载k8s.gcr.io源下镜像
sudo kubeadm init --pod-network-cidr=10.10.0.0/16 --kubernetes-version=v1.23.3 --image-repository registry.aliyuncs.com/google_containers

打印以下部分则表示成功
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 10.211.55.4:6443 --token h5n2rp.9ku02lfe2zz3wh2s \
	--discovery-token-ca-cert-hash sha256:b7dbc0040cfad408c54526b8c0352bf6cf985db6860c956865d813a41526afac 

## 重新创建token
``` 
kubeadm token create --print-join-command
```

## 安装网络插件
kubectl apply -f flannel.yaml

coredns负责域名解析;kube-proxy 负责实现service的配置 负载均衡和稳定的对外IP；cni插件负责实现IPperPod功能 



参考 https://cloud.tencent.com/developer/article/1888163 节点的添加和删除
https://blog.51cto.com/u_11726705/7395313 参考更新key
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo gpg --no-default-keyring --keyring gnupg-ring:/etc/apt/trusted.gpg.d/NAME.gpg --import

https://gist.github.com/islishude/231659cec0305ace090b933ce851994a 和本内容相似

https://blog.csdn.net/jmh1996/article/details/80432780
https://www.myfreax.com/adding-external-repositories-ubuntu/

```
// 要注意k8s版本和calico版本的兼容性
kubectl apply -f https://docs.projectcalico.org/archive/v3.24/manifests/calico.yaml 
# https://docs.projectcalico.org/manifests/calico.yaml
```
https://blog.51cto.com/omaidb/5264767 参考这里再实验一下
https://www.cnblogs.com/khtt/p/16563088.html 也可参考
再次运行

minikube安装
https://juejin.cn/post/7054077165888864292 国内 mibikube安