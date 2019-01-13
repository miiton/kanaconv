package kanaconv

import (
	"testing"
)

func TestHiraganaToKatakana(t *testing.T) {
	var result string
	result = HiraganaToKatakana("ひらがなをカタカナにへんかんします")
	if result != "ヒラガナヲカタカナニヘンカンシマス" {
		t.Errorf("fail! %s", result)
	}
	result = HiraganaToKatakana("ゔぁいおりん")
	if result != "ヴァイオリン" {
		t.Errorf("fail! %s", result)
	}
}

func BenchmarkHiraganaToKatakana(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HiraganaToKatakana("ひらがなをカタカナにへんかんします")
	}
}

func TestKatakanaToHiragana(t *testing.T) {
	var result string
	result = KatakanaToHiragana("カタカナヲひらがなニヘンカンシマス")
	if result != "かたかなをひらがなにへんかんします" {
		t.Errorf("fail! %s", result)
	}
	result = KatakanaToHiragana("ヴァイオリン")
	if result != "ゔぁいおりん" {
		t.Errorf("fail! %s", result)
	}
}

func BenchmarkKatakanaToHiragana(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KatakanaToHiragana("カタカナヲひらがなニヘンカンシマス")
	}
}

func TestHankakuToZenkaku(t *testing.T) {
	var result string
	result = HankakuToZenkaku("ﾊﾝｶｸｦ ｾﾞﾝｶｸﾆ ﾍﾝｶﾝｼﾏｽ0123456789")
	if result != "ハンカクヲ　ゼンカクニ　ヘンカンシマス０１２３４５６７８９" {
		t.Errorf("fail! %s", result)
	}
}

func BenchmarkHankakuToZenkaku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HankakuToZenkaku("ﾊﾝｶｸｦｾﾞﾝｶｸﾆﾍﾝｶﾝｼﾏｽ")
	}
}

func TestZenkakuToHankaku(t *testing.T) {
	var result string
	result = ZenkakuToHankaku("ゼンカクヲ　ハンカクニ　ヘンカンシマス０１２３４５６７８９")
	if result != "ｾﾞﾝｶｸｦ ﾊﾝｶｸﾆ ﾍﾝｶﾝｼﾏｽ0123456789" {
		t.Errorf("fail! %s", result)
	}
}

func BenchmarkZenkakuToHankaku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZenkakuToHankaku("ゼンカクヲハンカクニヘンカンシマス")
	}
}

func TestSmartConv(t *testing.T) {
	var result string
	result = SmartConv("ｶﾀｶﾅは全角に統一ＥＩＳＵＵＪＩ＋＝−，記号は　半角に 統一します")
	if result != "カタカナは全角に統一EISUUJI+=-,記号は 半角に 統一します" {
		t.Errorf("fail! %s", result)
	}
}

func BenchmarkSmartConv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SmartConv("ｶﾀｶﾅは全角に統一ＥＩＳＵＵＪＩ＋＝−，記号は半角に統一します")
	}
}
