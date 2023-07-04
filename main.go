package main

import (
	"fmt"
	"time"
)

var dolar int = 26
var euro int = 28
var bakiye int = 0
var krediMiktari int = 0
var dolarMiktari int = 0
var euroMiktari int = 0

func paraYatir() { // 1
	fmt.Println("Yatırmak istediğiniz tutarı giriniz : ")
	var yatirilanPara int
	fmt.Scanln(&yatirilanPara)
	bakiye += yatirilanPara
	fmt.Println("İşleminiz Yapılıyor...")
	time.Sleep(3 * time.Second)
	fmt.Println("Yeni bakiyeniz : ", bakiye)
	time.Sleep(1 * time.Second)

}

func Kredi() { // 2
	if bakiye < 3000 {
		fmt.Println("Kredi çekebilmek için en az 3000 TL bakiyeniz olmalıdır ! ")
		time.Sleep(1 * time.Second)
	} else if bakiye >= 3000 {
		fmt.Println("Çekmek İstediğiniz Kredi Tutarını Giriniz : ")
		fmt.Scanln(&krediMiktari)
		bakiye += krediMiktari
		fmt.Println("Kredi Tutarınız Hesabınıza Yatırıldı ! ")
		time.Sleep(1 * time.Second)
		fmt.Println("Yeni bakiyeniz : ", bakiye)
		time.Sleep(1 * time.Second)
	}

}
func krediSorgula() { // 3
	if krediMiktari == 0 {
		fmt.Println("Aktif Bir Kredi Borcunuz Bulunmamaktadır.")
		time.Sleep(1 * time.Second)
	} else if krediMiktari > 0 {
		fmt.Println("Aktif Kredi Borcunuz : ", krediMiktari)
		time.Sleep(1 * time.Second)
	}
}

func krediOde() { // 4
	if krediMiktari <= 0 {
		fmt.Println("Aktif bir kredi borcunuz bulunmamaktadır.")
		time.Sleep(1 * time.Second)
	} else if krediMiktari > 0 {
		fmt.Println("Ödemek istediğiniz kredi tutarını giriniz : ")
		var krediOdenen int
		fmt.Scanln(&krediOdenen)
		if krediOdenen > bakiye {
			fmt.Println("Yetersiz Bakiye ! ")
			time.Sleep(1 * time.Second)
		} else if krediOdenen <= bakiye {
			if krediOdenen > krediMiktari {
				fmt.Println("Lütfen Geçerli Bir Değer Giriniz ! ")
				time.Sleep(1 * time.Second)
			} else if krediOdenen <= krediMiktari {
				krediMiktari -= krediOdenen
				bakiye -= krediOdenen
				fmt.Println("İşleminiz Gerçekleştiriliyor...")
				time.Sleep(3 * time.Second)
				fmt.Println("İşleminiz Başarıyla Gerçekleştirildi !")
				time.Sleep(1 * time.Second)
				fmt.Println("Kalan Kredi Tutarınız : ", krediMiktari)
			}
		}
	}
}

func bozdurDolar() {
	if dolarMiktari == 0 {
		fmt.Println("Yetersiz Bakiye ! ")
		time.Sleep(1 * time.Second)
	} else if dolarMiktari > 0 {
		fmt.Println("Bozdurmak istediğiniz dolar miktarını giriniz : ")
		var bozdurulanDolar int
		fmt.Scanln(&bozdurulanDolar)
		if bozdurulanDolar > dolarMiktari {
			fmt.Println("Yetersiz Bakiye ! ")
		} else if bozdurulanDolar <= dolarMiktari {
			dolarMiktari -= bozdurulanDolar
			bakiye += bozdurulanDolar * dolar
			fmt.Println(bozdurulanDolar, "dolar bozduruldu")
			time.Sleep(1 * time.Second)
			fmt.Println("Yeni Bakiyeniz : ", bakiye)
			time.Sleep(1 * time.Second)
		}
	}
}

func bozdurEuro() {
	if euroMiktari == 0 {
		fmt.Println("Yetersiz Bakiye ! ")
		time.Sleep(1 * time.Second)
	} else if euroMiktari > 0 {
		fmt.Println("Bozdurmak istediğiniz euro miktarını giriniz : ")
		var bozdurulanEuro int
		fmt.Scanln(&bozdurulanEuro)
		if bozdurulanEuro > euroMiktari {
			fmt.Println("Yetersiz Bakiye ! ")
			time.Sleep(1 * time.Second)
		} else if bozdurulanEuro <= euroMiktari {
			euroMiktari -= bozdurulanEuro
			bakiye += bozdurulanEuro * euro
			fmt.Println(bozdurulanEuro, "euro bozduruldu.")
			time.Sleep(1 * time.Second)
			fmt.Println("Yeni bakiyeniz : ", bakiye)
			time.Sleep(1 * time.Second)
		}
	}
}

func paraCek() { // 5
	fmt.Println("Çekmek İstediğiniz Para Miktarını Giriniz: ")
	var cekilenPara int
	fmt.Scanln(&cekilenPara)
	if bakiye < cekilenPara {
		fmt.Println("Bakiyeniz Yetersiz ! ")
		time.Sleep(1 * time.Second)
	} else if bakiye >= cekilenPara {
		bakiye -= cekilenPara
		fmt.Println("İşleminiz Başarılı ! ")
		time.Sleep(1 * time.Second)
		fmt.Println("Hesabınızdan Çekilen Tutar : ", cekilenPara)
		time.Sleep(1 * time.Second)
		fmt.Println("Hesabınızda Kalan Tutar : ", bakiye)
	}
}

