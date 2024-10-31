package usecase

import (
	"FirstHttp/entity"
	"FirstHttp/pkg/sortfuncs"
	"strconv"
	"strings"
)

func QuickSortUsecase(unsortedString string) ([]int, error) {
	var unsortedIntArray []int
	unsortedStrings := strings.Split(unsortedString, " ")
	for _, s := range unsortedStrings {
		num, err := strconv.Atoi(s)
		if err == nil {
			unsortedIntArray = append(unsortedIntArray, num)
		}
	}
	sortedArray := sortfuncs.MyQuickSort(unsortedIntArray)
	return sortedArray, nil
}

func GenerateUserPass(login string) (entity.User, error) {
	u := entity.User{Login: login, PassHash: "pass_hash", Secret: "123"}
	return u, nil
}
