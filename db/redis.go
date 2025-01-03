package db

func (d *DB) RDB_Get(key string) (string, error) {

	return d.Rdc.Get(d.Ctx, key).Result()
}
func (d *DB) RDB_Pub(channel string, message interface{}) {
	err := d.Rdc.Publish(d.Ctx, channel, message).Err()
	if err != nil {
		panic(err)
	}
}
