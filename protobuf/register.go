package protobuf

func init() {

	_list = make(map[string]modelCreator)

	_list["parser_manhua_page"] = func() Model {
		return new(ParserManhuaPage)
	}

	_list["parser_manhua_listing"] = func() Model {
		return new(ParserManhuaListing)
	}

	_list["parser_article_listing"] = func() Model {
		return new(ParserArticleListing)
	}

	_list["parser_article_page"] = func() Model {
		return new(ParserArticlePage)
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

	_list["saver_fs"] = func() Model {
		return new(SaverFs)
	}
	_list["saver_article"] = func() Model {
		return new(SaverArticle)
	}

	_list["checker_fs"] = func() Model {
		return new(CheckerFs)
	}

	_list["checker_manhua_chapter"] = func() Model {
		return new(CheckerManhuaChapter)
	}

	_list["indexer_manhua_chapter"] = func() Model {
		return new(IndexManhuaChapter)
	}

	_list["indexer_search_item"] = func() Model {
		return new(IndexSearchItem)
	}

	_list["cacher_image"] = func() Model {
		return new(CacherImage)
	}

}
