package jsonstruct

//WordMeaning : word and its meaning in single JSON
type WordMeaning struct {
	Word         string
	WordID       string
	Meaning      string
	SimilarWords []string
}
