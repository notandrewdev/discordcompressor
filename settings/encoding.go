package settings

var Encoding = initEncoding()

func initEncoding() *encoding {
	return &encoding{
		TmixDownFPS:         false,
		HalveDownFPS:        true,
		SizeTargetMB:        8,
		BitrateTargetMult:   1,
		BitrateLimitMax:     12500,
		BitrateLimitMin:     500,
		BitrateTargets: []*Target{{
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 1080,
						FPSMax:  60,
					},
					{
						Focus:   "resolution",
						VResMax: 2160,
						FPSMax:  30,
					},
				},
				BitrateMin:   7500,
				Encoder:      "x264",
				AudioEncoder: "aac2",
				Preset:       "medium",
			}, {
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 1080,
						FPSMax:  60,
					},
					{
						Focus:   "resolution",
						VResMax: 1440,
						FPSMax:  30,
					},
				},
				BitrateMin:   5000,
				Encoder:      "x264",
				AudioEncoder: "aac2",
				Preset:       "slow",
			}, {
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 1080,
						FPSMax:  60,
					},
					{
						Focus:   "resolution",
						VResMax: 1440,
						FPSMax:  30,
					},
				},
				BitrateMin:   3750,
				Encoder:      "x264",
				AudioEncoder: "aac2",
				Preset:       "slower",
			}, {
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 900,
						FPSMax:  60,
					},
					{
						Focus:   "resolution",
						VResMax: 1080,
						FPSMax:  30,
					},
				},
				BitrateMin:   2500,
				Encoder:      "x264",
				AudioEncoder: "aac2",
				Preset:       "veryslow",
			}, {
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 720,
						FPSMax:  60,
					},
					{
						Focus:   "resolution",
						VResMax: 1080,
						FPSMax:  30,
					},
				},
				BitrateMin:   1500,
				Encoder:      "x264",
				AudioEncoder: "aac1.5",
				Preset:       "veryslow",
			}, {
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 720,
						FPSMax:  60,
					},
					{
						Focus:   "resolution",
						VResMax: 900,
						FPSMax:  30,
					},
				},
				BitrateMin:   1000,
				Encoder:      "x264",
				AudioEncoder: "aac1.5",
				Preset:       "veryslow",
			}, {
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 540,
						FPSMax:  60,
					},
					{
						Focus:   "resolution",
						VResMax: 720,
						FPSMax:  30,
					},
				},
				BitrateMin:   650,
				Encoder:      "x264",
				AudioEncoder: "aac1.1",
				Preset:       "veryslow",
			}, {
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 360,
						FPSMax:  60,
					},
					{
						Focus:   "resolution",
						VResMax: 540,
						FPSMax:  30,
					},
				},
				BitrateMin:   400,
				Encoder:      "x264",
				AudioEncoder: "aac1.1",
				Preset:       "veryslow",
			}, {
				Limits: []*Limits{
					{
						Focus:   "framerate",
						VResMax: 360,
						FPSMax:  30,
					},
					{
						Focus:   "resolution",
						VResMax: 540,
						FPSMax:  15,
					},
				},
				BitrateMin:   0,
				Encoder:      "x264",
				AudioEncoder: "aac1.1",
				Preset:       "veryslow",
			},
		},
		Encoders: []*Encoder{
			{
				Name:         "x264",
				Encoder:      "libx264",
				CodecName:    "h264",
				Options:      "-x264-params qpmin=20",
				Keyint:       10,
				PresetCmd:    "-preset",
				TwoPass:      true,
				PassCmd:      "-pass",
				Container:    "mp4",
			},
		},
		AudioEncoders: []*AudioEncoder{
			{
				Name:         "aac1.5",
				Encoder:      "aac",
				CodecName:    "aac",
				Options:      "-qscale:a 1.5",
				UsesBitrate:  false,
				MaxBitrate:   192,
				MinBitrate:   128,
				BitratePerc:  12,
			},
			{
				Name:         "aac1.1",
				Encoder:      "aac",
				CodecName:    "aac",
				Options:      "-qscale:a 1.1",
				UsesBitrate:  false,
				MaxBitrate:   192,
				MinBitrate:   128,
				BitratePerc:  12,
			},
			{
				Name:         "aac2",
				Encoder:      "aac",
				CodecName:    "aac",
				Options:      "-qscale:a 2",
				UsesBitrate:  false,
				MaxBitrate:   192,
				MinBitrate:   128,
				BitratePerc:  12,
			},
		},
	}
	}

type encoding struct {
	TmixDownFPS           bool
	HalveDownFPS          bool
	SizeTargetMB          float64
	BitrateTargetMult     float64
	BitrateLimitMax       int
	BitrateLimitMin       int
	BitrateTargets        []*Target
	Encoders              []*Encoder
	AudioEncoders         []*AudioEncoder
}

type Target struct {
	Limits       []*Limits
	BitrateMin   int
	Encoder      string
	AudioEncoder string
	Preset       string
}

type Limits struct {
	Focus   string
	VResMax int
	FPSMax  int
}

type Encoder struct {
	Name         string
	Encoder      string
	CodecName    string
	Options      string
	Keyint       int
	PresetCmd    string
	TwoPass      bool
	PassCmd      string
	Container    string
}

type AudioEncoder struct {
	Name         string
	Encoder      string
	CodecName    string
	Options      string
	UsesBitrate  bool
	MaxBitrate   int
	MinBitrate   int
	BitratePerc  int
}