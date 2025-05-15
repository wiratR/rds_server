sample json

```json
{
    "appid":"mch38806",
    "auth_code":12345,
    "channel":"wechat",
    "fee_type":"THB",
    "mch_order_no":"20240718020717107",
    "nonce_str":"OsD0operator_id=",
    "time_stamp":"20240718020717",
    "total_fee":100,
}


{
"appid":"mch38806",
"channel":"promtpay",
"fee_type":"THB",
"mch_order_no":"20240725100814042",
"nonce_str":"iQig",
"time_stamp":"20240725100814",
"total_fee":"100",
}
```

https://gateway.ksher.com/demo_sign.html

https://api.ksher.net/KsherAPI/dev/signature_algo.html#_verify_signature



### Run backend server

```bash
cd ./backend

docker-compose up -d --build
```

### Run wallet app client

```bash
cd wallet_app
flutter run
```
### this for swagger API

```bash
http://127.0.0.1:8000/swagger/index.html
```

### this for databse 

```bash
POSTGRES_HOST=127.0.0.1
POSTGRES_PORT=6500
POSTGRES_USER=admin
POSTGRES_PASSWORD=P@ssw0rd
POSTGRES_DB=rds_db
```