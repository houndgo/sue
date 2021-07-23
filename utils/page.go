package utils


func PageDefaultMap(page int ,size int ,mp map[string]interface{}) map[string]interface{} {
	if mp["page_index"] == nil {
		mp["page_index"] = page
	}

	if mp["page_size"] == nil {
		mp["page_size"] = size
	}
	return mp
}