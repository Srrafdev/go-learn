package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := i; j < len(podium); j++ {
			podium[i][0], podium[j][0] = podium[j][0], podium[i][0]
		}
	}
	return podium
}
