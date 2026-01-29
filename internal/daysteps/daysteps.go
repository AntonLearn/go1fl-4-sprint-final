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

var (
	ErrNoParse    = errors.New("ошибка парсинга входящей строки")
	ErrNoInt      = errors.New("ошибка преобразования количества шагов в число")
	ErrNoMoreZero = errors.New("ошибка меньше или равно нулю")
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	input := strings.Split(data, ",")
	if len(input) != 2 {
		return 0, 0, ErrNoParse
	}
	steps, err := strconv.Atoi(input[0])
	if err != nil {
		return 0, 0, ErrNoInt
	}
	if steps <= 0 {
		return 0, 0, ErrNoMoreZero
	}
	duration, err := time.ParseDuration(input[1])
	if err != nil {
		return 0, 0, ErrNoParse
	}
	if duration <= 0 {
		return 0, 0, ErrNoMoreZero
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	distance := stepLength * float64(steps) / mInKm
	ccal := 0.0
	ccal, err = spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, ccal)
}
