package model

type Translations struct {
	Profile    string
	Experience string
	Projects   string
	Education  string
	Courses    string
	Skills     string
	Languages  string
	Contact    string
}

func GetTranslations(lang string) Translations {
	switch lang {
	case "ru":
		return Translations{
			Profile:    "Профиль",
			Experience: "Опыт работы",
			Projects:   "Проекты",
			Education:  "Образование",
			Courses:    "Курсы",
			Skills:     "Навыки",
			Languages:  "Языки",
			Contact:    "Контакты",
		}
	default:
		return Translations{
			Profile:    "Profile",
			Experience: "Experience",
			Projects:   "Projects",
			Education:  "Education",
			Courses:    "Courses",
			Skills:     "Skills",
			Languages:  "Languages",
			Contact:    "Contact",
		}
	}
}
