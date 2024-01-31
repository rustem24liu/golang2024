package api

type Band struct {
	Id           string `json:"id"`
	Created_year int    `json:"created_year"`
	Name         string `json:"name"`
	Musics       *Music `json:"musics"`
	Country      string `json:"country"`
	Head         *Head  `json:"head"`
}

type Music struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Duration string `json:"duration"`
}

type Head struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var bands = []Band{
	{Id: "1", Created_year: 1994, Name: "Rammstein", Musics: &Music{Name: "Sonne", Year: 2001, Duration: "4:32"}},
	{Id: "2", Created_year: 1994, Name: "AC/DC", Musics: &Music{Name: "Highway to Hell", Year: 1979, Duration: "3:28"}},
	{Id: "3", Created_year: 1991, Name: "Radiohead", Musics: &Music{Name: "Creep", Year: 1992, Duration: "3:56"}},
	{Id: "4", Created_year: 1983, Name: "Red Hot Chili Peppers", Musics: &Music{Name: "Under the Bridge", Year: 1992, Duration: "4:24"}},
	{Id: "5", Created_year: 2002, Name: "Arctic Monkeys", Musics: &Music{Name: "Do I Wanna Know?", Year: 2013, Duration: "4:32"}},
	{Id: "6", Created_year: 1973, Name: "Led Zeppelin", Musics: &Music{Name: "Stairway to Heaven", Year: 1971, Duration: "8:02"}},
	{Id: "7", Created_year: 1980, Name: "Queen", Musics: &Music{Name: "Bohemian Rhapsody", Year: 1975, Duration: "5:55"}},
	{Id: "8", Created_year: 1962, Name: "The Beatles", Musics: &Music{Name: "Hey Jude", Year: 1968, Duration: "7:11"}},
	{Id: "9", Created_year: 1996, Name: "Oasis", Musics: &Music{Name: "Wonderwall", Year: 1995, Duration: "4:18"}},
	{Id: "10", Created_year: 1985, Name: "Guns N' Roses", Musics: &Music{Name: "Sweet Child o' Mine", Year: 1987, Duration: "5:56"}},
}
