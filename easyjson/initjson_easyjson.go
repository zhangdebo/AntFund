// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package easyjson

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson1add0730DecodeBookStackEasyjson(in *jlexer.Lexer, out *URIBroker) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "favicon.ico.url":
			out.FaviconIcoURL = string(in.String())
		case "app.404.url":
			out.App404URL = string(in.String())
		case "zdrmdata.rest.url":
			out.ZdrmdataRestURL = string(in.String())
		case "app.errorpage.url":
			out.AppErrorpageURL = string(in.String())
		case "authcenter.url":
			out.AuthcenterURL = string(in.String())
		case "app.goto.url":
			out.AppGotoURL = string(in.String())
		case "bumng.url":
			out.BumngURL = string(in.String())
		case "omeo.check.url":
			out.OmeoCheckURL = string(in.String())
		case "omeo.get.url":
			out.OmeoGetURL = string(in.String())
		case "assets.url":
			out.AssetsURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1add0730EncodeBookStackEasyjson(out *jwriter.Writer, in URIBroker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"favicon.ico.url\":"
		out.RawString(prefix[1:])
		out.String(string(in.FaviconIcoURL))
	}
	{
		const prefix string = ",\"app.404.url\":"
		out.RawString(prefix)
		out.String(string(in.App404URL))
	}
	{
		const prefix string = ",\"zdrmdata.rest.url\":"
		out.RawString(prefix)
		out.String(string(in.ZdrmdataRestURL))
	}
	{
		const prefix string = ",\"app.errorpage.url\":"
		out.RawString(prefix)
		out.String(string(in.AppErrorpageURL))
	}
	{
		const prefix string = ",\"authcenter.url\":"
		out.RawString(prefix)
		out.String(string(in.AuthcenterURL))
	}
	{
		const prefix string = ",\"app.goto.url\":"
		out.RawString(prefix)
		out.String(string(in.AppGotoURL))
	}
	{
		const prefix string = ",\"bumng.url\":"
		out.RawString(prefix)
		out.String(string(in.BumngURL))
	}
	{
		const prefix string = ",\"omeo.check.url\":"
		out.RawString(prefix)
		out.String(string(in.OmeoCheckURL))
	}
	{
		const prefix string = ",\"omeo.get.url\":"
		out.RawString(prefix)
		out.String(string(in.OmeoGetURL))
	}
	{
		const prefix string = ",\"assets.url\":"
		out.RawString(prefix)
		out.String(string(in.AssetsURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v URIBroker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1add0730EncodeBookStackEasyjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v URIBroker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1add0730EncodeBookStackEasyjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *URIBroker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1add0730DecodeBookStackEasyjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *URIBroker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1add0730DecodeBookStackEasyjson(l, v)
}
func easyjson1add0730DecodeBookStackEasyjson1(in *jlexer.Lexer, out *TitleInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fundLimit":
			out.FundLimit = string(in.String())
		case "netValue":
			out.NetValue = string(in.String())
		case "netValueDate":
			out.NetValueDate = string(in.String())
		case "profitSevenDays":
			out.ProfitSevenDays = string(in.String())
		case "profitTenThousand":
			out.ProfitTenThousand = string(in.String())
		case "dayOfGrowth":
			out.DayOfGrowth = string(in.String())
		case "lastWeek":
			out.LastWeek = string(in.String())
		case "riskEvaluation":
			out.RiskEvaluation = string(in.String())
		case "establishmentDate":
			out.EstablishmentDate = string(in.String())
		case "assetSize":
			out.AssetSize = string(in.String())
		case "fundManagerName":
			out.FundManagerName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1add0730EncodeBookStackEasyjson1(out *jwriter.Writer, in TitleInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fundLimit\":"
		out.RawString(prefix[1:])
		out.String(string(in.FundLimit))
	}
	{
		const prefix string = ",\"netValue\":"
		out.RawString(prefix)
		out.String(string(in.NetValue))
	}
	{
		const prefix string = ",\"netValueDate\":"
		out.RawString(prefix)
		out.String(string(in.NetValueDate))
	}
	{
		const prefix string = ",\"profitSevenDays\":"
		out.RawString(prefix)
		out.String(string(in.ProfitSevenDays))
	}
	{
		const prefix string = ",\"profitTenThousand\":"
		out.RawString(prefix)
		out.String(string(in.ProfitTenThousand))
	}
	{
		const prefix string = ",\"dayOfGrowth\":"
		out.RawString(prefix)
		out.String(string(in.DayOfGrowth))
	}
	{
		const prefix string = ",\"lastWeek\":"
		out.RawString(prefix)
		out.String(string(in.LastWeek))
	}
	{
		const prefix string = ",\"riskEvaluation\":"
		out.RawString(prefix)
		out.String(string(in.RiskEvaluation))
	}
	{
		const prefix string = ",\"establishmentDate\":"
		out.RawString(prefix)
		out.String(string(in.EstablishmentDate))
	}
	{
		const prefix string = ",\"assetSize\":"
		out.RawString(prefix)
		out.String(string(in.AssetSize))
	}
	{
		const prefix string = ",\"fundManagerName\":"
		out.RawString(prefix)
		out.String(string(in.FundManagerName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TitleInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1add0730EncodeBookStackEasyjson1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TitleInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1add0730EncodeBookStackEasyjson1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TitleInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1add0730DecodeBookStackEasyjson1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TitleInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1add0730DecodeBookStackEasyjson1(l, v)
}
func easyjson1add0730DecodeBookStackEasyjson2(in *jlexer.Lexer, out *MaterialInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "productId":
			out.ProductID = string(in.String())
		case "fundCode":
			out.FundCode = string(in.String())
		case "fundType":
			out.FundType = string(in.String())
		case "titleInfo":
			(out.TitleInfo).UnmarshalEasyJSON(in)
		case "fundBrief":
			(out.FundBrief).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1add0730EncodeBookStackEasyjson2(out *jwriter.Writer, in MaterialInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"productId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ProductID))
	}
	{
		const prefix string = ",\"fundCode\":"
		out.RawString(prefix)
		out.String(string(in.FundCode))
	}
	{
		const prefix string = ",\"fundType\":"
		out.RawString(prefix)
		out.String(string(in.FundType))
	}
	{
		const prefix string = ",\"titleInfo\":"
		out.RawString(prefix)
		(in.TitleInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"fundBrief\":"
		out.RawString(prefix)
		(in.FundBrief).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MaterialInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1add0730EncodeBookStackEasyjson2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MaterialInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1add0730EncodeBookStackEasyjson2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MaterialInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1add0730DecodeBookStackEasyjson2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MaterialInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1add0730DecodeBookStackEasyjson2(l, v)
}
func easyjson1add0730DecodeBookStackEasyjson3(in *jlexer.Lexer, out *JsonInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			out.Success = bool(in.Bool())
		case "message":
			out.Message = string(in.String())
		case "materialInfo":
			(out.MaterialInfo).UnmarshalEasyJSON(in)
		case "isLogin":
			out.IsLogin = bool(in.Bool())
		case "csrf":
			out.Csrf = string(in.String())
		case "pageName":
			out.PageName = string(in.String())
		case "uriBroker":
			(out.URIBroker).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1add0730EncodeBookStackEasyjson3(out *jwriter.Writer, in JsonInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"success\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Success))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"materialInfo\":"
		out.RawString(prefix)
		(in.MaterialInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"isLogin\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsLogin))
	}
	{
		const prefix string = ",\"csrf\":"
		out.RawString(prefix)
		out.String(string(in.Csrf))
	}
	{
		const prefix string = ",\"pageName\":"
		out.RawString(prefix)
		out.String(string(in.PageName))
	}
	{
		const prefix string = ",\"uriBroker\":"
		out.RawString(prefix)
		(in.URIBroker).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JsonInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1add0730EncodeBookStackEasyjson3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JsonInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1add0730EncodeBookStackEasyjson3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JsonInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1add0730DecodeBookStackEasyjson3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JsonInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1add0730DecodeBookStackEasyjson3(l, v)
}
func easyjson1add0730DecodeBookStackEasyjson4(in *jlexer.Lexer, out *GeneralInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fundName":
			out.FundName = string(in.String())
		case "establishmentDate":
			out.EstablishmentDate = string(in.String())
		case "fundCode":
			out.FundCode = string(in.String())
		case "assetSize":
			out.AssetSize = string(in.String())
		case "fundCompanyName":
			out.FundCompanyName = string(in.String())
		case "trusteeName":
			out.TrusteeName = string(in.String())
		case "fundManagerBackground":
			out.FundManagerBackground = string(in.String())
		case "fundManagerInfoList":
			if in.IsNull() {
				in.Skip()
				out.FundManagerInfoList = nil
			} else {
				in.Delim('[')
				if out.FundManagerInfoList == nil {
					if !in.IsDelim(']') {
						out.FundManagerInfoList = make([]FundManagerInfoList, 0, 1)
					} else {
						out.FundManagerInfoList = []FundManagerInfoList{}
					}
				} else {
					out.FundManagerInfoList = (out.FundManagerInfoList)[:0]
				}
				for !in.IsDelim(']') {
					var v1 FundManagerInfoList
					(v1).UnmarshalEasyJSON(in)
					out.FundManagerInfoList = append(out.FundManagerInfoList, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "investPhilosophy":
			out.InvestPhilosophy = string(in.String())
		case "investStrategy":
			out.InvestStrategy = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1add0730EncodeBookStackEasyjson4(out *jwriter.Writer, in GeneralInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fundName\":"
		out.RawString(prefix[1:])
		out.String(string(in.FundName))
	}
	{
		const prefix string = ",\"establishmentDate\":"
		out.RawString(prefix)
		out.String(string(in.EstablishmentDate))
	}
	{
		const prefix string = ",\"fundCode\":"
		out.RawString(prefix)
		out.String(string(in.FundCode))
	}
	{
		const prefix string = ",\"assetSize\":"
		out.RawString(prefix)
		out.String(string(in.AssetSize))
	}
	{
		const prefix string = ",\"fundCompanyName\":"
		out.RawString(prefix)
		out.String(string(in.FundCompanyName))
	}
	{
		const prefix string = ",\"trusteeName\":"
		out.RawString(prefix)
		out.String(string(in.TrusteeName))
	}
	{
		const prefix string = ",\"fundManagerBackground\":"
		out.RawString(prefix)
		out.String(string(in.FundManagerBackground))
	}
	{
		const prefix string = ",\"fundManagerInfoList\":"
		out.RawString(prefix)
		if in.FundManagerInfoList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.FundManagerInfoList {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"investPhilosophy\":"
		out.RawString(prefix)
		out.String(string(in.InvestPhilosophy))
	}
	{
		const prefix string = ",\"investStrategy\":"
		out.RawString(prefix)
		out.String(string(in.InvestStrategy))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GeneralInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1add0730EncodeBookStackEasyjson4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GeneralInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1add0730EncodeBookStackEasyjson4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GeneralInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1add0730DecodeBookStackEasyjson4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GeneralInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1add0730DecodeBookStackEasyjson4(l, v)
}
func easyjson1add0730DecodeBookStackEasyjson5(in *jlexer.Lexer, out *FundManagerInfoList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "key":
			out.Key = string(in.String())
		case "fundName":
			out.FundName = string(in.String())
		case "officeDate":
			out.OfficeDate = string(in.String())
		case "earnings":
			out.Earnings = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1add0730EncodeBookStackEasyjson5(out *jwriter.Writer, in FundManagerInfoList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix[1:])
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"fundName\":"
		out.RawString(prefix)
		out.String(string(in.FundName))
	}
	{
		const prefix string = ",\"officeDate\":"
		out.RawString(prefix)
		out.String(string(in.OfficeDate))
	}
	{
		const prefix string = ",\"earnings\":"
		out.RawString(prefix)
		out.String(string(in.Earnings))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FundManagerInfoList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1add0730EncodeBookStackEasyjson5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FundManagerInfoList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1add0730EncodeBookStackEasyjson5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FundManagerInfoList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1add0730DecodeBookStackEasyjson5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FundManagerInfoList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1add0730DecodeBookStackEasyjson5(l, v)
}
func easyjson1add0730DecodeBookStackEasyjson6(in *jlexer.Lexer, out *FundBrief) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fundNameAbbr":
			out.FundNameAbbr = string(in.String())
		case "fundName":
			out.FundName = string(in.String())
		case "fundCode":
			out.FundCode = string(in.String())
		case "establishmentDate":
			out.EstablishmentDate = string(in.String())
		case "shareSize":
			out.ShareSize = string(in.String())
		case "assetSize":
			out.AssetSize = string(in.String())
		case "fundManagerName":
			out.FundManagerName = string(in.String())
		case "saleStatus":
			out.SaleStatus = string(in.String())
		case "fundCompanyName":
			out.FundCompanyName = string(in.String())
		case "trusteeName":
			out.TrusteeName = string(in.String())
		case "manageRate":
			out.ManageRate = string(in.String())
		case "trusteeRate":
			out.TrusteeRate = string(in.String())
		case "purchaseMinMount":
			out.PurchaseMinMount = string(in.String())
		case "redeemMinMount":
			out.RedeemMinMount = string(in.String())
		case "purchaseRatio":
			out.PurchaseRatio = string(in.String())
		case "redeemRatio":
			out.RedeemRatio = string(in.String())
		case "generalInfo":
			(out.GeneralInfo).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1add0730EncodeBookStackEasyjson6(out *jwriter.Writer, in FundBrief) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fundNameAbbr\":"
		out.RawString(prefix[1:])
		out.String(string(in.FundNameAbbr))
	}
	{
		const prefix string = ",\"fundName\":"
		out.RawString(prefix)
		out.String(string(in.FundName))
	}
	{
		const prefix string = ",\"fundCode\":"
		out.RawString(prefix)
		out.String(string(in.FundCode))
	}
	{
		const prefix string = ",\"establishmentDate\":"
		out.RawString(prefix)
		out.String(string(in.EstablishmentDate))
	}
	{
		const prefix string = ",\"shareSize\":"
		out.RawString(prefix)
		out.String(string(in.ShareSize))
	}
	{
		const prefix string = ",\"assetSize\":"
		out.RawString(prefix)
		out.String(string(in.AssetSize))
	}
	{
		const prefix string = ",\"fundManagerName\":"
		out.RawString(prefix)
		out.String(string(in.FundManagerName))
	}
	{
		const prefix string = ",\"saleStatus\":"
		out.RawString(prefix)
		out.String(string(in.SaleStatus))
	}
	{
		const prefix string = ",\"fundCompanyName\":"
		out.RawString(prefix)
		out.String(string(in.FundCompanyName))
	}
	{
		const prefix string = ",\"trusteeName\":"
		out.RawString(prefix)
		out.String(string(in.TrusteeName))
	}
	{
		const prefix string = ",\"manageRate\":"
		out.RawString(prefix)
		out.String(string(in.ManageRate))
	}
	{
		const prefix string = ",\"trusteeRate\":"
		out.RawString(prefix)
		out.String(string(in.TrusteeRate))
	}
	{
		const prefix string = ",\"purchaseMinMount\":"
		out.RawString(prefix)
		out.String(string(in.PurchaseMinMount))
	}
	{
		const prefix string = ",\"redeemMinMount\":"
		out.RawString(prefix)
		out.String(string(in.RedeemMinMount))
	}
	{
		const prefix string = ",\"purchaseRatio\":"
		out.RawString(prefix)
		out.String(string(in.PurchaseRatio))
	}
	{
		const prefix string = ",\"redeemRatio\":"
		out.RawString(prefix)
		out.String(string(in.RedeemRatio))
	}
	{
		const prefix string = ",\"generalInfo\":"
		out.RawString(prefix)
		(in.GeneralInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FundBrief) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1add0730EncodeBookStackEasyjson6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FundBrief) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1add0730EncodeBookStackEasyjson6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FundBrief) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1add0730DecodeBookStackEasyjson6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FundBrief) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1add0730DecodeBookStackEasyjson6(l, v)
}
