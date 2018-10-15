package protobuf

func init() {
	_list = make(map[string]modelCreator)

	_list["parser_manhua_page"] = func() Model {
		return new(ParserManhuaPage)
	}

	_list["parser_manhua_listing"] = func() Model {
		return new(ParserManhuaListing)
	}

	_list["saver_manhua_book"] = func() Model {
		return new(SaverManhuaBook)
	}

	_list["saver_manhua_chapter"] = func() Model {
		return new(SaverManhuaChapter)
	}

	_list["saver_manhua_image"] = func() Model {
		return new(SaverManhuaImage)
	}

	_list["indexer_manhua_chapter"] = func() Model {
		return new(IndexManhuaChapter)
	}
}
