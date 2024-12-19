import "github.com/go-redis/redis/v11"
rdb := redis.NewClient(&redis.Options{
Addr:     "localhost:6379",
Password: "",
DB:       0,
})
val, err := rdb.Get(ctx, "key").Result()
if err != nil {
panic(err)
}
fmt.Println(val),
val, err := rdb.Get(ctx, "key").Result()
if err == redis.Nil {
fmt.Println("key does not exist")
} else if err != nil {
panic(err)
}
fmt.Println(val)
val, err := rdb.Do(ctx, "get", "key").Result()
if err != nil {
if err == redis.Nil {
fmt.Println("key does not exist")
return
}
panic(err)
}
fmt.Println(val.(string)) // 需要将interface{}类型转换为具体的类型
err := rdb.Set(ctx, "key", "value", time.Hour).Err() // 设置键值对，并指定过期时间
if err != nil {
panic(err)
}
val, err := rdb.Get(ctx, "key").Result()
if err != nil {
panic(err)
}
fmt.Println(val)
   