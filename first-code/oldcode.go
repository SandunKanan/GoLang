package main

import "fmt"

// /* Declare a single variable */
// var a int

// /* Declare a group of variables */
// var (
// 	b bool
// 	c float32
// 	d string
// )

// func main() {
// 	a=42
// 	b, c = true, 32.0
// 	d="string"
// 	fmt.Println(a, b, c, d)
// }



func main() {
	// a:=42					// Initialize and assign to a single variable
	// b, c := true, 32.0		// Initialize and assign to multiple variables
	// d:="string"
	// fmt.Println(a, b, c, d)

		/* User specified types */
		const a int32 = 12			// 32-bit integer
		const b float32 = 20.5		// 32-bit float
		var c complex128 = 1 + 4i	// 128-bit complex number
		var d uint16 = 14			// 16-bit unsigned integer
	
		n := 42						// int	
		pi := 3.14					// float64
		x, y := true, false			// bool
		z := "Go is awesome"		// string
		
		// print out types of the variables declared above
		fmt.Printf("user-specified types: \n%T %T %T %T \n", a, b, c, d)
		fmt.Printf("default types: \n%T %T %T %T %T\n", n, pi, x, y, z)

		/* Define an array of 4 strings - size of an array is fixed */
	var DeploymentOptions = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	// looping through i
	for i := 0; i < len(DeploymentOptions); i++ {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}

	// loop through array
	for index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}
	x2:=[]float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average(x2))

	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go":2012,
	}

	/* Loop through each entry and output the name and release year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}

	fmt.Printf("C was released in %d\n", firstReleases["C"])

	count := 1
	for count < 5 {
		count += count
	}
	fmt.Println(count) // 8
}



func average(x []float64) (avg float64) {
	total:=0.0
	if len(x) == 0 {
		avg=0
	} else {
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}
	return
}