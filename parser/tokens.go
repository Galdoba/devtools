package parser

import "strings"

func Ident() ParserFunc {
	//return JOIN(LIT(Alpha), WHILE(OR(LIT(Alpha), LIT(Digit))))
	return JOIN(LIT(Alpha), WHILE(OR(LIT(Alpha), LIT(Digit))))
}

func IDENT_D() ParserFunc {
	//return JOIN(LIT(Alpha), WHILE(OR(LIT(Alpha), LIT(Digit))))
	return JOIN(LIT(Alpha), WHILE(OR(LIT(Alpha), LIT(Digit))))
}

func Alpha(chr byte) bool {
	return LowerAlpha(chr) || UpperAlpha(chr)
}

func LowerAlpha(chr byte) bool {
	if chr >= 'a' && chr <= 'z' {
		return true
	}

	return false
}

func UpperAlpha(chr byte) bool {
	if chr >= 'A' && chr <= 'Z' {
		return true
	}
	return false
}

func Digit(chr byte) bool {
	if chr >= '0' && chr <= '9' {
		return true
	}
	return false
}

func HashAlpha(chr byte) bool {
	if chr >= '0' && chr <= '9' && chr >= 'A' && chr <= 'Z' && chr != 'I' && chr != 'O' {
		return true
	}
	return false
}

const (
	lowerLetterEng = "abcdefghijklmnopqrstuvwxyz"
	upperLetterEng = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetterRus = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	upperLetterRus = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	digits         = "0123456789"
	bracketsOpen   = "([{"
	bracketsClose  = ")]}"
	signs          = "+-*/="
)

func lowerLetter(chr byte) bool {
	low_letter := lowerLetterEng + lowerLetterRus
	return strings.Contains(low_letter, string(chr))
}

/*
//примерсоздания своих алфавитов
func EHex(chr byte) bool {
	const alph = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	return strings.Contains(alph, string(chr))
}

func Space(chr byte) bool {
	const alph = " 	\r\n\t"
	return strings.Contains(alph, string(chr))
}

func UniversalProfile() ParserFunc {
	return JOIN(LIT(EHex), LIT(EHex), LIT(EHex), LIT(EHex), LIT(EHex), LIT(EHex), LIT(EHex), '-', LIT(EHex))
}
*/
/*
глиф = любое печатное что-то
идеограмма = обозначение идеи, например любой тип разделения между словаит (пробел, таб, перенос строки)
*/
