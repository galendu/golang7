package main

import (
	"errors"
	"fmt"
)

// 1.ʵ��һ���������������ɸ�float64���ò�������������������Щ�����˻��ĵ���������Ϊ0ʱ����error
func q1_1(fa ...float64) (float64,error) {

  if len(fa)<1{
    return 0, errors.New("����Ϊ0")
  }
  tmpData:=1.
  for _,value:=range fa{
    if value == 0{
      return 0, errors.New("��������0")
    }
    tmpData=tmpData*value
  }
  ret := 1 / tmpData
  return ret,nil

}
func q1_2_1(cal ...float64) (float64){
  l:=len(cal)
  if l == 1{
    return cal[0]
  }
  return cal[l-1]*q1_2_1(cal[:l-1]...)


}

func q1_2(fa ...float64) (float64,error){
  if len(fa)<1{
    return -1,nil
  }
  ret:=q1_2_1(fa...)
  if ret == 0{
    return 0,errors.New("the result is 0 or �����0")
  }
  
  return 1/ret,nil

}

// 3.���������ӿڣ���������ж���ٶ���һ���ṹ�壺���ܣ�ͬʱʵ�����������ӿ�
type (
  Fish interface{
    youyong() string
}
  Reptile interface{
    move() string
  }
  
)
type Qingwa struct{
  name string

}
func (qw Qingwa) Youyong() string{
  return fmt.Sprintf("%s is youyong\n",qw.name)
}
func (qw Qingwa) Move() string { 
  return fmt.Sprintf("%s is moving\n",qw.name)

}
// 4.ʵ�ֺ���func square(num interface{}) interface{}������һ��interface{}��ƽ����interface{}������4�����ͣ�float32��float64��int��byte 
type Cal interface{
  Area(num interface{}) interface{}

}
type Square struct{
  X interface{}
}
func (s Square) Area(num interface{}) interface{}{
   v:=num
  switch v.(type){
  case int:
    v:=num.(int)
    return v * v
  case float32:
    v:=num.(float32)
    return v * v    
  case float64:
    v:=num.(float64)
    return v * v   
  case byte:
    v:=num.(byte)
    return v * v   
  default:
    return errors.New("�������Ͳ��Ϸ�")
}
}
 




func main(){
  // 1.ʵ��һ���������������ɸ�float64���ò�������������������Щ�����˻��ĵ���������Ϊ0ʱ����error   
  qrr:=[]float64{0.20,0.0221,0.1228}
  q_re,err1 :=q1_1(qrr...)
  if err1 == nil {
   fmt.Printf("Q1 result is %10.10f\n",q_re) 
  }else {
    fmt.Printf("ERROR����������0")
  }

  // 2.�����õݹ�ʵ��
  q_re2,err2:=q1_2(qrr...)
  if err2 == nil {
     fmt.Printf("Q2 result is %10.10f\n",q_re2) 
  }

  // 3.���������ӿڣ���������ж���ٶ���һ���ṹ�壺���ܣ�ͬʱʵ�����������ӿ�
   fish :=Qingwa{name: "xiaoyu"}
   fmt.Printf("Q3 result is:%s",fish.Move())      
   fmt.Printf("Q3 result is:%s",fish.Youyong())
   
  // 4.ʵ�ֺ���func square(num interface{}) interface{}������һ��interface{}��ƽ����interface{}������4�����ͣ�float32��float64��int��by 
  a:=Square{X:5.5}
  fmt.Println("Q4������Ϊ:",a.Area(a.X))


}