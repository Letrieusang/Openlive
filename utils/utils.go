package utils

import (
	"api-openlive/common"
	"api-openlive/config"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math"
	"math/big"
	"os"
	"strings"
	"time"
)

var (
	DESTINATION  = "storage/"
	allowedTypes = []binding.Binding{
		binding.Query,
		binding.Form,
		binding.FormPost,
		binding.FormMultipart,
	}
	OrderBy = []string{
		"ASC",
		"DESC",
	}

	AllowedTypes = map[string][]string{
		common.AVATAR: {
			"image/png",
			"image/jpeg",
			"image/jpg",
			"image/heic",
			"image/heif",
		},
		common.COVER: {
			"image/png",
			"image/jpeg",
			"image/jpg",
			"image/heic",
			"image/heif",
		},
		common.COLLECTION: {
			"image/png",
			"image/jpeg",
			"image/jpg",
			"image/heic",
			"image/heif",
		},
		common.ITEM: {
			"image/jpg",
			"image/png",
			"image/jpeg",
			"image/gif",
			"image/svg+xml",
			"video/mp4",
			"video/mov",
			"audio/mpeg",
			"video/webm",
			"audio/webm",
			"audio/wav",
			"audio/ogg",
			"video/ogg",
			"model/gltf+json",
			"model/gltf-binary",
			"image/heic",
			"image/heif",
			"video/quicktime",
		},
		common.ITEM_IMAGE: {
			"image/png",
			"image/jpeg",
			"image/jpg",
			"image/gif",
			"image/heic",
			"image/heif",
		},
	}

	ErrBindingFail = errors.New(fmt.Sprintf("Bind function only allows %v\\n", allowedTypes))
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func BindQuery(data interface{}, bindingType binding.Binding, ctx *gin.Context) error {
	ok := false
	for _, b := range allowedTypes {
		if b == bindingType {
			ok = true
			break
		}
	}
	if !ok {
		return ErrBindingFail
	}
	_ = ctx.MustBindWith(&data, bindingType)
	return nil
}

func InStringSlice(arr []string, e string) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}
	return false
}

func CommonHandle(rq interface{}, ctx *gin.Context) *common.Response {
	if err := ctx.ShouldBindJSON(&rq); err != nil {
		return common.NewResponse(1, "Can not parse request body", err.Error())

	}
	validate := validator.New()
	err := validate.Struct(rq)
	if err != nil {
		return common.NewResponse(1, "Validate request error", err.Error())
	}
	return nil
}

func MoveFile(sourcePath, destPath string) error {
	var arrayRawData = strings.Split(sourcePath, "/tmp/")
	if len(arrayRawData) < 2 {
		sourcePath = arrayRawData[0]
	} else {
		sourcePath = "tmp/" + arrayRawData[1]
	}
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	return nil
}

func GetFileName(path string) string {
	arr := strings.Split(path, "/")
	fileName := arr[len(arr)-1]
	return fileName
}

func ConvertEToString(num string) string {
	flt, _, err := big.ParseFloat(num, 10, 0, big.ToNearestEven)
	if err != nil {
		panic(err)
	}

	if flt.IsInt() {
		var i = new(big.Int)
		i, _ = flt.Int(i)
		return i.String()
	}

	return fmt.Sprintf("%f", flt)
}

func BuildRefLink(code string) string {
	link := common.REF_LINK_URL_DEV
	if os.Getenv("ENV") == "prod" {
		link = common.REF_LINK_URL_PROD
	}

	return fmt.Sprintf("%v?ref_code=%v", link, code)
}

func RebuildUrlPath(urlPath string) string {
	var arrayRawData = strings.Split(urlPath, "/storage/")
	if len(arrayRawData) < 2 {
		return urlPath
	}
	return config.GetServer().GetApiDomain() + "storage/" + arrayRawData[1]
}

func GetMD5Hash(text string) string {
	key := []byte(fmt.Sprintf("%v.%v", common.HASH_PASSWORD, text))
	return fmt.Sprintf("%x", md5.Sum(key))
}

func BuildDateFilterByDay(diffDay int) int64 {
	diffDay = diffDay * -1
	timeNow := time.Now()
	to, _ := time.Parse(common.ISO8601, timeNow.String()[0:10])
	startDate := to.AddDate(0, 0, diffDay)
	return startDate.Unix()
}

func ConvertStructToMap(v interface{}) map[string]int64 {
	bin, _ := json.Marshal(v)
	result := make(map[string]int64)
	_ = json.Unmarshal(bin, &result)
	return result
}

func ConvertStructToMapInterface(v interface{}) map[string]interface{} {
	bin, _ := json.Marshal(v)
	result := make(map[string]interface{})
	_ = json.Unmarshal(bin, &result)
	return result
}

func ConvertTimeFromMiliSec(miliSec int64) time.Time {
	return time.Unix(0, miliSec*int64(time.Millisecond))
}

func BuildFilterJson(param string, key string) string {
	// JSON_CONTAINS(test, '{"name":"Size","value":"22"}')
	constQuery := "JSON_CONTAINS(%v, '{\"name\":\"%v\",\"value\":\"%v\"}')"
	paramArr := strings.Split(param, ",")
	arrQuery := make([]string, 0, len(paramArr))
	for _, v := range paramArr {
		queryParam := strings.Split(v, ":")
		if len(queryParam) < 2 {
			continue
		}
		arrQuery = append(arrQuery, fmt.Sprintf(constQuery, key, queryParam[0], queryParam[1]))
	}

	return strings.Join(arrQuery, " OR ")
}

func ConvertPrice(price int64, rate float64) float64 {
	return float64(price) * rate
}

func InTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

func PercentChangeInt64(last, now int64) float64 {
	if last == 0 && now == 0 {
		return 0
	}

	if last == 0 {
		return 100
	}

	result := float64(now-last) / float64(last) * 100
	return math.Round(result*100) / 100
}

func PercentChangeFloat64(last, now float64) float64 {
	if last == 0 && now == 0 {
		return 0
	}

	if last == 0 {
		return 100
	}

	result := (now - last) / last * 100
	return math.Round(result*100) / 100
}

func PercentChangeString(last, now string) float64 {
	lastF, _, err := big.ParseFloat(last, 10, 0, big.ToNearestEven)
	if err != nil {
		panic(err)
	}

	nowF, _, err := big.ParseFloat(now, 10, 0, big.ToNearestEven)
	if err != nil {
		panic(err)
	}

	if lastF.Cmp(big.NewFloat(0)) == 0 {
		return 100
	}

	sub := new(big.Float).Sub(nowF, lastF)
	div := new(big.Float).Quo(sub, lastF)
	mul := new(big.Float).Mul(div, big.NewFloat(100))
	result, _ := mul.Float64()

	return result
}

func GetConvertedVolume(volume map[string]float64) float64 {
	result := float64(0)
	for k, v := range volume {
		result += v * rate[k]
	}

	return math.Round(result*100) / 100
}

func GetConvertedNow(price float64) float64 {
	now := time.Now().Format(common.ISO8601)
	result := price * rate[now]
	return math.Round(result*100) / 100
}

func GetConvertedFloorPrice(price float64) float64 {
	now := time.Now().Format(common.ISO8601)
	return rate[now] * price
}
