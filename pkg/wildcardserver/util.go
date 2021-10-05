package wildcardserver

import (
	"math"
	"net"
	"reflect"
	"strings"
)

type IPNotFoundErr struct {
	error
}
func findIPv4InDomain(domain string) (net.IP, error) {

	domainFrags := strings.Split(domain, ".")
	reverseAny(domainFrags)
	
	// ip by bar split
	barSplitIpFound := true
	barSplitIp, barSplitIpIndex, err := findIPv4ByBarSplit(domainFrags)
	if err != nil {
		barSplitIpFound = false
	}
	// ip by point split
	pointSplitIpFound := true
	pointSplitIp, pointSplitIpIndex, err := findIPv4ByStringSlice(domainFrags)
	if err != nil {
		pointSplitIpFound = false
	}

	err = IPNotFoundErr{}
	retIp := net.IP{}
	index := math.MaxInt
	if barSplitIpFound && barSplitIpIndex < index {
		index = barSplitIpIndex
		retIp = barSplitIp
		err = nil
	}
	if pointSplitIpFound && pointSplitIpIndex < index {
		index = barSplitIpIndex
		retIp = pointSplitIp
		err = nil
	}
	
	return retIp, err
}

func findIPv4ByStringSlice(domainFrags []string) (net.IP, int, error) {
	for cursor := 0; cursor <= len(domainFrags) - 4; cursor++ {
		ip, err := getIPv4ByStringSlice(domainFrags[cursor: cursor + 4])
		if err == nil {
			return ip, cursor, err
		}
	}
	return nil, 0, IPNotFoundErr{}
}

func getIPv4ByStringSlice(slice []string) (net.IP, error) {
	IPSlice := make([]string, len(slice))
	copy(IPSlice, slice)
	reverseAny(IPSlice)

	ipStr := strings.Join(IPSlice, ".")
	if ip := net.ParseIP(ipStr).To4(); ip != nil {
		return ip, nil
	}
	return nil, IPNotFoundErr{}
}

func findIPv4ByBarSplit(domainFrags []string) (net.IP, int, error) {
	for index, frag := range domainFrags {
		ip, err := getIPv4ByBarSplit(frag)
		if err == nil {
			return ip, index, err
		}
	}
	return nil, 0, IPNotFoundErr{}
}

func getIPv4ByBarSplit(subDomain string) (net.IP, error) {
	ipStr := strings.ReplaceAll(subDomain, "-", ".")
	if ip := net.ParseIP(ipStr).To4(); ip != nil {
		return ip, nil
	}
	return nil, IPNotFoundErr{}
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
