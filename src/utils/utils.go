package utils

import (
	"bytes"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	log "github.com/sirupsen/logrus"
)

const (
	SslPort     = ":443"
	Https       = "https://"
	HoursPerDay = 24
)

func GetEthBLockNumber(url string) (uint64, error) {
	httpClient := &http.Client{
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Timeout:       http.DefaultClient.Timeout,
	}

	var jsonStr = []byte(`{"id":1, "jsonrpc": "2.0", "method": "eth_blockNumber", "params": []}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	block, err := jsonparser.GetString(body, "result")
	if err != nil {
		return 0, err
	}
	result, err := ConvertHexToDecimal(block)

	return uint64(result), err
}

func GetEthSyncing(url string) (bool, error) {
	httpClient := &http.Client{
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Timeout:       http.DefaultClient.Timeout,
	}

	var jsonStr = []byte(`{"id":1, "jsonrpc": "2.0", "method": "eth_syncing", "params": []}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	isSyncing, err := jsonparser.GetBoolean(body, "result")
	if err != nil {
		return isSyncing, err
	}
	return false, err
}

func ConvertHexToDecimal(hex string) (int64, error) {
	res, err := strconv.ParseInt(hexaNumberToInteger(hex), 16, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

func ParseBoolToInt(value bool) uint64 {
	if value {
		return 1
	} else {
		return 0
	}
}

func GetSslExpireDays(domain string) (uint64, error) {
	domain = parseDomainForSslQuery(domain)
	conn1, err := tls.Dial("tcp", domain, nil)
	if err != nil {
		log.Error("Server doesn't support SSL certificate err: " + domain + err.Error())
		return 0, err
	}
	date := conn1.ConnectionState().PeerCertificates[0].NotAfter
	certificationExpireHours := date.Sub(time.Now()).Hours()
	return uint64(certificationExpireHours) / HoursPerDay, nil
}

func parseDomainForSslQuery(domain string) string {
	if strings.HasPrefix(domain, Https) {
		domain = strings.Split(domain, Https)[1]
		domain = strings.Split(domain, "/")[0]
		return domain + SslPort
	}
	return domain
}

func IsPortOpenOutside(ip string, port int) bool {
	address := net.JoinHostPort(ip, string(port))
	// 3 second timeout
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return false
	} else {
		if conn != nil {
			_ = conn.Close()
			return true
		} else {
			return false
		}
	}
}
