package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных: %q", data)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("не удалось определить количество шагов: %w", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("не удалось определить продолжительность прогулки: %w", err)
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность прогулки должна быть больше нуля")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
