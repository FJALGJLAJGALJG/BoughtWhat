package main

import "fmt"

func main() {

	boughtWhat()
}

func boughtWhat() {

	things := []string{"帽子", "发夹", "裙子", "手套"}
	// fmt.Println(len(things))
	for i := 0; i < len(things); i++ {
		for j := 0; j < len(things); j++ {
			if i != j {
				for k := 0; k < len(things); k++ {
					if (i != k) && (j != k) {
						if (boolType(things[j] != "手套", things[k] != "发夹") == 1) &&
							(boolType(things[i] != "发夹", things[k] != "裙子") == 1) &&
							(boolType(things[i] != "帽子", things[k] == "裙子") == 1) {
							fmt.Print("【小丽买了" + things[i] + ",")
							fmt.Print("小玲买了" + things[j] + ",")
							fmt.Println("小娟买了" + things[k]+"】")
						}
					}
				}
			}
		}
	}
}

func boolType(falg1 bool, falg2 bool) int {

	if falg1 == true && falg2 == true {
		return 2
	} else if falg1 == false && falg2 == false {
		return 0
	} else {
		return 1
	}
}

/**
小丽、小玲、小娟三个人一起去商场里买东西。
她们都买了各自需要的东西，有帽子，发夹，裙子，手套等，而且每个人买的东西还不同。
有一个人问她们三个都买了什么，小丽说：“小玲买的不是手套，小娟买的不是发夹。”
小玲说：“小丽买的不是发夹，小娟买的不是裙子。”小娟说：“小丽买的不是帽子，小娟买的是裙子。”
她们三个人，每个人说的话都是有一半是真的，一半是假的。那么，她们分别买了什么东西？
*/
