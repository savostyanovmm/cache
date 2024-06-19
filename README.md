# In-memory cache

Реализовано 4 метода для in-memory cache:
- New() - создание нового кеш объекта
- Set() - запись значения value в кеш по ключу key
- Get() - получение значения value из кеша по ключу key
- Delete() - удаление по ключу key

Пример использования

    // Создаем контейнер с временем жизни по-умолчанию равным 5 минут и удалением просроченного кеша каждые 10 минут
    cache := memorycache.New(5 * time.Minute, 10 * time.Minute)

    // Установить кеш с ключем "myKey" и временем жизни 5 минут
    cache.Set("myKey", "My value", 5 * time.Minute)

    // Получить кеш с ключом "myKey"
    i := cache.Get("myKey")
