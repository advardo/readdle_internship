package main
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "time"
    "strconv"
)


type Holidays struct {
    Name string
    Date string
}

func main() {
	t := time.Now()
	year := t.Year()
	resp, err := http.Get("https://date.nager.at/api/v2/PublicHolidays/"+strconv.Itoa(year)+"/UA") 
  	if err != nil { 
        fmt.Println(err) 
        return
    }

    defer resp.Body.Close()
    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    	fmt.Println(err)
    }

	var holidays []Holidays
    json.Unmarshal(bytes, &holidays)
    for l := range holidays {   
    	holidayDate, _ := time.Parse("2006-01-02", holidays[l].Date)
    	rounded := time.Date(holidayDate.Year(), holidayDate.Month(), holidayDate.Day(), 0, 0, 0, 0, time.UTC)
   		rounded1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
		days := rounded.Sub(rounded1).Hours() / 24
    	dayOfWeek := int(holidayDate.Weekday())

    	if days < 0 {
    		continue
    	} else if days == 0 {
    		fmt.Printf("Today is ")
    		wday(dayOfWeek, holidays[l].Name, holidays[l].Date)
    		break
    	} else {
    		fmt.Printf("Today is not holliday, but the nearest is ")
    		wday(dayOfWeek, holidays[l].Name, holidays[l].Date)
    		break
    	}
    }
}

func wday (dayOfWeek int, name string, date string) {
	wdate, _ := time.Parse("2006-01-02", date)
	tomorrow := wdate.Add(24 * time.Hour)
	daybefore := wdate.Add(-24 * time.Hour)
	dby := wdate.Add(-48 * time.Hour)
	dat := wdate.Add(48 * time.Hour)

	switch dayOfWeek {
	case 0:
		fmt.Println(name, ",", wdate.Month(), ",", wdate.Day(), "and the weekend will last 3 days:", daybefore.Month(),daybefore.Day(), "-", tomorrow.Month(), tomorrow.Day() )
	case 1:
		fmt.Println(name, ",", wdate.Month(), ",", wdate.Day(), "and the weekend will last 3 days", dby.Month(),dby.Day(), "-", wdate.Month(), wdate.Day())
	case 6:
		fmt.Println(name, ",", wdate.Month(), ",", wdate.Day(), "and the weekend will last 3 days", wdate.Month(),wdate.Day(), "-", dat.Month(), dat.Day())
	case 7: 
		fmt.Println(name, ",", wdate.Month(), ",", wdate.Day(), "and the weekend will last 3 days",daybefore.Month(),daybefore.Day(), "-", dat.Month(), dat.Day())
	default: 
		 fmt.Println(name, ",", wdate.Month(), ",", wdate.Day(),)
	}
}