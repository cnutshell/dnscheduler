## 项目介绍

`DN Store` 负责拉起/停止 `DN Shard`，并定期上报 DN Shard 运行状态。

包 `github.com/cnutshell/dnscheduler` 实现了 DN Shard 调度有关的逻辑。

调度逻辑在函数 `ScheduleDN`  中，输入参数为：

- expected - 期望运行的 DN Shard 数量
- stores - 集群中 DN Store 运行状态

```go
func ScheduleDN(expected uint32, stores []*DNStore) (*DNOperator, error)
```

根据集群中实际运行的 DN Shard 数量，以及期望运行的 DN Shard 数量，每次调用 `ScheduleDN` 会生成一个调度命令 `DNOperator`。

通过调度命令 `DNOperator` 指出哪一个 DN Store 需要执行什么样的命令（拉起 or 停止）。

## 代码结构

先过滤不符合条件的 DN Store，再对符合条件的 DN Store 进行打分，经过归一化后选择打分最高的。

- 过滤器逻辑在文件 `filter.go` 中；
- 打分逻辑在文件 `score.go` 中；
- 包入口函数在文件 `scheduler.go` 中；

```bash
.
├── scheduler.go  # DN Store 调度
├── scheduler_test.go
├── filter.go     # DN Store 过滤
├── filter_test.go
├── score.go      # DN Store 打分
└── score_test.go
```

- [ ] 实现仅关心集群中运行的 DN Shard 数量，通过对比期望运行的数量确定调度逻辑。实现不关心调度命令如何递送，亦不关心调度命令送达后 DN Store 如何拉起/停止 DN Shard。

## 更新

- DN Shard 同 DN Store 为 1:1
- DN Shard 为有状态服务，通过 ID 识别，不能单纯的基于数量调度
- DN Shard 采用 uint64 类型的 ID
- DN Store 采用 uuid 类型的 ID
- 调度逻辑内部维护 Operator，针对某一个 DN Shard，调度命令只发一次
- 一次调用可能有多个情况需要调度
- 参考复用 pd 框架，尽量统一两者的调度
