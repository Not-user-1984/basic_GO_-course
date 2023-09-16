package slice
 import "fmt"

var weekTempArr = [7]int{1,2 ,3 ,4 , 5 ,6 ,7}
var  numbers = []int{1, 2, 3, 4, 5}

var  p = &numbers
func DemoSlice() {
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4] 
	weekTempSlice := weekTempArr[:]
	weekTempSlice = append(weekTempSlice, 7)


	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
    fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2 
    fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6 
    fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7 
}

func InsreaseSlice() {
	for i := 0 ; i < 3 ; i++ {
		*p = append(*p,i+1)
	}
	fmt.Println(numbers) 
}

func Append_and_reverce() {
	dim := 100
	s := make([]int, 0 , dim)

	for i := 0; i < dim ; i++ {
		s = append(s, i+1)
	}
	
	s = append(s[:10], s[dim-10:]...)
	dim =len(s)
	for i := range s[:dim/2] {
		s[i], s[dim-i-1] = s[dim-i-1] ,s[i]
	}
	fmt.Println(s)
}
// Здесь dim предполагается, что это длина слайса s.

// s[:dim/2] создает новый слайс, содержащий первую половину элементов слайса s.
//  Из-за деления на 2 мы берем только элементы от начала и до середины слайса (не включая середину).
// for i := range ... используется для итерации по слайсу s[:dim/2].
//  В этой строке i будет принимать значения от 0 до dim/2 - 1,
//   которые являются индексами элементов первой половины слайса.
// Затем внутри цикла мы обмениваем элементы слайса местами, чтобы выполнить его обратный переворот:

// s[i], s[dim-i-1] = s[dim-i-1], s[i] - это короткий синтаксис обмена значений двух переменных в Go.
//  Мы меняем местами элемент s[i] из первой половины слайса с элементом s[dim-i-1] из второй половины слайса.
//   Таким образом, происходит переворот первой половины слайса s.
// После выполнения этого кода, первая половина слайса s будет развернута в обратном порядке.
// Обратите внимание, что этот код меняет оригинальный слайс s.
//  Если вы хотите сохранить оригинальный слайс и получить развернутую копию,
//   вам нужно сначала создать новый слайс и затем скопировать значения в обратном порядке.


func  String_reverce() {
	input := "строка есть строка "
	n := 0

	runes := make([]rune,len(input))

	for _, r := range input {
		runes[n] = r
		n++
	}

	runes = runes[0:n]

	for i := 0; i < n/2 ; i++ {
		runes[i],runes[n-1-i] = runes[n-1-i], runes[i]
	}
	output := string(runes)
	fmt.Println(output)
}
// input := "строка есть строка ": Исходная строка, которую мы хотим развернуть.
// runes := make([]rune, len(input)): Создается слайс рун runes,
// в котором будут храниться символы из исходной строки input.

// for _, r := range input { ... }:
// Итерируемся по исходной строке input с помощью range,
// где r представляет собой текущий символ (руну) в исходной строке.
// runes[n] = r: Каждый символ r добавляется в слайс рун runes в позицию n.
// n++: Переменная n используется для отслеживания количества символов,
//  добавленных в слайс runes.

// runes = runes[0:n]:
// Слайс рун runes обрезается до размера n,
// чтобы удалить ненужные элементы,
// которые могли быть выделены на дополнительных итерациях цикла for.

// for i := 0; i < n/2; i++ { ... }:
//  Цикл выполняет обратный переворот слайса runes,
//   чтобы получить обратный порядок символов.

// runes[i], runes[n-1-i] = runes[n-1-i], runes[i]:
// Здесь происходит обмен значениями между элементами слайса runes.
// Таким образом, слайс runes переворачивается.

// output := string(runes): Создается новая строка output из развернутого слайса рун runes.

// fmt.Println(output): Результат выводится в консоль.







