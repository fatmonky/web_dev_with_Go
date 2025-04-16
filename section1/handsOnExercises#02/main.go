package main

import "fmt"

// ex1
// func main(){
// 	s := []int{3,4,5}
// 	fmt.Println(s)
// 	for i, _ := range s{
// 		fmt.Println(i)
// 	}
// 	for i, v := range s{
// 		fmt.Println(i, v)
// 	}
// }

// ex2
// func main() {
// 	m := map[string]int{
// 		"George": 88,
// 		"PJ":     44,
// 		"Qihui":  42,
// 	}
// 	fmt.Println(m)
// 	for k, _ := range m {
// 		fmt.Println(k)
// 	}
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}
// }

// ex3-6
// type person struct {
// 	fName   string
// 	lName   string
// 	favFood []string
// }

// func (p person) walk() string {
// 	return fmt.Sprintf("%v is walking.\n", p.fName)
// }

// func main() {
// 	p1 := person{"George", "Washington", []string{"peas", "carrots", "steak"}}
// 	fmt.Println(p1)                //ex4
// 	fmt.Println(p1.fName)          //ex4
// 	fmt.Println(p1.favFood)        //ex5
// 	for _, v := range p1.favFood { //ex5
// 		fmt.Println(v)
// 	}
// 	s := p1.walk()
// 	fmt.Println(s)

// }

// ex7-10
// type vehicle struct {
// 	doors int
// 	color string
// }

// type truck struct {
// 	vehicle
// 	fourWheel bool
// }

// type sedan struct {
// 	vehicle
// 	luxury bool
// }

// type transportation interface {
// 	transportationDevice() string
// }

// func report(t transportation) {
// 	fmt.Println(t.transportationDevice())
// }

// func (s sedan) transportationDevice() string {
// 	return fmt.Sprintf("This %v sedan with %v doors goes vroom!\n", s.color, s.doors)
// }

// func (t truck) transportationDevice() string {
// 	return fmt.Sprintf("This %v truck with %v doors goes vroom!\n", t.color, t.doors)
// }

// func main() {
// 	t1 := truck{
// 		vehicle{
// 			4,
// 			"red",
// 		},
// 		true,
// 	}

// 	s1 := sedan{
// 		vehicle{
// 			4,
// 			"blue",
// 		},
// 		false,
// 	}
// 	fmt.Println(t1)
// 	fmt.Println(s1)
// 	fmt.Println(t1.vehicle.color)
// 	fmt.Println(s1.luxury)
// 	report(s1)
// 	report(t1)
// }

//ex11-16

// type gator int
// type flamingo bool
// type swampCreature interface {
// 	greeting()
// }

// var g1 gator

// func (g gator) greeting() {
// 	fmt.Println("Hello, I am a gator.")
// }

// func (f flamingo) greeting() {
// 	fmt.Println("Hello, I am pink and beautiful and wonderful.")
// }

// func bayou(s swampCreature) {
// 	s.greeting()
// }

// func main() {
// 	g1 = 3
// 	fmt.Println(g1)
// 	fmt.Printf("The type of g1 is %T\n", g1)
// 	var x int
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)
// 	x = int(g1)
// 	fmt.Println(x)
// 	g1.greeting()
// 	f1 := flamingo(true)
// 	bayou(g1)
// 	bayou(f1)
// }

// ex17
func main() {
	s := "I'm sorry dave I can't do that"
	fmt.Println(s)
	t := []byte(s)
	fmt.Println(t)
	sb := string(t)
	fmt.Println(sb)
	fmt.Println(sb[:14])
	fmt.Println(sb[10:22])
	fmt.Println(sb[17:])
	for _, v := range s {
		fmt.Println(string(v))
	}
}
