## 使用方法
随便访问一下福利中心  
抓包`https://h5.if.qidian.com`任意链接中的以下四个值:  
+ header中的  
  `SDKSign`
+ header->cookie中的  
  `QDInfo`
  `ywkey` (每月需要更新一次)
  `ywguid`   
放入`cmd/main.go`中
## 其他注意事项
+ 其实可以瞬间执行完,但是不知道他有没有风险控制,暂定15s间隔(`task.go L24`)
+ 可能会导致任何的损失(如封号),概不负责
## 破解流程如下:
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
# Todo
+ [ ] ywkey 每月需要更新,目前没有自动每月更新
+ MoreRewardTab
+ [x] 104=前往游戏中心玩游戏10分钟奖励10点币
  + 目前没有破解imei相关的加密,多账号时只能使用多台手机的cookie,否则后台会不计算时长
+ [ ] 103=前往游戏中心任意一款游戏充值1次奖励30点币
+ [ ] 121=签到互动多重福利(微博)/登陆携程领积分当钱花(好像只能领一次)
+ [ ] 222=打开推送通知，次日（24h）后可领取奖励(好像只能领一次)
