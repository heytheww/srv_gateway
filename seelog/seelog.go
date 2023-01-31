package seelog

import log "github.com/cihub/seelog"

type SeeLog struct {
	FileName string
}

func (s *SeeLog) NewLog() (log.LoggerInterface, error) {
	logger, err := log.LoggerFromConfigAsFile(s.FileName)

	if err != nil {
		return nil, err
	}

	log.ReplaceLogger(logger)
	return logger, nil
}
