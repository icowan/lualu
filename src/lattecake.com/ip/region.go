package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

type Citys []City

type City struct {
	Code string
	Name string
}

func main() {
	fileName := "/Users/dudulu/Workspace/GitHub/go/src/lattecake.com/ip/region.txt"
	fl, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fileName, err)
		return
	}

	defer fl.Close()

	sql := []string{}

	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		py := pinyin.NewArgs()
		//py.Heteronym = true

		spell := ""
		spell_all := ""

		for _, item := range pinyin.Pinyin(scanner.Text()[6:], py) {
			spell = spell + item[0][:1]
			spell_all = spell_all + item[0] + " "
		}
		spell = strings.ToUpper(spell)
		spell_all = strings.Trim(spell_all, " ")
		if scanner.Text()[2:6] == "0000" {
			sql = append(sql, "INSERT INTO `citys` (`code`, `name`, `province`, `city`, `county`, `spell`, `spell_all`) VALUES ("+scanner.Text()[:6]+", '"+scanner.Text()[6:]+"', "+scanner.Text()[:2]+"0000, 0, 0, '"+spell+"', '"+spell_all+"');")
			continue
		}
		if scanner.Text()[4:6] == "00" {
			sql = append(sql, "INSERT INTO `citys` (`code`, `name`, `province`, `city`, `county`, `spell`, `spell_all`) VALUES ("+scanner.Text()[:6]+", '"+scanner.Text()[6:]+"', "+scanner.Text()[:2]+"0000, "+scanner.Text()[0:4]+"00, 0, '"+spell+"', '"+spell_all+"');")
			continue
		}
		sql = append(sql, "INSERT INTO `citys` (`code`, `name`, `province`, `city`, `county`, `spell`, `spell_all`) VALUES ("+scanner.Text()[:6]+", '"+scanner.Text()[6:]+"', "+scanner.Text()[:2]+"0000, "+scanner.Text()[0:4]+"00, "+scanner.Text()[:6]+", '"+spell+"', '"+spell_all+"');")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("/Users/dudulu/Workspace/GitHub/go/src/lattecake.com/ip/region.sql")
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range sql {
		fmt.Println(item)
		n3, err := f.WriteString(item + "\n")
		if err != nil {
			fmt.Printf("wrote %d bytes\n", n3)
		}
	}
	f.Sync()

	fmt.Println(len(sql))

}

func (t Citys) Len() int { return len(t) }
func (t Citys) Less(i, j int) bool {
	return t[i].Code > t[i].Code
}
func (t Citys) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
