package services

import (
	"strings"
)

type CompatibilityInfo struct {
	Percentage  int
	Description string
}

const (
	DualityRelation       = 80
	ActivationRelation    = 60
	SemiDualityRelation   = 50
	MirageRelation        = 40
	MirrorRelation        = 50
	IdentityRelation      = 40
	BusinessRelation      = 40
	QuasiIdentityRelation = 40
	ContrastRelation      = 20
	SocialOrderRelation   = 20
	KindredRelation       = 40
	OrderRelation         = 20
	ControlRelation       = 20
	SuperEgoRelation      = 40
	RevisionRelation      = 20
	ConflictRelation      = 0
)

var staticCompatibility = map[string]CompatibilityInfo{
	"ENTP_ENTP": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ENTP_ISFP": {
		Percentage:  DualityRelation,
		Description: ""},
	"ENTP_ESFJ": {
		Percentage:  ActivationRelation,
		Description: ""},
	"ENTP_INTJ": {
		Percentage:  MirrorRelation,
		Description: ""},
	"ENTP_ENFJ": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ENTP_ISTJ": {
		Percentage:  RevisionRelation,
		Description: ""},
	"ENTP_ESTP": {
		Percentage:  BusinessRelation,
		Description: ""},
	"ENTP_INFP": {
		Percentage:  MirageRelation,
		Description: ""},
	"ENTP_ESFP": {
		Percentage:  SuperEgoRelation,
		Description: ""},
	"ENTP_INTP": {
		Percentage:  ContrastRelation,
		Description: ""},
	"ENTP_ENTJ": {
		Percentage:  QuasiIdentityRelation,
		Description: ""},
	"ENTP_ISFJ": {
		Percentage:  ConflictRelation,
		Description: ""},
	"ENTP_ESTJ": {
		Percentage:  OrderRelation,
		Description: ""},
	"ENTP_INFJ": {
		Percentage:  ControlRelation,
		Description: ""},
	"ENTP_ENFP": {
		Percentage:  KindredRelation,
		Description: ""},
	"ENTP_ISTP": {
		Percentage:  SemiDualityRelation,
		Description: ""},
	"ISFP_ISFP": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ISFP_ESFJ": {
		Percentage:  MirrorRelation,
		Description: ""},
	"ISFP_INTJ": {
		Percentage:  ActivationRelation,
		Description: ""},
	"ISFP_ENFJ": {
		Percentage:  RevisionRelation,
		Description: ""},
	"ISFP_ISTJ": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ISFP_ESTP": {
		Percentage:  MirageRelation,
		Description: ""},
	"ISFP_INFP": {
		Percentage:  BusinessRelation,
		Description: ""},
	"ISFP_ESFP": {
		Percentage:  ContrastRelation,
		Description: ""},
	"ISFP_INTP": {
		Percentage:  SuperEgoRelation,
		Description: ""},
	"ISFP_ENTJ": {
		Percentage:  ConflictRelation,
		Description: ""},
	"ISFP_ISFJ": {
		Percentage:  QuasiIdentityRelation,
		Description: ""},
	"ISFP_ESTJ": {
		Percentage:  ControlRelation,
		Description: ""},
	"ISFP_INFJ": {
		Percentage:  OrderRelation,
		Description: ""},
	"ISFP_ENFP": {
		Percentage:  SemiDualityRelation,
		Description: ""},
	"ISFP_ISTP": {
		Percentage:  KindredRelation,
		Description: ""},
	"ESFJ_ESFJ": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ESFJ_INTJ": {
		Percentage:  DualityRelation,
		Description: ""},
	"ESFJ_ENFJ": {
		Percentage:  KindredRelation,
		Description: ""},
	"ESFJ_ISTJ": {
		Percentage:  SemiDualityRelation,
		Description: ""},
	"ESFJ_ESTP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ESFJ_INFP": {
		Percentage:  ControlRelation,
		Description: ""},
	"ESFJ_ESFP": {
		Percentage:  QuasiIdentityRelation,
		Description: ""},
	"ESFJ_INTP": {
		Percentage:  ConflictRelation,
		Description: ""},
	"ESFJ_ENTJ": {
		Percentage:  SuperEgoRelation,
		Description: ""},
	"ESFJ_ISFJ": {
		Percentage:  ContrastRelation,
		Description: ""},
	"ESFJ_ESTJ": {
		Percentage:  BusinessRelation,
		Description: ""},
	"ESFJ_INFJ": {
		Percentage:  MirageRelation,
		Description: ""},
	"ESFJ_ENFP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ESFJ_ISTP": {
		Percentage:  ControlRelation,
		Description: ""},
	"INTJ_INTJ": {
		Percentage:  IdentityRelation,
		Description: ""},
	"INTJ_ENFJ": {
		Percentage:  SemiDualityRelation,
		Description: ""},
	"INTJ_ISTJ": {
		Percentage:  KindredRelation,
		Description: ""},
	"INTJ_ESTP": {
		Percentage:  ControlRelation,
		Description: ""},
	"INTJ_INFP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"INTJ_ESFP": {
		Percentage:  ConflictRelation,
		Description: ""},
	"INTJ_INTP": {
		Percentage:  QuasiIdentityRelation,
		Description: ""},
	"INTJ_ENTJ": {
		Percentage:  ContrastRelation,
		Description: ""},
	"INTJ_ISFJ": {
		Percentage:  SuperEgoRelation,
		Description: ""},
	"INTJ_ESTJ": {
		Percentage:  98,
		Description: ""},
	"INTJ_INFJ": {
		Percentage:  BusinessRelation,
		Description: ""},
	"INTJ_ENFP": {
		Percentage:  ControlRelation,
		Description: ""},
	"INTJ_ISTP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ENFJ_ENFJ": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ENFJ_ISTJ": {
		Percentage:  DualityRelation,
		Description: ""},
	"ENFJ_ESTP": {
		Percentage:  ActivationRelation,
		Description: ""},
	"ENFJ_INFP": {
		Percentage:  MirrorRelation,
		Description: ""},
	"ENFJ_ESFP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ENFJ_INTP": {
		Percentage:  ControlRelation,
		Description: ""},
	"ENFJ_ENTJ": {
		Percentage:  BusinessRelation,
		Description: ""},
	"ENFJ_ISFJ": {
		Percentage:  MirageRelation,
		Description: ""},
	"ENFJ_ESTJ": {
		Percentage:  SuperEgoRelation,
		Description: ""},
	"ENFJ_INFJ": {
		Percentage:  ContrastRelation,
		Description: ""},
	"ENFJ_ENFP": {
		Percentage:  QuasiIdentityRelation,
		Description: ""},
	"ENFJ_ISTP": {
		Percentage:  ConflictRelation,
		Description: ""},
	"ISTJ_ISTJ": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ISTJ_ESTP": {
		Percentage:  MirrorRelation,
		Description: ""},
	"ISTJ_INFP": {
		Percentage:  ActivationRelation,
		Description: ""},
	"ISTJ_ESFP": {
		Percentage:  RevisionRelation,
		Description: ""},
	"ISTJ_INTP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ISTJ_ENTJ": {
		Percentage:  MirageRelation,
		Description: ""},
	"ISTJ_ISFJ": {
		Percentage:  BusinessRelation,
		Description: ""},
	"ISTJ_ESTJ": {
		Percentage:  ContrastRelation,
		Description: ""},
	"ISTJ_INFJ": {
		Percentage:  SuperEgoRelation,
		Description: ""},
	"ISTJ_ENFP": {
		Percentage:  ConflictRelation,
		Description: ""},
	"ISTJ_ISTP": {
		Percentage:  QuasiIdentityRelation,
		Description: ""},
	"ESTP_ESTP": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ESTP_INFP": {
		Percentage:  DualityRelation,
		Description: ""},
	"ESTP_ESFP": {
		Percentage:  KindredRelation,
		Description: ""},
	"ESTP_INTP": {
		Percentage:  SemiDualityRelation,
		Description: ""},
	"ESTP_ENTJ": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ESTP_ISFJ": {
		Percentage:  RevisionRelation,
		Description: ""},
	"ESTP_ESTJ": {
		Percentage:  QuasiIdentityRelation,
		Description: ""},
	"ESTP_INFJ": {
		Percentage:  91,
		Description: ""},
	"ESTP_ENFP": {
		Percentage:  DualityRelation,
		Description: ""},
	"ESTP_ISTP": {
		Percentage:  ContrastRelation,
		Description: ""},
	"INFP_INFP": {
		Percentage:  IdentityRelation,
		Description: ""},
	"INFP_ESFP": {
		Percentage:  SemiDualityRelation,
		Description: ""},
	"INFP_INTP": {
		Percentage:  KindredRelation,
		Description: ""},
	"INFP_ENTJ": {
		Percentage:  RevisionRelation,
		Description: ""},
	"INFP_ISFJ": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"INFP_ESTJ": {
		Percentage:  ConflictRelation,
		Description: ""},
	"INFP_INFJ": {
		Percentage:  QuasiIdentityRelation,
		Description: ""},
	"INFP_ENFP": {
		Percentage:  ContrastRelation,
		Description: ""},
	"INFP_ISTP": {
		Percentage:  SuperEgoRelation,
		Description: ""},
	"ESFP_ESFP": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ESFP_INTP": {
		Percentage:  DualityRelation,
		Description: ""},
	"ESFP_ENTJ": {
		Percentage:  ActivationRelation,
		Description: ""},
	"ESFP_ISFJ": {
		Percentage:  MirrorRelation,
		Description: ""},
	"ESFP_ESTJ": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ESFP_INFJ": {
		Percentage:  RevisionRelation,
		Description: ""},
	"ESFP_ENFP": {
		Percentage:  BusinessRelation,
		Description: ""},
	"ESFP_ISTP": {
		Percentage:  MirageRelation,
		Description: ""},
	"INTP_INTP": {
		Percentage:  IdentityRelation,
		Description: ""},
	"INTP_ENTJ": {
		Percentage:  MirrorRelation,
		Description: ""},
	"INTP_ISFJ": {
		Percentage:  ActivationRelation,
		Description: ""},
	"INTP_ESTJ": {
		Percentage:  RevisionRelation,
		Description: ""},
	"INTP_INFJ": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"INTP_ENFP": {
		Percentage:  MirageRelation,
		Description: ""},
	"INTP_ISTP": {
		Percentage:  BusinessRelation,
		Description: ""},
	"ENTJ_ENTJ": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ENTJ_ISFJ": {
		Percentage:  DualityRelation,
		Description: ""},
	"ENTJ_ESTJ": {
		Percentage:  KindredRelation,
		Description: ""},
	"ENTJ_INFJ": {
		Percentage:  SemiDualityRelation,
		Description: ""},
	"ENTJ_ENFP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ENTJ_ISTP": {
		Percentage:  RevisionRelation,
		Description: ""},
	"ISFJ_ISFJ": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ISFJ_ESTJ": {
		Percentage:  SemiDualityRelation,
		Description: ""},
	"ISFJ_INFJ": {
		Percentage:  KindredRelation,
		Description: ""},
	"ISFJ_ENFP": {
		Percentage:  RevisionRelation,
		Description: ""},
	"ISFJ_ISTP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"ESTJ_ESTJ": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ESTJ_INFJ": {
		Percentage:  KindredRelation,
		Description: ""},
	"ESTJ_ENFP": {
		Percentage:  RevisionRelation,
		Description: ""},
	"ESTJ_ISTP": {
		Percentage:  SocialOrderRelation,
		Description: ""},
	"INFJ_INFJ": {
		Percentage:  IdentityRelation,
		Description: ""},
	"INFJ_ENFP": {
		Percentage:  MirrorRelation,
		Description: ""},
	"INFJ_ISTP": {
		Percentage:  ActivationRelation,
		Description: ""},
	"ENFP_ENFP": {
		Percentage:  IdentityRelation,
		Description: ""},
	"ENFP_ISTP": {
		Percentage:  DualityRelation,
		Description: ""},
	"ISTP_ISTP": {
		Percentage:  IdentityRelation,
		Description: ""},
}

