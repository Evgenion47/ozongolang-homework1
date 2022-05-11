package lastlettergame

func search(pokemon []string, lenght int) {
	if lenght > len(longestChain) {
		longestChain = nil
		for i := 0; i < lenght; i++ {
			longestChain = append(longestChain, pokemon[i])
		}
	}
	for i := lenght; i < len(pokemon); i++ {
		if pokemon[i][0] == pokemon[lenght-1][len(pokemon[lenght-1])-1] {
			pokemon[i], pokemon[lenght] = pokemon[lenght], pokemon[i]
			search(pokemon, lenght+1)
			pokemon[i], pokemon[lenght] = pokemon[lenght], pokemon[i]
		}
	}
}

var longestChain []string

func Sequence(dic []string) []string {
	for i := range dic {
		dic[0], dic[i] = dic[i], dic[0]
		search(dic, 1)
		dic[0], dic[i] = dic[i], dic[0]
	}
	return longestChain
}
