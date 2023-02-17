# cache

`Set(key string, value interface{})` - запись значения `value` в кеш по ключу `key` </br>
`Get(key string)`
`Delete(key)`


    func main() {
        cache := cache.New()

        cache.Set("userId", 42)
        userId := cache.Get("userId")

        fmt.Println(userId)

        cache.Delete("userId")
        userId := cache.Get("userId")

        fmt.Println(userId)
    }