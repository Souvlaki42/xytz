package config

func GetDefault() *Config {
	return &Config{
		SearchLimit:         25,
		DefaultDownloadPath: "~/Videos",
		DefaultQuality:      "best",
		SortByDefault:       "relevance",
		EmbedSubtitles:      false,
		EmbedMetadata:       true,
		EmbedChapters:       true,
		CookiesBrowser:      "",
		CookiesFile:         "",
	}
}
