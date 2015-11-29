package main

import ("fmt"
"math"
"strings"
"sort"
"os"
"log"
"io/ioutil"
"strconv"
"net/http"
"time")

func main(){
	
	fmt.Println("Hello World") //Hello World

	var age int = 40

	var favNum float64 = 1.6180339

	fmt.Println(age, favNum) //40 1.6180339

	var numOne = 1.000
	var num99 = 0.9999
	fmt.Println(numOne - num99) //9.999999999998899e-05

	fmt.Println("6 + 4 =", 6+4)	//6 + 4 = 10
	fmt.Println("6 - 4 =", 6-4) //6 - 4 = 2
	fmt.Println("6 * 4 =", 6*4) //6 * 4 = 24
	fmt.Println("6 / 4 =", 6/4) //6 / 4 = 1
	fmt.Println("6 % 4 =", 6%4) //6 % 4 = 2

	var myName string = "Derek Banas"
	fmt.Println(len(myName)) //11
	fmt.Println(myName + " is a robot") //Derek Banas is a robot

	fmt.Println("\n")
	fmt.Println("My name is Anjana Rajagopal \n I like tacos very much \n \n") //My name is Anjana Rajagopal
																			   //I like tacos very much

	const pi float64 = 3.14159265
	fmt.Printf("%.3f \n", pi) //3.142 - %.3f returns the float and rounds it to 3 decimal places
	fmt.Printf("%T \n", pi) //float64 - %T returns the type of the value

	var isOver40 bool = true
	fmt.Printf("%t \n", isOver40) //true - %t returns the word true or false

	fmt.Printf("%d \n", 100) //100 - %d returns base 10 (integer itself)
	fmt.Printf("%b \n", 100) //1100100 - %b returns base 2 (binary)
	fmt.Printf("%c \n", 44)  //, - %c  returns the character represented by the corresponding Unicode code
	fmt.Printf("%x \n", 17)  //11 - %x returns base 16 with lower-case letters for a-f
	fmt.Printf("%e \n", pi)  //3.141593e+00 - %e returns number in scientific notation

	fmt.Println("true && false = ", true && false) // true && false = false
	fmt.Println("true || false = ", true || false) // true || false = true
	fmt.Println("!true = ", !true) //!true = false
	fmt.Println("!false = ", !false) //!false = true
	fmt.Println("\n")

	i := 1
	for i <= 10 {
		fmt.Println(i) //Prints nums 1-10
		i++
		//1
		//2
		//3
		//4
		//5
		//6
		//7
		//8
		//9
		//10
	}
	fmt.Println("\n")

	for j := 0; j < 5; j++ { //Prints nums 0-4
		fmt.Println(j)
		//0
		//1
		//2
		//3
		//4
	}
	fmt.Println("\n")

	yourAge := 18

	if yourAge >= 16 { //You Can Drive
		fmt.Println("You Can Drive")
	} else {
		fmt.Println("You Can't Drive")
	}

	if yourAge >= 16 { //You Can Drive
		fmt.Println("You Can Drive")
	} else if yourAge >= 18 {
		fmt.Println("You Can Vote")
	} else {
		fmt.Println("You Can Have Fun")
	}

	yourAge_2 := 5
	switch yourAge_2 {
		case 16 : fmt.Println("Go Drive")
		case 18 : fmt.Println("Go Vote")
		default : fmt.Println("Go Have Fun")
	}
	fmt.Println("\n")

	var favNums2[5] float64
	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618

	fmt.Println(favNums2[3]) //3.141

	favNums3 := [5]float64 {1, 2, 3, 4, 5}

	for i, value := range favNums3 {
		fmt.Println(value, i)
		//1 0
		//2 1
		//3 2
		//4 3
		//5 4
	}

	fmt.Println("\n")

	numSlice := []int {5, 4, 3, 2, 1} //numSlice = {5, 4, 3, 2, 1}
	numSlice2 := numSlice[3:5] //numSlice2 = {2, 1}
	fmt.Println("numSlice2[0] =", numSlice2[0]) //numSlice2[0] = 2
	fmt.Println("numSlice2[1] =", numSlice2[1]) //numSlice2[1] = 1

	numSlice3 := make([]int, 5, 10) //has room for 10 elements
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3[0]) //5
	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3[6]) //-1

	fmt.Println("\n")
	presAge := make(map[string] int)
	presAge["TheodoreRoosevelt"] = 42
	fmt.Println(presAge["TheodoreRoosevelt"]) //42
	fmt.Println(len(presAge)) //1
	presAge["John F. Kennedy"] = 43
	fmt.Println(len(presAge)) //2
	delete(presAge, "John F. Kennedy")
	fmt.Println(len(presAge)) // 1

	fmt.Println("\n")

	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum :", addThemUp(listNums)) //Sum : 15

	fmt.Println("\n")
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2) //6 7

	fmt.Println("\n")
	fmt.Println(subtractThem(1, 2, 3, 4, 5)) //-15

	fmt.Println("\n")

	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}

	fmt.Println(doubleNum()) //6
	fmt.Println(doubleNum()) //12

	fmt.Println("\n")
	fmt.Println(factorial(3)) //6

	fmt.Println("\n")
	defer printTwo() // 2
	printOne() // 1

	fmt.Println("\n")
	fmt.Println(safeDiv(3, 0)) // 0
	fmt.Println(safeDiv(3, 2)) //1

	fmt.Println("\n")
	demPanic()

	fmt.Println("\n")
	x := 0
	changeXVal(x)
	//x = 2
	//Memory Address for x = 0x1026804o...
	fmt.Println("x =", x)
	fmt.Println("Memory Address for x =", &x)

	fmt.Println("\n")
	yPtr := new(int)
	changeYValNow(yPtr)

	fmt.Println("y =", *yPtr) //y = 100

	fmt.Println("\n")
	rect1 := Rectangle{0, 50, 10, 10}
	fmt.Println("Rectangle is", rect1.width, "wide") //Rectangle is 10 wide
	fmt.Println("Area of rectangle =", rect1.area()) //Area of rectangle = 100

	fmt.Println("\n")
	rect := Rectangle{20, 50}
	circ := Circle{4}

	fmt.Println("Rectangle Area =", getArea(rect)) //Rectangle Area = 1000
	fmt.Println("Circle Area =", getArea(circ)) //Circle Area = 50.26548245743669

	fmt.Println("\n")
	sampString := "Hello World"
	fmt.Prinln(string.Contains(sampString, "lo")) //true
	fmt.Prinln(string.Index(sampString, "lo")) //3
	fmt.Prinln(string.Count(sampString, "l")) //3
	fmt.Prinln(string.Replace(sampString, "l", "x", 3)) //Hexxo Worxd

	fmt.Println("\n")
	csvString := "1, 2, 3, 4, 5, 6"
	fmt.Println(strings.Split(csvString, ",")) //[1 2 3 4 5 6]
	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters: ", listOfLetters) //Letters: [a b c]
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")
	fmt.Println(listOfNums) //3, 2, 1

	fmt.Println("\n")
	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is some random text")
	file.Close()
	stream, er := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString) //This is some random text

	fmt.Println("\n")
	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	fmt.Println(float64(randInt)) //5
	fmt.Println(int(randFloat)) //10

	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt) //100

	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat) //250.5

	fmt.Println("\n")

	http.HandleFunc("/", handler)

	http.HandleFunc("/earth", handler2)

	http.ListenAndServe(":8080", nil)


	//0 : 0
	//1 : 0
	//2 : 0
	//3 : 0
	//4 : 0
	//5 : 0
	//6 : 0
	//7 : 0
	//8 : 0
	//9 : 0
	//0 : 1
	//1 : 1
	//2 : 1
	//3 : 1
	//4 : 1
	//5 : 1
	//6 : 1
	//7 : 1
	//8 : 1
	//9 : 1
	//...
	//8 : 9
	//9 : 9
	//1 : 9
	for i := 0; i < 10; i++ {
		go count(i)
	}
	time.Sleep(time.Millisecond * 11000)
}

