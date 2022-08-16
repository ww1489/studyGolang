package wlog

// NewLog 结构体构造函数
func NewLog(level string) Logger {
	tpm, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	return Logger{
		level: tpm,
	}
}

func (l Logger) Debug(format string, a ...interface{}) {
	if l.level <= DEBUG {
		l.message("DEBUG", format, a...)
	}
}

func (l Logger) Info(format string, a ...interface{}) {
	if l.level <= INFO {
		l.message("INFO", format, a...)
	}
}

func (l Logger) Worning(format string, a ...interface{}) {
	if l.level <= WORNING {
		l.message("WORNING", format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.level <= ERROR {
		l.message("ERROR", format, a...)
	}
}

func (l Logger) Fatal(format string, a ...interface{}) {
	if l.level <= FATAL {
		l.message("FATAL", format, a...)
	}
}
