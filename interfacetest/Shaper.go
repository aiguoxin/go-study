package interfacetest

type Shaper interface {
	Area() float32
	Ok() float32
}

type Square struct {
	//对外，属性都需要大写才行
	Side float32
}

/**
实现接口
不带*的参数相当于值传递，带*的相当于引用传递
 */
func (sq *Square)Area() float32 {
	return sq.Side * sq.Side
}

//值传递，调用后，不会改变sq中Side值，和java类似
func (sq Square)Ok()  float32 {
	sq.Side = sq.Side+1
	return sq.Side
}

