package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func get_sum() {
	sum := 0
	a := [8][5]int{}
	b := [8][5]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			a[i][j] = rand.Intn(1000)
			b[i][j] = rand.Intn(1000)
			sum += a[i][j] + b[i][j]

		}
	}
	fmt.Printf("����������ͽ��Ϊ: %d", sum)
}

func season_if(m int) {
	if m >= 2 && m <= 4 {
		fmt.Printf("%d�µ�ǰ�����Ǵ���\n", m)
	} else if m >= 5 && m <= 7 {
		fmt.Printf("%d�µ�ǰ�������ļ�\n", m)
	} else if m >= 8 && m <= 10 {
		fmt.Printf("%d�µ�ǰ�������＾\n", m)
	} else {
		fmt.Printf("%d�µ�ǰ�����Ƕ���\n", m)
	}

}

func season_sw(m int) {
	switch {
	case m >= 2 && m <= 4:
		fmt.Printf("%d�µ�ǰ�����Ǵ���\n", m)
	case m >= 5 && m <= 7:
		fmt.Printf("%d�µ�ǰ�������ļ�\n", m)
	case m >= 8 && m <= 10:
		fmt.Printf("%d�µ�ǰ�������＾\n", m)
	default:
		fmt.Printf("%d�µ�ǰ�����Ƕ���\n", m)

	}
}

type student struct {
	name                 string
	yuwen, shuxue, waiyu int
}

func get_all_stu_avg(num int) {
	//����ѧ������
	stu := student{}
	s := make([]student, 0, num)
	sum_y := 0
	sum_w := 0
	sum_s := 0
	avg_60 := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		stu.yuwen = rand.Intn(100) + 1
		stu.shuxue = rand.Intn(100) + 1
		stu.waiyu = rand.Intn(100) + 1
		stu.name = strconv.Itoa(i)
		s = append(s, stu)
		avg := (s[i].yuwen + s[i].shuxue + s[i].waiyu) / 3
		fmt.Printf("ѧ��%2d �ɼ��� ���� %3d ��ѧ %3d ���� %3d  ƽ���֣�%3d \n", i, s[i].yuwen, s[i].shuxue, s[i].waiyu, avg)
		sum_y += s[i].yuwen
		sum_s += s[i].shuxue
		sum_w += s[i].waiyu
		if avg < 60 {
			avg_60 += 1
		}

	}
	fmt.Printf("���ĵ�ƽ������  %5d\n", sum_y/num)
	fmt.Printf("��ѧ��ƽ������  %5d\n", sum_s/num)
	fmt.Printf("�����ƽ������  %5d\n", sum_w/num)
	fmt.Printf("ȫ�����ƽ����60��ͬѧ��%d��\n", avg_60)

}

func main() {

	// 1.�����ʼ������8*5�ľ�������������ĺͣ���Ԫ����ӣ�
	get_sum()
	// 2.�����·ݣ��ж������ĸ����ڡ��ֱ���if��switchʵ��
	season_if(5)
	season_sw(12)
	// 3.����һ��student�ṹ�壬�������������������ſεĳɼ�����һ��slice����һ�����ͬѧ����ÿλͬѧ��ƽ���ֺ����������ſε�ƽ���֣�ȫ��ͬѧƽ���ֵ���60���м�λ
	get_all_stu_avg(50)
}