func satinalEuro() { // 6
	fmt.Println("Satın almak istediğiniz euro miktarını giriniz : ")
	var istenenEuro int
	fmt.Scanln(&istenenEuro)
	var gidenEuro int
	gidenEuro = istenenEuro * euro
	if bakiye < gidenEuro {
		fmt.Println("Bakiyeniz Yetersiz ! ")
		time.Sleep(1 * time.Second)
	} else if bakiye >= gidenEuro {
		bakiye -= gidenEuro
		euroMiktari += istenenEuro
		fmt.Println("İşleminiz Başarılı ! ")
		time.Sleep(1 * time.Second)
		fmt.Println("Hesabınızdaki Euro Miktarı : ", istenenEuro)
		time.Sleep(1 * time.Second)
		fmt.Println("Hesabınızda kalan para miktarı : ", bakiye)
		time.Sleep(1 * time.Second)
	}
}

func paraGonder() { // 7
	if bakiye <= 0 {
		fmt.Println("Yetersiz Bakiye")
	} else if bakiye > 0 {
		fmt.Println("Göndermek istediğiniz para miktarını giriniz : ")
		var gondermekistenenMiktar int
		fmt.Scanln(&gondermekistenenMiktar)
		bakiye -= gondermekistenenMiktar
		time.Sleep(1 * time.Second)
		fmt.Println("Kime Göndermek İstiyorsunuz ?")
		var gondermekistenenKisi string
		fmt.Scanln(&gondermekistenenKisi)
		fmt.Println("İşleminiz Gerçekleştiriliyor...")
		time.Sleep(1 * time.Second)
		fmt.Println("İşleminiz Gerçekleştiriliyor...")
		time.Sleep(1 * time.Second)
		fmt.Println("İşleminiz Gerçekleştiriliyor...")
		time.Sleep(1 * time.Second)
		fmt.Println("İşleminiz Gerçekleştiriliyor...")
		time.Sleep(1 * time.Second)
		fmt.Println("İşleminiz Gerçekleştiriliyor...")
		time.Sleep(1 * time.Second)
		fmt.Println(gondermekistenenKisi, "adlı kişiye", gondermekistenenMiktar, "TL para gönderilmiştir !")
		time.Sleep(1 * time.Second)
		fmt.Println("Yeni Bakiyeniz : ", bakiye)
		time.Sleep(1 * time.Second)
	}
}

func satinalDolar() { // 8
	fmt.Println("Satın almak istediğiniz dolar miktarını giriniz : ")
	var istenenDolar int
	fmt.Scanln(&istenenDolar)
	var gidenDolar int
	gidenDolar = istenenDolar * dolar
	if bakiye < gidenDolar {
		fmt.Println("Bakiyeniz yetersiz !")
		time.Sleep((1 * time.Second))
	} else if bakiye >= gidenDolar {
		bakiye -= gidenDolar
		dolarMiktari += istenenDolar
		fmt.Println("İşleminiz Başarıyla Gerçekleşti !")
		time.Sleep(1 * time.Second)
		fmt.Println("Hesabınızdaki Yeni TL miktarı : ", bakiye)
		time.Sleep(1 * time.Second)
		fmt.Println("Hesabınızdaki yeni Dolar miktarı : ", istenenDolar)
	}
}

func sorgula() {
	fmt.Println("Hesabınızdaki TL Miktarı : ", bakiye)
	time.Sleep(1 * time.Second)
	fmt.Println("Hesabınızdaki Dolar Miktarı : ", dolarMiktari)
	time.Sleep(1 * time.Second)
	fmt.Println("Hesabınızdaki Euro Miktarı : ", euroMiktari)
	time.Sleep(3 * time.Second)
}

func main() {
	for {
		fmt.Println("Yapmak istediğiniz işlem : ")
		fmt.Println("Para Yatırmak İçin '1'")
		fmt.Println("Kredi Çekmek İçin '2'")
		fmt.Println("Kredi Sorgulamak İçin '3'")
		fmt.Println("Kredi Ödemek İçin '4'")
		fmt.Println("Para Çekmek İçin '5'")
		fmt.Println("Euro Satın Almak İçin '6'")
		fmt.Println("Para Göndermek İçin '7'")
		fmt.Println("Dolar Satın Almak İçin '8'")
		fmt.Println("Çıkış Yapmak İçin 'Q'")
		fmt.Println("Dolar Bozdurmak İçin '9'")
		fmt.Println("Euro Bozdurmak İçin '10'")
		fmt.Println("Hesabınızdaki Para Türlerini Görüntülemek İçin '11'")
		var secim string
		fmt.Scanln(&secim)
		if secim == "1" {
			paraYatir()
		} else if secim == "2" {
			Kredi()
		} else if secim == "3" {
			krediSorgula()
		} else if secim == "4" {
			krediOde()
		} else if secim == "5" {
			paraCek()
		} else if secim == "6" {
			satinalEuro()
		} else if secim == "7" {
			paraGonder()
		} else if secim == "8" {
			satinalDolar()
		} else if secim == "Q" {
			break
		} else if secim == "9" {
			bozdurDolar()
		} else if secim == "10" {
			bozdurEuro()
		} else if secim == "11" {
			sorgula()
		}
	}
}
