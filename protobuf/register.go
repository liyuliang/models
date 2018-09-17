package protobuf

import "models/protobuf/task"

func init() {
	_list = make(map[string]modelCreator)

	_list["manhua"] = func() Model {
		return new(Manhua)
	}

	_list["saver_manhua_site"] = func() Model {
		return new(task.SaverManhuaSite)
	}

	_list["saver_manhua_chapter"] = func() Model {
		return new(task.SaverManhuaChapter)
	}

	_list["saver_manhua_image"] = func() Model {
		return new(task.SaverManhuaImage)
	}

	_list["indexer_manhua_chapter"] = func() Model {
		return new(task.IndexManhuaChapter)
	}
}
