package main

// В задании не указано могут ли быть дубликаты:
// 1) в исходных слайсах
// 2) в результирующем слайсе
//
// Предполагаю что:
// 1) исходные слайсы могут быть какие угодно, кроме nil
// 2) в результирующем слайсе не должно быть дубликатов
// также, предполагаю, что порядок элементов в слайсах не важен

func Intersection(a []int, b []int) (bool, []int) {
	if len(b) < len(a) {
		return Intersection(b, a)
	} 
	// len A always <= then len B
	elementsA := make(map[int]bool, len(a))
	elementsAandB := make(map[int]bool, len(a))
	
	for _, value := range a {
		elementsA[value] = true
	}
	
	for _, value := range b {
		if elementsA[value] {
			elementsAandB[value] = true
		}
	}
	
	intersection := make([]int, 0, len(elementsAandB))
	for element := range elementsAandB {
		intersection = append(intersection, element)
	}

	return (len(intersection)>0), intersection
}