package main //piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				for s := '9'; s >= '0'; s-- {
					if (i*10)+(j-48) > (k*10)+(s-48) {
						priz01(i)
						priz01(j)
						priz01(' ')
						priz01(k)
						priz01(s)
						if i == '0' && j == '1' && k == '0' && s == '0' {
							break
						}
						priz01(',')
						priz01(' ')
					}
				}
			}
		}
	}
}

func priz01(s rune) {
	z01.PrintRune(s)
}
