package positiontypeenum

type positionTypeMapEnums map[string]interface{}
var (
	positionTypeEnums = []positionTypeMapEnums{
		map[string]interface{}{
			"key":"teacher",
			"text":"Teacher",
		},
		map[string]interface{}{
			"key":"staff",
			"text":"Staff",
		},
		map[string]interface{}{
			"key":"homeroom-teacher",
			"text":"Homeroom teacher",
		},
	}
)

func GetKey(index int) string{
	value := positionTypeEnums[index]

	return value["key"].(string)
}

func GetText(index int) string{
	value := positionTypeEnums[index]

	return value["text"].(string)
}

func GetEnums() []positionTypeMapEnums{
	return positionTypeEnums
}

func GetEnum(index int) positionTypeMapEnums{
	return positionTypeEnums[index]
}
