package main

func main() {

	traceSlice := func(s[] int) {
		println("traceSlice trace starts")
		for _, v := range s {
			println(v);
		}
		println("traceSlice trace ends")
	}

	array := []int{1, 2, 3, 4, 5}
	slice := array[1:3]

	traceSlice(array)
	traceSlice(slice)

	array[1] = 6

	traceSlice(array)
	traceSlice(slice)
}
