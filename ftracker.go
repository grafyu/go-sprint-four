package ftracker

import (
    "fmt"
    "math"
)

// Основные константы, необходимые для расчетов.
const (
    lenStep   = 0.65  // средняя длина шага в метрах.
    mInKm     = 1000  // количество метров в километре.
    minInH    = 60    // количество минут в часе.
    kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
    cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию (км), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
    return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки. (км/ч)
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки (часы).
func meanSpeed(action int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    distance := distance(action)
    return distance / duration
}

// WalkingSpentCalories возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки (часы).
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
    calories := 0.0
    distance := distance(action)
    speed := meanSpeed(action, duration)

    switch {
	case trainingType == "Бег":
		calories = RunningSpentCalories(action, weight, duration)
	case trainingType == "Ходьба":
		calories = WalkingSpentCalories(action, duration, weight, height)
	case trainingType == "Плавание":
		calories = SwimmingSpentCalories(lengthPool, countPool, duration, weight)
	default:
		return "неизвестный тип тренировки"
	}

    templString := "Тип тренировки: %s\nДлительность: %.2f ч.\n" + 
                    "Дистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"
    return fmt.Sprintf(templString, trainingType, duration, distance, speed, calories)
}

// WalkingSpentCalories возвращает количе\ство потраченных колорий при беге.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки (часы).
func RunningSpentCalories(action int, weight, duration float64) float64 {
    // Константы для расчета калорий, расходуемых при беге.
    const (
        speedMult  = 18   // множитель средней скорости.
        speedShift = 1.79 // среднее количество сжигаемых калорий при беге.
    )

    meanSpeed := meanSpeed(action, duration)
    return speedMult * meanSpeed * speedShift * weight / mInKm * duration * minInH
}

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки (часы).
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
    // Константы для расчета калорий, расходуемых при ходьбе.
    const (
        weightMult = 0.035 // множитель массы тела.
        heightMult = 0.029 // множитель роста.
    )

    meanSpeed := meanSpeed(action, duration)*kmhInMsec
    duration = duration * minInH
    meanSpeedSqr := math.Pow(meanSpeed, 2)
    return (weightMult + heightMult * (meanSpeedSqr / height)) * weight * duration
}

// swimMeanSpeed возвращает среднюю скорость при плавании (км/ч).
//
// Параметры:
//
// lengthPool int — длина бассейна (метры).
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки (часы).
func swimMeanSpeed(lengthPool, countPool int, duration float64) float64 {

    if duration == 0 {
        return 0
    }
    return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна (метры).
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки (часы).
// weight float64 — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
    // Константы для расчета калорий, расходуемых при плавании.
    const (
        speedShift = 1.1  // среднее количество сжигаемых колорий при плавании относительно скорости.
        weightMult = 2    // множитель веса при плавании.
    )

    meanSpeed := swimMeanSpeed(lengthPool, countPool, duration)
    return (meanSpeed + speedShift) * weightMult * weight * duration
}
