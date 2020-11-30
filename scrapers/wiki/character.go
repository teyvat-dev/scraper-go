package wikicharacterscraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/imdario/mergo"
	wikicharacterscrapertypes "teyvat.dev/scraper-go/scrapers/wiki/types"
)

// Scrape character data from wiki
func Scrape(w http.ResponseWriter, r *http.Request) string {
	// characters := make([]*wikicharacterscrapertypes.CharacterPrisma, 0)
	characterTableInfos := make([]*wikicharacterscrapertypes.CharacterTableInfo, 0)
	characterProfileInfos := make([]*wikicharacterscrapertypes.CharacterProfileInfo, 0)

	// Instantiate default collector
	tableCollector := colly.NewCollector(
		colly.Async(true),
	)

	tableCollector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	profileCollector := tableCollector.Clone()

	storyCollector := tableCollector.Clone()

	voicelinesCollector := tableCollector.Clone()

	// Extract Data from table
	tableCollector.OnHTML(".article-table tbody tr", func(e *colly.HTMLElement) {
		name := e.ChildText("td:nth-child(3) a")

		if name == "" {
			return
		}

		rarity, rarityErr := strconv.Atoi(e.ChildText("td:nth-child(1)"))

		if rarityErr != nil {
			return
		}

		temp := &wikicharacterscrapertypes.CharacterTableInfo{
			Rarity:  rarity,
			Image:   e.ChildAttr("td:nth-child(2) a", "href"),
			Name:    name,
			Element: e.ChildText("td:nth-child(4) a:last-of-type"),
			Weapon:  e.ChildText("td:nth-child(5) a:last-of-type"),
			Sex:     e.ChildText("td:nth-child(6)"),
			Nation:  e.ChildText("td:nth-child(7) a:first-of-type"),
		}

		link := fmt.Sprintf("https://genshin-impact.fandom.com%v", e.ChildAttr("td:nth-child(3) a", "href"))

		profileCollector.Visit(link)

		storyCollector.Visit(fmt.Sprintf("%v/Story", link))

		voicelinesCollector.Visit(fmt.Sprintf("%v/Voicelines", link))

		characterTableInfos = append(characterTableInfos, temp)
	})

	profileCollector.OnHTML("div .WikiaPageContentWrapper", func(e *colly.HTMLElement) {

		name := e.DOM.Find("h1#firstHeading").Text()
		image, _ := e.DOM.Find("div#pi-tab-0 img.pi-image-thumbnail").Attr("src")
		introduction := e.DOM.Find("h3 span#Introduction").Parent().Next().Text()
		personality := e.DOM.Find("h3 span#Personality").Parent().Next().Text()

		bio := e.DOM.Find("div.pi-section-content[data-ref=\"0\"]")
		birthday := bio.Find("div.pi-item[data-source=\"birthday\"] div").Text()
		constellation := bio.Find("div.pi-item[data-source=\"constellation\"] div").Text()
		affiliation := bio.Find("div.pi-item[data-source=\"affiliation\"] div").Text()
		dish := bio.Find("div.pi-item[data-source=\"dish\"] div").Text()

		voiceActors := e.DOM.Find("div.pi-section-content[data-ref=\"1\"]")
		voiceEN := voiceActors.Find("div.pi-item[data-source=\"voiceEN\"] div").Text()
		voiceCN := voiceActors.Find("div.pi-item[data-source=\"voiceCN\"] div").Text()
		voiceJP := voiceActors.Find("div.pi-item[data-source=\"voiceJP\"] div").Text()
		voiceKR := voiceActors.Find("div.pi-item[data-source=\"voiceKR\"] div").Text()

		talentTable := e.DOM.Find("table.wikitable:nth-of-type(1)")

		// TODO: Talents have their own page that can be scraped for skill attributes!
		// https://genshin-impact.fandom.com/wiki/Sharpshooter
		talents := make([]*wikicharacterscrapertypes.CharacterTalent, 0)
		talentTable.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {
			if selection.Find("td:nth-of-type(2)").Text() == "None" {
				return
			}

			if len(selection.Find("td:nth-of-type(1)").Nodes) == 0 {
				return
			}

			reg, err := regexp.Compile("[^a-zA-Z]+")
			if err != nil {
				log.Fatal(err)
			}

			Type := reg.ReplaceAllString(strings.ReplaceAll(strings.Split(selection.Find("td:nth-of-type(1)").Text(), "-")[0], " ", ""), "")
			Name := selection.Find("td:nth-of-type(2)").Text()
			Icon, _ := selection.Find("td:nth-of-type(3) a img").Attr("data-src") // Do we want to catch this err?
			Info := selection.Find("td:nth-of-type(4)").Text()                    // Can be parsed better

			temp := &wikicharacterscrapertypes.CharacterTalent{
				Type: Type,
				Name: Name,
				Icon: Icon,
				Info: Info,
			}

			talents = append(talents, temp)
		})
		// Level Mats
		// levelMaterialTable := e.DOM.Find("table.wikitable:nth-of-type(2)")
		constellationsTable := e.DOM.Find("table.wikitable:nth-of-type(3)")
		constellations := make([]*wikicharacterscrapertypes.CharacterConstellation, 0)
		constellationsTable.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {
			if len(selection.Find("td:nth-of-type(1)").Nodes) == 0 {
				return
			}

			Level, _ := strconv.Atoi(strings.ReplaceAll(selection.Find("th:nth-of-type(1)").Text(), "\n", ""))
			Name := selection.Find("td:nth-of-type(1)").Text()
			Effect := selection.Find("td:nth-of-type(2)").Text() // Can be parsed better

			temp := &wikicharacterscrapertypes.CharacterConstellation{
				Level:  Level,
				Name:   Name,
				Effect: Effect,
			}

			constellations = append(constellations, temp)
		})

		// Ascensions // TODO
		// ascensionsTable := e.DOM.Find("table.wikitable:nth-of-type(4)")
		// Stats (dont grab this from wiki) maybe?

		characterProfileInfos = append(characterProfileInfos, &wikicharacterscrapertypes.CharacterProfileInfo{
			Name:           name,
			Image:          image,
			Introduction:   introduction,
			Personality:    personality,
			Birthday:       birthday,
			Constellation:  constellation,
			Affiliation:    affiliation,
			Dish:           dish,
			VoiceEN:        voiceEN,
			VoiceCN:        voiceCN,
			VoiceJP:        voiceJP,
			VoiceKR:        voiceKR,
			Talents:        talents,
			Constellations: constellations,
		})
	})

	storyCollector.OnHTML("div .WikiaPageContentWrapper", func(e *colly.HTMLElement) {
		// Scrape Storys
	})

	voicelinesCollector.OnHTML("div .WikiaPageContentWrapper", func(e *colly.HTMLElement) {
		// Scrape Voice Lines
	})

	// tableCollector.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("TableCollector: Visiting", r.URL.String())
	// })

	// profileCollector.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("ProfileCollector: Visiting", r.URL.String())
	// })

	// storyCollector.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("StoryCollector: Visiting", r.URL.String())
	// })

	// voicelinesCollector.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("VoicelinesCollector: Visiting", r.URL.String())
	// })

	tableCollector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	profileCollector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	storyCollector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	voicelinesCollector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Init scraping
	tableCollector.Visit("https://genshin-impact.fandom.com/wiki/Characters/List#")

	// Wait until threads are finished
	tableCollector.Wait()
	profileCollector.Wait()
	storyCollector.Wait()
	voicelinesCollector.Wait()

	characters := make(wikicharacterscrapertypes.CharacterPrismaPayload)

	for _, char := range characterTableInfos {
		characters[char.Name] = wikicharacterscrapertypes.CharacterPrisma{
			Name:   char.Name,
			Rarity: char.Rarity,
			CharacterProfile: wikicharacterscrapertypes.CharacterProfilePrisma{
				Region: char.Nation,
				Vision: char.Element,
			},
			Elements: []string{char.Element},
			Weapon:   char.Weapon,
		}
	}

	for _, char := range characterProfileInfos {
		parsedTalents := make([]*wikicharacterscrapertypes.CharacterTalentsPrisma, 0)

		for _, talent := range char.Talents {
			parsedTalents = append(parsedTalents, &wikicharacterscrapertypes.CharacterTalentsPrisma{
				Name:        talent.Name,
				Description: talent.Info,
				Type:        talent.Type,
			})
		}
		temp := wikicharacterscrapertypes.CharacterPrisma{
			Name:           char.Name,
			Constellations: char.Constellations,
			CharacterProfile: wikicharacterscrapertypes.CharacterProfilePrisma{
				Affiliation:   char.Affiliation,
				Birthday:      char.Birthday,
				Constellation: char.Constellation,
				SpecialtyDish: char.Dish,
				Overview:      char.Personality,
				// Story: char.Story,
				// VoiceLines: char.VoiceLines,
				VoiceActor: wikicharacterscrapertypes.CharacterProfileVoiceActorPrisma{
					EN: char.VoiceEN,
					CN: char.VoiceCN,
					JP: char.VoiceJP,
					KR: char.VoiceKR,
				},
			},
			Overview: char.Introduction,
			Talents:  parsedTalents,
		}

		mergo.Merge(&temp, characters[char.Name])

		characters[char.Name] = temp
	}

	b, err := json.Marshal(characters)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", string(b))
}
