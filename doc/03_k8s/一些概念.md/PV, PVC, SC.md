PVC可以理解为持久化存储的“接口”，它提供了对某种持久化存储的描述，但不提供具体的实现；而这个持久化存储的实现部分则由PV负责完成。

# pv
这个API对象主要定义的是一个持久化存储在宿主机上的目录，比如一个NFS的挂载目录。
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nfs-pv-test1
spec:
  storageClassName: manual
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteMany
  nfs:
    server: 10.211.55.5
    path: "/mnt/nfs"
```
如果有多个pv声明了相同的nfs路径，那么最终在pod中使用时会共享这个文件夹。

# pvc
描述的Pod所希望使用的持久化存储的属性。

```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nfs-pv-test1
spec:
  accessModes:
    - ReadWriteMany
  storageClassName: manual
  resources:
    requests:
      storage: 1Gi
```

## 使用pvc
1. 在pod层声明使用什么存储,并起名name
2. 在container中声明使用pod中哪个name,和container中的路径

```
apiVersion: v1
kind: Pod
metadata:
  name: pod-use-nfs-pvc
  labels:
    role: web-frontend
spec:
  containers:
  - name: web
    image: nginx
    ports:
      - name: web
        containerPort: 80
    volumeMounts:
        - name: nfs-mount1 # 挂载到容器路径
          mountPath: "/usr/share/nginx/html"

  volumes: # 声明使用哪个pvc
  - name: nfs-mount1
    persistentVolumeClaim:
      claimName: nfs-pv-test1

```

进入到pod中，在mountPath目录下创建文件后在nfs的文件下查看


# sc
storageClass可以作为创建pv的模版，方便在大规模集群中自动创建pv来匹配pvc。若上述pv和pvc中storageClassName没有对应的api对象，则这个过程绑定叫做Static Provisioning，若有对应的api对象，这个过程Dynamic Provisioning。
storageClass对象会定义如下两个部分内容：
- 第一，PV 的属性。比如，存储类型、Volume 的大小等等。
- 第二，创建这种 PV 需要用到的存储插件。比如，Ceph 等等。

使用storageClass需要自定义的Provisioner控制器。
## NFS Subdir External Provisioner
### helm方法
```
// 新增repo
helm repo add nfs-subdir-external-provisioner https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/
helm repo update

// 生成value文件
helm show values nfs-subdir-external-provisioner/nfs-subdir-external-provisioner > values.yaml

// 修改value.yaml
nfs:
  server: 
  path: /mnt/nfs
  mountOptions: 10.211.55.5
  volumeName: nfs-subdir-external-provisioner-root
  # Reclaim policy for the main nfs volume
  reclaimPolicy: Retain

// 检查最终的yaml文件
helm install -f values.yaml nfs nfs-subdir-external-provisioner/nfs-subdir-external-provisioner --dry-run

// 进行部署
kubectl create ns nfs-provisioner
helm install -f values.yaml nfs nfs-subdir-external-provisioner/nfs-subdir-external-provisioner -n nfs-provisioner

// 校验 可以看到一个nfs-client的storageClass
kubectl get sc
```


# 挂载逻辑
在使用nfs的pvc时，当pod被分配到某个node上后，kubectl需要在宿主机上创建volume，将nfs端共享目录挂载到宿主机的volume。后续在kubectl在创建pod时就可以直接将volume挂载到pod中。
kubectl在挂载远程nfs端时，需要宿主机有nfs客户端工具，所以发生如下报错时，是没有相关工具，导致挂载命令执行失败。
```
Normal   Scheduled    15s               default-scheduler  Successfully assigned default/pod-use-nfs-pvc-v2 to node-3
  Warning  FailedMount  7s (x5 over 14s)  kubelet            MountVolume.SetUp failed for volume "nfs-pv-test2" : mount failed: exit status 32
Mounting command: mount
Mounting arguments: -t nfs 10.211.55.5:/mnt/nfs /var/lib/kubelet/pods/706a913e-9881-4a61-9381-1caf0b20f6f4/volumes/kubernetes.io~nfs/nfs-pv-test2
Output: mount: /var/lib/kubelet/pods/706a913e-9881-4a61-9381-1caf0b20f6f4/volumes/kubernetes.io~nfs/nfs-pv-test2: bad option; for several filesystems (e.g. nfs, cifs) you might need a /sbin/mount.<type> helper program.
```

# CSI插件kaifa
flex volume方式可以简单的时间pv功能，但是provision需要csi模式才能做到。
csi主要需要两方面，其一是controller来监听pvc，根据需求创建pv，在进行绑定时需要kubelet上进行两阶段操作，attach和mount后才能提供一个在宿主机上持久化volume。后续起pod时就可以直接使用这个volume了。