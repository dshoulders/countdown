package utils

func ChunkArray(chunkCount int, arr []string) [][]string {
	var divided [][]string

	chunkSize := (len(arr) + chunkCount - 1) / chunkCount

	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize

		if end > len(arr) {
			end = len(arr)
		}

		divided = append(divided, arr[i:end])
	}

	return divided
}

func GetLongest(arr []string) string {
	var longest string

	for _, word := range arr {
		var stringLength = len(word)
		if stringLength > len(longest) {
			longest = word
		}
	}

	return longest
}
