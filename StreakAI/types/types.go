package types

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Response struct {
	Solutions [][]int `json:"solutions"`
}
