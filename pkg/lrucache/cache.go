package lrucache

// LRUCache interface
type LRUCache interface {

	// Add Добавляет новое значение с ключом в кеш (с наивысшим приоритетом), возвращает true, если все прошло успешно
	// В случае дублирования ключа вернуть false
	// В случае превышения размера - вытесняется наименее приоритетный элемент
	Add(key, value string) bool

	// Get Возвращает значение под ключом и флаг его наличия в кеше
	// В случае наличия в кеше элемента повышает его приоритет
	Get(key string) (value string, ok bool)

	// Remove Удаляет элемент из кеша, в случае успеха возврашает true, в случае отсутствия элемента - false
	Remove(key string) (ok bool)
}
