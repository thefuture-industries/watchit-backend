package machinelearning

import (
	"math"
)

// Структура для хранения результатов анализа тона
type ToneAnalysis struct {
	Positive   float64
	Negative   float64
	Horror     float64
	TotalScore float64
}

// Анализ эмоционального тона текста
// ---------------------------------
func AnalyzeTone(str string) ToneAnalysis {
	words := Stemming(str)
	var analysis ToneAnalysis

	if len(words) == 0 {
		return analysis
	}

	// Словарь эмоциональных тонов
	var positiveWordWeights = map[string]float64{"good": 2.734, "great": 2.627, "awesome": 3.484, "amazing": 2.466, "excellent": 1.757, "wonderful": 3.64, "fantastic": 2.7342, "beautiful": 1.6, "brilliant": 1.3, "perfect": 1.5, "happy": 1.1, "joy": 1.2, "love": 1.3, "best": 1.4, "exciting": 1.2, "fun": 1.1, "funny": 1.1, "nice": 0.9, "pleasant": 0.9, "charming": 1.0, "masterpiece": 2.0, "innovative": 1.4, "creative": 1.3, "engaging": 1.2, "powerful": 1.4, "impressive": 1.3, "memorable": 1.3, "outstanding": 1.5, "captivating": 1.4, "stunning": 1.4, "inspiring": 1.3, "magical": 1.3, "original": 1.4, "clever": 1.2, "authentic": 1.3, "fresh": 1.241, "dynamic": 1.2, "smooth": 1.1, "solid": 1.5, "christmas": 1.453, "festive": 1.334, "celebration": 1.7233, "merry": 1.464, "jolly": 1.362, "cheerful": 1.3, "joyful": 4.443, "entertaining": 3.2843, "enjoyable": 3.5623, "delightful": 1.633, "thrilling": 5.7345, "adventure": 3.35123, "superb": 3.845, "incredible": 2.8340, "spectacular": 3.0340, "remarkable": 2.6340, "touching": 2.4340, "heartwarming": 2.6340, "uplifting": 2.6340, "moving": 2.4340, "emotional": 2.4340, "amaz": 2.8340, "excell": 2.6340, "wonder": 2.6340, "fantast": 2.8340, "brillian": 2.6340, "graceful": 2.4340, "triumph": 2.8340, "masterful": 3.0340, "enchanting": 2.8340, "polished": 2.6340, "refined": 2.4340, "wholesome": 2.4340, "refreshing": 2.6340, "compelling": 2.8340, "passionate": 2.6340, "glorious": 2.8340, "marvelous": 2.8340, "splendid": 2.6340, "elegant": 2.4340, "exquisite": 3.0340, "divine": 2.8340, "radiant": 2.6340, "vibrant": 2.6340, "harmonious": 2.4340, "sublime": 2.8340, "peaceful": 2.4340, "legendary": 2.8340, "extraordinary": 2.8340, "devoted": 2.4340, "beloved": 2.6340, "heroic": 2.8340, "determined": 2.4340, "successful": 2.6340, "revered": 2.8340, "victorious": 2.8340, "sacred": 2.6340}
	var negativeWordWeights = map[string]float64{"bad": -1.2341, "terrible": -3.8342, "horrible": -2.9343, "awful": -1.6344, "poor": -2.0345, "worst": -3.9346, "boring": -1.4347, "stupid": -2.6348, "dull": -1.2349, "ugly": -2.4350, "sad": -1.8351, "hate": -3.8352, "disappointing": -2.4353, "weak": -1.0354, "wrong": -2.2355, "waste": -3.6356, "annoying": -1.4357, "confusing": -2.2358, "painful": -3.4359, "mediocre": -1.0360, "cliche": -2.4361, "predictable": -1.2362, "shallow": -2.4363, "amateur": -3.6364, "messy": -1.2365, "slow": -2.0366, "cheap": -1.4367, "derivative": -2.2368, "forgettable": -3.6369, "incoherent": -2.6370, "flat": -1.0371, "pointless": -3.6372, "awkward": -1.2373, "ridiculous": -2.4374, "unwatchable": -3.9375, "amateurish": -2.6376, "failed": -3.8377, "flawed": -1.4378, "generic": -2.2379, "lifeless": -3.6380, "disastrous": -3.8381, "pathetic": -2.6382, "abysmal": -3.8383, "atrocious": -3.9384, "incompetent": -2.6385, "worthless": -3.8386, "tedious": -1.4387, "tasteless": -2.4388, "unpleasant": -1.2389, "inadequate": -2.4390, "offensive": -3.6391, "misguided": -1.4392, "lackluster": -2.4393, "stale": -1.2394, "tiresome": -2.4395, "unbearable": -3.8396, "dreadful": -2.6397, "insipid": -1.4398, "monotonous": -2.4399, "repulsive": -3.8400, "depressing": -2.6401, "gloomy": -1.4402, "miserable": -3.8403, "tragic": -2.6404, "heartbreaking": -3.6405, "chaotic": -1.4406, "terribl": -3.8407, "horribl": -3.8408, "disappoint": -2.6409, "bore": -1.6410, "brutal": -3.6411, "violent": -3.6412, "devastating": -3.8413, "ruthless": -2.6414, "sinister": -3.6415, "deadly": -3.6416, "dangerous": -2.4417, "threatening": -2.4418, "destructive": -3.6419, "vile": -3.8420, "twisted": -2.6421, "corrupt": -3.6422, "hostile": -2.4423, "grim": -2.4424, "lethal": -3.6425, "menacing": -2.6426, "savage": -3.8427, "merciless": -3.8428, "vicious": -3.8429, "cruel": -3.6430}
	var horrorWordWeights = map[string]float64{"scary": 2.3341, "frightening": 3.3342, "terrifying": 2.4343, "creepy": 1.2344, "haunting": 3.3345, "disturbing": 2.2346, "intense": 3.3347, "suspenseful": 2.4348, "dark": 1.1349, "grim": 2.1350, "eerie": 3.2351, "spooky": 1.1352, "horrific": 2.3353, "chilling": 3.3354, "sinister": 2.2355, "macabre": 1.2356, "unsettling": 2.2357, "tense": 3.3358, "atmospheric": 1.2359, "brutal": 2.3360, "mysterious": 3.2361, "supernatural": 2.3362, "ghostly": 1.3363, "haunted": 3.4364, "violent": 2.3365, "bloody": 1.3366, "gore": 3.4367, "suspense": 2.3368, "thrilling": 1.3369, "menacing": 3.3370, "monster": 2.3371, "creature": 1.2372, "demon": 3.4373, "zombie": 2.3374, "vampire": 1.3375, "ghost": 3.3376, "evil": 2.4377, "terr": 1.4378, "haunt": 3.3379, "frighten": 2.3380, "chill": 1.3381, "nightmare": 3.4382, "deadly": 2.3383, "panic": 1.3384, "trapped": 3.2385, "lethal": 2.3386, "twisted": 1.3387, "savage": 3.4388, "merciless": 2.4389, "vicious": 1.4390, "lurking": 3.3391, "stalking": 2.3392, "ominous": 1.3393, "devilish": 3.3394, "fiendish": 2.3395, "wicked": 1.3396, "cursed": 3.3397, "demonic": 2.4398, "hellish": 1.4399, "grotesque": 3.3400, "malevolent": 2.3401}

	// Счетчик эмоциональных тонов
	for _, word := range words {
		// Позитивные слова
		if weight, exists := positiveWordWeights[word]; exists {
			analysis.Positive += weight
		}

		// Негативные слова
		if weight, exists := negativeWordWeights[word]; exists {
			analysis.Negative += weight
		}

		// Horrors слова
		if weight, exists := horrorWordWeights[word]; exists {
			analysis.Horror += weight
		}
	}

	// Нормализуем значения относительно количества слов
	analysis.Positive /= float64(len(words))
	analysis.Negative /= float64(len(words))
	analysis.Horror /= float64(len(words))

	// Вычисляем общий тон
	analysis.TotalScore = analysis.Positive - analysis.Negative + (analysis.Horror * 0.5)

	// Возвращение результата
	return analysis
}

