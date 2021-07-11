package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var russianText = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var foreignText = `Some folks are born made to wave the flag
	They're red, white and blue
	And when the band plays "Hail to the Chief"
	They point the cannon at you, Lord
	It ain't me, it ain't me
	I ain't no senator's son, son
	It ain't me, it ain't me
	I ain't no fortunate one
	Some folks are born silver spoon in hand
	Lord, don't they help themselves, yeah
	But when the taxman comes to the door
	The house look a like a rummage sale
	It ain't me, it ain't me
	I ain't no millionaire's son, no, no
	It ain't me, it ain't me
	I ain't no fortunate one
	Yeah, some folks inherit star-spangled eyes
	They send you down to war
	And when you ask 'em, "How much should we give?"
	They only answer, "More, more, more"
	It ain't me, it ain't me
	I ain't no military son, son
	It ain't me, it ain't me
	I ain't no fortunate one, one
	It ain't me, it ain't me
	I ain't no fortunate one
	It ain't me, it ain't me
	I ain't no fortunate one`
var numText = `1 2 2 3 3 2 3 4 5 5 55 555`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("text with numbers", func(t *testing.T) {
		require.Equal(t, Top10(numText), []string{
			"2",   // 3
			"3",   // 3
			"5",   // 2
			"1",   // 1
			"4",   // 1
			"55",  // 1
			"555", // 1
		})
	})

	t.Run("english text", func(t *testing.T) {
		require.Equal(t, Top10(foreignText), []string{
			"ain't",     // 24
			"no",        // 9
			"I",         // 8
			"It",        // 8
			"it",        // 8
			"me",        // 8
			"me,",       // 8
			"the",       // 6
			"fortunate", // 5
			"one",       // 5
		})
	})

	t.Run("russian test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(russianText))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(russianText))
		}
	})
}
