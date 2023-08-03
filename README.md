# bread-app

## 動作環境

Windows 10 v22H2

## プロジェクト構成

.
├─*bread-app-api* graphQLサーバ
├─*bread-app-cli* パンの情報を取得し、DBに保存するCLI
├─*bread-app-common* 共通ライブラリ
├─*mock_request_bread-app-api* モックリクエスト用ファイル
└─*terraform* firestore用のterraform

### bread-app-api

bread-app-api
├─*domain*
│  ├─*dao* DBアクセスinterface
│  └─*entity* DB Entity
├─*graph*
│  ├─*model* graphQL model
│  ├─*resolvers* Query実装, DI管理
│  └─*services* ユースケース層
├─*infrastructure*
│  └─*dao* DBアクセス実装
├─*internal* gqlgen interface
└─*tools*

### bread-app-cli

bread-app-cli
├─*cmd* コマンド
├─*config* 設定関連
├─*domain*
│  ├─*dao* DBアクセスinterface
│  └─*entity* DB Entity
├─*infrastructure*
│  ├─*contentful* Contentfulアクセス
│  └─*dao* DBアクセス実装
└─*services* ユースケース層

### bread-app-common

bread-app-common
├─*core* firebaseクライアントをシングルトンで管理
└─*errors* bread-appエラー

## 環境構築

### terraform

作成済のfirestoreを使用する場合、9-3から実施してください.

1. terraformをインストール

    * 参考: [Install Terraform](https://developer.hashicorp.com/terraform/downloads)

2. gcloud CLIをインストール

    * 参考: [gcloud CLI をインストールする](https://cloud.google.com/sdk/docs/install?hl=ja)

3. Google認証情報のセット

    以下のコマンドを実行し、ブラウザからGoogle認証を行う.

    ```sh
    ~github.com/miyamo2theppl/bread-app$ gcloud auth application-default login
    ```

    ターミナルに認証ファイルの格納場所が表示されていれば成功.

    ```sh
    ~github.com/miyamo2theppl/bread-app$ Credentials saved to file: [{APPDATA}\gcloud\application_default_credentials.json]
    ```

4. 作業ディレクトリへ移動する.

    ```sh
    ~github.com/miyamo2theppl/bread-app/terraform$ cd ./terraform
    ```

5. ワークスペースの初期化.

    ```sh
    ~github.com/miyamo2theppl/bread-app/terraform$ terraform init
    ```

6. プロジェクトID採番用のタイムスタンプを変更する.

    対象ファイル: terraform.tfvars

    ```terraform.tfvars
    unix_timestamp = "{unix時間(整数)}"
    ```

7. 変更セットの確認.

    ```sh
    ~github.com/miyamo2theppl/bread-app/terraform$ terraform plan
    ```

8. デプロイ.

    ```sh
    ~github.com/miyamo2theppl/bread-app/terraform$ terraform apply
    ```

9. Firebaseの認証ファイルを取得.

    1. ブラウザからFireBaseコンソールを開き、以下のページを開く.

        Bread > プロジェクトの設定 > サービスアカウント > Firebase Admin SDK

    2. 「新しい秘密鍵を生成」ボタン押下.

    3. ダウンロードされたファイルを「firebase_credential.json」にリネームし、以下のディレクトリに格納.

     * Windows: %USERPROFILE%/.config/breadapp/

     * Linux/Unix/Mac: $HOME/.config/breadapp/

### bread-app-cli

1. 依存パッケージのインストール.

    ```sh
    ~github.com/miyamo2theppl/bread-app$ cd ./bread-app-cli&go mod tidy
    ```

2. ビルド.

    ```sh
    ~github.com/miyamo2theppl/bread-app/bread-app-cli$ cd ../&go build github.com/miyamo2theppl/bread-app/bread-app-cli
    ```

3. 設定ファイルを格納.

    * フォーマット

        ```json
        {
            "contentful": {
                "token": {Contentfulのアクセストークン}
            }
        }
        ```

    * 格納場所
        * Windows: %USERPROFILE%/.config/breadapp/

        * Linux/Unix/Mac: $HOME/.config/breadapp/

### bread-app-api

1. 依存パッケージのインストール.

    ```sh
    ~github.com/miyamo2theppl/bread-app$ cd ./bread-app-api&go mod tidy
    ```

2. ビルド.

    ```sh
    ~github.com/miyamo2theppl/bread-app/bread-app-api$ cd ../&go build github.com/miyamo2theppl/bread-app/bread-app-api
    ```

## 動作確認

### bread-app-cli

以下のコマンドを実行.

```sh
~github.com/miyamo2theppl/bread-app$ bread-app-cli fetch 6QRk7gQYmOyJ1eMG9H4jbB 41RUO5w4oIpNuwaqHuSwEc 4Li6w5uVbJNVXYVxWjWVoZ
```

また、任意のコマンドライン引数でContentfulから取得するデータを変更することも可能.

```sh
~github.com/miyamo2theppl/bread-app$ bread-app-cli fetch Bread1 Bread2 Bread3
```

### bread-app-api

1. Rest Client(VSCode拡張)のインストール.

    * インストール: [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

    * 参考: [VS Code上でHTTPリクエストを送信し、VS Code上でレスポンスを確認できる「REST Client」拡張の紹介](https://qiita.com/toshi0607/items/c4440d3fbfa72eac840c)

2. graphQLサーバを起動.

    1. 以下のコマンドを実行.

        ```sh
        ~github.com/miyamo2theppl/bread-app$ bread-app-api
        ```

    2. <http://localhost:8080/>にアクセスし、GraphQL playgroundが表示されていれば起動成功.

3. 以下のhttpファイルでgraphQLサーバにリクエストを行う.

    * mock_request_bread-app-api/bread.http

    * mock_request_bread-app-api/breads.http