// Расчет схожести эмоционального тона
// -----------------------------------
func CalculateToneSimilarity(tone1, tone2 ToneAnalysis) float64 {
	// Вычисляем разницу по каждому тону
	positiveDiff := math.Abs(tone1.Positive-tone2.Positive) * 4
	negativeDiff := math.Abs(tone1.Negative-tone2.Negative) * 4
	horrorDiff := math.Abs(tone1.Horror-tone2.Horror) * 4
	totalDiff := math.Abs(tone1.TotalScore-tone2.TotalScore) * 4

	similarity := 0.0 + ((positiveDiff + negativeDiff + horrorDiff + totalDiff) / 4.0)

	if similarity < 0 {
		return 0.0
	}

	if similarity > 1 {
		return 1.0
	}

	return similarity
}

// Функция для определения преобладающего тона
// -------------------------------------------
func GetDominantTone(analysis ToneAnalysis) string {
	if analysis.Horror > analysis.Positive && analysis.Horror > analysis.Negative {
		return "horror"
	}

	if analysis.Positive > analysis.Negative {
		return "positive"
	}

	return "negative"
}

// Функция для получения процентного соотношения тонов
// ---------------------------------------------------
func GetTonePercentages(analysis ToneAnalysis) map[string]float64 {
	total := analysis.Positive + analysis.Negative + analysis.Horror
	if total == 0 {
		return map[string]float64{
			"positive": 0.0,
			"negative": 0.0,
			"horror":   0.0,
		}
	}

	return map[string]float64{
		"positive": (analysis.Positive / total) * 100,
		"negative": (analysis.Negative / total) * 100,
		"horror":   (analysis.Horror / total) * 100,
	}
}