var zodiacCompatibility = map[string]map[string]int{
	"Овен": {
		"Овен":     20,
		"Телец":    5,
		"Близнецы": 15,
		"Рак":      0,
		"Лев":      20,
		"Дева":     0,
		"Весы":     5,
		"Скорпион": 5,
		"Стрелец":  20,
		"Козерог":  5,
		"Водолей":  15,
		"Рыбы":     0,
	},
	"Телец": {
		"Овен":     5,
		"Телец":    20,
		"Близнецы": 5,
		"Рак":      20,
		"Лев":      0,
		"Дева":     15,
		"Весы":     20,
		"Скорпион": 5,
		"Стрелец":  5,
		"Козерог":  20,
		"Водолей":  0,
		"Рыбы":     15,
	},
	"Близнецы": {
		"Овен":     15,
		"Телец":    5,
		"Близнецы": 20,
		"Рак":      5,
		"Лев":      15,
		"Дева":     0,
		"Весы":     20,
		"Скорпион": 5,
		"Стрелец":  5,
		"Козерог":  5,
		"Водолей":  20,
		"Рыбы":     0,
	},
	"Рак": {
		"Овен":     0,
		"Телец":    15,
		"Близнецы": 5,
		"Рак":      20,
		"Лев":      5,
		"Дева":     15,
		"Весы":     0,
		"Скорпион": 20,
		"Стрелец":  5,
		"Козерог":  5,
		"Водолей":  5,
		"Рыбы":     20,
	},
	"Лев": {
		"Овен":     20,
		"Телец":    0,
		"Близнецы": 15,
		"Рак":      5,
		"Лев":      20,
		"Дева":     5,
		"Весы":     15,
		"Скорпион": 0,
		"Стрелец":  20,
		"Козерог":  5,
		"Водолей":  5,
		"Рыбы":     5,
	},
	"Дева": {
		"Овен":     5,
		"Телец":    20,
		"Близнецы": 0,
		"Рак":      15,
		"Лев":      5,
		"Дева":     20,
		"Весы":     5,
		"Скорпион": 15,
		"Стрелец":  0,
		"Козерог":  20,
		"Водолей":  5,
		"Рыбы":     5,
	},
	"Весы": {
		"Овен":     5,
		"Телец":    5,
		"Близнецы": 20,
		"Рак":      0,
		"Лев":      15,
		"Дева":     5,
		"Весы":     20,
		"Скорпион": 5,
		"Стрелец":  15,
		"Козерог":  0,
		"Водолей":  20,
		"Рыбы":     5,
	},
	"Скорпион": {
		"Овен":     5,
		"Телец":    5,
		"Близнецы": 5,
		"Рак":      20,
		"Лев":      0,
		"Дева":     15,
		"Весы":     5,
		"Скорпион": 20,
		"Стрелец":  5,
		"Козерог":  15,
		"Водолей":  0,
		"Рыбы":     20,
	},
	"Стрелец": {
		"Овен":     20,
		"Телец":    5,
		"Близнецы": 5,
		"Рак":      5,
		"Лев":      20,
		"Дева":     0,
		"Весы":     15,
		"Скорпион": 5,
		"Стрелец":  20,
		"Козерог":  5,
		"Водолей":  15,
		"Рыбы":     0,
	},
	"Козерог": {
		"Овен":     0,
		"Телец":    20,
		"Близнецы": 5,
		"Рак":      5,
		"Лев":      5,
		"Дева":     20,
		"Весы":     0,
		"Скорпион": 15,
		"Стрелец":  5,
		"Козерог":  20,
		"Водолей":  5,
		"Рыбы":     15,
	},
	"Водолей": {
		"Овен":     15,
		"Телец":    0,
		"Близнецы": 20,
		"Рак":      5,
		"Лев":      5,
		"Дева":     5,
		"Весы":     20,
		"Скорпион": 0,
		"Стрелец":  15,
		"Козерог":  5,
		"Водолей":  20,
		"Рыбы":     0,
	},
	"Рыбы": {
		"Овен":     5,
		"Телец":    15,
		"Близнецы": 0,
		"Рак":      20,
		"Лев":      5,
		"Дева":     5,
		"Весы":     5,
		"Скорпион": 20,
		"Стрелец":  0,
		"Козерог":  15,
		"Водолей":  5,
		"Рыбы":     20,
	},
}

