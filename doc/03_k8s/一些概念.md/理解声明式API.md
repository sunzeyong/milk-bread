# 如何理解
1. 所谓“声明式”，指的就是我只需要提交一个定义好的 API 对象来“声明”，我所期望的状态是什么样子。
2. 声明式 API”允许有多个 API 写端，以 PATCH 的方式对 API 对象进行修改，而无需关心本地原始 YAML 文件的内容。
3. 有了上述两个能力，Kubernetes 项目才可以基于对 API 对象的增、删、改、查，在完全无需外界干预的情况下，完成对“实际状态”和“期望状态”的调谐（Reconcile）过程。

因为是可以有多个api写端，所以可以在API对象创建过程中，可以增加自定义的配置以patch方式添加到API对象的定义中。
在此过程中，就可以添加一些自定义控制器。

“声明式 API”并不像“命令式 API”那样有着明显的执行逻辑。这就使得基于声明式 API 的业务功能实现，往往需要通过控制器模式来“监视”API 对象的变化（比如，创建或者删除 Network），然后以此来决定实际要执行的具体工作。


# api对象如何定位
Metadata中有两个最重要的属性：Namespace和Name，分别定义了对象的Namespace归属及名字，这两个属性唯一定义了某个对象实例。