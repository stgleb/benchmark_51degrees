package main

import (
	"github.com/admanicpv/51degrees"
	"log"
)

var (
	FiftyoneDegreesProvider *fiftyonedegrees.FiftyoneDegreesProvider
	Properties              string = `Id, BrowserName, BrowserVersion, LayoutEngine, PlatformName, PlatformVendor, PlatformVersion, IsMobile,
		 IsCrawler, Html5, DeviceType, HardwareFamily, HardwareModel, HardwareName, HardwareVendor`
	userAgents = []string{
		"User-Agent Mozilla/5.0 (iPhone CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/57.0.2987.137 Mobile/14E304 Safari/602.1",
		"User-Agent Mozilla/5.0 (Windows NT 6.1; WOW64) SkypeUriPreview Preview/0.5", "Firefox", "Mac OS X",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.76 Safari/537.36",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:48.0) Gecko/20100101 Firefox/48.0",
	}
)

func InitFiftyonedegrees(FiftyonedegreesDbPath, properties string, poolSize, cacheSize int) (*fiftyonedegrees.FiftyoneDegreesProvider, error) {
	var err error
	FiftyoneDegreesProvider, err = fiftyonedegrees.NewFiftyoneDegreesProvider(FiftyonedegreesDbPath, properties, poolSize, cacheSize) //config.PackagePath + "/" + strings.Trim(config.Config.Fiftyonedegrees.FiftyonedegreesDbPath, "/"))
	if err != nil {
		log.Printf("51Degrees FiftyoneDegreesProvider error: %v", err.Error())
		return nil, err
	}

	return FiftyoneDegreesProvider, nil
}

func UserAgentDetectByString(ua string) string {
	uaJSON := FiftyoneDegreesProvider.Parse(ua)
	return uaJSON
}

func main() {
	var err error

	FiftyoneDegreesProvider, err = InitFiftyonedegrees("51Degrees-PremiumV3_2.dat", Properties, 16, 2)

	if err != nil {
		log.Fatal("51Degrees init failed! Terminating.")
	}

	for i := 0; i < 1000000000; i++ {
		go func() string {
			return UserAgentDetectByString(userAgents[i%5])
		}()
	}
}
