package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func PrintSparceTable(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	//Вводим данные

	var N, K int

	fmt.Scan(&N)
	//Массив чисел
	num_arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&num_arr[i])
	}

	fmt.Scan(&K)
	//Массив отрезков
	sec_arr := make([][2]int, K)
	for i := 0; i < K; i++ {
		fmt.Scan(&sec_arr[i][0], &sec_arr[i][1])
	}

	var q, pow int
	q = 1
	for pow = 1; pow <= N; pow *= 2 {
		q++
		//Ищем высоту твблицы
	}
	q -= 1

	//делаем разреженную таблицу
	sparce_table := make([][]int, q) //+1 т.к. нумерация с 1 а не 0
	fmt.Println("Старт расширение таблицы", q)
	//заполнение таблицы
	for i := 0; i < q; i++ {
		for j := 0; j < N+1; j++ {
			sparce_table[i] = append(sparce_table[i], -1)
		}
	}
	fmt.Println("Конец расширение таблицы")
	fmt.Println(sparce_table)

	//конфигурация: sparce_table[глубина][ширина]
	//[0, 0, 0, 0, 0]
	//[0, 0, 0, 0, 0]
	//[0, 0, 0, 0, 0]
	//[0, 0, 0, 0, 0]

	fmt.Println("Старт заполнение исходниками")

	for i := 1; i < N+1; i++ {
		sparce_table[0][i] = num_arr[i-1]
	}
	fmt.Println("Конец заполнение исходниками")
	fmt.Println(sparce_table)

	fmt.Println("Старт полное заполнение")

	for deep, dive := 0, 2; deep < q-1; deep++ {
		for i := 1; i <= N; i++ {
			maxi := -1
			pred := i
			for k := 1; k <= dive && i+dive <= N+1; k++ {
				if sparce_table[deep][i+dive-1] == -1 {
					maxi = -1
					break
				}
				maxi = max(maxi, sparce_table[deep][pred])
				pred++
			}
			sparce_table[deep+1][i] = maxi
		}
		dive *= 2
	}

	fmt.Println("Конец полное заполнение")

	PrintSparceTable(sparce_table)

}
