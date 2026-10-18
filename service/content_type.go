package service

type ContentType string

const (
	ContentText ContentType = "text/plain"
	ContentJson ContentType = "application/json"
)

func (ct ContentType) isValid() bool {
	switch ct {
	case ContentJson, ContentText:
		return true
	}
	return false
}
