package tournament

import (
    "bufio"
    "fmt"
    "io"
    "strings"
    "sort"
    "errors"
)

type TeamStats struct {
    MatchesPlayed int
	Wins          int
	Draws         int
	Losses        int
	Points        int
}

type Team map[string]*TeamStats 

func New() Team {
    return make(Team)
}

func (t Team) String() string {
    if len(t) == 0 {
        return ""
    }
    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P"))
    for name, stats := range t {
        sb.WriteString(fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", name, stats.MatchesPlayed, stats.Wins, stats.Draws, stats.Losses, stats.Points))
    }
    return sb.String()
}

func SortTeams(teams Team) Team{
	type teamEntry struct {
		Name  string
		Stats *TeamStats
	}

	var teamList []teamEntry
	for name, stats := range teams {
		teamList = append(teamList, teamEntry{Name: name, Stats: stats})
	}

	// Сортируем сначала по очкам, затем по имени
	sort.Slice(teamList, func(i, j int) bool {
		if teamList[i].Stats.Points == teamList[j].Stats.Points {
			return teamList[i].Name < teamList[j].Name // Сортировка по имени, если очки равны
		}
		return teamList[i].Stats.Points > teamList[j].Stats.Points // Сортировка по очкам
	})

	// Создаем новую отсортированную карту Team
	sortedTeams := make(Team)
	for _, entry := range teamList {
		sortedTeams[entry.Name] = entry.Stats
	}

	return sortedTeams
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
    teams := New()
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 || line[0] == '#' || line[0] == '\n' {
            continue
        }
        data := strings.Split(line, ";")
        if len(data) != 3 {
            return errors.New("")
        }
        team1 := data[0]
		team2 := data[1]
		result := data[2]

		if _, ok := teams[team1]; !ok {
			teams[team1] = &TeamStats{}
		}
		if _, ok := teams[team2]; !ok {
			teams[team2] = &TeamStats{}
		}

		teams[team1].MatchesPlayed++
		teams[team2].MatchesPlayed++

		switch result {
		case "win":
			teams[team1].Wins++
			teams[team1].Points += 3
			teams[team2].Losses++
		case "loss":
			teams[team2].Wins++
			teams[team2].Points += 3
			teams[team1].Losses++
		case "draw":
			teams[team1].Draws++
			teams[team2].Draws++
			teams[team1].Points++
			teams[team2].Points++
        default:
            return errors.New("")
		}
    }
    if err := scanner.Err(); err != nil {
		return err
	}

	// Write the results to the writer
	sortedTeams := SortTeams(teams)
    _, err := writer.Write([]byte(sortedTeams.String()))
	return err
}
