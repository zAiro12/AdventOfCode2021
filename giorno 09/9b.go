package main

import "fmt"

type punto struct {
	x, y int
}

func main() {
	var matrice [100][]int
	for i := 0; ; i++ {
		var in string
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		newin := []rune(in)
		for j := 0; j < len(newin); j++ {
			matrice[i] = append(matrice[i], int(newin[j])-48)
		}
	}
	punti := salva(matrice)
	fmt.Println(punti)

}

func salva(matrice [100][]int) []punto {
	var punti []punto
	var punto punto
	for i := 0; i < 100; i++ {
		for j := 0; j < len(matrice[i]); j++ {
			//controllo lato sx
			if i == 0 {
				if j == 0 {
					if matrice[i][j] < matrice[i][j+1] && matrice[i][j] < matrice[i+1][j] {
						punto.x = j
						punto.y = i
						punti = append(punti, punto)
					}
				} else if j == len(matrice[i])-1 {
					if matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i][j-1] {
						punto.x = j
						punto.y = i
						punti = append(punti, punto)
					}
				} else {
					if matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i][j-1] && matrice[i][j] < matrice[i][j+1] {
						punto.x = j
						punto.y = i
						punti = append(punti, punto)
					}
				}
				//controllo lato dx
			} else if i == len(matrice)-1 {
				if j == 0 {
					if matrice[i][j] < matrice[i][j+1] && matrice[i][j] < matrice[i+1][j] {
						punto.x = j
						punto.y = i
						punti = append(punti, punto)
					}
				} else if j == len(matrice[i])-1 {
					if matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i][j-1] {
						punto.x = j
						punto.y = i
						punti = append(punti, punto)
					}
				} else {
					if matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i][j-1] && matrice[i][j] < matrice[i][j+1] {
						punto.x = j
						punto.y = i
						punti = append(punti, punto)
					}
				}
				//controllo lato top
			} else if j == 0 {
				if matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i][j+1] {
					punto.x = j
					punto.y = i
					punti = append(punti, punto)
				}
				//controllo lato bot
			} else if j == len(matrice[i])-1 {
				if matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i][j-1] {
					punto.x = j
					punto.y = i
					punti = append(punti, punto)
				}
				//caso generale
			} else {
				if matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i][j+1] && matrice[i][j] < matrice[i][j-1] {
					punto.x = j
					punto.y = i
					punti = append(punti, punto)
				}
			}
		}
	}
	return punti
}

func cerca(punti []punto) {
	for i := 0; i < len(punti); i++ {
		//controllo angoli
		//00
		if punti[i].x == 0 {
			if punti[i].y == 0 {
				if punti[i+1].x != 9 {

				}
			}
		}

	}
	return
}
