package config

type Config struct {
	Weights                Weights    `yaml:"weights"`
	Thresholds             Thresholds `yaml:"thresholds"`
	BatchSizeLimits        Limits     `yaml:"batch_size_limits"`
	IntervalLimits         Limits     `yaml:"interval_limits"`
	SamplingInterval       int        `yaml:"sampling_interval"`
	ProcessingIntervalBase float64    `yaml:"processing_interval_base"`
	Constants              Constants  `yaml:"constants"`
	StaticBatchSize        int        `yaml:"static_batch_size"`
	WorkerCount            int        `yaml:"worker_count"`
}

type Weights struct {
	W1 float64 `yaml:"w1"`
	W2 float64 `yaml:"w2"`
	W3 float64 `yaml:"w3"`
	W4 float64 `yaml:"w4"`
}

type Thresholds struct {
	Priority float64 `yaml:"priority"`
}

type Limits struct {
	Min float64 `yaml:"min"`
	Max float64 `yaml:"max"`
}

type Constants struct {
	Alpha float64 `yaml:"alpha"`
	Beta  float64 `yaml:"beta"`
	Gamma float64 `yaml:"gamma"`
	C     float64 `yaml:"c"`
}

func LoadConfig(path string) *Config {

	var cfg = Config{
		Weights: Weights{
			W1: 1,
			W2: 2,
			W3: 1,
			W4: 1,
		},
		Thresholds: Thresholds{
			Priority: 16,
		},
		BatchSizeLimits: Limits{
			Min: 10,
			Max: 100,
		},
		IntervalLimits: Limits{
			Min: 1.0,
			Max: 10.0,
		},
		SamplingInterval:       5000,
		ProcessingIntervalBase: 100.0,
		Constants: Constants{
			Alpha: 1.0,
			Beta:  0.0,
			Gamma: 1.0,
			C:     100.0,
		},
		StaticBatchSize: 50,
		WorkerCount:     4,
	}

	return &cfg
}
