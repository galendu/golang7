package main //把int32.go归属到mypath包中
import ("fmt" //表示I/O输入输出的功能包
        "math" //表示基本的数学运算包
		"strings" //表示字符串的基本操作包
		)//引入包 fmt math strings
//输出一个int32对应的二进制表示
func BinaryFormat(n int32) string{
	a := uint32(n) //将传进来的int32类型数值转换为无符号的uint32类型数值后赋值给a
	sb := strings.Builder{} //拼接字符串赋值给sb
	c := uint32(math.Pow(2,31)) //求2的31次幂 转换为无符号的uint32类型数值后赋值给c c=10000000000000000000000000000000
	for i := 0; i<32; i++{
		if a & c > 0 {
			sb.WriteString("1")
		}else{
			sb.WriteString("0")
		}
		c >>= 1 //c向右移动一位后赋值给c
	}
	return sb.String()

}

func main(){
	fmt.Println(BinaryFormat(233))
	fmt.Println(BinaryFormat(-233))
	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(111111))
}