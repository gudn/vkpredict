// Интерфейс препоцессоров
//
// Формат позволяет использовать большинство стандратных функций, к примеру
// strings.ToLower или strings.TrimSpace
package preprocessing

type Preprocessor func(input string) string
