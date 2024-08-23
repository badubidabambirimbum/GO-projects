package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

const Pi = 3.14 // константа

type vehicle struct {
	class string
}

type car struct { // Структура (с маленькой буквы - не экспортируется)
	name     string // Если бы структура была в другом файле, то к значениям с маленькой буквы можно было бы обратиться только через set и get функции
	topSpeed int
	vehicle
}

type Number interface { // Интерфейс (с большой буквы - экспортируется)
	up2()
}

type myInt int       // Свой тип
type myFloat float64 // Свой тип

func (m *myInt) up2() { // Метод для определенного типа (myInt)
	*m *= 2
}

func (m *myFloat) up2() { // myFloat
	*m *= 2.0
}

// Так как в методах используется указатель, то и используемыми типами будут указатели
func myPrint(num Number) {
	num.up2()
	intMy, ok := num.(*myInt) // утверждение типа
	if ok {
		fmt.Println("int :", *intMy)
	}
	floatMy, ok := num.(*myFloat) // утверждение типа
	if ok {
		fmt.Println("float :", *floatMy)
	}
}

var f func(int, int) int = func(x int, y int) int { return x * y } // Анонимная функция

func main() {
	defer fmt.Println("Program has been finished") // Выполняется после окончания программы

	var valInt int = 5                      // Целое число
	var valFloat float64 = 5.0              // Число с плавающей запятой
	var valString string = "555 - 55 = 500" // Строка
	var valBool bool = true                 // Логический тип
	var sInt *int = &valInt                 // Ссылка
	valR := 'Ц'                             // Руна (для многобайтовых символов)
	var valB byte = 100                     // Байт (для однобайтовых символов)

	fmt.Printf("valInt: %d, valFloat: %0.2f, valString: %s, valBool: %t, *sInt: %v, valR: %v, valB: %v\n", valInt, valFloat, valString, valBool, *sInt, valR, valB)          // Printf форматирует и сразу выводит
	newS := fmt.Sprintf("valInt: %d, valFloat: %0.2f, valString: %s, valBool: %t, *sInt: %v, valR: %v, valB: %v\n", valInt, valFloat, valString, valBool, *sInt, valR, valB) // Sprintf форматирует и возвращает значение!
	fmt.Println(newS)
	fmt.Println(utf8.RuneCountInString(valString)) // Длина строки

	fmt.Println("-----------------------------------------------")
	fmt.Printf("%s : ", valString)
	for i := 0; i < len(valString); i++ { // valString[i] - байт
		fmt.Printf("%s", string(valString[i])) // печатает байты. Если в строке будут символы, которые весят больше 1 байта, то будут проблемы. В таком случае лучше использовать руны
	}

	fmt.Println("\n-----------------------------------------------")
	Rstring := "байты"
	fmt.Printf("%s : ", Rstring)
	for i := 0; i < len(Rstring); i++ {
		fmt.Printf("%s", string(Rstring[i])) // Русские буквы печатаются плохо, так как они весят 2 байта
	}
	fmt.Println("\n-----------------------------------------------")

	for _, char := range Rstring { // char - руна
		fmt.Printf("%s", string(char))
	}

	fmt.Println("\n-----------------------------------------------")

	fmt.Println(valInt % 4) // Остаток от деления

	var arr = [3]int{1, 2, 3} // Массив
	var seg []int             // Сегмент (можно make не использовать)
	// seg := make([]int, 3) // Сегмент
	seg = append(seg, 3, 2, 1)
	seg = append(seg, []int{6, 4, 5}...)
	fmt.Println("arr:", arr)
	fmt.Println("seg:", seg)
	sort.Slice(seg, func(i, j int) bool { // Сортировка
		return seg[i] < seg[j]
	})
	fmt.Println("seg sort:", seg)

	newMap := map[string]int{"Alex": 22, "Ivan": 17} // Карта (make обязателен, если пустая карта создается. Пример: var newMap map[string]int. Правильно newMap := make(map[string]int))
	fmt.Println("\nmap:", newMap)

	myStruct := car{name: "Corvette", topSpeed: 337} // Структура
	myStruct.class = "Car"                           // обращаемся напрямую к переменной из другой структуры
	fmt.Printf("\nstruct: %#v\n\n", myStruct)

	myVal1 := myInt(4) // Собственный тип
	myVal1.up2()       // Вызов метода

	myVal2 := myFloat(4.0)
	myVal2.up2()

	myPrint(&myVal1) //  Из-за того, что метод up2 получает указатель, мы должны здесь тоже передавать указатель
	myPrint(&myVal2)

	fmt.Printf("\nАнонимная функция: %d\n\n", f(5, 4)) // Анонимная функция

	value := 5
	binary := fmt.Sprintf("%b", value) // преобразование в двоичное число
	n := len(binary)
	binary2 := strings.Repeat("1", n)                                 // Создание строки "1" n повторов
	value1, _ := strconv.ParseInt(binary, 2, 64)                      // преобразование двоичного числа в int64
	value2, _ := strconv.ParseInt(binary2, 2, 64)                     // преобразование двоичного числа в int64
	fmt.Printf("XOR : %d ^ %d = %d\n", value1, value2, value1^value2) // XOR

	f := 5.6
	fmt.Printf("из float %0.1f в int %d\n", f, int(f)) // Преобразование float в int (обязательно надо просвоить float64 переменной)

	sss := "-1.5"
	ints, _ := strconv.Atoi(sss) // из строки в int
	fmt.Println(ints)

	floats, _ := strconv.ParseFloat(sss, 64) // из строки в float64
	fmt.Println(floats)

	panic("panic") // прерывает программу
}
