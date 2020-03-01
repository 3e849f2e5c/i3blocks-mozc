# i3blocks Mozc

`i3blocks`に`Mozc`ステータスを表示

## 推奨環境

- fcitx-mozc
- [fcitx-dbus-status](https://github.com/clear-code/fcitx-dbus-status)

## 使用例

`i3blocks`の設定ファイルで

```
[mozc_status]
command=~/.config/i3blocks/blocks/mozc_status
label=<span color="#cc294c"> </span>
interval=persist
MOZC_DISPLAY_MODE=0
MOZC_DISPLAY_ROMAJI=0
min_width=" あ"
```

![](/preview/preview.png)

# i3blocks Mozc

Shows the status of Mozc in i3blocks

## Requirements

- fcitx-mozc
- [fcitx-dbus-status](https://github.com/clear-code/fcitx-dbus-status)

## Example

In i3blocks config file

```
[mozc_status]
command=~/.config/i3blocks/blocks/mozc_status
label=<span color="#cc294c"> </span>
interval=persist
MOZC_DISPLAY_MODE=0
MOZC_DISPLAY_ROMAJI=0
min_width=" あ"
```

![](/preview/preview.png)
