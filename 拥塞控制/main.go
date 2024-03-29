package 拥塞控制

//拥塞控制：
//1.当拥塞窗口小于Threshold,处于慢启动状态,拥塞窗口的大小为指数增长
//2.当拥塞窗口大小大于等于Threshold的时候，进入拥塞避免状态,拥塞窗口呈线性增长
//3.loss事件发生的时候，分情况
//4.如果是3次ack，说明此时拥塞窗口还有一部分空间，那么threshold变为之前拥塞窗口的1/2,拥塞窗口也变为之前的1/2
//5.如果是timeout(超时),说明此时拥塞比较严重，拥塞窗口直接设为1
