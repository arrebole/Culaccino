package sql

// Exists 查询多个
func Exists(query ...string) int64 {
	switch len(query) {
	case 1:
		return existStorage(query...)
	case 2:
		return existRepo(query...)
	case 3:
		return existChapter(query...)
	default:
		return 0
	}
}

func existStorage(query ...string) int64 {
	return StorageDB.Exists(query[0]).Val()
}

func existRepo(query ...string) int64 {
	return RepoDB.Exists(query[0] + "/" + query[1]).Val()
}
func existChapter(query ...string) int64 {
	return ChapterDB.Exists(query[0] + "/" + query[1] + "/" + query[2]).Val()
}
