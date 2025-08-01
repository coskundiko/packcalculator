package helper

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	LOG "log"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

func TempDir(dir, pattern string) (name string, err error) {

	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	err = os.Mkdir("tmp/"+hex.EncodeToString(randBytes), 0600)
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Panic().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v", fn, line, err)
	}
	return name, err
}

func CheckError(err error, successMsg ...string) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Panic().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v", fn, line, err)
	} else if len(successMsg) > 0 && successMsg[0] != "" {
		log.Info().Msgf("%v", successMsg[0])
	}
}

func DurationToHours(years int, months int, days int) float64 {
	return time.Since(time.Now().AddDate(-years, -months, -days)).Hours()
}

func YoloToCoco(Boxes interface{}) map[string]float64 {
	boxes := Boxes.(primitive.A)
	x_min := boxes[0].(float64)
	y_min := boxes[1].(float64)
	x_max := boxes[2].(float64)
	y_max := boxes[3].(float64)
	box_width := x_max - x_min
	box_height := y_max - y_min
	return map[string]float64{
		"width":  box_width,
		"height": box_height,
		"x":      x_min,
		"y":      y_min,
	}
}

func NumCPU() int {
	numGoroutines := runtime.NumCPU()
	if numGoroutines > 24 {
		numGoroutines = 15
	} else if numGoroutines > 12 {
		numGoroutines = 8
	} else if numGoroutines > 7 {
		numGoroutines = 4
	} else if numGoroutines > 1 {
		numGoroutines = 2
	}
	return numGoroutines
}

func CheckErrorf(err error, format string, v ...interface{}) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Panic().Msgf("[error] %s:%d %s - %v", fn, line, err, fmt.Sprintf(format, v...))
	}
}

func RunMethodByName(any interface{}, name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

func PrintIfError(err error, successMsg ...string) error {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[line] %s:%d %v", fn, line)
	} else if len(successMsg) > 0 && successMsg[0] != "" {
		log.Info().Msgf("%v", successMsg[0])
	}
	return err
}

func PrintIfErrorf(err error, format string, v ...interface{}) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v", fn, line, fmt.Sprintf(format, v...))
	}
}

func LogReturnError(err error, message interface{}) error {
	_, fn, line, _ := runtime.Caller(1)
	log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v %v", fn, line, message, err)
	return fmt.Errorf("[error] %s:%d %v %v", fn, line, message, err)
}

func LogReturnErrorf(err error, format string, v ...interface{}) error {
	_, fn, line, _ := runtime.Caller(1)
	log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v %v", fn, line, fmt.Sprintf(format, v...), err)
	return fmt.Errorf("[error] %s:%d %v %v", fn, line, fmt.Sprintf(format, v...), err)
}

func PrintError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v", fn, line, err)
}

func Print(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v", fn, line, err)
}

func PrintErrorf(err error, format string, v ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v", fn, line, fmt.Sprintf(format, v...))
}

func DebugIfError(err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Debug().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v", fn, line, err)
	}
}

func DebugError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Debug().Err(err).Msgf("[error] %s:%d %v", fn, line, err)
}
func Debug(data any) {
	_, fn, line, _ := runtime.Caller(1)
	log.Debug().Msgf("[debug] %v - [line] %s:%d ", data, fn, line)
}

func LogReturnErrorClientGRPC(code codes.Code, err error, message string) error {
	_, fn, line, _ := runtime.Caller(1)
	log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v %v", fn, line, message, err)
	return status.Errorf(code, "[error] %s:%d %v %v", fn, line, message, err)
}

func ObjectIDToString(id interface{}) string {
	return id.(primitive.ObjectID).Hex()
}

func CheckErrorExit(err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Error().Stack().Err(errors.Errorf("%v", err)).Msgf("[error] %s:%d %v", fn, line, err)
		os.Exit(2)
	}
}

func Shellout(shell, command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shell, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if stderr.String() != "" {
		log.Debug().Msg(stderr.String())
	}
	return stdout.String(), stderr.String(), err
}

func ShelloutWithContext(ctx context.Context, shell, command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, shell, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if stderr.String() != "" {
		log.Debug().Msg(stderr.String())
	}
	ctx.Done()
	return stdout.String(), stderr.String(), err
}

func CheckFolders(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			LOG.Fatal(err)
		}
		log.Debug().Err(err).Msg("created: " + path)
	}
}

func CheckIfFileExists(path string) error {
	_, fn, line, _ := runtime.Caller(1)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Debug().Err(err).Msgf("%s:%d", fn, line)
		return err
	}
	return nil
}

func SerializeAny(msg interface{}) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func DeserializeAny(b []byte) (msg interface{}, err error) {
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err = decoder.Decode(&msg)
	return msg, err
}

func Serialize(msg *map[string]string) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func Deserialize(b []byte) (msg map[string]string, err error) {
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err = decoder.Decode(&msg)
	return msg, err
}

//func CompareKBwithMB(compare float64, with string) bool {
//	a := bytesize.New(compare)
//	b, err := bytesize.Parse(with)
//	CheckError(err)
//	if a > b {
//		return true
//	}
//	return false
//
//}

func ParallelizeFunc(functions ...func()) {
	n := runtime.NumCPU()
	if cap(functions) > n {
		_, fn, line, _ := runtime.Caller(1)
		log.Fatal().Msgf("[error] %s:%d number of concurrent functions cannot be more than the number of CPU cores: %v, given: %v", fn, line, n, cap(functions))
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))
	defer waitGroup.Wait()
	for _, function := range functions {
		go func(do func()) {
			defer waitGroup.Done()
			do()
		}(function)
	}
}

func InArray(val interface{}, array interface{}) (exists bool) {
	if reflect.TypeOf(array).Kind() == reflect.Slice {
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}

func IsValueInList(value string, list []string) bool {
	values := make([]interface{}, 0, len(list))
	for _, v := range list {
		values = append(values, v)
		if value == v {
			return true
		}
	}
	return false
}

func FormFileName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Replace(name, "/", "-", -1)
	name = strings.Replace(name, " ", "_", -1)
	name = strings.Replace(name, "\\", "-", -1)
	return name
}

func RegularFileName(str string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		LOG.Fatal(err)
	}
	return strings.Replace(strings.Replace(re.ReplaceAllString(strings.Replace(strings.Replace(str, " ", "_TRANSFRITZFILENAMESPACE_", -1), ".", "_TRANSFRITZFILENAMEDOT_", -1), ""), "_TRANSFRITZFILENAMESPACE_", " ", -1), "_TRANSFRITZFILENAMEDOT_", ".", -1)
}
