package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// 都道府県ごとのCityIDマップ
type CityIDMap map[string]string

// 都道府県のCityIDのマップを格納する
type PrefectureMap map[string]CityIDMap

// 都道府県ごとのCityIDを読み込む関数
func loadCityIDs(filePath string) (PrefectureMap, error) {
	// JSONファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ファイルを開くことができません: %v", err)
	}
	defer file.Close()

	// ファイルの内容を読み取る
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ファイルの読み取りエラー: %v", err)
	}

	// JSONデータをマップにデコード
	var prefectureMap PrefectureMap
	err = json.Unmarshal(bytes, &prefectureMap)
	if err != nil {
		return nil, fmt.Errorf("JSONのデコードエラー: %v", err)
	}

	return prefectureMap, nil
}

// 天気予報の構造体
type Forecast struct {
	PublicTime       string `json:"publicTimeFormatted"`
	PublishingOffice string `json:"publishingOffice"`
	Title            string `json:"title"`
	Description      struct {
		BodyText string `json:"bodyText"`
		Text     string `json:"text"`
	} `json:"description"`
	Forecasts []struct {
		Date      string `json:"date"`
		DateLabel string `json:"dateLabel"`
		Telop     string `json:"telop"`
		Detail    struct {
			Weather string `json:"weather"`
			Wind    string `json:"wind"`
			Wave    string `json:"wave"`
		} `json:"detail"`
		Temperature struct {
			Min struct {
				Celsius string `json:"celsius"`
			} `json:"min"`
			Max struct {
				Celsius string `json:"celsius"`
			} `json:"max"`
		} `json:"temperature"`
		ChanceOfRain struct {
			T00_06 string `json:"T00_06"`
			T06_12 string `json:"T06_12"`
			T12_18 string `json:"T12_18"`
			T18_24 string `json:"T18_24"`
		} `json:"chanceOfRain"`
		Image struct {
			Title string `json:"title"`
		} `json:"image"`
	} `json:"forecasts"`
	Copyright struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"copyright"`
}

func main() {
	// city_ids.jsonファイルを読み込む
	prefectureMap, err := loadCityIDs("./city_ids.json")
	if err != nil {
		log.Fatalf("CityIDの読み込みエラー: %v", err)
	}

	// 都道府県名を入力してもらう
	fmt.Print("APIのリクエスト元の状況により一部nullとなってしまい、表示できない可能性があります。\n")
	fmt.Println("検索したい都道府県を入力してください：")
	var prefecture string
	fmt.Scanln(&prefecture)

	// 入力された都道府県が存在するか確認
	prefecture = strings.TrimSpace(prefecture)
	cityMap, exists := prefectureMap[prefecture]
	if !exists {
		log.Fatalf("%sという都道府県は見つかりません。", prefecture)
	}

	// 都道府県内の市区町村を表示
	fmt.Printf("%s の市区町村:\n", prefecture)
	for cityName, cityID := range cityMap {
		fmt.Printf("%s: %s\n", cityName, cityID)
	}

	// 市区町村名を入力してもらう
	fmt.Println("検索したい市区町村を市区町村名で入力してください：")
	var city string
	fmt.Scanln(&city)

	// 入力された市区町村名に対応するIDを取得
	city = strings.TrimSpace(city)
	cityID, exists := cityMap[city]
	if !exists {
		log.Fatalf("%sという市区町村は見つかりません。", city)
	}

	// IDをURLに入れてAPIリクエストを送る
	url := fmt.Sprintf("https://weather.tsukumijima.net/api/forecast/city/%s", cityID)
	fmt.Printf("リクエストURL: %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("HTTPリクエストエラー: %s", err)
	}
	defer response.Body.Close()

	// レスポンスのボディを読み込む
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("レスポンス読み取りエラー: %s", err)
	}

	// JSONデータを構造体にデコード
	var forecastData Forecast
	if err := json.Unmarshal(body, &forecastData); err != nil {
		log.Fatalf("JSONのデコードエラー: %s", err)
	}

	// 予報の発表日と発表気象台を表示
	fmt.Printf("予報の発表日: %s\n", forecastData.PublicTime)
	fmt.Printf("予報発表気象台: %s\n", forecastData.PublishingOffice)

	// 今日と明日の天気予報を分ける
	var todayForecast, tomorrowForecast *struct {
		Date      string `json:"date"`
		DateLabel string `json:"dateLabel"`
		Telop     string `json:"telop"`
		Detail    struct {
			Weather string `json:"weather"`
			Wind    string `json:"wind"`
			Wave    string `json:"wave"`
		} `json:"detail"`
		Temperature struct {
			Min struct {
				Celsius string `json:"celsius"`
			} `json:"min"`
			Max struct {
				Celsius string `json:"celsius"`
			} `json:"max"`
		} `json:"temperature"`
		ChanceOfRain struct {
			T00_06 string `json:"T00_06"`
			T06_12 string `json:"T06_12"`
			T12_18 string `json:"T12_18"`
			T18_24 string `json:"T18_24"`
		} `json:"chanceOfRain"`
		Image struct {
			Title string `json:"title"`
		} `json:"image"`
	}
	for _, forecast := range forecastData.Forecasts {
		if forecast.DateLabel == "今日" {
			todayForecast = &forecast
		}
		if forecast.DateLabel == "明日" {
			tomorrowForecast = &forecast
		}
	}

	// 今日の天気予報を表示
	fmt.Println("\n今日の天気予報:")
	printWeather(todayForecast)

	// 明日の天気予報を表示
	fmt.Println("\n明日の天気予報:")
	printWeather(tomorrowForecast)

	// 著作権情報を表示
	fmt.Printf("\n%s\n", forecastData.Copyright.Title)
	fmt.Printf("詳細はこちら: %s\n", forecastData.Copyright.Link)
}

// 天気予報を表示する関数
func printWeather(forecast *struct {
	Date      string `json:"date"`
	DateLabel string `json:"dateLabel"`
	Telop     string `json:"telop"`
	Detail    struct {
		Weather string `json:"weather"`
		Wind    string `json:"wind"`
		Wave    string `json:"wave"`
	} `json:"detail"`
	Temperature struct {
		Min struct {
			Celsius string `json:"celsius"`
		} `json:"min"`
		Max struct {
			Celsius string `json:"celsius"`
		} `json:"max"`
	} `json:"temperature"`
	ChanceOfRain struct {
		T00_06 string `json:"T00_06"`
		T06_12 string `json:"T06_12"`
		T12_18 string `json:"T12_18"`
		T18_24 string `json:"T18_24"`
	} `json:"chanceOfRain"`
	Image struct {
		Title string `json:"title"`
	} `json:"image"`
}) {
	if forecast == nil {
		fmt.Println("予報データがありません")
		return
	}

	// 予報データを表示
	fmt.Printf("日付: %s\n", forecast.Date)
	fmt.Printf("天気: %s\n", forecast.Telop)
	fmt.Printf("詳細: %s\n", forecast.Detail.Weather)
	fmt.Printf("風: %s\n", forecast.Detail.Wind)
	fmt.Printf("波: %s\n", forecast.Detail.Wave)
	fmt.Printf("最低気温: %s℃\n", forecast.Temperature.Min.Celsius)
	fmt.Printf("最高気温: %s℃\n", forecast.Temperature.Max.Celsius)
	fmt.Print("降水確率:\n")
	fmt.Printf(" 0～6時: %s\n", forecast.ChanceOfRain.T00_06)
	fmt.Printf(" 6～12時: %s\n", forecast.ChanceOfRain.T06_12)
	fmt.Printf(" 12～18時: %s\n", forecast.ChanceOfRain.T12_18)
	fmt.Printf(" 18～24時: %s\n", forecast.ChanceOfRain.T18_24)
}
