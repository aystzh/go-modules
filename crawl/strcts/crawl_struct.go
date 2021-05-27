package strcts

//参数结构体
// {"id":"YTHSHH001","idType":"USERID","index":1,"productType":"","rows":100,"status":"0"}
type RequestParams struct {
	Id          string `json:"id"`
	IdType      string `json:"idType"`
	Index       int64  `json:"index"`
	ProductType string `json:"productType"`
	Rows        int64  `json:"rows"`
	Status      string `json:"status"`
}

//一级表单返回参数json结构体
type PostData struct {
	Code string `json:"code"` // 000000
	Data struct {
		AccountID   int64  `json:"accountId"`   // 2.1015601e+07
		AccountName string `json:"accountName"` // 一体化保险商城
		Level       string `json:"level"`       // 0
		ProductList struct {
			Data []struct {
				ActiveStatus string `json:"activeStatus"` // 0
				CBigImg      string `json:"cBigImg"`      // /images/kabTyLogo_b.png
				CCode        string `json:"cCode"`        // 2041TY
				CDesc        string `json:"cDesc"`        // 癌症医疗保险
				CImgURL      string `json:"cImgUrl"`      // /images/kabT_s.png
				CKind        string `json:"cKind"`        // healthy
				CName        string `json:"cName"`        // 抗癌保特药版
				CProdExtend  struct {
					YTHMALL struct {
						TAG      []string `json:"TAG"`      // 最高100%赔付
						Desc     string   `json:"desc"`     // “抗癌+特药”有效保障传统放化疗+新型靶向药品癌症治疗方式。
						NfeeDesc string   `json:"nfeeDesc"` // ￥48<span class='priceClass'>/年起</span>
					} `json:"YTHMALL"`
				} `json:"cProdExtend"`
				NPrice     float64 `json:"nPrice"`     // 48
				NProductID int64   `json:"nProductId"` // 143
				NSort      int64   `json:"nSort"`      // 2
			} `json:"data"`
			Index int64 `json:"index"` // 1
			Total int64 `json:"total"` // 22
		} `json:"productList"`
		StoreName interface{} `json:"storeName"` // <nil>
		TagID     string      `json:"tagId"`     // 0
	} `json:"data"`
	Message string `json:"message"` // 加载产品列表完成
	Success bool   `json:"success"` // true
}

//二级表单返回参数结构体
type GetData struct {
	Code string `json:"code"` // 000000
	Data struct {
		ID             int64  `json:"id"`          // 1
		Code           string `json:"code"`        // cyz001
		Desc           string `json:"desc"`        // 一人投保全家可享
		Itemname       string `json:"itemname"`    // 家庭综合险
		Name           string `json:"name"`        // 我爱我家
		NeedConfirm    bool   `json:"needConfirm"` // true
		ProdObjectList []struct {
			ID          string `json:"id,,omitempty"` //
			Code        string `json:"code"`          // POLICYHOLDER
			DataType    string `json:"dataType"`      // single
			Desc        string `json:"desc"`          //
			ElementList []struct {
				Code  string `json:"code"`  // name
				Name  string `json:"name"`  // 姓名
				Need  bool   `json:"need"`  // true
				Show  bool   `json:"show"`  // true
				Type  string `json:"type"`  // Name
				Value string `json:"value"` //
			} `json:"elementList"`
			Input             bool   `json:"input"`             // true
			Name              string `json:"name"`              // 投保人信息
			Need              bool   `json:"need"`              // true
			SubProdObjectList string `json:"subProdObjectList"` //
		} `json:"prodObjectList"`
		ProdSchemeList []struct {
			YTHMALL struct {
				SpecialSecurity []struct {
					Name  string `json:"name"`  // 保障额度
					Value string `json:"value"` // 整体累计保额超<span class='qytext'>200万元</span>
				} `json:"SpecialSecurity"`
				Desc []struct {
					Name  string `json:"name"`  // 年龄限制
					Value string `json:"value"` // 意外伤害<br>限30天-65周岁
				} `json:"desc"`
				ProdFeatures []struct {
					Value string `json:"value"` // 每天不到7毛钱，享高达<span class='qytext'>200万元</span>保额；<span class='qytext'>百万家财险、高额责任险、多种意外险</span>全方位保障，守护全家；
				} `json:"prodFeatures"`
				SellPoint struct {
					LeftBig    string `json:"leftBig"`    // 保额高达<span class='numStyle'>200</span>万
					LeftSmall  string `json:"leftSmall"`  // 每天仅0.66元起
					RightBig   string `json:"rightBig"`   // 一人投保全家可享
					RightSmall string `json:"rightSmall"` // 家财、责任、意外全方位守护
				} `json:"sellPoint"`
				Slogan string `json:"slogan"` // 为家庭财产、家庭成员提供全方位、综合保障。
			} `json:"YTHMALL"`
			BelongCode string `json:"belongCode"` // property
			Desc       string `json:"desc"`       // 100元
			Extra      []struct {
				Name      string   `json:"name"`      // 特别约定
				Value     string   `json:"value"`     //
				ValueList []string `json:"valueList"` // 家财险主险及附加险每次事故绝对免赔额200元；
			} `json:"extra"`
			Freeway   string `json:"freeway"` // 家财险主险及附加险每次事故绝对免赔额200元；
			Introduce struct {
				ItemList []struct {
					Name  string `json:"name"`  // 产品介绍
					Value string `json:"value"` // 为家庭财产、家庭成员提供全方位、综合保障。
				} `json:"itemList"`
				ObjList []struct {
					Name  string `json:"name"`  // 适用人群
					Value string `json:"value"` // 适用于个人及家庭的自有或使用的<br>家庭财产
				} `json:"objList"`
			} `json:"introduce"`
			Morefreeway bool   `json:"morefreeway"` // true
			Name        string `json:"name"`        // 优选款
			Notice      []struct {
				Name  string `json:"name"`  // 投保须知1
				Value string `json:"value"` // 1、本保险支持在32个地区销售，包括北京、天津、河北、山西、内蒙、辽宁、吉林、黑龙江、上海、江苏、浙江、安徽、福建、江西、山东、河南、湖北、湖南、广东、广西、重庆、四川、云南、陕西、贵州、甘肃、新疆、宁波、青岛、深圳、大连、厦门。
			} `json:"notice"`
			Premium        int64 `json:"premium"` // 100
			ProdObjectData struct {
				COVERAGE []struct {
					ID               int64  `json:"id"`               // 1
					Code             string `json:"code"`             // CP00000001
					DeductibleAmount int64  `json:"deductibleAmount"` // 0
					DeductibleRate   int64  `json:"deductibleRate"`   // 0
					Desc             string `json:"desc"`             // 50万
					Name             string `json:"name"`             // 房屋主体
					PerSumInsured    int64  `json:"perSumInsured"`    // 0
					Premium          int64  `json:"premium"`          // 0
					SumInsured       int64  `json:"sumInsured"`       // 500000
				} `json:"COVERAGE"`
			} `json:"prodObjectData"`
			SchemeID int64 `json:"schemeId"` // 1
		} `json:"prodSchemeList"`
		ProductType string `json:"productType"` // property
		ShowWSTB    bool   `json:"showWSTB"`    // true
		URL         string `json:"url"`         // /images/anju_b.jpg
	} `json:"data"`
	Message string `json:"message"` // 加载产品完成
	Success bool   `json:"success"` // true
}
