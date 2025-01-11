package sign

import (
	"fmt"
	"testing"
)

type SDKTestCase struct {
	SDKSign string
	Uri     string
}

const PhoneQDInfo = "SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWviV9DlRVtRllf2xwn3SDILgUVhzxJXrTRNcGeynaP07zVZ5qe7MsQgKlxQdWxM6mFdhlYjvxrV+vON3pGlpR6i99QzTXesSmhLrNUXyxfExycfosXSayIx7cg++mgVmMuYlzq0lHdkLtE9Xy/osz0yxcsC+f+qlmWM7h/koIE014cWbRBsHjwBQJJkU6h2fa5lDRNhYZLnfQ=="
const PhoneSDKSign = "fwU0VSlfsV/NtCFBjpJarbYpi9mlbLU/EDzhOVoz2RdtheX+SLpjTy8L2+gA InschgJSs1O5vbtFpSZ6+GPI8iEd6QhtwlTz8ODKLNM1r+aH0A8sY5+lP6la DPt/GpDgPvW5ZvKHiqnIqFEJHRoPYEshR2+cAq03JfcYLPvSfE7DpuHLVA2F mRtLGCdVWmTujc/5Lb+/Cmk="

var SDKTestCases = []SDKTestCase{
	SDKTestCase{
		SDKSign: "fwU0VSlfsV/NtCFBjpJarbYpi9mlbLU/EDzhOVoz2RdtheX+SLpjTy8L2+gA InschgJSs1O5vbtFpSZ6+GPI8iEd6QhtwlTz8ODKLNM1r+aH0A8sY5+lP6la DPt/GpDgPvW5ZvKHiqnIqFEJHRoPYEshR2+cAq03JfcYLPvSfE7DpuHLVA2F mRtLGCdVWmTujc/5Lb+/Cmk=",
		Uri:     "https://h5.if.qidian.com/argus/api/v2/checkin/checkin",
	},
	{
		SDKSign: "fwU0VSlfsV/NtCFBjpJarYVUiYqJXavL/LAgO3IV/xUuHaTBj607dzH/S0kA eLzq4qWCJScF1reZMyzMpPDPlxQyitc41n+g3bhPhEBEt18HjXk9LWBV0spB oSuafeDBstlCN0ZL0Dh1D71dG5ln/pui76DopJLGSz3jaGCXKBHJkn2kDaSW eqvggFjm8zGBA7YjK2QUl5E=",
		Uri:     "https://h5.if.qidian.com/argus/api/v2/video/adv/mainPage",
	},
	{
		SDKSign: "fwU0VSlfsV/NtCFBjpJarebgN6GCNJijvqoi7Fp77n+TQH4qyVwdiMaZ5J44 e6JHXNj9mfUS5lMGivZtsljj1/Pi8ojlRiYReSL+bWIgn1I+q3IQKfepVbgS mXKTXoXjBkBIx7vTrQwIZUYuH5T9a54ixOluPscfLr020RlX3pjxOfIOOTjK ITIHBvtgyknsKSETU1Ok3x0=",
		Uri:     "https://h5.if.qidian.com/argus/api/v1/video/adv/finishWatch?taskId=932958800725147650&BanId=0&BanMessage=&CaptchaAId=&CaptchaType=0&CaptchaURL=&Challenge=&Gt=&NewCaptcha=0&Offline=0&PhoneNumber=&SessionKey=",
	},
}

func TestQDInfo(t *testing.T) {
	meta, err := NewMeta(PhoneQDInfo, PhoneSDKSign)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", meta)

}

func TestSDKSign(t *testing.T) {
	for _, testCase := range SDKTestCases {
		uri := hash(normParams(testCase.Uri))
		meta, err := NewMeta(PhoneQDInfo, testCase.SDKSign)
		if err != nil {
			t.Fatal(err)
		}
		get, err := meta.SdkRW.SDkGet(FiledHashUrl)
		if err != nil {
			return
		}
		if uri != get {
			t.Error("not equal")
		}
		get, err = meta.SdkRW.SDkGet(FiledHashSignatures)
		fmt.Printf("%v\n", get)
	}
}
func TestDecrypt(t *testing.T) {
	sign, err := DecryptSDKSign(PhoneSDKSign)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", sign)
	info, err := DecryptQDInfo(PhoneQDInfo)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", info)
}
