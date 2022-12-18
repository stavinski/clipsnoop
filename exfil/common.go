package exfil

type ExfilInterface interface {
	Write(content string)
}
