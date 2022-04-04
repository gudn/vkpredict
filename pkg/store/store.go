// Интерфейс хранилища
package store

import "errors"

var (
	ErrStoreIsNil = errors.New("store: store is nil")
)

type ID string

// Пустой ID
var None ID

type Adder interface {
	// Атомарно добавить элементы для заданный ID. Если какой-то ID является None,
	// то необходимо сгенерировать уникальный. Вызывающая сторона отвечает за то,
	// что длины этих двух слайсов равны
	Add(iids []ID, items []string) ([]ID, error)
}

type Store interface {
	Adder
	// Атомарно получить элементы по заданным ID. В случае отсутствия документа
	// просто исплючить его из мапы, ошибку выкидывать только при внутренних
	// ошибках реализации
	Get(ids []ID) (map[ID]string, error)
}
