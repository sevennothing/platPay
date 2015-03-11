/* test data type with golang*/

package main

import "fmt"

func main(){
	fmt.Println("golang data type test....");
	//bool
	var boollen bool;
	boollen = true;
	fmt.Println("bool : ",boollen);
	boollen = false;
	fmt.Println("bool : ",boollen);
	// uint8
	var uint8_x uint8;
	fmt.Println("uint8 : ",uint8_x);
	uint8_x = 0xff
	fmt.Println("uint8 : ",uint8_x);
	// uint16
	var uint16_x uint16;
	fmt.Println("uint16 : ",uint16_x);
	uint16_x = 0xffff
	fmt.Println("uint16 : ",uint16_x);
	// uint32
	var uint32_x uint32;
	fmt.Println("uint32 : ",uint32_x);
	uint32_x = 0xffffffff
	fmt.Println("uint32 : ",uint32_x);
	// int8
	var int8_x int8;
	fmt.Println("int8 : ",int8_x);
	int8_x = -100
	//int8_x = 0x85
	fmt.Println("int8 : ",int8_x);
	// float32
	var float32_x float32;
	fmt.Println("float32 : ",float32_x);
	float32_x = -100.123456789
	float32_x = -0.123456789
	fmt.Println("float32 : ",float32_x);
	// string
	var str string;
	fmt.Println("string : ",str);
	str = "string -0.123456789"
	fmt.Println("string : ",str);
	fmt.Printf("str is of Type %T\n", str);
	// auto type
	var str2 = "this a init sting";
	fmt.Println("str2 : ",str2);
	fmt.Printf("str2 is of Type %T\n", str2);
	var value = 10;
	fmt.Println("value : ",value);
	fmt.Printf("value is of Type %T\n", value);
	//========================================
	// const 
	const PI float32 = 3.1415926
	//const R	int8 = 10;
	const R	float32 = 10;
	var area float32;
	area = PI * R * R;
	fmt.Println("r=10, and ricle is : ",area);

}
