package log

func Prefix(newprefix ...string) Logger {
	return Logger{
		prefixes: newprefix,
	}
}
