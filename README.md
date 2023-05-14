# 💤Do-not-use-PC-CLI
これは指定時間帯に、パソコンを使用できなくするコマンドラインツールです。   
## ℹ仕様
指定された時間帯になると自動的にシャットダウンします。   
「Windows起動時に自動起動」にこのプログラムを設定しておくと、手動で起動してしまっても、強制的にシャットダウンされます。   
現在、Windowsのみで動作します。   

## 🔽インストール
```
git clone https://github.com/PenguinCabinet/Do-not-use-PC-CLI
cd Do-not-use-PC-CLI
go build -ldflags -H=windowsgui
New-Item setting.txt
```
setting.txtにPCを使用できなくする時間帯を指定してください。   
例:15:00から22:00の間、パソコンを使用できなくします。
```
15:00
22:00
```
例:22:00から**翌朝**の6:00の間、パソコンを使用できなくします。
```
22:00
6:00
```

[Windowsの起動時に自動的に実行するアプリとして、Do-not-use-PC-CLI.exeを追加してください。](https://support.microsoft.com/ja-jp/windows/windows-10-%E3%81%AE%E8%B5%B7%E5%8B%95%E6%99%82%E3%81%AB%E8%87%AA%E5%8B%95%E7%9A%84%E3%81%AB%E5%AE%9F%E8%A1%8C%E3%81%99%E3%82%8B%E3%82%A2%E3%83%97%E3%83%AA%E3%82%92%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B-150da165-dcd9-7230-517b-cf3c295d89dd#:~:text=%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%AE%E5%A0%B4%E6%89%80%E3%82%92%E9%96%8B,%E3%81%AB%E8%B2%BC%E3%82%8A%E4%BB%98%E3%81%91%E3%81%BE%E3%81%99%E3%80%82)    
スタートアップ フォルダーにDo-not-use-PC-CLI.exeのショートカットをおいて下さい。   
