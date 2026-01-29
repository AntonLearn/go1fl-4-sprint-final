package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	ErrNoParse        = errors.New("ошибка парсинга входящей строки")
	ErrNoInt          = errors.New("ошибка преобразования количества шагов в число")
	ErrNoMoreZero     = errors.New("ошибка меньше или равно нулю")
	ErrNoTypeTraining = errors.New("ошибка не указан тип тренировки")
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	input := strings.Split(data, ",")
	if len(input) != 3 {
		return 0, "", 0, ErrNoParse
	}
	steps, err := strconv.Atoi(input[0])
	if err != nil {
		return 0, "", 0, ErrNoInt
	}
	if steps <= 0 {
		return 0, "", 0, ErrNoMoreZero
	}
	typeTraining := input[1]
	if typeTraining == "" {
		return 0, "", 0, ErrNoTypeTraining
	}
	duration, err := time.ParseDuration(input[2])
	if err != nil {
		return 0, "", 0, ErrNoParse
	}
	if duration <= 0 {
		return 0, "", 0, ErrNoMoreZero
	}
	return steps, typeTraining, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLenght := height * stepLengthCoefficient
	distanceInM := float64(steps) * stepLenght
	return distanceInM / float64(mInKm)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceInKm := distance(steps, height)
	durationInHour := duration.Hours()
	return distanceInKm / durationInHour
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeTraining, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	ccal := 0.0
	switch typeTraining {
	case "Ходьба":
		ccal, err = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		ccal, err = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		log.Println(err)
		return "", err
	}
	distance := distance(steps, height)
	meanSpeed := meanSpeed(steps, height, duration)
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeTraining, float64(duration.Hours()), distance, meanSpeed, ccal), err
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0.0, ErrNoMoreZero
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	ccalRunning, err := RunningSpentCalories(steps, weight, height, duration)
	return ccalRunning * walkingCaloriesCoefficient, err
}
