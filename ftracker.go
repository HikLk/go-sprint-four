// Package ftracker реализует функции для подсчета затраченных калориев для разных видов активностей.
package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	LenStep   = 0.65  // средняя длина шага.
	MInKm     = 1000  // количество метров в километре.
	MinInH    = 60    // количество минут в часе.
	KmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	CmInM     = 100   // количество сантиметров в метре.
)

// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func Distance(action int) float64 {
	return float64(action) * LenStep / MInKm
}

// MeanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
func MeanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	distance := Distance(action)
	return distance / duration
}

// ShowTrainInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
func ShowTrainInfo(
	action int,
	trainingType string,
	duration, weight, height float64,
	lengthPool, countPool int) string {
	switch {
	case trainingType == "Бег":
		distance := Distance(action)
		speed := MeanSpeed(action, duration)
		calories := RunSpentCal(action, weight, duration)
		return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, trainingType, duration, distance, speed, calories)
	case trainingType == "Ходьба":
		distance := Distance(action)
		speed := MeanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height)
		return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, trainingType, duration, distance, speed, calories)
	case trainingType == "Плавание":
		distance := Distance(action)
		speed := SwimmingMeanSpeed(lengthPool, countPool, duration)
		calories := SwimSpentCal(lengthPool, countPool, duration, weight)
		return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	RunningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
	RunningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)

// RunSpentCal возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
func RunSpentCal(action int, weight, duration float64) float64 {
	scaledSpeed := RunningCaloriesMeanSpeedMultiplier * MeanSpeed(action, duration) * RunningCaloriesMeanSpeedShift
	return scaledSpeed * weight * duration * MinInH / MInKm
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	WalkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	WalkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	walkSpeed := math.Pow(MeanSpeed(action, duration)*KmhInMsec, 2)
	heightMsec := height / CmInM
	return (WalkingCaloriesWeightMultiplier*weight + (walkSpeed/heightMsec)*WalkingSpeedHeightMultiplier*weight) * duration * MinInH
}

// Константы для расчета калорий, расходуемых при плавании.
const (
	SwimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых колорий при плавании относительно скорости.
	SwimmingCaloriesWeightMultiplier = 2   // множитель веса при плавании.
)

// SwimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func SwimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / MInKm / duration
}

// SwimSpentCal возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
func SwimSpentCal(lengthPool, countPool int, duration, weight float64) float64 {
	shiftedSpeed := SwimmingMeanSpeed(lengthPool, countPool, duration) + SwimmingCaloriesMeanSpeedShift
	return shiftedSpeed * SwimmingCaloriesWeightMultiplier * weight * duration
}
