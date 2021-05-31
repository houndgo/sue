package page

func CreatePageQuery (maps map[string]interface{}) {
	if _, pageOk := maps["page_size"];!pageOk{
		maps["page_size"]= 10
	}
	if _, pageIndexOk := maps["page_index"];!pageIndexOk{
		maps["page_index"]= 1
	}
}