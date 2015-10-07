package jsonapi

type Links struct {
	Self    string `json:"self"`
	Related string `json:"related"`
}

type Data struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Relationship struct {
	Links Links `json:"links"`
	Data  Data  `json:"data"`
}
