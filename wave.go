package main

import "encoding/json"

func ParseWave(b []byte) (*JhoveWave, error) {
	jhoveWave := JhoveWave{}
	if err := json.Unmarshal(b, &jhoveWave); err != nil {
		return nil, err
	}
	return &jhoveWave, nil
}

type JhoveWave struct {
	Jhove struct {
		Name          string `json:"name"`
		Release       string `json:"release"`
		Date          string `json:"date"`
		ExecutionTime string `json:"executionTime"`
		RepInfo       []struct {
			URI             string `json:"uri"`
			ReportingModule struct {
				Name    string `json:"name"`
				Release string `json:"release"`
				Date    string `json:"date"`
			} `json:"reportingModule"`
			LastModified string   `json:"lastModified"`
			Size         int      `json:"size"`
			Format       string   `json:"format"`
			Status       string   `json:"status"`
			SigMatch     []string `json:"sigMatch"`
			MimeType     string   `json:"mimeType"`
			Profiles     []string `json:"profiles"`
			Properties   []struct {
				WAVEMetadata []struct {
					AESAudioMetadata struct {
						AesAnalogDigitalFlag string `json:"aes:analogDigitalFlag"`
						AesSchemaVersion     string `json:"aes:schemaVersion"`
						AesFormat            string `json:"aes:format"`
						AesAudioDataEncoding string `json:"aes:audioDataEncoding"`
						AesByteOrder         string `json:"aes:byteOrder"`
						AesFirstSampleOffset int    `json:"aes:firstSampleOffset"`
						AesUse               struct {
							AesUseType   string `json:"aes:useType"`
							AesOtherType string `json:"aes:otherType"`
						} `json:"aes:use"`
						AesPrimaryIdentifier     string `json:"aes:primaryIdentifier"`
						AesPrimaryIdentifierType string `json:"aes:primaryIdentifierType"`
						AesFace                  struct {
							AesTimeline struct {
								TcfStartTime struct {
									TcfFrameCount   int    `json:"tcf:frameCount"`
									TcfTimeBase     int    `json:"tcf:timeBase"`
									TcfVideoField   string `json:"tcf:videoField"`
									TcfCountingMode string `json:"tcf:countingMode"`
									TcfHours        int    `json:"tcf:hours"`
									TcfMinutes      int    `json:"tcf:minutes"`
									TcfSeconds      int    `json:"tcf:seconds"`
									TcfFrames       int    `json:"tcf:frames"`
									TcfSamples      struct {
										TcfSampleRate      string `json:"tcf:sampleRate"`
										TcfNumberOfSamples int    `json:"tcf:numberOfSamples"`
									} `json:"tcf:samples"`
									TcfFilmFraming struct {
										TcfFraming     string `json:"tcf:framing"`
										TcfFramingType string `json:"tcf:framingType"`
									} `json:"tcf:filmFraming"`
								} `json:"tcf:startTime"`
								TcfDuration struct {
									TcfFrameCount   int    `json:"tcf:frameCount"`
									TcfTimeBase     int    `json:"tcf:timeBase"`
									TcfVideoField   string `json:"tcf:videoField"`
									TcfCountingMode string `json:"tcf:countingMode"`
									TcfHours        int    `json:"tcf:hours"`
									TcfMinutes      int    `json:"tcf:minutes"`
									TcfSeconds      int    `json:"tcf:seconds"`
									TcfFrames       int    `json:"tcf:frames"`
									TcfSamples      struct {
										TcfSampleRate      string `json:"tcf:sampleRate"`
										TcfNumberOfSamples int    `json:"tcf:numberOfSamples"`
									} `json:"tcf:samples"`
									TcfFilmFraming struct {
										TcfFraming     string `json:"tcf:framing"`
										TcfFramingType string `json:"tcf:framingType"`
									} `json:"tcf:filmFraming"`
								} `json:"tcf:duration"`
							} `json:"aes:timeline"`
							AesNumChannels int `json:"aes:numChannels"`
							AesStreams     []struct {
								AesChannelNum  int    `json:"aes:channelNum"`
								AesMapLocation string `json:"aes:mapLocation"`
							} `json:"aes:streams"`
						} `json:"aes:face"`
						AesFormatList []struct {
							AesBitDepth         int     `json:"aes:bitDepth"`
							AesSampleRate       float64 `json:"aes:sampleRate"`
							AesBitrateReduction struct {
								AesCodecName                      string `json:"aes:codecName"`
								AesCodecNameVersion               string `json:"aes:codecNameVersion"`
								AesCodecCreatorApplication        string `json:"aes:codecCreatorApplication"`
								AesCodecCreatorApplicationVersion string `json:"aes:codecCreatorApplicationVersion"`
								AesCodecQuality                   string `json:"aes:codecQuality"`
								AesDataRate                       string `json:"aes:dataRate"`
								AesDataRateMode                   string `json:"aes:dataRateMode"`
							} `json:"aes:bitrateReduction"`
						} `json:"aes:formatList"`
					} `json:"AESAudioMetadata,omitempty"`
					CompressionCode       string   `json:"CompressionCode,omitempty"`
					AverageBytesPerSecond int      `json:"AverageBytesPerSecond,omitempty"`
					BlockAlign            int      `json:"BlockAlign,omitempty"`
					ExtraFormatBytes      []string `json:"ExtraFormatBytes,omitempty"`
					Fact                  struct {
						Size int `json:"Size"`
					} `json:"Fact,omitempty"`
					Data struct {
						DataLength int `json:"DataLength"`
					} `json:"Data,omitempty"`
				} `json:"WAVEMetadata"`
			} `json:"properties"`
		} `json:"repInfo"`
	} `json:"jhove"`
}
