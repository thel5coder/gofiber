package religionenum

type religionMapEnums map[string]interface{}
var (
	religionEnums = []religionMapEnums{
		map[string]interface{}{
			"key":"islam",
			"text":"Islam",
		},
		map[string]interface{}{
			"key":"protestan",
			"text":"Protestan",
		},
		map[string]interface{}{
			"key":"katolik",
			"text":"Katolik",
		},
		map[string]interface{}{
			"key":"Hindu",
			"text":"Hindu",
		},
		map[string]interface{}{
			"key":"buddha",
			"text":"Buddha",
		},
		map[string]interface{}{
			"key":"khonghucu",
			"text":"Khonghucu",
		},
	}
)

func GetKey(index int) string{
	value := religionEnums[index]

	return value["key"].(string)
}

func GetText(index int) string{
	value := religionEnums[index]

	return value["text"].(string)
}

func GetEnums() []religionMapEnums{
	return religionEnums
}

func GetEnum(index int) religionMapEnums{
	return religionEnums[index]
}