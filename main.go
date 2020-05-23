package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"math/rand"
	"time"
)

const listsize = 9

type GameData struct {
	status bool
	count int
}

type List [listsize]int

func main(){
	list := List{1,2,3,4,5,6,7,8,0}
	index := 8
	var data GameData
	var actArr []int

	clear()
	list.broken()
	for !(data.status) {

		clear()
		list.show()
		actArr = readAction()
		index = list.change(actArr, index)

		if ( list[8] == 0 ) {
			data.status = list.check()
		}
	}
	clear()
	list.show()
	fmt.Printf("Congratulation! Good job!!!\n")
}

func (l *List)check() bool {
	for i := range l {
		if (( i < 8 )  && ( l[i] != i+1 )) {
			return false
		}
	}

	return true
}

func (l *List)broken(){
	rand.Seed(time.Now().UnixNano())
	index := 0
	for i := 0; i < 10; i++ {
		index = rand.Intn(6)
		l[index], l[index+1], l[index+2] = l[index+2], l[index], l[index+1]
	}		
}

func (l *List)change(actArr []int, index int) int {

	for _, v := range actArr {
		switch v {
			case 97  : // action A
				index = l.replace(index, index+1)
			case 100 : // action D
				index = l.replace(index, index-1)
			case 115 : // action S
				index = l.replace(index, index-3)
			case 119 : // action W				
				index = l.replace(index, index+3)
		}		
	}
	return index
}

func readAction() []int {

	var actArr []int
	i := 0
	action := bufio.NewReader(os.Stdin)

	for {
		b, err := action.ReadByte()

		if ((err != nil) || (i > 9)) {
			break
		}

		switch b {
			case 97, 100, 115, 119 :
				actArr = append(actArr, int(b))
				i++
			case 10 :
				return actArr
		}
	}
	return actArr
}

func (l *List)replace(ia, ib int) int {
	testa := ((ia >=0 ) && (ia < 9))
	testb := ((ib >=0 ) && (ib < 9))
	testab := ( ia != ib)

	if testa && testb && testab {
		switch ia {
			case 2, 3 :
				if (( ib == 3 ) || ( ib == 2 )){
					return ia		
				}
			case 5, 6 :
				if (( ib == 6 ) || ( ib == 5 )){
					return ia		
				}
		}

		l[ia], l[ib] = l[ib], l[ia]
		return ib
	}
	return ia
}

func (l List)show(){
	fmt.Println("========================================")
	fmt.Printf("\t%v\t%v\t%v\n\n",l[0], l[1], l[2])
	fmt.Printf("\t%v\t%v\t%v\n\n",l[3], l[4], l[5])
	fmt.Printf("\t%v\t%v\t%v\n",  l[6], l[7], l[8])
	fmt.Println("========================================")			
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)
}

func clear() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()	
}

//abdroid test 

