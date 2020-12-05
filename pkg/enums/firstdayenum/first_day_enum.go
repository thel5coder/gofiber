package firstdayenum

var (
	firstDayEnum = []map[string]interface{}{
		{
			"key":"monday",
			"text":"Monday",
		},
		{
			"key":"sunday",
			"text":"Sunday",
		},
	}
)

func GetEnumFromKey(key string) map[string]interface{}{
	for _,dayEnum := range firstDayEnum {
		if dayEnum["key"] == key {
			return dayEnum
		}
	}
	return nil
}

func GetEnums() []map[string]interface{}{
	return firstDayEnum
}

func GetKeyFromKey(key string) string {
	for _,dayEnum := range firstDayEnum {
		if dayEnum["key"] == key {
			return dayEnum["key"].(string)
		}
	}
	return ""
}