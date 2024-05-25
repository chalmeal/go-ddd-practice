**go-ddd-practice**

### 構成
```
├── .vscode
├── pkg
|　　　├── controller
|　　　|　　  ├── response
|　　　|　　  |　　　└── response.go
|　　　|　　  ├── app_controllers.go
|　　　|　　  ├── master_controllers.go
|　　　|　　  └── routers.go
|　　　├── error
|　　　|　　  └── errors.go
|　　　├── infrastracture
|　　　|　　  ├── db
|　　　|　　  |　　　├── db.go
|　　　|　　  |　　　└── tx.go
|　　　|　　  ├── repository
|　　　|　　  |　　　└── account_repository.go
|　　　|　　  └── security
|　　　|　　   　　　└── security.go
|　　　├── model
|　　　| 　　 └── accounts.go
|　　　├── test
|　　　|　　  └── master_test.go
|　　　└── usecase
|　　　 　　  ├── dto
|　　　 　　  |　　　├── app_dto.go
|　　　 　　  |　　　├── dto.go
|　　　 　　  |　　　└── master_dto.go
|　　　 　　  ├── app_service.go
|　　　  　　 └── master_service.go
├── .env
├── go.mod
├── go.sum
└── main.go
```

### モデル
**accounts**
| 論理名           | 物理名       |  主   | 型       | 必須  | 桁数  | 一意  | 備考                |
| ---------------- | ------------ | :---: | -------- | :---: | :---: | :---: | ------------------- |
| ID               | id           |  〇   | int      |  〇   |       |  〇   |                     |
| 作成日時         | created_at   |       | datetime |       |       |       |                     |
| 更新日時         | updated_at   |       | datetime |       |       |       |                     |
| 削除日時         | deleted_at   |       | datetime |       |       |       |                     |
| パスワード       | password     |       | varchar  |  〇   |  500  |       | SHA-256でハッシュ化 |
| アクセストークン | access_token |       | varchar  |       |  500  |       |                     |
| アカウント名     | name         |       | varchar  |  〇   |  50   |       |                     |
| メールアドレス   | email        |       | varchar  |  〇   |  50   |       |                     |

### ユースケース
* IDはgorm.ModelのAIをそのまま利用
* 返すアクセストークンはJWTなどで生成せず、ハードコードして一応ハッシュ化


#### ログイン
| **Login** |                                                                                                  |
| --------- | ------------------------------------------------------------------------------------------------ |
| メソッド  | POST                                                                                             |
| URI       | /api/login                                                                                       |
| 説明      | IDとパスワードを入力してログインする。IDとパスワードが一致すれば、アクセストークンが発行される。 |

```
{
  "id": 1,
  "password": password
}
```

#### アカウント取得(1件)
| **GetAccount** |                           |
| -------------- | ------------------------- |
| メソッド       | GET                       |
| URI            | /api/master/:id           |
| 説明           | アカウントを1件取得する。 |

#### アカウント取得(全件)
| **GetAccountAll** |                            |
| ----------------- | -------------------------- |
| メソッド          | GET                        |
| URI               | /api/master                |
| 説明              | アカウントを全件取得する。 |

#### アカウント登録
| **RegisterAccount** |                        |
| ------------------- | ---------------------- |
| メソッド            | POST                   |
| URI                 | /api/master            |
| 説明                | アカウントを登録する。 |

```
{
  "name": "山田 太郎",
  "password": "Yamada_Taro01",
  "email": "yamada@example.com"
}
```

#### アカウント編集
| **EditAccount** |                        |
| --------------- | ---------------------- |
| メソッド        | PUT                    |
| URI             | /api/master/edit       |
| 説明            | アカウントを編集する。 |

```
{
  "id": "1",
  "name": "山田 太郎",
  "email": "yamada@sample.com"
}
```

#### アカウント削除
| **DeleteAccount** |                            |
| ----------------- | -------------------------- |
| メソッド          | DELETE                     |
| URI               | /api/master/delete         |
| 説明              | アカウントを論理削除する。 |

```
{
  "id": "1",
}
```

### 参考
* [ddd_on_golang_sample](https://github.com/yu-croco/ddd_on_golang_sample)
* [ざっくりDDD・クリーンアーキテクチャにおける各層の責務を理解したい①（ドメイン層・ユースケース層編）](https://qiita.com/kotobuki5991/items/22712c7d761c659a784f)
* [ドメイン層の実装](https://terasolunaorg.github.io/guideline/current/ja/ImplementationAtEachLayer/DomainLayer.html)
* [sample-boot-jpa](https://github.com/jkazama/sample-boot-jpa)