package floorAccess

import (
	"fmt"
	"strings"
)

func GetEmployeeID() {
	fmt.Println("Поднесите скуд")
	var id string
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Ошибка ID:", err.Error())
		return
	}
	floor := 0
	if strings.HasPrefix(id, "1") {
		fmt.Print("Вы имейте доступ к 1 этажу\nВыберите этаж: ")
		_, err = fmt.Scan(&floor)
		switch floor {
		case 1:
			fmt.Println(" Добро пожаловать на 1-й этаж!")
		default:
			fmt.Println("Нет доступа!")
		}

	} else if strings.HasPrefix(id, "2") {
		fmt.Print("Вы имейте доступ к 1 и 2 этажу\nВыберите этаж: ")
		_, err = fmt.Scan(&floor)
		switch floor {
		case 1:
			fmt.Println("Добро пожаловать на 1-й этаж!")
		case 2:
			fmt.Println("Добро пожаловать на 2-й этаж!")
		default:
			fmt.Println("Нет доступа!")
		}

	} else if strings.HasPrefix(id, "3") {
		fmt.Print("Вы имейте доступ к 1,2 и 3 этажу\nВыберите этаж: ")
		_, err = fmt.Scan(&floor)
		switch floor {
		case 1:
			fmt.Println("Добро пожаловать на 1-й этаж!")
		case 2:
			fmt.Println("Добро пожаловать на 2-й этаж!")
		case 3:
			fmt.Println("Добро пожаловать на 3-й этаж!")
		default:
			fmt.Println("Нет доступа!")
		}

	} else if strings.HasPrefix(id, "4") {
		fmt.Print("Вы имейте доступ к 1,2,3 и 4 этажу\nВыберите этаж: ")
		_, err = fmt.Scan(&floor)
		switch floor {
		case 1:
			fmt.Println("Добро пожаловать на 1-й этаж!")
		case 2:
			fmt.Println("Добро пожаловать на 2-й этаж!")
		case 3:
			fmt.Println("Добро пожаловать на 3-й этаж!")
		case 4:
			fmt.Println("Добро пожаловать на 4-й этаж!")
		default:
			fmt.Println("Нет доступа!")
		}

	} else {
		fmt.Println("Неверный ID")
		return
	}

}
