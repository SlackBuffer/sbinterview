package sort

import (
	"sort"
)

type Student struct {
	Name    string
	ID      int
	Score1  int
	Score2  int
	Ranking int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	// 排名展示顺序的权重：Score1 > Score2 > ID（ID 小在前）
	return s[i].Score1 > s[j].Score1 || (s[i].Score1 == s[j].Score1 && s[i].Score2 > s[j].Score2) || (s[i].Score1 == s[j].Score1 && s[i].Score2 == s[j].Score2 && s[i].ID < s[j].ID)
}

func SortStudents(s Students) Students {
	if len(s) == 0 {
		return nil
	}
	sort.Sort(s)

	s[0].Ranking = 1
	for i := 1; i < len(s); i++ {
		if s[i].Score1 == s[i-1].Score1 && s[i].Score2 == s[i-1].Score2 {
			s[i].Ranking = s[i-1].Ranking
		} else {
			s[i].Ranking = i + 1
		}
	}
	return s
}
