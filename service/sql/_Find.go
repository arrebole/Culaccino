package sql

// Find 查询多个
func Find(out interface{}, query ...string) bool {
	switch out.(type){
	
	case *module.Storage:
		return findStorage(out.(*module.Storage),query...)
	
	case: *module.Repo:
		return findRepo(out.(*module.Repo),query...)
	
	case *module.Chapter:
		return findChapter(out.(*module.Chapter),query...)
		
	default:
		return false
	}
}

func findStorage(arg *module.Storage,query ...string)bool{
	return false
}

func findRepo(arg *module.Repo,query ...string)bool{
	return false
}
func findChapter(arg *module.Chapter,query ...string)bool{
	return false
}