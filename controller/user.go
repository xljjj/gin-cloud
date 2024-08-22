package controller

import (
	"CloudDrive/model"
	"CloudDrive/redis"
	"CloudDrive/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Welcome 欢迎页
func Welcome(c *gin.Context) {
	// 测试GO-template
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title": "欢迎使用 GO 网盘",
	})
}

// PrivateInfo 私密信息
type PrivateInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}

// QQUserInfo QQ用户信息
type QQUserInfo struct {
	Nickname  string
	FigureUrl string `json:"figure_url"`
}

// Login 登录
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// HandleLogin 处理登录
func HandleLogin(c *gin.Context) {
	state := "xxxxxxx"
	url := "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" +
		viper.GetString("qq.appID") +
		"&redirect_uri=" +
		viper.GetString("qq.redirectURL") +
		"&state=" + state
	c.Redirect(http.StatusMovedPermanently, url)
}

// GetQQToken 获取access_token
func GetQQToken(c *gin.Context) {
	code := c.Query("code")
	loginUrl := "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=" +
		viper.GetString("qq.appID") + "&client_secret=" +
		viper.GetString("qq.appKey") + "&redirect_uri=" +
		viper.GetString("redirectURL") +
		"&code=" + code
	response, err := http.Get(loginUrl)
	if err != nil {
		fmt.Println("请求错误", err.Error())
		return
	}
	defer response.Body.Close()
	bs, _ := io.ReadAll(response.Body)
	body := string(bs)
	resultMap := util.ConvertToMap(body)
	info := &PrivateInfo{}
	info.AccessToken = resultMap["access_token"]
	info.RefreshToken = resultMap["refresh_token"]
	info.ExpiresIn = resultMap["expires_in"]
	GetOpenId(info, c)
}

// GetOpenId 获取QQ的openId
func GetOpenId(info *PrivateInfo, c *gin.Context) {
	resp, err := http.Get(fmt.Sprintf("%s?access_token=%s", "https://graph.qq.com/oauth2.0/me", info.AccessToken))
	if err != nil {
		fmt.Println("GetOpenId Err", err.Error())
		return
	}
	defer resp.Body.Close()

	bs, _ := io.ReadAll(resp.Body)
	body := string(bs)
	info.OpenId = body[45:77]

	GetUserInfo(info, c)
}

// GetUserInfo 获取QQ用户信息
func GetUserInfo(info *PrivateInfo, c *gin.Context) {
	params := url.Values{}
	params.Add("access_token", info.AccessToken)
	params.Add("openid", info.OpenId)
	params.Add("oauth_consumer_key", viper.GetString("qq.appID"))

	uri := fmt.Sprintf("https://graph.qq.com/user/get_user_info?%s", params.Encode())
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("GetUserInfo Err:", err.Error())
		return
	}
	defer resp.Body.Close()

	bs, _ := io.ReadAll(resp.Body)

	LoginSucceed(string(bs), info.OpenId, c)
}

// LoginSucceed 登录成功后的处理
func LoginSucceed(userInfo, openId string, c *gin.Context) {
	var qqUserInfo QQUserInfo
	//将数据转为结构体
	if err := json.Unmarshal([]byte(userInfo), &qqUserInfo); err != nil {
		fmt.Println("转换json失败", err.Error())
		return
	}

	//创建一个token
	hashToken := util.Md5Encode("token" + string(time.Now().Unix()) + openId)
	//存入redis
	if err := redis.SetKey(c, hashToken, openId, time.Hour*24); err != nil {
		fmt.Println("Redis Set Err:", err.Error())
		return
	}
	//设置cookie
	c.SetCookie("Token", hashToken, 3600*24, "/", "pyxgo.cn", false, true)

	if ok := model.UserExists(openId); ok { //用户存在直接登录
		//登录成功重定向到首页
		c.Redirect(http.StatusMovedPermanently, "/cloud/index")
	} else {
		model.CreateUser(openId, qqUserInfo.Nickname, qqUserInfo.FigureUrl)
		//登录成功重定向到首页
		c.Redirect(http.StatusMovedPermanently, "/cloud/index")
	}
}

// Logout 退出登录
func Logout(c *gin.Context) {
	token, err := c.Cookie("Token")
	if err != nil {
		fmt.Println("cookie", err.Error())
	}

	if err := redis.DeleteKey(c, token); err != nil {
		fmt.Println("Del Redis Err:", err.Error())
	}

	c.SetCookie("Token", "", 0, "/", "pyxgo.cn", false, false)
	c.Redirect(http.StatusFound, "/")
}
