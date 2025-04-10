package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Group struct {
	size  int
	index int
}

type Room struct {
	computers int
	index     int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	groups := make([]Group, n)
	rooms := make([]Room, m)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Считываем группы
	for i := 0; i < n; i++ {
		scanner.Scan()
		var size int
		fmt.Sscanf(scanner.Text(), "%d", &size)
		groups[i] = Group{size: size + 1, index: i} // size + 1 для учителя
	}

	// Считываем аудитории
	for i := 0; i < m; i++ {
		scanner.Scan()
		var comps int
		fmt.Sscanf(scanner.Text(), "%d", &comps)
		rooms[i] = Room{computers: comps, index: i}
	}

	// Сортируем группы по размеру
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].size < groups[j].size
	})

	// Сортируем аудитории по количеству компьютеров
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i].computers < rooms[j].computers
	})

	// Ответ: для каждой группы - номер аудитории (или 0)
	answer := make([]int, n)
	roomIdx := 0

	for _, group := range groups {
		// Найти первую подходящую аудиторию
		for roomIdx < m && rooms[roomIdx].computers < group.size {
			roomIdx++
		}
		if roomIdx < m {
			// Назначаем аудиторию
			answer[group.index] = rooms[roomIdx].index + 1 // +1 потому что нумерация с 1
			roomIdx++
		} else {
			// Нет подходящей аудитории
			answer[group.index] = 0
		}
	}

	// Подсчитываем количество распределённых групп
	count := 0
	for _, v := range answer {
		if v != 0 {
			count++
		}
	}

	// Вывод
	fmt.Println(count)
	for i, v := range answer {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
