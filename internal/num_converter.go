package internal

var romanInput = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

func ConvertRomanToArab(r string) int {
	var arabicNumber int
	switch r {
	case "I":
		arabicNumber = 1
	case "II":
		arabicNumber = 2
	case "III":
		arabicNumber = 3
	case "IV":
		arabicNumber = 4
	case "V":
		arabicNumber = 5
	case "VI":
		arabicNumber = 6
	case "VII":
		arabicNumber = 7
	case "VIII":
		arabicNumber = 8
	case "IX":
		arabicNumber = 9
	case "X":
		arabicNumber = 10
	}
	return arabicNumber
}

func ConvertArabToRoman(num int) string {
	var romanNumber string
	switch num {
	case 1:
		romanNumber = "I"
	case 2:
		romanNumber = "II"
	case 3:
		romanNumber = "III"
	case 4:
		romanNumber = "IV"
	case 5:
		romanNumber = "V"
	case 6:
		romanNumber = "VI"
	case 7:
		romanNumber = "VII"
	case 8:
		romanNumber = "VIII"
	case 9:
		romanNumber = "IX"
	case 10:
		romanNumber = "X"
	case 11:
		romanNumber = "XI"
	case 12:
		romanNumber = "XII"
	case 13:
		romanNumber = "XIII"
	case 14:
		romanNumber = "XIV"
	case 15:
		romanNumber = "XV"
	case 16:
		romanNumber = "XVI"
	case 17:
		romanNumber = "XVII"
	case 18:
		romanNumber = "XVIII"
	case 19:
		romanNumber = "XIX"
	case 20:
		romanNumber = "XX"
	case 21:
		romanNumber = "XXI"
	case 22:
		romanNumber = "XXII"
	case 23:
		romanNumber = "XXIII"
	case 24:
		romanNumber = "XXIV"
	case 25:
		romanNumber = "XXV"
	case 26:
		romanNumber = "XXVI"
	case 27:
		romanNumber = "XXVII"
	case 28:
		romanNumber = "XXVIII"
	case 29:
		romanNumber = "XXIX"
	case 30:
		romanNumber = "XXX"
	case 31:
		romanNumber = "XXXI"
	case 32:
		romanNumber = "XXXII"
	case 33:
		romanNumber = "XXXIII"
	case 34:
		romanNumber = "XXXIV"
	case 35:
		romanNumber = "XXXV"
	case 36:
		romanNumber = "XXXVI"
	case 37:
		romanNumber = "XXXVII"
	case 38:
		romanNumber = "XXXVIII"
	case 39:
		romanNumber = "XXXIX"
	case 40:
		romanNumber = "XL"
	case 41:
		romanNumber = "XLI"
	case 42:
		romanNumber = "XLII"
	case 43:
		romanNumber = "XLIII"
	case 44:
		romanNumber = "XLIV"
	case 45:
		romanNumber = "XLV"
	case 46:
		romanNumber = "XLVI"
	case 47:
		romanNumber = "XLVII"
	case 48:
		romanNumber = "XLVIII"
	case 49:
		romanNumber = "XLIX"
	case 50:
		romanNumber = "L"
	case 51:
		romanNumber = "LI"
	case 52:
		romanNumber = "LII"
	case 53:
		romanNumber = "LIII"
	case 54:
		romanNumber = "LIV"
	case 55:
		romanNumber = "LV"
	case 56:
		romanNumber = "LVI"
	case 57:
		romanNumber = "LVII"
	case 58:
		romanNumber = "LVIII"
	case 59:
		romanNumber = "LIX"
	case 60:
		romanNumber = "LX"
	case 61:
		romanNumber = "LXI"
	case 62:
		romanNumber = "LXII"
	case 63:
		romanNumber = "LXIII"
	case 64:
		romanNumber = "LXIV"
	case 65:
		romanNumber = "LXV"
	case 66:
		romanNumber = "LXVI"
	case 67:
		romanNumber = "LXVII"
	case 68:
		romanNumber = "LXVIII"
	case 69:
		romanNumber = "LXIX"
	case 70:
		romanNumber = "LXX"
	case 71:
		romanNumber = "LXXI"
	case 72:
		romanNumber = "LXXII"
	case 73:
		romanNumber = "LXXIII"
	case 74:
		romanNumber = "LXXIV"
	case 75:
		romanNumber = "LXXV"
	case 76:
		romanNumber = "LXXVI"
	case 77:
		romanNumber = "LXXVII"
	case 78:
		romanNumber = "LXXVIII"
	case 79:
		romanNumber = "LXXIX"
	case 80:
		romanNumber = "LXXX"
	case 81:
		romanNumber = "LXXXI"
	case 82:
		romanNumber = "LXXXII"
	case 83:
		romanNumber = "LXXXIII"
	case 84:
		romanNumber = "LXXXIV"
	case 85:
		romanNumber = "LXXXV"
	case 86:
		romanNumber = "LXXXVI"
	case 87:
		romanNumber = "LXXXVII"
	case 88:
		romanNumber = "LXXXVIII"
	case 89:
		romanNumber = "LXXXIX"
	case 90:
		romanNumber = "XC"
	case 91:
		romanNumber = "XCI"
	case 92:
		romanNumber = "XCII"
	case 93:
		romanNumber = "XCIII"
	case 94:
		romanNumber = "XCIV"
	case 95:
		romanNumber = "XCV"
	case 96:
		romanNumber = "XCVI"
	case 97:
		romanNumber = "XCVII"
	case 98:
		romanNumber = "XCVIII"
	case 99:
		romanNumber = "XCIX"
	case 100:
		romanNumber = "C"
	}
	return romanNumber
}

func isRoman(num string) bool {
	for _, n := range romanInput {
		if n == num {
			return true
		}
	}
	return false
}
