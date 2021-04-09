package utils

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// Convert string to int
func ConvStrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err, "Failed to convert string to int")
	}
	return i
}

// Log error when error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg+": ", err)
	}
}

// Exec shell script file
func ExecShellScriptFile(pathFile string) {
	cmd := exec.Command(pathFile, "-c")
	err := cmd.Run()
	if err != nil {
		log.Println("Exec failed:" + err.Error())
	}
}

// Hash value with MD5
func HashWithMd5(value []byte) string {
	hash := md5.New()
	hash.Write(value)
	return hex.EncodeToString(hash.Sum(nil))
}

// Path to check if does not exist THEN make folder
func MkDirUpload(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}
}

// Read csv from csv file
func ReadCsv(pathFile string) (rows [][]string, err error) {
	f, err := os.Open(pathFile)
	if err != nil {
		log.Println(err.Error())
	}

	csvr := csv.NewReader(f)
	defer f.Close()
	//csvr.FieldsPerRecord += 1
	rows, err = csvr.ReadAll()
	if err != nil {
		log.Println(err.Error())
	}

	return rows, err
}

// Upload single file
func UploadSingleFile(c *gin.Context, keyName string) (err error) {
	file, _ := c.FormFile(keyName)
	dst := "./storages/uploads/" + file.Filename

	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// Get current time
func GetCurrentTime(format string) string {
	now := time.Now()
	// 2006-01-02 15:04:05
	// 2006-01-02
	return now.Format(format)
}

func WriteLogToFile(code int, msg string) {
	log.WithFields(log.Fields{
		"code": code,
	}).Info(msg)
}
