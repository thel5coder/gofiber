package subjectcolorenum

var (
	subjectColorEnums = []string{"2A8AF3", "FFD251", "9E52D9", "EE2D4F", "70C49C", "F69107", "00C3B0", "D18ACE", "8EACCC", "C7614B", "3EC9D2",
		"C0D24E", "E01ED8", "D69377", "B8BF87", "00C667", "C1B174", "B695CF", "D97F8F", "C79F11"}
)

func GetColorEnumByIndex(index int) string {
	count := len(subjectColorEnums)
	if index == count-1 {
		return subjectColorEnums[0]
	}

	return subjectColorEnums[index]
}

func GetColorIndex(color string) int {
	for index, subjectColorEnum := range subjectColorEnums {
		if color == subjectColorEnum {
			return index
		}
	}

	return 0
}
