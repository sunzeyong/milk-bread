# custom resource definition
有了CRD 就可以创建CR了 但是只是服务端识别了 客户端如何操作还需要编码实现

# custom resource

自定义控制器里面有informer和同的handler handler中内容就是要进行调谐的逻辑 根据实际状态和期望状态对比来操作
所谓的 Informer，就是一个自带缓存和索引机制，可以触发 Handler 的客户端库。这个本地缓存在 Kubernetes 中一般被称为 Store，索引一般被称为 Index。
admission controller是在创建pod之前的操作 和 自定义控制器有点区别
operate本质上是和controller是一个原理 就是获取定义的API对象信息 然后进行操作