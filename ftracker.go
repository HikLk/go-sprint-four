package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func Distance(action int) float64 {
	return float64(action) * lenStep / mInKm
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
func ShowTrainInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	// ваш код здесь
	switch {
	case trainingType == "Бег":
		distance := Distance(action)            // вызовите здесь необходимую функцию
		speed := MeanSpeed(action, duration)    // вызовите здесь необходимую функцию
		calories := R(action, weight, duration) // вызовите здесь необходимую функцию
		return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, trainingType, duration, distance, speed, calories)
	case trainingType == "Ходьба":
		distance := Distance(action)                               // вызовите здесь необходимую функцию
		speed := MeanSpeed(action, duration)                       // вызовите здесь необходимую функцию
		calories := WalkSpentCal(action, weight, duration, height) // вызовите здесь необходимую функцию
		return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, trainingType, duration, distance, speed, calories)
	case trainingType == "Плавание":
		distance := Distance(action)                                      // вызовите здесь необходимую функцию
		speed := SwimSpentCal(lengthPool, countPool, duration, weight)    // вызовите здесь необходимую функцию
		calories := SwimSpentCal(lengthPool, countPool, duration, weight) // вызовите здесь необходимую функцию
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
	runCalMeanSpeedMulti = 18   // множитель средней скорости.
	runCalMeanSpeedShift = 1.79 // среднее количество сжигаемых калорий при беге.
)

// R возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
func R(action int, weight, duration float64) float64 {
	// ваш код здесь
	return (runCalMeanSpeedMulti * MeanSpeed(action, duration) * runCalMeanSpeedShift) * weight / mInKm * duration * minInH
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkCalWeightMulti   = 0.035 // множитель массы тела.
	walkSpeedHeightMulti = 0.029 // множитель роста.
)

// WalkSpentCal возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkSpentCal(action int, duration, weight, height float64) float64 {
	// ваш код здесь
	return (walkCalWeightMulti*weight + (math.Pow(MeanSpeed(action, duration)*kmhInMsec, 2)/height*cmInM)*walkSpeedHeightMulti*weight) * duration * minInH
}

// Константы для расчета калорий, расходуемых при плавании.
const (
	swimCalMeanSpeedShift = 1.1 // среднее количество сжигаемых колорий при плавании относительно скорости.
	swimCalWeightMulti    = 2   // множитель веса при плавании.
)

// SwimMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func SwimMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
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
	// ваш код здесь
	return SwimMeanSpeed(lengthPool, countPool, duration) * swimCalWeightMulti * weight * duration
}
