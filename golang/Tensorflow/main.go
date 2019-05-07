package main

import (
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	tensorflowOP "github.com/tensorflow/tensorflow/tensorflow/go/op"
)

//https://log.zvz.im/2018/07/15/go-tensorflow/

func main() {

	// 让我们描述我们的需求：创建图
	// 我们想要定义两个运行时使用的 placeholder
	// 第一个 placeholder A 是 [2, 2] 整数张量
	// 第二个 placeholder x 是 [2, 1] 整数张量
	// 然后计算 Y = Ax
	// 创建图的节点：一个空节点，作为图的根节点
	root := tensorflowOP.NewScope()
	// 定义两个占位符
	A := tensorflowOP.Placeholder(root, tf.Int64, tensorflowOP.PlaceholderShape(tf.MakeShape(2, 2)))
	x := tensorflowOP.Placeholder(root, tf.Int64, tensorflowOP.PlaceholderShape(tf.MakeShape(2, 1)))
	// 定义可以接受 A & x 作为输入的操作节点
	product := tensorflowOP.MatMul(root, A, x)
	// 每次我们将 `Scope` 传入一个操作时，我们都将这个操作置于这个作用域内。
	// 如你所见，我们有一个通过 NewScope 创建空域：
	// 这个空域是我们所创建的图的根，我们用 “/”表示它。
	// 现在我们让 tensorflow 通过我们的定义来构建图。
	// 实体的图是通过我们用域和操作组合起来定义的“抽象”图生成的。
	graph, err := root.Finalize()
	if err != nil {
		// 处理这个错误没有什么用处
		// 如果我们对图的定义做错了，我们只能手动修正这些定义。
		// 它很想一个 SQL 查询过程：如果查询语句错了，我们只能重写它
		panic(err.Error())
	}
	// 至此：我们的图定义语法上就没有问题了。
	// 我们现在可以将其放入一个 Session 中使用了。
	var sess *tf.Session
	sess, err = tf.NewSession(graph, &tf.SessionOptions{})
	if err != nil {
		panic(err.Error())
	}
	// 为了使用占位符，我们必须创建含有数值的张量传入网络中
	var matrix, column *tf.Tensor
	// A = [ [1, 2], [-1, -2] ]
	if matrix, err = tf.NewTensor([2][2]int64{{1, 2}, {-1, -2}}); err != nil {
		panic(err.Error())
	}
	// x = [ [10], [100] ]
	if column, err = tf.NewTensor([2][1]int64{{10}, {100}}); err != nil {
		panic(err.Error())
	}
	var results []*tf.Tensor
	if results, err = sess.Run(map[tf.Output]*tf.Tensor{
		A: matrix,
		x: column,
	}, []tf.Output{product}, nil); err != nil {
		panic(err.Error())
	}
	for _, result := range results {
		fmt.Println(result.Value().([][]int64))
	}
}
