package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	daysInMonth = map[time.Month]int{
		time.January:   31,
		time.February:  28,
		time.March:     31,
		time.April:     30,
		time.May:       31,
		time.June:      30,
		time.July:      31,
		time.August:    31,
		time.September: 30,
		time.October:   31,
		time.November:  30,
		time.December:  31,
	}
	dayIndex = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	date := randDate()
	ds := date.Format("January 02, 2006")
	answer := strings.ToLower(date.Format("Monday"))
	fmt.Println(ds)
	fmt.Println("Type 'help' for help/info")
	for {
		fmt.Print("\n > ")
		var input string
		fmt.Scanln(&input)
		switch strings.ToLower(input) {
		case answer:
			fmt.Printf("\nCorrect! %s was on a %s%s\n", ds, strings.ToUpper(answer[:1]), answer[1:])
			os.Exit(0)
		case "help":
			fmt.Println(helpText)
		case "days":
			fmt.Println(helpDays)
		case "anchor":
			fmt.Println(helpAnchor)
		case "century_anchor":
			centuryAnchor(date)
		case "year_anchor":
			fmt.Printf("The anchor day for %d is %s\n", date.Year(), anchorDay(date.Year()))
		default:
			fmt.Printf("\n%q is incorrect\n", input)
		}
	}
}

func centuryAnchor(date time.Time) {
	y := date.Year()
	cent := ((y / 100) * 100)
	fmt.Printf("The anchor day for the %d's is %s\n", cent, anchorDay(cent))
}

func anchorDay(y int) string {
	i := 2 + 5*(y%4) + 4*(y%100) + 6*(y%400)
	return dayIndex[i%7]
}

func randDate() time.Time {
	y := rand.Intn(3000)
	m := time.Month(rand.Intn(11) + 1)
	if m == time.February && isLeap(y) {
		daysInMonth[m]++
	}
	d := rand.Intn(daysInMonth[m]-1) + 1
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

const (
	helpText = `
Type 'days' for a list of Doomsdays
Type 'anchor' for help on finding anchors
Type 'century_anchor' to show the century anchor for the date
Type 'year_anchor' to show the doomsday for the date's year`

	helpDays = `
| Month |  Month/Day  |      Mnemonic      |             All days             |
| ----- | ----------- | ------------------ | -------------------------------- |
| Jan   | 1/3, 1/4    | 3rd in 3, 4th in 4 | 3, 10, 17, 24, 31 (+1 for leap)  |
| Feb   | 2/28, 2/29  | last day of Feb    | 0, 7, 14, 21, 28 (+1 for leap)   |
| Mar   | 3/14        | last day of Feb    | 0, 7, 14, 21, 28                 |
| Apr   | 4/4         | evens double       | 4, 11, 18, 25, 32                |
| May   | 5/9         | 9-to-5 at 7-11     | 2, 9, 16, 23, 30                 |
| Jun   | 6/6         | evens double       | 6, 13, 20, 27                    |
| Jul   | 7/11        | 9-to-5 at 7-11     | 4, 11, 18, 25, 32                |
| Aug   | 8/8         | evens double       | 1, 8, 15, 22, 29                 |
| Sep   | 9/5         | 9-to-5 at 7-11     | 5, 12, 19, 26                    |
| Oct   | 10/10       | evens double       | 3, 10, 17, 24, 31                |
| Nov   | 11/7        | 9-to-5 at 7-11     | 0, 7, 14, 21, 28                 |
| Dec   | 12/12       | evens double       | 5, 12, 19, 26                    |`

	helpAnchor = `
Find anchor day for the century:
  Let c = year // 100
  Let r = c % 4
    anchor = 5 * r % 7 + 2
      -OR-
    r=0, Tuesday
    r=1, Sunday
    r=2, Friday
    r=3, Wednesday

Find anchor day for the year:
  Let y = last 2 digits of year
    anchor = ((y // 12) + (y % 12) + (y % 12 // 4)) % 7 + century_anchor
      -OR-
    a = y // 12
    b = y % 12
    c = b // 4
    d = (a + b + c) % 7
    anchor = century_anchor + d

| Century | Anchor |     Mnemonic     |
| ------- | ------ | ---------------- |
|  1600s  |  Tue   |                  |
|  1700s  |  Sun   |                  |
|  1800s  |  Fri   |                  |
|  1900s  |  Wed   | We-in-dis-day    |
|  2000s  |  Tue   | Y-Tue-K          |
|  2100s  |  Sun   | 21-day is Sunday |
|  2200s  |  Fri   |                  |`
)