func GetCompatibilityData(myMBTI, otherMBTI string) (int, string) {
	key1 := strings.ToUpper(myMBTI) + "_" + strings.ToUpper(otherMBTI)
	key2 := strings.ToUpper(otherMBTI) + "_" + strings.ToUpper(myMBTI)
	if data, ok := staticCompatibility[key1]; ok {
		return data.Percentage, data.Description
	}
	if data, ok := staticCompatibility[key2]; ok {
		return data.Percentage, data.Description
	}
	return 0, "Данные о совместимости отсутствуют."
}

func GetZodiacCompatibility(myZodiac, otherZodiac string) int {
	myZodiac = strings.TrimSpace(myZodiac)
	otherZodiac = strings.TrimSpace(otherZodiac)
	if compat, ok := zodiacCompatibility[myZodiac][otherZodiac]; ok {
		return compat
	}
	return 0
}

var personalitySelfDescriptions = map[string]string{
	"INFJ": "Глубокий, чуткий и проницательный человек.\nУ вас есть уникальная способность ощущать эмоциональное состояние и физические потребности своего партнера.",
	"ENFJ": "Вы очень жизнерадостная личность, которая умеет привнести в отношения ощущения новизны и праздника.",
	"INFP": "Вы очень глубокая и искренняя натура, всегда готовы подстроиться под других, проявить уступчивость.\nНо если в процессе взаимодействия обнаруживается несовпадение ключевых ценностных ориентиров, то Вы вмиг превращаетесь в жестких и резких защитников своей позиции.",
	"ENFP": "Вы глубокий, теплый и нежный человек.\nВы буквально окутываете партнера вниманием и любовью.\n",
	"INTJ": "Вы очень созидательная натура.\nК построению отношений вы подходите бережно, прислушиваясь и внимательно приглядываясь к партнеру.",
	"ENTJ": "К отношениям Вы подходите весьма серьезно и с большой ответственностью.\nУ Вас очень высокие стандарты для себя, и такие же высокие требования к партнеру.",
	"INTP": "Вы можете быть чрезвычайно преданными в отношениях и глубоко влюбленными в своего партнера.\n ВЫ любяите заботиться о других, но Вам трудно открыто выражать это. Вам сложно понять свои новые и, возможно, сильные эмоции, и еще труднее выразить их.",
	"ENTP": "В отношениях Вы можете быть прекрасным партнером.\nВы спокойны и гибки, с Вами легко поладить и договориться.\nОднако если партнер будет Вас сдерживать, предпочитая оставаться на одном месте, Вы не будете колебаться: Вы просто уйдете.",
	"ISTJ": "Вы очень надежны в отношениях.\nВы серьезно относитесь к обязательствам и хотите, чтобы Ваши партнеры знали, что на Вас всегда можно положиться.",
	"ESTJ": "Вы прагматичны и продуктивны.\nСделаете почти все для человека, которого любите. Они стремятся быть уравновешенными, разумными и лояльными. Вы убеждены, что Ваша любовь очевидна в действиях, а не в том, чтобы выражать ее вслух.",
	"ISTP": "Вы предпочитаете наслаждаться настоящим моментом и не склонны брать на себя долгосрочные обязательства.\nВы больше всего цените собственную автономию и личное пространство.\nВы способны испытывать глубокие сильные чувства, но не всегда умеете выражать их. Вы не любите озвучивать то, что считаете понятным без слов.",
	"ESTP": "Вы живете настоящим моментом.\nМожете быть безрассудными в своей постоянной потребности в новых впечатлениях и новых завоеваниях.\nВремя от времени можете быть преданными партнерами, но выполняете такие обязательства довольно редко.",
	"ISFJ": "Вы придаете большое значение своим личным отношениям.\nВы очень щедрый и любящий человек.\nЗабота о других и служение естественны для Вас.",
	"ESFJ": "Вы чрезвычайно лояльны, заботливы и ответственны в отношениях.\nПризнательность — это величайший подарок, который могут дать Вам Ваши партнеры.",
	"ISFP": "Вы теплый и щедрый человек.\nУ Вас есть глубина эмоций и заботы, которые часто не очевидны для других, кроме тех, кто очень хорошо Вас знает.",
	"ESFP": "Вы относятся к своим любовным отношениям с размахом — подобно тому, как Вы относитесь к своей жизни в целом.\nВы любите быть влюбленными и стараетесь максимально использовать каждый момент.\nИ в то же время Вы старательно избегаете всяческих обязательств и обещаний.",
}

func GetPersonalitySelfDescription(mbti string) string {
	if desc, ok := personalitySelfDescriptions[strings.ToUpper(mbti)]; ok {
		return desc
	}
	return "Описание для данного типа личности отсутствует."
}
