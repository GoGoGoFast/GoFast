package idcardutil

import (
	"errors"
	"regexp"
	"sync"
	"time"
)

// Province 省、市、区的结构体
type Province struct {
	Code  string `json:"code"`
	Label string `json:"label"`
	List  []City `json:"list"`
}

type City struct {
	Code  string     `json:"code"`
	Label string     `json:"label"`
	List  []District `json:"list"`
}

type District struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

var (
	once      sync.Once
	provinces []Province
)

func loadProvinces() {
	for _, p := range idCardJSONData {
		province := Province{
			Code:  p["code"].(string),
			Label: p["label"].(string),
		}

		for _, c := range p["list"].([]map[string]interface{}) {
			city := City{
				Code:  c["code"].(string),
				Label: c["label"].(string),
			}

			for _, d := range c["list"].([]map[string]interface{}) {
				district := District{
					Code:  d["code"].(string),
					Label: d["label"].(string),
				}
				city.List = append(city.List, district)
			}
			province.List = append(province.List, city)
		}
		provinces = append(provinces, province)
	}
}

var idCard18Pattern = regexp.MustCompile(`^\d{17}(\d|X)$`)

// IsValidCard 验证身份证是否合法
func IsValidCard(idCard string) bool {
	return len(idCard) == 18 && idCard18Pattern.MatchString(idCard) && validate18IDCard(idCard)
}

// 验证18位身份证
func validate18IDCard(idCard string) bool {
	// 校验位计算
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checksum := 0
	for i := 0; i < 17; i++ {
		num := int(idCard[i] - '0')
		checksum += num * weights[i]
	}
	checkDigits := "10X98765432"
	checkDigit := checkDigits[checksum%11]

	return checkDigit == idCard[17]
}

// GetBirthByIdCard 获取生日
func GetBirthByIdCard(idCard string) (string, error) {
	if !IsValidCard(idCard) {
		return "", errors.New("无效的身份证号码")
	}
	return idCard[6:14], nil
}

// GetAgeByIdCard 获取年龄
func GetAgeByIdCard(idCard string, currentDate time.Time) (int, error) {
	birthStr, err := GetBirthByIdCard(idCard)
	if err != nil {
		return 0, err
	}

	birthDate, err := time.Parse("20060102", birthStr)
	if err != nil {
		return 0, err
	}

	age := currentDate.Year() - birthDate.Year()
	if currentDate.YearDay() < birthDate.YearDay() {
		age--
	}
	return age, nil
}

// GetGenderByIdCard 获取性别
func GetGenderByIdCard(idCard string) (string, error) {
	if !IsValidCard(idCard) {
		return "", errors.New("无效的身份证号码")
	}

	genderCode := idCard[16]

	if (genderCode-'0')%2 == 1 {
		return "男", nil
	}
	return "女", nil
}

// GetRegionByIdCard 获取省份和城市
func GetRegionByIdCard(idCard string) (string, error) {
	if !IsValidCard(idCard) {
		return "", errors.New("无效的身份证号码")
	}

	code := idCard[:6]

	once.Do(loadProvinces)

	for _, province := range provinces {
		if province.Code[:2] == code[:2] { // 省级匹配
			for _, city := range province.List {
				if city.Code[:4] == code[:4] { // 市级匹配
					for _, district := range city.List {
						if district.Code == code { // 区级匹配
							return province.Label + " " + city.Label + " " + district.Label, nil
						}
					}
					return province.Label + " " + city.Label, nil
				}
			}
			return province.Label, nil
		}
	}
	return "未知", nil
}
