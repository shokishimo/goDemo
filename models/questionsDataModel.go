package questionsDataModel

type Data struct {
	Question string
	Answer   string
}

func (d *Data) setData(q string, a string) {
	d.Question = q
	d.Answer = a
}
