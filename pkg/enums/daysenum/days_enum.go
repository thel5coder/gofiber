package daysenum

var (
	dayEnums = []map[string]interface{}{
		{
			"key":"monday",
			"text":"Monday",
		},
		{
			"key":"tuesday",
			"text":"Tuesday",
		},
		{
			"key":"wednesday",
			"text":"Wednesday",
		},
		{
			"key":"thursday",
			"text":"Thursday",
		},
		{
			"key":"friday",
			"text":"Friday",
		},
		{
			"key":"saturday",
			"text":"Saturday",
		},
		{
			"key":"sunday",
			"text":"Sunday",
		},
	}
)

func GetEnumFromKey(key string) map[string]interface{}{
	for _,dayEnum := range dayEnums {
		if dayEnum["key"] == key {
			return dayEnum
		}
	}
	return nil
}

func GetEnums() []map[string]interface{}{
	return dayEnums
}

func GetKeyFromKey(key string) string {
	for _,dayEnum := range dayEnums {
		if dayEnum["key"] == key {
			return dayEnum["key"].(string)
		}
	}
	return ""
}
