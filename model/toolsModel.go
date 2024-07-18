package model

type Tools struct {
	Tools map[string]string
}

func (t *Tools) ToolsChoice() []string {
	array := make([]string, 0)
	for k := range t.Tools {
		return append(array, k)
	}

	return nil
}
