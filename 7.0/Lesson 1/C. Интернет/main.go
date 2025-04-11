package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// структура карты
type Card struct {
	Minutes int
	Cost    int
	Quality float64
}

// сортировка по убыванию качества карты
func BubbleSort(arr []Card) []Card {
	for j := len(arr) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i].Quality < arr[i+1].Quality {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	//свободное время в интернете
	m, _ := strconv.Atoi(scanner.Text())

	//Массив со структурами карт
	CardShop := make([]Card, 30)

	//загружаем значения минут
	for i := 0; i < 30; i++ {
		scanner.Scan()
		CardShop[i].Minutes, _ = strconv.Atoi(scanner.Text())
	}

	//загружаем значения цен
	for i := 0; i < 30; i++ {
		CardShop[i].Cost = int(math.Pow(2, float64(i)))
	}

	//загружаем качество карт мин/руб
	for i := 0; i < 30; i++ {
		CardShop[i].Quality = float64(CardShop[i].Minutes) / float64(CardShop[i].Cost)
	}

	SortedCardShop := BubbleSort(CardShop)

	SpendedMoney, Total, Time := 0, int(math.Pow(2, 31)), 0

	//идем по маскимально выгодным карточкам и пробуем их покупать,
	//если удаётся из максимально выгодных карт собрать нужное время,
	//то заканчиваем подбор и выводим результат.
	//Если удаётся собрать больше минут по цене ниже, чем была
	//до этого, то мы это делаем и изеняем итоговую цену
	for i := 0; i < 30; i++ {

		for Time+SortedCardShop[i].Minutes < m {
			Time += SortedCardShop[i].Minutes
			SpendedMoney += SortedCardShop[i].Cost
		}
		if Time+SortedCardShop[i].Minutes == m {
			SpendedMoney += SortedCardShop[i].Cost
			if Total > SpendedMoney {
				Total = SpendedMoney
			}
			break
		} else {
			if Total > SpendedMoney+SortedCardShop[i].Cost {
				Total = SpendedMoney + SortedCardShop[i].Cost
			}
		}
	}

	fmt.Println(Total)

}
