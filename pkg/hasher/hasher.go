package hasher

import (
	"strings"

	"github.com/catinello/base62"
)
const URL_LENGTH = 10
/*
Описание алгоритма:
Есть два варианта, которыми можно решить задачу. Первый вариант полный URL провести сквозь hash функцию, но мы столкнемся с коллизиями, когда две разные ссылки
приведут к одному и тому же хэшу. Второй вариант (который выбрал я) это взять id ссылки в базе и заэнкодить в base62(количество латинских букв в верхнем и нижнем регистре + цифры). Если в результате мы получим строку, длиной меньше 10, то можем добавить "_" подчеркивание, пока длина не станет равна 10.
*/
func HashURL(id int) string {
	hashURL := base62.Encode(id)	
	// Так как у нас base62, то каждый символ равен 1 байту, следовательно мы можем сравнивать длину строки с нужной длиной(10) и добавлять подчеркивание, в качестве пустого символа.
	if len(hashURL) < 10 {
		count := 10 - len(hashURL)
		emptyString := strings.Repeat("_", count)
		hashURL += emptyString
	}
	return hashURL
}