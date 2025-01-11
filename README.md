# 老版本的破解流程如下:
https://www.jianshu.com/p/58ec69e04983
QDInfo 算法如下:
```python
import base64
from Crypto.Cipher import DES3

key = b'0821CAAD409B84020821CAAD'
data = base64.b64decode('SO+aPyWTJ02k4C9FkkB29fACDXIsJx4pAGbhVI07D8hjHPOEsCFgpJ99gS3kYIjunO+UrcWbhPgIlUSo3XxdoisFnouWF80qfP+9nYAPZWuWE/x7ukJhxq8DEJW+n90UAoC6t3e9KFYaJ/yFFUfggDVS6xpzIkTxCCDps2WxRBcdvOXoA5I5/i3jrw8wJqw0DmbxzkSOoKB1T5VHx/VjWCoYTuW8fA5DlGMQL+4lQldYUANNM1Aarp6oD16p7Rqc9JpGyHOOnKF3tDxv8vGv0ElZszGBKKqK70o3d0OzvfmgFhyXErR92g==')
cryptor = DES3.new(key, DES3.MODE_CBC, b'00000000')
print(cryptor.decrypt(data))
```
SDKSign 算法如下:
```python
import base64

from Crypto.Cipher import DES3

key = b'8YV#U2Butm,VutR2B_W[*}6t'
data = base64.b64decode('fwU0VSlfsV/NtCFBjpJarbYpi9mlbLU/EDzhOVoz2RdtheX+SLpjTy8L2+gA InschgJSs1O5vbtFpSZ6+GPI8iEd6QhtwlTz8ODKLNM1r+aH0A8sY5+lP6la DPt/GpDgPvW5ZvKHiqnIqFEJHRoPYEshR2+cAq03JfcYLPvSfE7DpuHLVA2F mRtLGCdVWmTujc/5Lb+/Cmk=')
cryptor = DES3.new(key, DES3.MODE_CBC, b'01234567')
print(cryptor.decrypt(data))
```

# 有什么用
