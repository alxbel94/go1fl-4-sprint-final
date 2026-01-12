package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	refactoredData := strings.Split(data, ",")
	if len(refactoredData) == 2 {
		steps, err := strconv.Atoi(refactoredData[0])
		if err != nil || steps <= 0 {
			if err == nil {
				err = errors.New("неверные шаги")
			}
			fmt.Println("Ошибка:", err)
			log.Println(err)
			return 0, 0, err
		}
		duration, err := time.ParseDuration(refactoredData[1])
		if err != nil || duration <= 0 {
			if err == nil {
				err = errors.New("неверная продолжительность")
			}
			fmt.Println("Ошибка:", err)
			log.Println(err)
			return 0, 0, err
		}

		return steps, duration, nil
	}
	err := errors.New("количество элементов в слайсе не равно 2")
	fmt.Println("Ошибка:", err)
	log.Println(err)
	return 0, 0, err
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil || steps <= 0 {
		fmt.Println("Ошибка при парсинге данных:", err)
		return ""
	}
	distance := (float64(steps) * stepLength) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps,
		distance,
		calories,
	)

}
