// Code generated by goctl. DO NOT EDIT.
package types

type GetUpdateResponse struct {
	Version   string `json:"version"`
	Link      string `json:"link"`
	ChangeLog string `json:"changeLog"`
}

type RefreshUpdateReq struct {
	Secret string `json:"secret"`
}

type RefreshUpdateResp struct {
	Code int `json:"code"`
}
