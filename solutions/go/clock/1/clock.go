package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
    hour int
    minute int
}

func New(h, m int) Clock {
    H := (h + m / 60) % 24
    for H < 0 {
        H += 24
    }
    M := m % 60
    for M < 0 {
        M += 60
        H -= 1
    }
    for H < 0 {
        H += 24
    }
    return Clock{hour: H, minute: M}
    
}

func (c Clock) Add(m int) Clock {
	total_minutes := c.minute + m;
    hour := (c.hour + total_minutes / 60) % 24;
    for hour < 0 {
        hour += 24;
    }
    minute := total_minutes % 60;
    for minute < 0 {
        minute += 60;
        hour -= 1;
    }
    for hour < 0 {
        hour += 24;
    }
    return Clock{hour: hour, minute: minute}
}

func (c Clock) Subtract(m int) Clock {
	totalMinutes := c.minute - m
    hour := (c.hour + totalMinutes/60) % 24
    // Обработка отрицательных значений
    for hour < 0 {
        hour += 24
    }
    minute := totalMinutes % 60
    // Если минуты отрицательные, то корректируем
    if minute < 0 {
        minute += 60
        hour -= 1
    }
    // Обработка отрицательных значений часов
    for hour < 0 {
        hour += 24
    }
    return Clock{hour: hour, minute: minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
