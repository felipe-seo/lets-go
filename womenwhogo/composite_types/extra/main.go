package main

import "fmt"

func main() {
	//exercise 1
	yellowTeam := [5]string{
		"Fernando",
		"João",
		"Lucia",
		"Mariana",
		"Ana",
	}
	redTeam := [5]string{
		"Helena",
		"Jonas",
		"José",
		"Juliana",
	}
	fmt.Println("Yellows: ", yellowTeam, "\nReds: ", redTeam)
	fmt.Printf("%T", yellowTeam)

	//exercise 2
	yellowTeamSlice := yellowTeam[:]
	redTeamSlice := redTeam[:]

	redTeamSlice = append(redTeamSlice, "Luis Inácio")
	fmt.Println("\n\nYellows: ", yellowTeamSlice, "\nReds: ", redTeamSlice)
	fmt.Printf("%T\n", yellowTeamSlice)

	//exercise 3
	cidades := map[string]string{
		"São Paulo":      "Brazil",
		"Belo horizonte": "Brazil",
		"Recife":         "Brazil",
		"New York":       "Brazil",
		"Yokohama":       "Japan",
		"Nagoya":         "Japan",
		"Osaka":          "Japan",
		"New york":       "United States",
		"New Orleans":    "United States",
		"Detroit":        "United States",
	}
	frequency := make(map[string]int)
	fmt.Println(cidades)

	for _, v := range cidades {
		frequency[v] += 1
	}
	fmt.Println("\n", frequency)
}
