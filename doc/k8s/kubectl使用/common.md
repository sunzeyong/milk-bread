# 常用命令

## kubectl apply
```
// 创建单个API对象
kubectl apply -f example.yaml

// 创建文件夹下全部API对象
kubectl apply -f <directory>
```

## kubectl get
```
// 获取全部的pod
kubectl get pods

// -o wide输出包含一些附加信息，-w表示开启资源对象更新的监控
kubectl get pod -o wide -w
```

## kubectl describe 
```
// 列出全部pod的详情
kubectl describe pod

kubectl describe pod <pod-name>
```

## kubectl delete
```
kubectl delete -f pod.yaml

kubectl delete pods,services -l <label-key>=<label-value>
```

## kubectl exec
```
// 在pod中第一个容器中执行命令
kubectl exec <pod-name> -- date

// 指定容器执行命令
kubectl exec <pod-name> -c <container-name> -- date

// 进入交互终端界面
kubectl exec -it <pod-name> -- /bin/bash
```

## kubectl logs
```
// 打印容器中日志，若不添加-c 则默认第一个, -f 持续输出, --tail 10 输出末尾10行
kubectl logs -f <pod-name> -c <container-name> --tail 10
```


## kubectl edit
```
kubectl edit <resource>/<name>
eg: kubectl edit deploy/ngx-dep

```


## kubectl api-resources
可以查看当前kubectl版本支持的所有资源类型