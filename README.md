# In-memory cache

Реализовано 4 метода для in-memory cache:
- New() - создание нового кеш объекта
- Set() - запись значения value в кеш по ключу key
- Get() - получение значения value из кеша по ключу key
- Delete() - удаление по ключу key

Пример использования

    func main() {
    	cache := cache.New()
    
    	cache.Set("userId", 42)
    	userId, _ := cache.Get("userId")
    
    	fmt.Println(userId)
    
    	cache.Delete("userId")
    	userId, _ = cache.Get("userId")
    
    	fmt.Println(userId)
    }
