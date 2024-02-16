package util

func CalcOffset(pageIndex, pageSize int) int {
	return (pageIndex - 1) * pageSize
}
