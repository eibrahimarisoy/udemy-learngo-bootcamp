package main

import (
	"fmt"
)

func main() {
	text := `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	// 1. Method
	// for i := 0; i < len(text); {
	// 	r, size := utf8.DecodeRuneInString(text[i:])

	// 	fmt.Printf("%c", r)

	// 	i += size
	// } // bu şekilde düzgün çıktı alamayız We must use runes to iterate over a string.

	// 2. Method
	for _, r := range text {
		fmt.Printf("%c", r)
	}
	fmt.Println()

}
