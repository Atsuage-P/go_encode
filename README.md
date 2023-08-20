# Base64によるエンコードとソートの練習
内容は[Qiita](https://qiita.com/igossou/items/209beb900481f68bfe40)にまとめた

## やったこと
- `application/json`リクエストをGo構造体にデコード
- JSONのBase64エンコード
- エンコードしたJSONのBase64デコード
- 日付によるソート

## メモ
- `json.NewDecoder`と`json.Unmarshal`はどちらもJSONをGoのデータ構造に変換する

- `json.NewDecoder`の使い所
  - JSONを逐次的に読み込んで処理するため大規模なJSONデータに適する
  - `io.Reader`インターフェースを実装する入力ストリーム(ファイル、ネットワーク接続、httpリクエスト)などのデータを変換する

- `json.Unmarshal`の使い所
  - JSONを一度にメモリに読み込んで処理するため小規模なJSONデータに適する
  - バイトスライスや文字列といった既存のデータを直接変換する

- `time.Time`型の日時比較は`Equal, After, Before`が便利そう
