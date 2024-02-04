package BuildCategory

type BuildCategory struct {
	CodeId string
	Base64 string
}

func ResponseBuildCategoryImage(codeId string, base64 string) BuildCategory {
	return BuildCategory{
		CodeId: codeId,
		Base64: base64,
	}
}
