package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {

	if n == 1 {

		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(i)
			if i != '9' {

				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}

	if n == 3 {

		for i := '0'; i <= '7'; i++ {
			for j := i + 1; j <= '8'; j++ {
				for k := j + 1; k <= '9'; k++ {

					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)

					if i != '7' || j != '8' || k != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}

	if n == 9 {

		for a := '0'; a <= '1'; a++ {
			for b := a + 1; b <= '2'; b++ {
				for c := b + 1; c <= '3'; c++ {
					for d := c + 1; d <= '4'; d++ {
						for e := d + 1; e <= '5'; e++ {
							for f := e + 1; f <= '6'; f++ {
								for g := f + 1; g <= '7'; g++ {
									for h := g + 1; h <= '8'; h++ {
										for i := h + 1; i <= '9'; i++ {
											z01.PrintRune(a)
											z01.PrintRune(b)
											z01.PrintRune(c)
											z01.PrintRune(d)
											z01.PrintRune(e)
											z01.PrintRune(f)
											z01.PrintRune(g)
											z01.PrintRune(h)
											z01.PrintRune(i)
											if a != '1' || b != '2' || c != '3' || d != '4' || e != '5' || f != '6' || g != '7' || h != '8' || i != '9' {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
