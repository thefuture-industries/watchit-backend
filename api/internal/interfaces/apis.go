package interfaces

type IApis interface {
	// Из string массива жанров в массив цифр
	ArrayGenreIDS(str string) ([]int, error)
}
