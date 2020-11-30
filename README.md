## Teyvat Dev Scraper Go

Microservice scrapers writen in go to aid the automation of teyvat dev's database. Scraping tasks will be run daily.

### Why

With genshin impact updating consistantly there is a need for automation, these scrapers help with that task, although not perfect will assist with the upkeap of up to date information in the API

### Why GO?

It's fast and is well suited for microservices. Also can be used as a decent learning opportunity

### How to contribute

You can contribute by cloning the repo and start hacking away, to test locally you can run the following command and visit the endpoint of the desired scraped page.

```
go run cmd/main.go

and visit
localhost:8080/<scraper to run>

eg. 
localhost:8080/WikiScrapeCharacters
```

I am also super greatful for any tips and tricks for go development since I am very new to it!

### Completion List

- [ ] Artifacts
- [ ] ArtifactSets
- [ ] - Characters **Current Development**
- [ ] CharacterAscensions
- [ ] CharacterAscensionMaterials
- [ ] CharacterProfiles **Current Development**
- [ ] CommonAscensionMaterials
- [ ] CommonMaterials
- [ ] Consumeables
- [ ] ConsumeableRecipes
- [ ] CookingMaterials
- [ ] CraftingMaterials
- [ ] Domains
- [ ] Elements
- [ ] ForgeRecipes
- [ ] Regions
- [ ] Talents **Current Development**
- [ ] TalentLevelUpMaterials
- [ ] Weapons
- [ ] WeaponAscensions
- [ ] WeaponAscensionMaterials
- [ ] WeaponEnhancementMaterials
