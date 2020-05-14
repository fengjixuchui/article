# 基于mqtt的推送
mqtt跟grpc和websocket等不同，mqtt是我在工作当中没有使用到的技术，所以这篇关于mqtt的文章可能会更加详细一些，可能不仅限于mqtt做推送，还会研究一些mqtt的应用场景等。

就推送而言，mqtt绝对是一个非常值得研究的协议，之前文章里面的grpc，websocket有两个方面都没有考虑，第一方面，没有考虑消息丢失的情况，像股票行情推送这种，丢了就丢了，没什么大不了的，推送的准确性不是非常重要，及时性比较重要；第二个方面，grpc和websocket的发布和订阅实际上只有broker和subscriber的代码，发布者没有关注，实际上有些场景是需要订阅者的同时也是发布者，比如说两台设备相互通信，或者两个客户端相互publish消息。如果考虑这两个方面，websocket或者tcp长连接代码写起来就远远不止这么简单了。
