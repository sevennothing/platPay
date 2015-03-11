/* test condition with golang*/

package main

import "fmt"

func main(){
	fmt.Println("golang conditon test....");
	//if..else

	// for
	var i int8 = 1;
	for i=80; i<100; i++ {
		fmt.Println(i);
		if( i > 90){
			fmt.Println(i," biger than 90")
			break
		}else{
			fmt.Println(i," letter than 90");
			continue
		}
	}

	var ret int = select_test(10,100);
	fmt.Println(ret);

}

func select_test(initValue, endValue int) int {
	var result int
	var i int = 0;
	select {
	case i = initValue:
		fmt.Println("initValue = ",i);
		break;
	case i > endValue:
		fmt.Println("endValue = ",i);
		break;
	default:
		i++;
	}
	return result
}
