package main // piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	idx := 0
	number := 0
	str := ""
	s := ""
	for _, charNbr := range nbr {
		for idxBasefrom, charBasefrom := range baseFrom {
			if charNbr == charBasefrom {
				idx = idxBasefrom
			}
		}
		number = (number * len(baseFrom)) + idx
	}
	bt := 0
	for number > 0 {

		bt = number % len(baseTo)

		str += string(baseTo[bt])

		number /= len(baseTo)

	}
	for i := len(str) - 1; i >= 0; i-- {
		s += string(str[i])
	}
	return s
}
