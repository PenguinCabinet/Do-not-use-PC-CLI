# 💤Do-not-use-PC-CLI
これは指定時間帯にのみ、パソコンを使用できるようにするコマンドラインツールです。   
## 仕様
指定された時間帯以外になると自動的にシャットダウンします。   
「Windows起動時に自動起動」にこのプログラムを設定しておくと、手動で再度起動してしまっても、強制的にシャットダウンされます。   
現在、Windowsのみで動作します。   

## 🔽インストール
```
git clone https://github.com/PenguinCabinet/Do-not-use-PC-CLI
cd Do-not-use-PC-CLI
go build -ldflags -H=windowsgui
New-Item setting.yaml
```
setting.yamlにPCを使用できる時間帯を指定してください。   

例:   
* 水曜日以外では、8:00から12:00、14:00から20:00の間、パソコンを使用できるようにします。(それ以外の時間帯はパソコンを使用できません)   
* 水曜日では、8:00から12:00、14:00から22:00の間、パソコンを使用できるようにします。(それ以外の時間帯はパソコンを使用できません)    
```yaml
rules:
  - 
    if: 
      weeks: ["Mon","Tue","Thu","Fri","Sat","Sun"]
    apply:
      allowtimes:
        -
          start: 
            hours: 8
            minutes: 0
          end: 
            hours: 12
            minutes: 0
        - 
          start: 
            hours: 14
            minutes: 0
          end: 
            hours: 20
            minutes: 0
  - 
    if: 
      weeks: ["Wed"]
    apply:
      allowtimes:
        -
          start: 
            hours: 8
            minutes: 0
          end: 
            hours: 12
            minutes: 0
        - 
          start: 
            hours: 14
            minutes: 0
          end: 
            hours: 22
            minutes: 0
```

Windowsの起動と同時に実行するアプリとして、タスクスケジューラにDo-not-use-PC-CLI.exeを追加してください。
