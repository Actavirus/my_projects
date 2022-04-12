package dates

const DaysInWeek int = 7 //объявляем константу
func WeeksToDays(weeks int) int { //получаем количество недель
	return weeks * DaysInWeek // умножаем на количество дней в неделе, чтобы вычислить общее количество дней
}
func DaysToWeeks(days int) float64 { // получаем количество дней
	return float64(days) / float64(DaysInWeek) //делим на количество дней в неделе, чтобы получить количество недель
}
