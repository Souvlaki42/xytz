package config

type QualityPreset struct {
	Name        string
	Description string
	Format      string
}

var QualityPresets = []QualityPreset{
	{
		Name:        "best",
		Description: "Best available quality",
		Format:      "bv*+ba/b",
	},
	{
		Name:        "4k",
		Description: "4k (3840x2160)",
		Format:      "bestvideo[height<=2160]+bestaudio/best[height<=2160]",
	},
	{
		Name:        "2k",
		Description: "2k (2560x1440)",
		Format:      "bestvideo[height<=1440]+bestaudio/best[height<=1440]",
	},
	{
		Name:        "1080p",
		Description: "Full HD (1920x1080)",
		Format:      "bestvideo[height<=1080]+bestaudio/best[height<=1080]",
	},
	{
		Name:        "720p",
		Description: "HD (1280x720)",
		Format:      "bestvideo[height<=720]+bestaudio/best[height<=720]",
	},
	{
		Name:        "480p",
		Description: "SD (854x480)",
		Format:      "bestvideo[height<=480]+bestaudio/best[height<=480]",
	},
	{
		Name:        "360p",
		Description: "Low quality (640x360)",
		Format:      "bestvideo[height<=360]+bestaudio/best[height<=360]",
	},
}

func PresetNames() []string {
	names := make([]string, len(QualityPresets))
	for i, p := range QualityPresets {
		names[i] = p.Name
	}

	return names
}

func GetPresetByName(name string) *QualityPreset {
	for _, p := range QualityPresets {
		if p.Name == name {
			return &p
		}
	}

	return nil
}

func IsValidPreset(name string) bool {
	return GetPresetByName(name) != nil
}

func ResolveQuality(quality string) string {
	if quality == "" {
		return QualityPresets[0].Format
	}

	preset := GetPresetByName(quality)
	if preset != nil {
		return preset.Format
	}

	return quality
}