func count(id int){
	for i := 0; i < 10; i++{
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}

type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}

func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}

func changeYValNow(yPtr *int){
	*yPtr = 100
}

func safeDiv(num1, num2 int) int {
	defer func(){
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}

func printOne(){
	fmt.Println(1)
}

func printTwo(){
	fmt.Println(2)
}

func addThemUp(numbers [] float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func next2Values(number int) (int, int) {
	return number+1, number+2
}

func subtractThem(args ... int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}


//factorial(3)
//3 * factorial(2) == 3 * 2 = 6
//2 * factorial(1) == 2 * 1 = 2
//factorial(0) == 1
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)
}

func demPanic(){
	defer func(){
		fmt.Println(recover())
	}()

	panic("PANIC")
}

func changeXVal(x *int){
	*x = 2
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

//OUTPUT:
//Hello World
//40 1.6180339
//9.999999999998899e-05
//6 + 4 = 10
//6 - 4 = 2
//6 * 4 = 24
//6 / 4 = 1
//6 % 4 = 2
//11
//Derek Banas is a robot
//
//My name is Anjana Rajagopal
//I like tacos very much
//
//3.142
//float64
//true
//100
//1100100
//,
//11
//3.141593e+00
//true && false = false
//true || false = true
//!true = false
//!false = true
//
//1
//2
//3
//4
//5
//6
//7
//8
//9
//10
//
//0
//1
//2
//3
//4
//
//You Can Drive
//You Can Drive
//Go Have Fun
//
//3.141
//1 0
//2 1
//3 2
//4 3
//5 4
//
//numSlice2[0] = 2
//numSlice2[1] = 1
//5
//-1
//
//42
//1
//2
//1
//
//Sum : 15
//
//6 7
//
//-15
//
//6
//12
//
//6
//1
//2
//
//0
//1
//
//x = 2
//Memory Address for x = 0x1026804o...
//
//y = 100
//
//Rectangle is 10 wide
//Area of rectangle = 100
//
//Rectangle Area = 1000
//
//true
//3
//3
//Hexxo Worxd
//
//[1 2 3 4 5 6]
//Letters: [a b c]
//3, 2, 1
//
//This is some random text
//
//5
//10
//100
//250.5