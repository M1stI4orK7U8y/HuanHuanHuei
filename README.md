# HuanHuanHuei (換換匯)

數位貨幣互相換匯的小東西

A digital token exchanger

# 功能清單
* [x]  btc <-> eth 互換 ( btc and eth bidirection exchange)
* [ ]  ERC20 <-> any 互換 ( ERC20 and any support token bidirection exchange )
* [ ]  Log紀錄 ( enhanced log funcion )
* [ ]  CI/CD

# 對外API Restful API for user agent

* 查詢紀錄 query record GET: /api/v1/record/:id 

* 查詢多個紀錄 query records POST: /api/v1/record 

| key | description |
| ------ | ------ |
| ids | txid1 |
| ids | txid2 | 

* 請求換匯 request exchange token POST: /api/v1/huanhuanhuei

| key | description | example |
| ------ | ------ | ------ | 
| txid | 傳送給官方地址的交易hash (txhash of tx which is send amount to official address) | "fa9ff813525f7e2e2901ae2634ff634d5c7aff3f6399347f564f6b554b1344d8" |
| receiver | 目標貨幣的收款地址 receiver address of target token| "0xA1820474b14aA475c8dCc8133773243DD4c98722"
| from | 來源貨幣代碼(詳見表) code of source token | 0 |
| to | 目標貨幣代碼 code of target token | 1 |

## 貨幣代碼 token code
| token | value | 
| ------ | ------ | 
| BTC | 0 | 
| ETH | 1 | 

# 架構圖 Architecture

請參照architecture.png

# 用法 Usage

## 事前準備 Prepare
- BTC官方地址 BTC official address
- ETH官方地址 ETH official address
- config.ymal: 可從config-template.yaml修改 can modify from config-template.yaml

## 貨幣A -> 貨幣B token A -> token B
- 用戶傳送交易給官方指定貨幣帳戶 User send a tx to specific token official address
- 用戶端呼叫/api/v1/huanhuanhuei  User agent call api

## 查詢狀態 query record
- 用戶端呼叫/api/v1/record/:txid  User agent call api with id
- 用戶端呼叫/api/v1/record  User agent call api with ids

# License
The HuanHuanHuei is licensed under the GNU Lesser General Public License v3.0, also included in our repository in the LICENSE file.
