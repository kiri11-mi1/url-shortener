package shorten

const ALPHABET = "qwertyuiopasdfghjkzxcvbnmQWERTYUPASDFGHJKLZXCVBNM123456789"

var ALPHABET_LEN = uint32(len(ALPHABET))

func Shorten(id uint32) string {
	var nums []uint32
	var result string
	for id > 0 {
		nums = append(nums, id%ALPHABET_LEN)
		id /= ALPHABET_LEN
	}
	for _, index := range getReverseArray(nums) {
		result += string(ALPHABET[index])
	}
	return result
}

func getReverseArray(array []uint32) []uint32 {
	var newArray []uint32
	for i := len(array) - 1; i > -1; i-- {
		newArray = append(newArray, array[i])
	}
	return newArray
}
