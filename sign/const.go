package sign

const signatures = "308202253082018ea00302010202044e239460300d06092a864886f70d0101050500305731173015060355040a0c0ec386c3b0c2b5c3a3c396c390c38e311d301b060355040b0c14c386c3b0c2b5c3a3c396c390c38ec384c38dc3b8311d301b06035504030c14c386c3b0c2b5c3a3c396c390c38ec384c38dc3b8301e170d3131303731383032303331325a170d3431303731303032303331325a305731173015060355040a0c0ec386c3b0c2b5c3a3c396c390c38e311d301b060355040b0c14c386c3b0c2b5c3a3c396c390c38ec384c38dc3b8311d301b06035504030c14c386c3b0c2b5c3a3c396c390c38ec384c38dc3b830819f300d06092a864886f70d010101050003818d0030818902818100a3d47f8bfd8d54de1dfbc40a9caa88a43845e287e8f40da2056be126b17233669806bfa60799b3d1364e79a78f355fd4f72278650b377e5acc317ff4b2b3821351bcc735543dab0796c716f769c3a28fedc3bca7780e5fff6c87779f3f3cdec6e888b4d21de27df9e7c21fc8a8d9164bfafac6df7d843e59b88ec740fc52a3c50203010001300d06092a864886f70d0101050500038181001f7946581b8812961a383b2d860b89c3f79002d46feb96f2a505bdae57097a070f3533c42fc3e329846886281a2fbd5c87685f59ab6dd71cc98af24256d2fbf980ded749e2c35eb0151ffde993193eace0b4681be4bcee5f663dd71dd06ab64958e02a60d6a69f21290cb496dd8784a4c31ebadb1b3cc5cb0feebdaa2f686ee2"

var (
	SDKSignPass = []byte("8YV#U2Butm,VutR2B_W[*}6t")
	SDKSignIV   = []byte("01234567")
	QDInfoPass  = []byte("0821CAAD409B84020821CAAD")
	QDInfoIV    = []byte("00000000")
)
var hashedSignatures = hash(signatures)

const UNKNOWNStr = "未知字符串"
const UNKNOWNInt = "未知int"

const FiledTimestamp = "时间戳"
const FiledAppversion = "App版本号"
const FiledHashUrl = "Hash(URL)"
const FiledHashSignatures = "hash(Signatures)"
const FiledQIMEI = "QIMEI"
const FiledUUID = "系统UUID"
const FiledSystem = "系统设备类型"

var QDInfoStruct = []string{
	FiledUUID,       //SRV	U55d98d04b89bfe8c10001ab17310
	FiledAppversion, //7.9.384
	"系统分辨率",         //1080
	"系统版本号",         //2296
	"source",        //1002140
	"系统安卓版本号",       //14
	UNKNOWNInt,      //1
	FiledSystem,     //2106118C
	"系统版本Code",      //1466
	"source2",       //1000014
	UNKNOWNInt,      //4
	"用户token",       //344203808
	FiledTimestamp,  //1735543229551
	UNKNOWNInt,      //0
	FiledQIMEI,      //c517bf9e55d98d04b89bfe8c10001ab17310
	UNKNOWNStr,      //be939f2c6192e1ae
	UNKNOWNStr,      //
	UNKNOWNStr,      //45fff87862ce33d8
	UNKNOWNStr,      //caa8f0c6ffb801753eddeb87100017417310
	UNKNOWNInt,      //0
}

var SDKSignStruct = []string{
	"标志符号",              //qYJ]Q9FYhq?
	FiledTimestamp,      //1736561971779
	UNKNOWNStr,          //344203808
	FiledQIMEI,          //c517bf9e55d98d04b89bfe8c10001ab17310
	UNKNOWNInt,          //1
	"App版本号",            //7.9.384
	UNKNOWNInt,          //0
	FiledHashUrl,        //d41d8cd98f00b204e9800998ecf8427e
	FiledHashSignatures, //f189adc92b816b3e9da29ea304d4a7e4
}
