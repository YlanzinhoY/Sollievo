package model

type Tools struct {
	Tools map[string]string
}

func (t *Tools) ToolsChoice() ([]string, error) {
	array := make([]string, 0, 10)
	for k := range t.Tools {
		array = append(array, k)
	}

	return array, nil
}
